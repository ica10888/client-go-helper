package kubectl

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	_ "k8s.io/kubernetes/pkg/kubectl/cmd/cp"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	_ "unsafe"
)

func (i *Pod) Cp(srcPath string, destPath string) (error) {
	restconfig, err, coreclient := InitRestClient()

	reader, writer := io.Pipe()
	// strip trailing slash (if any)
	if destPath != "/" && strings.HasSuffix(string(destPath[len(destPath)-1]), "/") {
		destPath = destPath[:len(destPath)-1]
	}
	if err := checkDestinationIsDir(i, destPath); err == nil {
		// If no error, destPath was found to be a directory.
		// Copy specified src into it
		destPath = destPath + "/" + path.Base(srcPath)
	}
	go func() {
		defer writer.Close()
		err := cpMakeTar(srcPath, destPath, writer)
		cmdutil.CheckErr(err)
	}()
	var cmdArr []string

	cmdArr = []string{"tar", "-xf", "-"}
	destDir := path.Dir(destPath)
	if len(destDir) > 0 {
		cmdArr = append(cmdArr, "-C", destDir)
	}
	// Prepare the API URL used to execute another process within the Pod.  In
	// this case, we'll run a remote shell.
	req := coreclient.RESTClient().
		Post().
		Namespace(i.Namespace).
		Resource("pods").
		Name(i.Name).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: i.ContainerName,
			Command:   cmdArr,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(restconfig, "POST", req.URL())
	if err != nil {
		log.Fatalf("error %s\n", err)
		return err
	}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  reader,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    false,
	})
	if err != nil {
		log.Fatalf("error %s\n", err)
		return err
	}
	return nil
}


func checkDestinationIsDir(i *Pod, destPath string) error {
	return i.Exec([]string{"test", "-d", destPath})
}

//go:linkname cpMakeTar k8s.io/kubernetes/pkg/kubectl/cmd/cp.makeTar
func cpMakeTar(srcPath, destPath string, writer io.Writer) error


func (i *Pod) Cp2(srcPath string, destPath string) (error) {
	clientset, _ := InitClient()
	restconfig, _, _ := InitRestClient()
	u := clientset.CoreV1().RESTClient().Get().
		Namespace(i.Namespace).
		Name(i.Name).
		Resource("pods").
		SubResource("exec").
		Param("command", "/bin/cat").
		Param("command", srcPath).
		Param("container", i.ContainerName).
		Param("stderr", "true").
		Param("stdout", "true").URL()

	switch u.Scheme {
	case "https":
		u.Scheme = "wss"
	case "http":
		u.Scheme = "ws"
	default:
		log.Fatalf("malformed URL %s", u.String())
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	tlsConfig, err := rest.TLSConfigFor(restconfig)
	if err != nil {
		log.Fatalf("failed to read tls config: %v", err)
	}

	dialer := &websocket.Dialer{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: tlsConfig,
	}

	c, _, err := dialer.Dial(u.String(), req.Header)
	if err != nil {
		log.Fatalf("failed to do req: %v", err)
	}
	defer c.Close()

	var res []byte
	for {
		msgT, p, err := c.ReadMessage()
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				break
			}
			fmt.Printf("err %T %v\n", err, err)
			break
		}
		if msgT != 2 {
			log.Fatalf("unknown message type %d", msgT)
		}
		res = append(res, p...)
	}
	fmt.Printf("body:\n%s\n", res)
	return nil
}
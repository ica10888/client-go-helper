package kubectl

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"

	_ "k8s.io/kubernetes/pkg/kubectl/cmd/cp"

)


//kubectl  describe pods
// from  https://stackoverflow.com/questions/40764400/rest-api-alternative-for-describe-command
// a describe is actually a combination of results from the pod and the events APIs:
func (i *Pod) Describe()(string ,error) {
	// Status
	podGetOpts := metav1.GetOptions{}
	//podLogOpts.Container = containerMeta.ContainerName
	clientset, e := InitClient()
	if e != nil{
		fmt.Errorf("something wrong happend ,%s",e)
	}
	pod,err := clientset.CoreV1().Pods(i.Namespace).Get(i.PodName, podGetOpts)
	if err != nil {
		log.Fatalf( "error in get pod events")
	}
	//parsing client-go data structures into K8s yaml spec
	json, err := json.Marshal(pod)
	if err != nil {
		fmt.Errorf("json Unmarshal err")
	}
	data ,_ := yaml.JSONToYAML(json)
	fmt.Print(string(data))
	// Events
	// Instantiate loader for kubeconfig file.
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	// Get a rest.Config from the kubeconfig file.  This will be passed into all
	// the client objects we create.
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return  "",err
	}

	// Create a Kubernetes core/v1 client.
	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		return  "",err
	}

	// Prepare the API URL used to execute another process within the Pod.  In
	// this case, we'll run a remote shell.
	req := coreclient.RESTClient().
		Get().
		Namespace(i.Namespace).
		Resource("events").
		Param("fieldSelector","involvedObject.name="+ i.PodName)

	rawJson ,_:= req.Do().Raw()
	rawYaml ,_ := yaml.JSONToYAML(rawJson)
	return string(rawYaml),nil
}

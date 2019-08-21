package kubectl

import (
	"github.com/ica10888/client-go-helper/pkg/kubectl/client"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

//kubectl  describe pods
// from  https://stackoverflow.com/questions/40764400/rest-api-alternative-for-describe-command
// a describe is actually a combination of results from the pod and the events APIs:
func (i *pod) Describe() (string, error) {
	// Status
	podGetOpts := metav1.GetOptions{}
	//podLogOpts.Container = containerMeta.ContainerName
	clientset, e := client.InitClient()
	if e != nil {
		return "", fmt.Errorf("something wrong happend ,%s", e)
	}
	pod, err := clientset.CoreV1().Pods(i.Namespace).Get(i.Name, podGetOpts)
	if err != nil {
		log.Fatalf("error in get pod events")
	}
	//parsing client-go data structures into K8s yaml spec
	json, err := json.Marshal(pod)
	if err != nil {
		return "", fmt.Errorf("json Unmarshal err")
	}
	data, _ := yaml.JSONToYAML(json)
	str1 := string(data)
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
		return "", err
	}

	// Create a Kubernetes core/v1 client.
	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		return "", err
	}

	// Prepare the API URL used to execute another process within the Pod.  In
	// this case, we'll run a remote shell.
	req := coreclient.RESTClient().
		Get().
		Namespace(i.Namespace).
		Resource("events").
		Param("fieldSelector", "involvedObject.name="+i.Name)

	rawJson, _ := req.Do().Raw()
	rawYaml, _ := yaml.JSONToYAML(rawJson)
	str2 := string(rawYaml)
	str := str1 + str2
	return str, nil
}

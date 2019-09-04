package kubectl

import (
	"github.com/ghodss/yaml"
	"github.com/ica10888/client-go-helper/pkg/kubectl/client"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)


func (i *pod) Describe() (coreV1.Pod, coreV1.EventList , error)  {
	//Pod
	podGetOpts := metav1.GetOptions{}
	clientset, err := client.InitClient()
	if err != nil {
		return coreV1.Pod{} ,coreV1.EventList{}, err
	}

	pod , err := clientset.
		CoreV1().
		Pods(i.Namespace).
		Get(i.Name, podGetOpts)

	if err != nil {
		return  coreV1.Pod{} ,coreV1.EventList{}, err
	}

	// Events
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return  *pod ,coreV1.EventList{}, err
	}

	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		return *pod ,coreV1.EventList{}, err
	}

	req := coreclient.
		RESTClient().
		Get().
		Namespace(i.Namespace).
		Resource("events").
		Param("fieldSelector", "involvedObject.name="+i.Name)

	eventsJson, _ := req.Do().Raw()
	eventsYaml, _ := yaml.JSONToYAML(eventsJson)
	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode([]byte(eventsYaml), nil, nil)
	events := obj.(*coreV1.EventList)
	if err != nil {
		return *pod ,coreV1.EventList{}, err
	}

	return *pod ,*events,nil
}

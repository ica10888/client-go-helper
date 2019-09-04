package kubectl

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/ica10888/client-go-helper/pkg/kubectl/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"k8s.io/client-go/kubernetes/scheme"
	coreV1 "k8s.io/api/core/v1"
)


func (i *pod) Describe() (string, error) {
	//Pod
	podGetOpts := metav1.GetOptions{}
	clientset, e := client.InitClient()
	if e != nil {
		return "", fmt.Errorf("something wrong happend ,%s", e)
	}

	_, err := clientset.
		CoreV1().
		Pods(i.Namespace).
		Get(i.Name, podGetOpts)

	if err != nil {
		log.Fatalf("error in get pod events")
	}

	// Events
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return "", err
	}

	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		return "", err
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
	log.Print(events)

	if err != nil {
		return "",err
	}


	return "", nil
}

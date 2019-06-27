package kubectl

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
	corev1 "k8s.io/api/core/v1"
)

func TestLogs(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	opts := corev1.PodLogOptions{}
	str,e := pod.Logs(&opts)
	if e !=nil {
		log.Print(e)
	}else {
		log.Print(str)
	}

}

func TestDescribe(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	str,e :=pod.Describe()
	if e !=nil {
		log.Print(e)
	}
	log.Print(str)
}

func TestExec(t *testing.T) {

	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	e := pod.Exec([]string{"ls","-al"})
	if e !=nil {
		log.Print(e)
	}
}

func TestCp(t *testing.T) {
	// kubectl cp /tmp/localfile  api-test-775cf487ff-7zhnj:/opt
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	e := pod.Cp("/tmp/localfile","/opt/")
		if e !=nil {
		log.Print(e)
	}
}


func TestGet(t *testing.T) {

	opts := v1.ListOptions{}
	kapi := Kubeapi{
		ApiVersion:   "extensions/v1beta1",
		Kind:  "Deployment",
		Yaml:   "",
	}
	//same result
	kapi2 := Kubeapi{
		ApiVersion:   "apps/v1beta2",
		Kind:  "Deployment",
		Yaml:   "",
	}
	items ,e := Get(&opts,&kapi,"dev")
	if e !=nil {
		log.Print(e)
	}
	log.Print(items)
	items2 ,e := Get(&opts,&kapi2,"dev")
	if e !=nil {
		log.Print(e)
	}
	log.Print(items2)
}
package kubectl

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestLogs(t *testing.T) {
	pod := pod{
		ContainerName: "",
		Name:          "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	opts := corev1.PodLogOptions{}
	str, err := pod.Logs(&opts)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(str)
	}

}

func TestDescribe(t *testing.T) {
	pod := Pod("hy-demo-front-686df5c58f-4v8nk","dev")
	pods, events, err := pod.Describe()
	if err != nil {
		log.Print(err)
	}
	log.Println(pods)
	log.Println(events)
	log.Println("=======================")
	node := Node("node13")
	nodes, events, err := node.Describe()
	if err != nil {
		log.Print(err)
	}
	log.Println(nodes)
	log.Println(events)



}

func TestExec(t *testing.T) {

	pod := pod{
		ContainerName: "",
		Name:          "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	err := pod.Exec([]string{"ls", "-al"})
	if err != nil {
		log.Print(err)
	}
}

func TestCp(t *testing.T) {
	// kubectl cp /tmp/local  api-test-77c6f9bf8c-nhhp5:/opt -n dev
	pod :=Pod("api-test-77c6f9bf8c-nhhp5","dev")
	err := pod.Cp().ToPod("/tmp/local", "/opt")
	if err != nil {
		log.Print(err)
	}
	log.Println("=======================")
	// kubectl cp  api-test-775cf487ff-7zhnj:/opt/app.jar /tmp
	err2 := pod.Cp().FromPod("/opt/app.jar", "/tmp")
	if err2 != nil {
		log.Print(err)
	}
}

func TestGetAll(t *testing.T) {
	pod1 := pod{
		ContainerName: "",
		Name:          "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	opts := v1.ListOptions{}
	items, err := pod1.GetAll(&opts)
	if err != nil {
		log.Print(err)
	}
	for _, v := range items {
		json, _ := json.Marshal(v)
		rawYaml, _ := yaml.JSONToYAML(json)
		log.Println(string(rawYaml))
	}
	log.Println("=======================")
	pod2 := pod{
		ContainerName: "",
		Name:          "api-test",
		Namespace:     "dev",
	}
	items2, err := pod2.GetAll(&opts)
	if err != nil {
		log.Print(err)
	}
	for _, v := range items2 {
		json, _ := json.Marshal(v)
		rawYaml, _ := yaml.JSONToYAML(json)
		log.Println(string(rawYaml))
	}
}

func TestGet(t *testing.T) {
	pod := pod{
		ContainerName: "",
		Name:          "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	opts := v1.GetOptions{}
	v1Pod, err := pod.Get(&opts)
	if err != nil {
		log.Print(err)
	}
	json, _ := json.Marshal(v1Pod)
	rawYaml, _ := yaml.JSONToYAML(json)
	log.Println(string(rawYaml))

}

/*
func TestTop(t *testing.T) {
	node := Node("node13")
	nodeMetrics, err := node.Top()
	if err != nil {
		log.Print(err)
	}
	log.Println(nodeMetrics)

}
*/

func TestScale(t *testing.T) {
	deployment := deployment{
		Name:          "api-test",
		Namespace:     "dev",
	}
	err := deployment.Scale(2)
	if err != nil {
		log.Print(err)
	}

}

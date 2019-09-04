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
	// kubectl cp /tmp/localfile  api-test-775cf487ff-7zhnj:/opt
	/*	pod := pod{
			ContainerName: "",
			Name:          "api-test-775cf487ff-7kmk2",
			Namespace:     "dev",
		}
		err := pod.Cp().ToPod("/opt/", "/tmp/localfile")
		if err != nil {
			log.Print(err)
		}*/

}

func TestCp2(t *testing.T) {
	// kubectl cp /tmp/localfile  api-test-775cf487ff-7zhnj:/opt
/*	pod := pod{
		ContainerName: "",
		Name:          "www-6df84b8b6f-w7lv6",
		Namespace:     "dev",
	}
	err := pod.Cp().FromPod("/tmp/", "/usr/src/app/web/server.js")
	if err != nil {
		log.Print(err)
	}
*/
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

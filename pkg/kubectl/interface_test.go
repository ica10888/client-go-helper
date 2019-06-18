package kubectl

import (
	"log"
	"testing"
)

func TestLogs(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		PodName:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	str,e := pod.Logs(false)
	if e !=nil {
		log.Print(e)
	}else {
		log.Print(str)
	}

}

func TestDescribe(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		PodName:       "api-test-775cf487ff-7zhnj",
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
		PodName:       "api-test-775cf487ff-7zhnj",
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
		PodName:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	e := pod.Cp("/tmp/localfile","/opt/")
		if e !=nil {
		log.Print(e)
	}
}



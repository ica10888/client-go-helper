package kubectl

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	corev1 "k8s.io/api/core/v1"

)



type podApi interface {
	coreApi
	Cp()
	Exec(cmd []string) error
	Logs(podLogOpts *corev1.PodLogOptions) (string, error)
	//Top()
	//Scale()
}


type nodeApi interface {
	coreApi
	//Top()
	//Drain()
	//Taint()
	//Cordon()
	//Uncordon()

}

//4 types of Rereplica Set , Node and Pod
type coreApi interface {
	baseKubeApi
	Describe() (interface{}, corev1.EventList , error)
}

type baseKubeApi interface {
	Get() error
	GetAll(opts *v1.ListOptions) (interface{}, error)
	Delete(opts *v1.DeleteOptions) error
	Patch(data string, pt *types.PatchType) (interface{}, error)
	//Label()
	//Annotate()
}

type kubectlApi interface {
	Apply(yaml string, namespace string) error
	Create(yaml string, namespace string) error
}


//used for apply/create
type kubeapi struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Yaml       string
}

//ckubectl cp
type copyer struct {
	pod *pod
}

func (i *pod) Cp() copyer {
	return copyer{i}
}

func (i copyer) ToPod(srcPath string, destPath string) error {
	return i.pod.copyToPod(srcPath, destPath)
}

func (i copyer) FromPod(srcPath string, destPath string) error {
	return i.pod.copyFromPod(srcPath, destPath)
}

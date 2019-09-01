package kubectl

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type delete interface {
	Delete(opts *v1.DeleteOptions) error
}

type describe interface {
	Describe() (string, error)
}

type get interface {
	Get() error
}

type exec interface {
	Exec(cmd []string) error
}

type cp interface {
	Cp(srcPath string, destPath string) error
}

type logs interface {
	Logs(podLogOpts *corev1.PodLogOptions) (string, error)
}

//used for apply/create
type kubeapi struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Yaml       string
}

//ckubectl cp
type Copyer struct {
	srcPath string
	pod     *pod
}

func (i *pod) Cp(srcPath string) Copyer {
	return Copyer{srcPath, i}
}

func (i *Copyer) ToPod(destPath string) error {
	return i.pod.copyToPod(i.srcPath, destPath)
}

func (i *Copyer) FromPod(destPath string) error {
	return i.pod.copyFromPod(i.srcPath, destPath)
}

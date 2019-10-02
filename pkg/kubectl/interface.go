package kubectl


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

package kubectl

type Pod struct {
	PodName string
	ContainerName string
	Namespace     string
}


type Describe interface {
	Describe()(string ,error)
}




package kubectl

type Pod struct {
	PodName string
	ContainerName string
	Namespace     string
}


type describe interface {
	Describe()(string ,error)
}




package pkg

type ContainerMeta struct {
	ContainerName string
	PodName       string
	Namespace     string
}

type Interface interface {

	Apply(yaml string,namespace string)  (error) //kubectl apply -f

	Logs(containerMeta *ContainerMeta) (error)  //kubectl logs

	LogsWithPrevious(containerMeta *ContainerMeta) (error)  //kubectl logs  --previous

	Describe(containerMeta *ContainerMeta) (error) //kubectl describe pods

	Exec(containerMeta *ContainerMeta, cmd []string) (error) //kubectl exec

	Cp(containerMeta *ContainerMeta,srcPath string,destPath string) (error) //kubectl cp
}


type kubectl struct {}





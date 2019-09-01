package kubectl

func Pod(name string, namespace string, ContainerName string) pod {
	return pod{name, namespace, ContainerName}
}

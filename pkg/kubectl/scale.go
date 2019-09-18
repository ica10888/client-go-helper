package kubectl

import (
	"github.com/ica10888/client-go-helper/pkg/kubectl/client"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
)

func (i *deployment) Scale(opts *autoscalingv1.Scale) ( error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	_, err = clientset.AppsV1().Deployments(i.Namespace).UpdateScale(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *replicaSet) Scale(opts *autoscalingv1.Scale) (error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	_, err = clientset.AppsV1().ReplicaSets(i.Namespace).UpdateScale(i.Name, opts)
	if err != nil {
		return  err
	}
	return  nil
}

func (i *statefulSet) Scale(opts *autoscalingv1.Scale) (error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	_, err = clientset.AppsV1().StatefulSets(i.Namespace).UpdateScale(i.Name, opts)
	if err != nil {
		return  err
	}
	return  nil
}

package kubectl

import (
	"github.com/ica10888/client-go-helper/pkg/kubectl/client"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *deployment) Scale(replicas int32) ( error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	opts, err := clientset.AppsV1().Deployments(i.Namespace).GetScale(i.Name,v1.GetOptions{})
	opts.Spec.Replicas = replicas;
	_, err = clientset.AppsV1().Deployments(i.Namespace).UpdateScale(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *replicaSet) Scale(replicas int32) ( error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	opts, err := clientset.AppsV1().ReplicaSets(i.Namespace).GetScale(i.Name,v1.GetOptions{})
	opts.Spec.Replicas = replicas;
	_, err = clientset.AppsV1().ReplicaSets(i.Namespace).UpdateScale(i.Name, opts)
	if err != nil {
		return  err
	}
	return  nil
}

func (i *statefulSet) Scale(replicas int32) ( error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	opts, err := clientset.AppsV1().StatefulSets(i.Namespace).GetScale(i.Name,v1.GetOptions{})
	opts.Spec.Replicas = replicas;
	_, err = clientset.AppsV1().StatefulSets(i.Namespace).UpdateScale(i.Name, opts)
	if err != nil {
		return  err
	}
	return  nil
}

func (i *replicationController) Scale(replicas int32) ( error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return  err
	}
	opts, err := clientset.CoreV1().ReplicationControllers(i.Namespace).GetScale(i.Name,v1.GetOptions{})
	opts.Spec.Replicas = replicas;
	_, err = clientset.CoreV1().ReplicationControllers(i.Namespace).UpdateScale(i.Name, opts)
	if err != nil {
		return  err
	}
	return  nil
}


package kubectl

import (
	"bytes"
	"io"
	corev1 "k8s.io/api/core/v1"
)

func (i *pod) Logs(podLogOpts *corev1.PodLogOptions) (string, error) {
	podLogOpts.Container = i.ContainerName
	clientset, err := InitClient()
	if err != nil {
		return "", err
	}
	req := clientset.CoreV1().Pods(i.Namespace).GetLogs(i.Name, podLogOpts)
	podLogs, err := req.Stream()
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", err
	}
	str := buf.String()
	return str, nil
}

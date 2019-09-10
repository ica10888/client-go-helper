package kubectl


import (
	"k8s.io/client-go/tools/clientcmd"
	 "k8s.io/metrics/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func (i *pod) Top() {
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}

	metrics, err := versioned.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	metrics.MetricsV1beta1().NodeMetricses().Get("your node name", metav1.GetOptions{})
	metrics.MetricsV1beta1().NodeMetricses().List(metav1.ListOptions{})
	metrics.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(metav1.ListOptions{})
	metrics.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).Get("your pod name", metav1.GetOptions{})
}
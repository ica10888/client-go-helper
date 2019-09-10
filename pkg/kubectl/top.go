package kubectl

//looping trying to add package
/*
import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)



func (i *pod) Top()( v1beta1.PodMetrics,error) {
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return v1beta1.PodMetrics{} ,err
	}

	metrics, err := versioned.NewForConfig(restconfig)
	if err != nil {
		return  v1beta1.PodMetrics{} ,err
	}

	podMetrics, err := metrics.
		MetricsV1beta1().
		PodMetricses(i.Namespace).
		Get(i.Name, metav1.GetOptions{})

	if err != nil {
		return v1beta1.PodMetrics{} ,err
	}
	return *podMetrics , nil
}



func (i *node) Top ()(v1beta1.NodeMetrics,error) {
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return v1beta1.NodeMetrics{} ,err
	}

	metrics, err := versioned.NewForConfig(restconfig)
	if err != nil {
		return  v1beta1.NodeMetrics{} ,err
	}

	nodeMetrics, err := metrics.
		MetricsV1beta1().
		NodeMetricses().
		Get(i.Name, metav1.GetOptions{})

	if err != nil {
		return v1beta1.NodeMetrics{} ,err
	}
	return *nodeMetrics , nil
}
*/
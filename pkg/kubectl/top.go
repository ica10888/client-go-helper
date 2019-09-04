package kubectl


/*import (
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func top() {
	var kubeconfig, master string //empty, assuming inClusterConfig
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil{
		panic(err)
	}

	mc, err := metrics.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	mc.MetricsV1beta1().NodeMetricses().Get("your node name", metav1.GetOptions{})
	mc.MetricsV1beta1().NodeMetricses().List(metav1.ListOptions{})
	mc.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(metav1.ListOptions{})
	mc.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).Get("your pod name", metav1.GetOptions{})
}*/
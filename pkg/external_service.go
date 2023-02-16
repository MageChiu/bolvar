package pkg

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
)

func buildKubeClient(kubeConf string) kubernetes.Interface {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConf)
	if err != nil {
		//panic(err.Error())
		klog.Warningf("build kube client error: %v", err)
		return nil
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func genService(name, namespace string, kubeClient kubernetes.Interface,
	destPath string) {
	clusterService := corev1.Service{}
	clusterService.Name = "test-gro"
	klog.Infof("svr:%v", clusterService)
	data, err := yaml.Marshal(clusterService)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s.yaml", destPath, "work-test"), data, 0666)
	if err != nil {
		panic(err)
	}

	//encoder := yamlutil.
	w, _ := os.Create(fmt.Sprintf("%s/%s.yaml", destPath, "work-test-2"))
	encoder := yaml.NewEncoder(w)
	encoder.Encode(clusterService)
}

func createExternalService(ctx context.Context,
	serviceName, externalServiceName,
	namespace string, kubeClient kubernetes.Interface) error {
	clusterService, err := kubeClient.CoreV1().Services(namespace).Get(ctx, serviceName, metav1.GetOptions{
		TypeMeta:        metav1.TypeMeta{},
		ResourceVersion: "",
	})
	if err != nil {
		return err
	}
	if clusterService.Spec.Type == corev1.ServiceTypeNodePort {
		return nil
	}

}

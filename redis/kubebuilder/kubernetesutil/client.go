package kubernetesutil

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func generateK8sConfig() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here
	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	return kubeConfig.ClientConfig()
}

func GenerateKubeClient() (kubernetes.Clientset, error) {
	config, cfgError := generateK8sConfig()
	if cfgError != nil {
		return kubernetes.Clientset{}, cfgError
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return kubernetes.Clientset{}, err
	}
	return *client, nil
}

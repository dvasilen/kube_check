package main

import (
	"fmt"

	// "k8s.io/kubernetes/pkg/client/clientcmd"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned"
	// "k8s.io/kubernetes/pkg/kubectl"
	"k8s.io/kubernetes/pkg/util/intstr"

	// "k8s.io/kubernetes/pkg/client/unversioned/c"
	// "k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	// cmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
)

// var CLIENT *unversioned.Client

func main() {
	// clientConfig := restclient.Config{}
	// clientConfig.Host = hostIp
	// clientConfig.
	/*	namespace := "zenlin-d"
		clientContext := cmdapi.NewContext()
		clientContext.Namespace = namespace
		// cl
		cmdapiConfig := cmdapi.NewConfig()

		// unversioned.config{}
		// api.

		clientConfig := clientcmd.NewNonInteractiveClientConfig(*cmdapiConfig, "default", &clientcmd.ConfigOverrides{
			Context: *clientContext,
		})
		// config, err := a.ClientConfig()
		// config.Host
		// api.
		// kubectl.SetOriginalConfiguration(info, original)
		unversionedConfig, err := clientConfig.ClientConfig()
		unversionedConfig.Host = "127.0.0.1:8080"
		client, err := unversioned.New(unversionedConfig)

		// client.RESTClient.Post().
		if err != nil {
			fmt.Printf("New unversioned client err: %v!\n", err.Error())
		}
	*/

	namespace := "aaaa"
	clientConfig := restclient.Config{}
	clientConfig.Host = "127.0.0.1:8080"
	client, err := unversioned.New(&clientConfig)
	if err != nil {
		fmt.Println(err)
	}

	service := &api.Service{
		ObjectMeta: api.ObjectMeta{
			Labels: map[string]string{},
		},
		Spec: api.ServiceSpec{
			Selector: map[string]string{},
		},
	}

	service.Spec.Ports = make([]api.ServicePort, 1)
	service.ObjectMeta.SetName("zenlin-default")
	// service.ObjectMeta.SetNamespace(api.NamespaceDefault)
	service.ObjectMeta.SetNamespace(namespace)
	service.ObjectMeta.Labels["app"] = namespace
	service.Spec.Ports[0] = api.ServicePort{Port: 6379, TargetPort: intstr.FromString("default")}
	service.Spec.Selector["app"] = namespace

	// client.ConfigMaps
	_, err = client.Namespaces().Get(namespace)
	if err != nil {
		namespaceObj := &api.Namespace{
			ObjectMeta: api.ObjectMeta{Name: namespace},
		}
		if _, err := client.Namespaces().Create(namespaceObj); err != nil {
			fmt.Printf("Create namespace err : %v\n", err)
			return
		}
		fmt.Println("creating namespace successful")
	}

	if _, err := client.Services(namespace).Create(service); err != nil {
		fmt.Printf("Create service err : %v\n", err)
		// return err
	}

	// CLIENT = client
}

package main

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/docker/docker/api/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//_ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	//_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	//_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"github.com/docker/docker/client"
)

func main() {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	configMap, err := clientset.CoreV1().ConfigMaps("default").Get(context.TODO(), "sample-configmap", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("This is the config map: %v", configMap)

	imagesToPull := strings.Split(configMap.Data["image"], ",")

	//node, err := clientset.CoreV1().Nodes().Get(context.TODO(), "gke-dev-cluster-default-pool-f3dc0af3-2glc", metav1.GetOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(node)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	for i := range imagesToPull {
		out, err := cli.ImagePull(context.TODO(), imagesToPull[i], types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}

		defer out.Close()
		if _, err := ioutil.ReadAll(out); err != nil {
			panic(err)
		}
	}

}

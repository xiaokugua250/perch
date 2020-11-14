package tools

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"

	"io/ioutil"

	"k8s.io/apimachinery/pkg/api/meta"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"

	"k8s.io/client-go/restmapper"
)

func TestResourceConfigParser(t *testing.T) {
	kubeconfig := "k8s_dev.config"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	dd, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	err = ResourceConfigParser(*c, dd, "k8s_demo.yaml")
	if err != nil {
		log.Println("----<>", err)
	}

}

func TestResourceConfigParser2(t *testing.T) {

	b, err := ioutil.ReadFile("k8s_demo.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q \n", string(b))

	config, err := clientcmd.BuildConfigFromFlags("", "k8s_dev.config")
	if err != nil {
		log.Fatal(err)
	}

	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	dd, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(b), 100)
	for {
		var rawObj runtime.RawExtension
		if err = decoder.Decode(&rawObj); err != nil {
			log.Println("err in decode rawobj", err)
			fmt.Printf(" decode result is %+v\n", rawObj)
			break
		}

		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			log.Fatal(err)
		}

		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		gr, err := restmapper.GetAPIGroupResources(c.Discovery())
		if err != nil {
			log.Fatal(err)
		}

		mapper := restmapper.NewDiscoveryRESTMapper(gr)
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			log.Fatal(err)
		}

		var dri dynamic.ResourceInterface
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace("default")
			}
			dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dd.Resource(mapping.Resource)
		}

		if _, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{}); err != nil {
			log.Fatal(err)
		}
	}
	if err != io.EOF {
		log.Fatal("eof ", err)
	}
}

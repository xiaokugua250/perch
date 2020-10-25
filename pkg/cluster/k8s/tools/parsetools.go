/**
目标是通过解析yaml 构造k8s 集群资源对象
https://stackoverflow.com/questions/58783939/using-client-go-to-kubectl-apply-against-the-kubernetes-api-directly-with-mult
 */
package tools

import (
	"bytes"
	"context"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/restmapper"
)



/**
解析YAML格式的配置信息，并利用k8sclient 进行资源创建

 */
func ResourceConfigParser(client kubernetes.Clientset,dynamicClient dynamic.Interface,yamlfile string) ( err error){


	fileBytes,err:= ioutil.ReadFile(yamlfile)
	if err!= nil{
		return err
	}

	log.Printf("%q \n", string(fileBytes))

	yamlDecoder:= yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(fileBytes),4096)
	for {
		var rawObj runtime.RawExtension
		if err = yamlDecoder.Decode(&rawObj); err != nil {
			break
		}
		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return err
		}

		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		gr, err := restmapper.GetAPIGroupResources(client.Discovery())
		if err != nil {
		return err
		}

		mapper := restmapper.NewDiscoveryRESTMapper(gr)
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			return err
		}

		var dri dynamic.ResourceInterface
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace("default")
			}
			dri = dynamicClient.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dynamicClient.Resource(mapping.Resource)
		}

		if _, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{}); err != nil {
			return err
		}
	}
	if err != io.EOF {
		return err
	}
	return nil



}

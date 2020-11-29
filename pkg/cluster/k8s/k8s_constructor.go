/**
直接通过yaml或json构造k8s 资源对象
ref:
1.https://github.com/kubernetes/client-go/issues/193
2.https://gist.github.com/pytimer/0ad436972a073bb37b8b6b8b474520fc
*/
package k8s

import "perch/pkg/cluster/k8s/tools"

/**

 */
func (k8sClientSet *ClientSet) K8SConstructorFileValidate(yamlFile string) error {

	return tools.ResourceConfigParser(*k8sClientSet.K8SClientSet, *k8sClientSet.K8sDynamitcClient, yamlFile)

}

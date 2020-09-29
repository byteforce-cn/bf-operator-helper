package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeConfigMap(instance interface{}, data map[string]string) *corev1.ConfigMap {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	name := objMeta.Name + ResourceTypePostfix(ResourceTypeConfigMap)
	wr := MakeOwnerReference(instance, name)
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       objMeta.Namespace,
			OwnerReferences: wr,
		},
		Data: data,
	}

	return cm
}

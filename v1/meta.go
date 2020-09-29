package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func MakeSimpleMetaObject(instance interface{}, wr []metav1.OwnerReference) metav1.ObjectMeta {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	labels := make(map[string]string, 0)
	objectMeta := makeMetaObj(objMeta.Name, objMeta.Namespace, labels, wr)
	return objectMeta
}

func MakeMetaObjectWithName(instance interface{}, name string, wr []metav1.OwnerReference) metav1.ObjectMeta {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	labels := make(map[string]string, 0)
	objectMeta := makeMetaObj(name, objMeta.Namespace, labels, wr)
	return objectMeta
}

func MakeMetaObjectWithNameAndLabels(instance interface{}, name string, labels map[string]string, wr []metav1.OwnerReference) metav1.ObjectMeta {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	objectMeta := makeMetaObj(name, objMeta.Namespace, labels, wr)
	return objectMeta
}

func makeMetaObj(name string, ns string, labels map[string]string, wr []metav1.OwnerReference) metav1.ObjectMeta {
	objectMeta := metav1.ObjectMeta{
		Name:            name,
		Namespace:       ns,
		OwnerReferences: wr,
		Labels:          labels,
	}
	return objectMeta
}

func MakeOwnerReference(instance interface{}, name string) []metav1.OwnerReference {
	typeMeta, objMeta := GetTypeMetaAndObjectMeta(instance)
	wr := []metav1.OwnerReference{
		{
			APIVersion: typeMeta.APIVersion,
			Kind:       typeMeta.Kind,
			Name:       name,
			UID:        objMeta.UID,
		},
	}
	return wr
}

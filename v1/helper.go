package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

func ResourceTypePostfix(resourceType ResourceType) string {
	if resourceType == ResourceTypeHeadlessService {
		return HeadlessServicePostfix
	}
	if resourceType == ResourceTypeService {
		return ServicePostfix
	}
	if resourceType == ResourceTypeConfigMap {
		return ConfigMapPostfix
	}
	if resourceType == ResourceTypeIngress {
		return IngressPostfix
	}
	return ""
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func GetTypeMetaAndObjectMeta(instance interface{}) (metav1.TypeMeta, metav1.ObjectMeta) {
	r := reflect.ValueOf(instance).Elem()
	typeMetaField := r.FieldByName("TypeMeta")
	objectMetaField := r.FieldByName("ObjectMeta")
	tmf := typeMetaField.Interface().(metav1.TypeMeta)
	om := objectMetaField.Interface().(metav1.ObjectMeta)
	return tmf, om
}

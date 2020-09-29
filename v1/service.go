package v1

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func HeadlessServiceExist(instance interface{}, client client.Client) bool {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	name := objMeta.Name + ResourceTypePostfix(ResourceTypeHeadlessService)
	return serviceExist(client, name, objMeta)
}

func ServiceExist(instance interface{}, client client.Client) bool {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	name := objMeta.Name + ResourceTypePostfix(ResourceTypeService)
	return serviceExist(client, name, objMeta);
}

func MakeHeadlessService(meta metav1.ObjectMeta, servicePorts []corev1.ServicePort, appSelector map[string]string) *corev1.Service {
	serviceSpec := makeServiceSpec(appSelector, servicePorts)
	serviceSpec.ClusterIP="None"
	service := makeService(meta, serviceSpec)
	return service
}

func MakeClusterIPService(meta metav1.ObjectMeta, servicePorts []corev1.ServicePort, appSelector map[string]string) *corev1.Service {

	serviceSpec := makeServiceSpec(appSelector, servicePorts)
	service := makeService(meta, serviceSpec)
	return service
}

func makeService(meta metav1.ObjectMeta, serviceSpec corev1.ServiceSpec) *corev1.Service {
	service := &corev1.Service{
		ObjectMeta: meta,
		Spec:       serviceSpec,
	}
	return service
}

func makeServiceSpec(appSelector map[string]string, servicePorts []corev1.ServicePort) corev1.ServiceSpec {
	serviceSpec := corev1.ServiceSpec{
		Selector: appSelector,
		Ports:    servicePorts,
	}
	return serviceSpec
}

func serviceExist(client client.Client, name string, objMeta metav1.ObjectMeta) bool {
	service := &corev1.Service{}
	if err := client.Get(context.Background(), types.NamespacedName{Name: name, Namespace: objMeta.Namespace}, service); err != nil {
		return false
	}
	return true
}

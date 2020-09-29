package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeRWOVolumnTemplates(instance interface{}, templateName string, spec StorageSpec) corev1.PersistentVolumeClaim {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)

	name := objMeta.Name + "-" + templateName + "-pvc"
	access := corev1.ReadWriteOnce
	pvc := makeVolumeTemplate(name, access, spec)
	return pvc
}

func MakeRWMVolumnTemplates(instance interface{}, templateName string, spec StorageSpec) corev1.PersistentVolumeClaim {
	_, objMeta := GetTypeMetaAndObjectMeta(instance)
	name := objMeta.Name + "-" + templateName + "-pvc"
	access := corev1.ReadWriteOnce
	pvc := makeVolumeTemplate(name, access, spec)
	return pvc
}

func makeVolumeTemplate(name string, access corev1.PersistentVolumeAccessMode, spec StorageSpec) corev1.PersistentVolumeClaim {
	pvc := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes:      []corev1.PersistentVolumeAccessMode{access},
			StorageClassName: &spec.Name,
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(spec.Size),
				},
			},
		},
	}
	return pvc
}

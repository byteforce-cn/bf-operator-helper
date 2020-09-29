package v1

import corev1 "k8s.io/api/core/v1"

func MakeContainerEnvs() []corev1.EnvVar {
	env := make([]corev1.EnvVar, 0)
	return env
}

func MakeContainerEnvVar(key string, val string) *corev1.EnvVar {
	env := &corev1.EnvVar{
		Name:  key,
		Value: val,
	}
	return env
}

func MakeContainerEnvPodIP(key string) *corev1.EnvVar {
	path := "status.podIP"
	return makeFieldRefEnvVar(key, path)
}

func MakeContainerEnvHostIP(key string) *corev1.EnvVar {
	path := "status.hostIP"
	return makeFieldRefEnvVar(key, path)
}

func MakeContainerEnvMetaName(key string) *corev1.EnvVar {
	path := "metadata.name"
	return makeFieldRefEnvVar(key, path)
}

func MakeContainerEnvMetaNameSpace(key string) *corev1.EnvVar {
	path := "metadata.namespace"
	return makeFieldRefEnvVar(key, path)
}

func makeFieldRefEnvVar(key string, path string) *corev1.EnvVar {
	env := &corev1.EnvVar{
		Name: key,
		ValueFrom: &corev1.EnvVarSource{
			FieldRef: &corev1.ObjectFieldSelector{
				FieldPath: path,
			},
		},
	}
	return env
}

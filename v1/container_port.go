package v1

import corev1 "k8s.io/api/core/v1"

func MakeContainerPorts(key string, val string) []corev1.ContainerPort {
	port := make([]corev1.ContainerPort, 0)
	return port
}

func MakeDefaultContainerPortVar(name string, port int32) *corev1.ContainerPort {
	protocol := corev1.ProtocolTCP
	containerPort := MakeContainerPort(name, port, protocol)
	return containerPort
}

func MakeUDPContainerPortVar(name string, port int32) *corev1.ContainerPort {
	protocol := corev1.ProtocolUDP
	containerPort := MakeContainerPort(name, port, protocol)
	return containerPort
}

func MakeContainerPort(name string, port int32, protocol corev1.Protocol) *corev1.ContainerPort {
	containerPort := &corev1.ContainerPort{
		Name:          name,
		ContainerPort: port,
		Protocol:      protocol,
	}
	return containerPort
}

func MakeContainerPortVar(name string, port int32, protocol corev1.Protocol) *corev1.ContainerPort {
	containerPort := &corev1.ContainerPort{
		Name:          name,
		ContainerPort: port,
		Protocol:      protocol,
	}
	return containerPort
}

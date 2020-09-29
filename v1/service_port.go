package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func MakeServicePorts() []corev1.ServicePort {
	env := make([]corev1.ServicePort, 0)
	return env
}

func MakeSimpleTCPServicePort(name string, port int32) corev1.ServicePort {
	p := corev1.ProtocolTCP
	sp := makeServicePort(name, port, port, p)
	return sp
}

func MakeSimpleUDPServicePort(name string, port int32) corev1.ServicePort {
	p := corev1.ProtocolUDP
	sp := makeServicePort(name, port, port, p)
	return sp
}

func MakeTCPServicePort(name string, port int32, targetPort int32) corev1.ServicePort {
	p := corev1.ProtocolTCP
	sp := makeServicePort(name, port, targetPort, p)
	return sp
}

func MakeUDPServicePort(name string, port int32, targetPort int32) corev1.ServicePort {
	p := corev1.ProtocolUDP
	sp := makeServicePort(name, port, targetPort, p)
	return sp
}

func makeServicePort(name string, port int32, targetPort int32, p corev1.Protocol) corev1.ServicePort {
	sp := corev1.ServicePort{
		Port:       port,
		Name:       name,
		Protocol:   p,
		TargetPort: intstr.IntOrString{IntVal: targetPort},
	}
	return sp
}

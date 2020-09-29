package v1

type StorageSpec struct {
	Name string `json:"className"`
	Size string `json:"size"`
}

var (
	ResourcePrefix         = "byteforce-"
	HeadlessServicePostfix = "-headless"
	ServicePostfix         = "-svc"
	ConfigMapPostfix       = "-cm"
	IngressPostfix         = "-ing"
)

type ResourceType string

const (
	ResourceTypeConfigMap       ResourceType = "ConfigMap"
	ResourceTypeHeadlessService ResourceType = "HeadlessService"
	ResourceTypeService         ResourceType = "Service"
	ResourceTypeIngress         ResourceType = "Ingress"
)
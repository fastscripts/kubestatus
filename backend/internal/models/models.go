package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Success   string
}

type Resources struct {
	Used     string `json:"used"`
	Capacity string `json:"capacity"`
}

type ClusterStatus struct {
	NodeCount int       `json:"node_count"`
	CPU       Resources `json:"cpu"`
	Memory    Resources `json:"memory"`
}

type Container struct {
	Name       string `json:"name,omitempty"`
	CPU        int    `json:"cpu,omitempty"`
	MEM        int    `json:"mem,omitempty"`
	MEMRequest int    `json:"memRequested,omitempty"`
	CPURequest int    `json:"cpuRequested,omitempty"`
	MEMLimit   int    `json:"memLimit,omitempty"`
	CPULimit   int    `json:"cpuLimit,omitempty"`
}

type Pod struct {
	Name       string      `json:"name"`
	Containers []Container `json:"children"`
}

type Namespace struct {
	Name string `json:"name"`
	Pods []Pod  `json:"children"`
}

type Data struct {
	Name       string      `json:"name"`
	Namespaces []Namespace `json:"children"`
}

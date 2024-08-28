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

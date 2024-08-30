package models

// PodList ist das ergebnis der Abfrage nach den Resourcen Limits
type PodList struct {
	Kind  string `json:"kind"`
	Items []struct {
		Metadata struct {
			Name         string `json:"name"`
			GenerateName string `json:"generateName"`
			Namespace    string `json:"namespace"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Name      string `json:"name"`
				Image     string `json:"image"`
				Resources struct {
					Limits struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
			} `json:"containers"`
		} `json:"spec"`
	} `json:"items"`
}

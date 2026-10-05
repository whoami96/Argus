package collector

type Alert struct {
	Name    string `json:"name"`
	State   string `json:"state"`
	Env     string `json:"env"`
	Source  string `json:"source"`
	Count   int    `json:"count"`
	Summary string `json:"summary"`
}

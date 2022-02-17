package internal

type Pokemon struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Stat  []Stat `json:"stat"`
	Url   string `json:"url"`
}

type Stat struct {
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

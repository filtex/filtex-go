package models

type Field struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Label     string   `json:"label"`
	Operators []string `json:"operators"`
	Values    []Lookup `json:"values"`
}

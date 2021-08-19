package model

type Network struct {
	ID   int64  `json:"id,omitemtpy"`
	Name string `json:"name,omitemtpy"`
	Ports []Port `json:"ports,omitemtpy"`
}

type Port struct {
	Type   string `json:"type,omitemtpy"`
	Target string `json:"target,omitemtpy"`
}


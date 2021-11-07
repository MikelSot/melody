package domain

type TypeMessage struct {
	ID   uint   `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type TypeMessages []TypeMessage

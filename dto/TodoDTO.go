package dto

import (
	"encoding/json"
	"io"
)

type TodoDTO struct {
	Status bool `json:"status,omitempty"`
}

func (td *TodoDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(td)
}

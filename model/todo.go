package model

import (
	"encoding/json"
	"io"
	"log"
)

type Todo struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Todos []Todo

func (t *Todos) ToJSON(w io.Writer) error {
	log.Printf("%#v", t)
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *Todo) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(t)
}

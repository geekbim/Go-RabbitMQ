package rabbitmq

import (
	"encoding/json"
)

type Todo struct {
	Id      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Created string `json:"created,omitempty"`
}

func (p *Todo) Serialize() ([]byte, error) {
	return json.Marshal(&Todo{
		Id:      p.Id,
		Name:    p.Name,
		Created: p.Created,
	})
}

func (u *Todo) UnSerialize(data []byte) error {
	err := json.Unmarshal(data, u)
	if err != nil {
		return err
	}
	return nil
}

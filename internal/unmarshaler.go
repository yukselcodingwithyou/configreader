package internal

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
)

type ConfigUnmarshaler interface {
	Unmarshal(content []byte, v interface{}) error
}

type jsonUnmarshaler struct {
}

func (j jsonUnmarshaler) Unmarshal(content []byte, v interface{}) error {
	err := json.Unmarshal(content, v)
	if err != nil {
		return errors.New("json unmarshal failed")
	}
	return nil
}

type yamlUnmarshaler struct {
}

func (y yamlUnmarshaler) Unmarshal(content []byte, v interface{}) error {
	err := yaml.Unmarshal(content, v)
	if err != nil {
		return errors.New("yaml unmarshal failed")
	}
	return nil
}

func NewUnmarshaler(contentType string) ConfigUnmarshaler {
	if contentType == "JSON" {
		return &jsonUnmarshaler{}
	} else {
		return &yamlUnmarshaler{}
	}
}

package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UnmarshalerTest struct {
	suite.Suite
	yamlUnmarshaler ConfigUnmarshaler
	jsonUnmarshaler ConfigUnmarshaler
}

func Test_RunUnmarshalerTestSuite(t *testing.T) {
	suite.Run(t, new(UnmarshalerTest))
}

func (h *UnmarshalerTest) SetupTest() {
	h.yamlUnmarshaler = &yamlUnmarshaler{}
	h.jsonUnmarshaler = &jsonUnmarshaler{}
}

func (h *UnmarshalerTest) TestUnmarshalJsonOK() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	content := []byte(`{"var1": true, "var2": "value"}`)

	err := h.jsonUnmarshaler.Unmarshal(content, &s)

	assert.Nil(t, err)
	assert.Equal(t, s.Var1, true)
	assert.Equal(t, s.Var2, "value")
}

func (h *UnmarshalerTest) TestUnmarshalJsonNotOK() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	expectedErr := errors.New("json unmarshal failed")
	content := []byte(``)

	err := h.jsonUnmarshaler.Unmarshal(content, &s)

	assert.Equal(t, expectedErr, err)
}

func (h *UnmarshalerTest) TestUnmarshalYamlOK() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	content := []byte(
		`
var1: true
var2: value
`)

	err := h.yamlUnmarshaler.Unmarshal(content, &s)

	assert.Nil(t, err)
	assert.Equal(t, s.Var1, true)
	assert.Equal(t, s.Var2, "value")
}

func (h *UnmarshalerTest) TestUnmarshalYamlNotOK() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	expectedErr := errors.New("yaml unmarshal failed")

	content := []byte(
		`
					var1: true
	var2: value
`)

	err := h.yamlUnmarshaler.Unmarshal(content, &s)

	assert.Equal(t, expectedErr, err)
}

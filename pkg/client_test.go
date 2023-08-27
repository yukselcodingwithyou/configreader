package pkg

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yukselcodingwithyou/configreader/internal/mocks"
	"testing"
)

type ConfigReaderTest struct {
	suite.Suite
	configInjectorFromLocal  *mocks.ConfigInjector
	configInjectorFromRemote *mocks.ConfigInjector
	jsonConfigUnmarshaler    *mocks.ConfigUnmarshaler
	yamlConfigUnmarshaler    *mocks.ConfigUnmarshaler

	localJsonConfigReader  ConfigReader
	remoteJsonConfigReader ConfigReader
	localYamlConfigReader  ConfigReader
	remoteYamlConfigReader ConfigReader
}

func Test_RunIConfigReaderTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigReaderTest))
}

func (h *ConfigReaderTest) SetupTest() {
	h.configInjectorFromLocal = new(mocks.ConfigInjector)
	h.configInjectorFromRemote = new(mocks.ConfigInjector)
	h.jsonConfigUnmarshaler = new(mocks.ConfigUnmarshaler)
	h.yamlConfigUnmarshaler = new(mocks.ConfigUnmarshaler)

	h.localJsonConfigReader = &configReader{
		injector:    h.configInjectorFromLocal,
		path:        "path",
		unmarshaler: h.jsonConfigUnmarshaler,
	}

	h.remoteJsonConfigReader = &configReader{
		injector:    h.configInjectorFromRemote,
		path:        "path",
		unmarshaler: h.jsonConfigUnmarshaler,
	}

	h.localYamlConfigReader = &configReader{
		injector:    h.configInjectorFromLocal,
		path:        "path",
		unmarshaler: h.yamlConfigUnmarshaler,
	}

	h.remoteYamlConfigReader = &configReader{
		injector:    h.configInjectorFromRemote,
		path:        "path",
		unmarshaler: h.yamlConfigUnmarshaler,
	}
}

func (h *ConfigReaderTest) TestReadJsonConfigFromLocalOK() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	content := []byte(`{"var1": true, "var2": "value"}`)

	h.configInjectorFromLocal.On("InjectFrom", "path").Return(content, nil)

	h.jsonConfigUnmarshaler.On("Unmarshal", content, s).Return(nil)

	err := h.localJsonConfigReader.Read(s)

	h.configInjectorFromLocal.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromLocal.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.jsonConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.jsonConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.Nil(t, err)
}

func (h *ConfigReaderTest) TestReadJsonConfigFromLocalInjectionError() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	expectedInjectionErr := ErrConfigContentRetrieveFailed
	h.configInjectorFromLocal.On("InjectFrom", "path").Return(nil, expectedInjectionErr)

	err := h.localJsonConfigReader.Read(s)

	h.configInjectorFromLocal.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromLocal.AssertNumberOfCalls(t, "InjectFrom", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedInjectionErr)
}

func (h *ConfigReaderTest) TestReadJsonConfigFromLocalUnmarshalError() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	content := []byte(`{"var1": true, "var2": "value"}`)

	expectedUnmarshalError := ErrConfigContentUnmarshalFailed

	h.configInjectorFromLocal.On("InjectFrom", "path").Return(content, nil)

	h.jsonConfigUnmarshaler.On("Unmarshal", content, s).Return(expectedUnmarshalError)

	err := h.localJsonConfigReader.Read(s)

	h.configInjectorFromLocal.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromLocal.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.jsonConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.jsonConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedUnmarshalError)
}

func (h *ConfigReaderTest) TestReadJsonConfigFromRemoteOK() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	content := []byte(`{"var1": true, "var2": "value"}`)

	h.configInjectorFromRemote.On("InjectFrom", "path").Return(content, nil)

	h.jsonConfigUnmarshaler.On("Unmarshal", content, s).Return(nil)

	err := h.remoteJsonConfigReader.Read(s)

	h.configInjectorFromRemote.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromRemote.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.jsonConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.jsonConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.Nil(t, err)
}

func (h *ConfigReaderTest) TestReadJsonConfigFromRemoteInjectionError() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	expectedInjectionErr := ErrConfigContentRetrieveFailed
	h.configInjectorFromRemote.On("InjectFrom", "path").Return(nil, expectedInjectionErr)

	err := h.remoteJsonConfigReader.Read(s)

	h.configInjectorFromRemote.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromRemote.AssertNumberOfCalls(t, "InjectFrom", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedInjectionErr)
}

func (h *ConfigReaderTest) TestReadJsonConfigFromRemoteUnmarshalError() {
	t := h.T()

	s := struct {
		Var1 bool   `json:"var1"`
		Var2 string `json:"var2"`
	}{}

	content := []byte(`{"var1": true, "var2": "value"}`)

	expectedUnmarshalError := ErrConfigContentUnmarshalFailed

	h.configInjectorFromRemote.On("InjectFrom", "path").Return(content, nil)

	h.jsonConfigUnmarshaler.On("Unmarshal", content, s).Return(expectedUnmarshalError)

	err := h.remoteJsonConfigReader.Read(s)

	h.configInjectorFromRemote.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromRemote.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.jsonConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.jsonConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedUnmarshalError)
}

func (h *ConfigReaderTest) TestReadYamlConfigFromLocalOK() {
	t := h.T()

	s := struct {
		Var1 bool   `yaml:"var1"`
		Var2 string `yaml:"var2"`
	}{}

	content := []byte(
		`
var1: true
var2: value
`)

	h.configInjectorFromLocal.On("InjectFrom", "path").Return(content, nil)

	h.yamlConfigUnmarshaler.On("Unmarshal", content, s).Return(nil)

	err := h.localYamlConfigReader.Read(s)

	h.configInjectorFromLocal.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromLocal.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.yamlConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.yamlConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.Nil(t, err)
}

func (h *ConfigReaderTest) TestReadYamlConfigFromLocalInjectionError() {
	t := h.T()

	s := struct {
		Var1 bool   `yaml:"var1"`
		Var2 string `yaml:"var2"`
	}{}

	expectedInjectionErr := ErrConfigContentRetrieveFailed
	h.configInjectorFromLocal.On("InjectFrom", "path").Return(nil, expectedInjectionErr)

	err := h.localYamlConfigReader.Read(s)

	h.configInjectorFromLocal.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromLocal.AssertNumberOfCalls(t, "InjectFrom", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedInjectionErr)
}

func (h *ConfigReaderTest) TestReadYamlConfigFromLocalUnmarshalError() {
	t := h.T()

	s := struct {
		Var1 bool   `yaml:"var1"`
		Var2 string `yaml:"var2"`
	}{}

	content := []byte(
		`
					var1: true
	var2: value
`)

	expectedUnmarshalError := ErrConfigContentUnmarshalFailed

	h.configInjectorFromLocal.On("InjectFrom", "path").Return(content, nil)

	h.yamlConfigUnmarshaler.On("Unmarshal", content, s).Return(expectedUnmarshalError)

	err := h.localYamlConfigReader.Read(s)

	h.configInjectorFromLocal.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromLocal.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.yamlConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.yamlConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedUnmarshalError)
}

func (h *ConfigReaderTest) TestReadYamlConfigFromRemoteOK() {
	t := h.T()

	s := struct {
		Var1 bool   `yaml:"var1"`
		Var2 string `yaml:"var2"`
	}{}

	content := []byte(
		`
var1: true
var2: value
`)

	h.configInjectorFromRemote.On("InjectFrom", "path").Return(content, nil)

	h.yamlConfigUnmarshaler.On("Unmarshal", content, s).Return(nil)

	err := h.remoteYamlConfigReader.Read(s)

	h.configInjectorFromRemote.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromRemote.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.yamlConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.yamlConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.Nil(t, err)
}

func (h *ConfigReaderTest) TestReadYamlConfigFromRemoteInjectionError() {
	t := h.T()

	s := struct {
		Var1 bool   `yaml:"var1"`
		Var2 string `yaml:"var2"`
	}{}

	expectedInjectionErr := ErrConfigContentRetrieveFailed
	h.configInjectorFromRemote.On("InjectFrom", "path").Return(nil, expectedInjectionErr)

	err := h.remoteYamlConfigReader.Read(s)

	h.configInjectorFromRemote.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromRemote.AssertNumberOfCalls(t, "InjectFrom", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedInjectionErr)
}

func (h *ConfigReaderTest) TestReadYamlConfigFromRemoteUnmarshalError() {
	t := h.T()

	s := struct {
		Var1 bool   `yaml:"var1"`
		Var2 string `yaml:"var2"`
	}{}

	content := []byte(
		`
			  var1: true
	var2: value
`)

	expectedUnmarshalError := ErrConfigContentUnmarshalFailed

	h.configInjectorFromRemote.On("InjectFrom", "path").Return(content, nil)

	h.yamlConfigUnmarshaler.On("Unmarshal", content, s).Return(expectedUnmarshalError)

	err := h.remoteYamlConfigReader.Read(s)

	h.configInjectorFromRemote.AssertCalled(t, "InjectFrom", "path")
	h.configInjectorFromRemote.AssertNumberOfCalls(t, "InjectFrom", 1)

	h.yamlConfigUnmarshaler.AssertCalled(t, "Unmarshal", content, s)
	h.yamlConfigUnmarshaler.AssertNumberOfCalls(t, "Unmarshal", 1)

	assert.NotNil(t, err)
	assert.Equal(t, err, expectedUnmarshalError)
}

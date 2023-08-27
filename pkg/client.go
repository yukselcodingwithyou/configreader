package pkg

import (
	"github.com/yukselcodingwithyou/configreader/internal"
	"os"
)

// ContentType indicates where the config will be from
// Supported values
//   - JSON
//   - YAML
type ContentType string

func (f ContentType) str() string {
	switch f {
	case JSON:
		return "JSON"
	case YAML:
		return "YAML"

	}
	return ""
}

const (
	JSON ContentType = "JSON"
	YAML ContentType = "YAML"
)

// From indicates where the config will be from
// Supported values
//   - LOCAL: local file path,
//   - REMOTE: a http server address
type From string

func (f From) str() string {
	switch f {
	case LOCAL:
		return "LOCAL"
	case REMOTE:
		return "REMOTE"

	}
	return ""
}

const (
	LOCAL  From = "LOCAL"
	REMOTE From = "REMOTE"
)

// ConfigReader reads the config into given interface
type ConfigReader interface {
	Read(v interface{}) error
}

type configReader struct {
	injector    internal.ConfigInjector
	path        string
	unmarshaler internal.ConfigUnmarshaler
}

// Read reads the config into given interface
func (c configReader) Read(v interface{}) error {
	content, err := c.injector.InjectFrom(c.path)
	if err != nil {
		return ErrConfigContentRetrieveFailed
	}
	err = c.unmarshaler.Unmarshal(content, v)
	if err != nil {
		return ErrConfigContentUnmarshalFailed
	}
	return nil
}

// NewConfigReader creates a new ConfigReader
// from: where the config will be retrieved
// contentType: content type of the config
// configVariableName: env variable of config (can be file path or http server address)
func NewConfigReader(from From, contentType ContentType, configVariableName string) ConfigReader {
	path := os.Getenv(configVariableName)
	injector := internal.NewConfigInjector(from.str())
	unmarshaler := internal.NewUnmarshaler(contentType.str())
	if from == LOCAL {
		return &configReader{injector: injector, path: path, unmarshaler: unmarshaler}
	} else {
		return &configReader{injector: injector, path: path, unmarshaler: unmarshaler}
	}
}

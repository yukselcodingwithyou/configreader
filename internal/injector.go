package internal

import (
	"errors"
	"os"
)

type ConfigInjector interface {
	InjectFrom(path string) ([]byte, error)
}

type fromLocalInjector struct {
}

func (f fromLocalInjector) InjectFrom(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("read data from local file failed")
	}
	return data, nil
}

type fromRemoteInjector struct {
	httpClient RemoteClient
}

func (f fromRemoteInjector) InjectFrom(path string) ([]byte, error) {
	return f.httpClient.Get(path)
}

func NewConfigInjector(from string) ConfigInjector {
	if from == "LOCAL" {
		return &fromLocalInjector{}
	} else {
		return &fromRemoteInjector{httpClient: newClient()}
	}
}

package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yukselcodingwithyou/configreader/internal/mocks"
	"testing"
)

type InjectorTest struct {
	suite.Suite
	client                   *mocks.RemoteClient
	configInjectorFromRemote ConfigInjector
	configInjectorFromLocal  ConfigInjector
}

func Test_RunInjectorTestSuite(t *testing.T) {
	suite.Run(t, new(InjectorTest))
}

func (h *InjectorTest) SetupTest() {
	h.client = new(mocks.RemoteClient)
	h.configInjectorFromRemote = &fromRemoteInjector{httpClient: h.client}
	h.configInjectorFromLocal = &fromLocalInjector{}
}

func (h *InjectorTest) TestInjectFromRemoteOK() {
	t := h.T()

	response := []byte(`{"var1": true, "var2": "value"}`)

	h.client.On("Get", "path").Return(response, nil)

	content, err := h.configInjectorFromRemote.InjectFrom("path")

	assert.Nil(t, err)
	assert.Equal(t, content, response)

	h.client.AssertCalled(t, "Get", "path")
	h.client.AssertNumberOfCalls(t, "Get", 1)
}

func (h *InjectorTest) TestInjectFromRemoteNotOK() {
	t := h.T()

	errMsg := "response acquire failed"
	expectedErr := errors.New(errMsg)

	h.client.On("Get", "path").Return(nil, expectedErr)

	content, err := h.configInjectorFromRemote.InjectFrom("path")

	assert.Nil(t, content)
	assert.Equal(t, expectedErr, err)

	h.client.AssertCalled(t, "Get", "path")
	h.client.AssertNumberOfCalls(t, "Get", 1)
}

func (h *InjectorTest) TestInjectFromLocalOK() {
	t := h.T()

	response := []byte(`{"var1": true, "var2": "value"}`)

	content, err := h.configInjectorFromLocal.InjectFrom("mocks/config.json")

	assert.Nil(t, err)
	assert.Equal(t, string(content), string(response))
}

func (h *InjectorTest) TestInjectFromLocalNotOK() {
	t := h.T()

	errMsg := "read data from local file failed"
	expectedErr := errors.New(errMsg)

	content, err := h.configInjectorFromLocal.InjectFrom("abc")

	assert.Nil(t, content)
	assert.Equal(t, err, expectedErr)
}

package pkg

import "errors"

var ErrConfigContentUnmarshalFailed = errors.New("config content unmarshal failed")
var ErrConfigContentRetrieveFailed = errors.New("config content retrieve failed")

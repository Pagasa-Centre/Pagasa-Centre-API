package request

import (
	"encoding/json"
	"errors"
	"io"
)

type Validator interface {
	Validate() error
}

func DecodeAndValidate(src io.ReadCloser, target Validator) error {
	if src == nil {
		return errors.New("empty body")
	}

	defer func(src io.ReadCloser) {
		_ = src.Close()
	}(src)

	if err := json.NewDecoder(src).Decode(target); err != nil {
		return err
	}

	return target.Validate()
}

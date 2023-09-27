package whisper

import (
	"encoding/json"
	"errors"
	"strings"
)

var ErrAudioFormat = errors.New("the audio file could not be decoded or its format is not supported")

type ErrorResult map[string]any

func (e *ErrorResult) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func errorHandler(err *ErrorResult) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "The audio file could not be decoded or its format is not supported") {
		return ErrAudioFormat
	}
	return err
}

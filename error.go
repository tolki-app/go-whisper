package whisper

import (
	"errors"
	"strings"
)

var ErrAudioFormat = errors.New("the audio file could not be decoded or its format is not supported")

type Error struct {
	Message string
	Type    string
	Param   *string
	Code    *int
}

func (e *Error) Error() string {
	return e.Message
}

func errorHandler(err *Error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Message, "The audio file could not be decoded or its format is not supported") {
		return ErrAudioFormat
	}
	return err
}

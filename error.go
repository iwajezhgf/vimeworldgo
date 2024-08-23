package vimeworldgo

import "fmt"

// https://vimeworld.github.io/api-docs/#commonerrors
type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code=%d, msg=%s", e.ErrorCode, e.ErrorMsg)
}

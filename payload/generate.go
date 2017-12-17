package payload

import (
	"errors"
	"log"
	"net/http"
)

type Generate struct {
	Filename string                 `json:"filename"`
	Data     map[string]interface{} `json:"data"`
}

func (g *Generate) Validate(r *http.Request) error {
	log.Println(g.Data, g.Filename)
	if g.Filename == "" {
		return errors.New("filename is required")
	}
	if len(g.Data) == 0 {
		return errors.New("data is required")
	}
	return nil
}

type SuccessResponse struct {
	Data string `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

package api

import (
	"io"
	"net/http"
)

func (a *Api) Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

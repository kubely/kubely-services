package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kubely/kubely-services/kedge"
)

type KedgeController struct {
}

type CreatePayload struct {
	Kedge []*kedge.Meta `json:"kedge"`
}

func GenerateDeploy(w http.ResponseWriter, r *http.Request) {
	var payload CreatePayload
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	// generate a file for each Meta
	for _, km := range payload.Kedge {
		err := km.CreateFile()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
	fileToDownload, err := kedge.Generate(payload.Kedge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(fileToDownload)
	err = kedge.Deploy(payload.Kedge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := ioutil.ReadFile(fileToDownload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// w.Header().Set("Content-Type", "text/oc-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// buf := bytes.NewBuffer(data)
	// dec := base64.NewDecoder(base64.StdEncoding, buf)
	fmt.Println("------------------", fileToDownload)
	// w.WriteHeader(http.StatusCreated)
	// w.Header().Set("Content-Type", "application/yml")
	// w.Header().Set("Content-Disposition", "Attachment")
	w.Header().Set("Content-Disposition", "attachment; filename=yes.png")
	w.Header().Set("Content-Type", "image/png")
	// w.Write(data)
	io.Copy(w, bytes.NewReader(data))
	// w.Header().Set("Content-Disposition", "Attachment")
	// http.ServeContent(w, req, name, modtime, content)
	// http.ServeContent(w, r, "ok", time.Now(), bytes.NewReader(data))
	// http.ServeFile(w, r, fileToDownload)
	return
}

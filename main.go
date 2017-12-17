package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kedge-trial/blockly-kedge/server/controllers"
)

const PORT = 9999

type PayloadValidation interface {
	Validate(r *http.Request) error
}

func decodeAndValidate(r *http.Request, v PayloadValidation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		log.Println("error in decode", err)
		return err
	}
	defer r.Body.Close()
	return v.Validate(r)
}

func prebuildChecks() bool {
	return verifyMinishift() && verifyKedge() && verifyKubectl() && verifyOC()
}

func verifyMinishift() bool {
	cmd := exec.Command("minishift", "status", "|", "grep", "Minishift:")
	op, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("error: verifyMinishift: ", err)
	}
	if strings.Contains(string(op), "Running") {
		cmd2 := exec.Command("minishift", "status", "|", "grep", "OpenShift:")
		op2, err := cmd2.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(op2), "Running") {
			return true
		}
		log.Fatal("Minishift is not running")
		return false
	}
	log.Fatal("OpenShift is not running")
	return false
}

func verifyKedge() bool {
	cmd := exec.Command("kedge", "--help")
	op, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("error: verifyKedge: ", err, strings.Join(cmd.Args, " "))
	}
	if strings.Contains(string(op), "Simple, Concise & Declarative Kubernetes Applications") {
		return true
	}
	log.Fatal("kedge is not available")
	return false
}

func verifyKubectl() bool {
	cmd := exec.Command("kubectl", "--help")
	op, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("error: verifyKubectl: ", err)
	}
	if strings.Contains(string(op), "controls the Kubernetes cluster manager") {
		return true
	}
	log.Fatal("kubectl is not available")
	return false
}

func verifyOC() bool {
	cmd := exec.Command("oc", "--help")
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("error: verifyOC: ", err)
	}
	return true
}

// todos
// enhance http response structure
func main() {
	prebuildChecks()
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("POST").Path("/generate").HandlerFunc(controllers.GenerateDeploy)
	log.Println("starting server on ", PORT)
	addr := fmt.Sprintf(":%d", PORT)
	log.Fatal(http.ListenAndServe(addr, router))
}

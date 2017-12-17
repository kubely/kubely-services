package kedge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type Meta struct {
	Filename string                 `json:"filename"`
	Data     map[string]interface{} `json:"data"`
	Filepath string
}

const (
	BASE_PATH  = "../"
	UPLOAD_DIR = "uploads/"
)

func (m *Meta) CreateFile() error {
	databytes, err := json.Marshal(m.Data)
	if err != nil {
		return errors.Wrap(err, "failed to marshel data")
	}
	m.Filepath = UPLOAD_DIR + m.Filename
	err = ioutil.WriteFile(m.Filepath, databytes, 0644)
	if err != nil {
		return errors.Wrap(err, "failed to write file")
	}
	return nil
}

func Generate(instances []*Meta) (string, error) {
	allArtifactsPath := UPLOAD_DIR + "kubernetes.yml"
	var query string
	for _, m := range instances {
		query += fmt.Sprintf(" -f %s ", m.Filepath)
	}
	cmdStr := fmt.Sprintf("kedge generate -v %s > %s", query, allArtifactsPath)

	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("kedge generate failed: %s\n", strings.Join(cmd.Args, " ")))
	}
	return allArtifactsPath, nil
}

func Deploy(instances []*Meta) error {
	var query string
	for _, m := range instances {
		query += fmt.Sprintf(" -f %s ", m.Filepath)
	}
	cmdStr := fmt.Sprintf("kedge apply %s", query)
	// cmdStr := fmt.Sprintf("kedge apply -f %s", allArtifactsPath)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("kedge apply failed: %s\n", strings.Join(cmd.Args, " ")))
	}
	return nil
}

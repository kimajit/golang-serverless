package parameter

import (
	"encoding/json"
	"log"
	"strings"
)

type Payload struct {
	Branch     string     `json:"branch"`
	Parameters Parameters `json:"parameters"`
}
type Parameters struct {
	ImageName   string `json:"imageName"`
	ImageTag    string `json:"imageTag"`
	ServiceName string `json:"serviceName"`
	ServicePath string `json:"servicePath"`
}

var (
// marshaler models.JSONMarshaler = new(json.Encoder)
// reader    models.Reader        = new(strings.Reader)
)

func GetParams(sec map[string]string, imageTag string, reponame string) (*strings.Reader, error) {
	rawPayload := Payload{Branch: sec["branchName"], Parameters: Parameters{ImageName: reponame, ImageTag: imageTag, ServiceName: sec["serviceName"], ServicePath: sec["servicePath"]}}
	jsonPayLoad, jsonErr := json.Marshal(rawPayload)
	if jsonErr != nil {
		return nil, jsonErr
	}
	payload := string(jsonPayLoad)
	log.Println("request payload:", payload)
	reader := strings.NewReader(payload)
	return reader, nil
}

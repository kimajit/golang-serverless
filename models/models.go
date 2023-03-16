package models

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type JSONMarshaler interface {
	MarshalJSON(any) ([]byte, error)
}

type JSONUnMarshaler interface {
	UnmarshalJSON([]byte, any) error
}

type Reader interface {
	NewReader(string) (reader *strings.Reader)
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type SecretsManagerClient interface {
	New(p client.ConfigProvider, cfgs ...*aws.Config) *secretsmanager.SecretsManager
	GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
}

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

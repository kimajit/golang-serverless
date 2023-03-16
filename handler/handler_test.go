package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type mockSecretsManagerClient struct{}

func (m *mockSecretsManagerClient) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	if *input.SecretId != secretsId {
		return nil, errors.New("secret not found")
	}
	// return some mock secret value
	return &secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(`{
            "baseUrl": "https://example.com/",
            "vcs": "github",
            "org": "my-org",
            "pipelineName": "my-pipeline",
            "circleCIToken": "my-token",
            "serviceName": "my-service",
            "branchName": "my-branch"
        }`),
	}, nil
}

type mockJSONMarshaler struct{}

func (m *mockJSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	return []byte(`{"mock": true}`), nil
}

func (m *mockJSONMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`{Branch: mybranch, Parameters: Parameters{ImageName: myimagename, ImageTag: myimageTag, ServiceName: myservicename, ServicePath: myservicepath}}`), nil
}

type mockStringReader struct{}

func (m *mockStringReader) ReadString(delim byte) (string, error) {
	return "mock string", nil
}

func (m *mockStringReader) NewReader() *strings.Reader {
	return strings.NewReader(string([]byte(`{Branch: mybranch, Parameters: Parameters{ImageName: myimagename, ImageTag: myimageTag, ServiceName: myservicename, ServicePath: myservicepath}}`)))
}

type mockClient struct{}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString("mock response body")),
	}, nil
}

// func TestHandler(t *testing.T) {
// 	t.Run("testing for handler function", func(t *testing.T) {
// 		secretsManagerClient = &mockSecretsManagerClient{}
// 		jsonMarshaler = &mockJSONMarshaler{}
// 		stringReader = &mockStringReader{}
// 		client = &mockClient{}
// 		event := events.CloudWatchEvent{
// 			Detail: json.RawMessage(`{
// 		"repository-name": "my-repo",
// 		"image-tags": ["my-tag"]
// 		}`),
// 		}
// 		err := Handler(context.Background(), event)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("testing for handler function", func(t *testing.T) {
// 		secretsManagerClient = &mockSecretsManagerClient{}
// 		jsonMarshaler = &mockJSONMarshaler{}
// 		stringReader = &mockStringReader{}
// 		client = &mockClient{}
// 		event := events.CloudWatchEvent{}
// 		err := Handler(context.Background(), event)
// 		assert.NotNil(t, err)
// 	})
// }

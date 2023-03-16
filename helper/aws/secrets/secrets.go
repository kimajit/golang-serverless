package secrets

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var (
// unmarshaler models.JSONUnMarshaler
// smclient models.SecretsManagerClient = new(secretsmanager.SecretsManager)
)

func GetAWSsecrets(secretId string) (map[string]string, error) {
	var secrets map[string]string
	session := createAwsSession()
	client := secretsmanager.New(session)
	response, secretsManagerErr := client.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	})
	if secretsManagerErr != nil {
		return nil, secretsManagerErr
	}
	jsonErr := json.Unmarshal([]byte(*response.SecretString), &secrets)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return secrets, nil
}

func createAwsSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

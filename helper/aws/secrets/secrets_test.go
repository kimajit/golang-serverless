package secrets

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/stretchr/testify/assert"
)

type mockSecretsManager struct {
	GetSecretValueFunc func(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
}

func (m *mockSecretsManager) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	return m.GetSecretValueFunc(input)
}

func TestGetAWSsecrets(t *testing.T) {
	t.Run("test for successfull secret retrival", func(t *testing.T) {
		expectedSecrets := map[string]string{
			"foo": "bar",
		}
		// mock := &mockSecretsManager{
		// 	GetSecretValueFunc: func(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
		// 		output := &secretsmanager.GetSecretValueOutput{
		// 			SecretString: aws.String(`{"foo": "bar"}`),
		// 		}
		// 		return output, nil
		// 	},
		// }
		secrets, err := GetAWSsecrets("test")
		assert.Nil(t, err)
		assert.NotNil(t, secrets)
		assert.Equal(t, expectedSecrets, secrets)
	})
	t.Run("test for failed secrets retrival", func(t *testing.T) {
		// mock := &mockSecretsManager{
		// 	GetSecretValueFunc: func(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
		// 		output := &secretsmanager.GetSecretValueOutput{
		// 			SecretString: aws.String(""),
		// 		}
		// 		return output, nil
		// 	},
		// }
		secrets, err := GetAWSsecrets("")
		assert.NotNil(t, err)
		assert.Nil(t, secrets)
	})
}

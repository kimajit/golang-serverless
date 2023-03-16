package parameter

import (
	"strings"

	"github.com/stretchr/testify/mock"
)

type MockMarshaler struct {
	mock.Mock
}

func (m *MockMarshaler) MarshalJSON(input interface{}) ([]byte, error) {
	args := m.Called(input)
	return args.Get(0).([]byte), args.Error(1)
}

type MockReader struct {
	mock.Mock
}

func (m *MockReader) NewReader() *strings.Reader {
	return &strings.Reader{}
}

// func TestGetParams(t *testing.T) {
// 	t.Run("Test for getting params", func(t *testing.T) {
// 		mockMarshaler := &MockMarshaler{}
// 		mockReader := &MockReader{}
// 		Secrets := map[string]string{
// 			"branchName":  "test",
// 			"serviceName": "test",
// 			"servicePath": "test",
// 		}
// 		reader, err := GetParams(Secrets, "imageTag", "reponame", mockMarshaler, mockReader)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, reader)
// 	})

// 	t.Run("test for failed get params function", func(t *testing.T) {
// 		mockMarshaler := &MockMarshaler{}
// 		mockReader := &MockReader{}
// 		secrets := map[string]string{}
// 		reader, err := GetParams(secrets, "", "", mockMarshaler, mockReader)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, reader)
// 		assert.Equal(t, 0, reader.Len())
// 	})
// }

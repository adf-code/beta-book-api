package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
)

type MinioClient struct {
	mock.Mock
}

func (m *MinioClient) UploadFile(ctx context.Context, file multipart.File, objectName string, size int64, contentType string) (string, error) {
	args := m.Called(ctx, file, objectName, size, contentType)
	return args.String(0), args.Error(1)
}

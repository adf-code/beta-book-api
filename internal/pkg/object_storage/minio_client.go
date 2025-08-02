package object_storage

import (
	"context"
	"fmt"
	"github.com/adf-code/beta-book-api/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog"
	"mime/multipart"
	"os"
)

type ObjectStorageClient interface {
	UploadFile(ctx context.Context, file multipart.File, objectName string, size int64, contentType string) (string, error)
}

type MinioClient struct {
	endpoint   string
	accessKey  string
	secretKey  string
	bucketName string
	logger     zerolog.Logger
	Client     *minio.Client
}

func NewMinioClient(cfg *config.AppConfig, logger zerolog.Logger) *MinioClient {
	return &MinioClient{
		endpoint:   cfg.MiniEndpoint,
		accessKey:  cfg.MinioAccessKey,
		secretKey:  cfg.MinioSecretKey,
		bucketName: cfg.MinioBucketName,
		logger:     logger,
	}
}

func (m *MinioClient) InitMinio() *MinioClient {
	client, err := minio.New(m.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.accessKey, m.secretKey, ""),
		Secure: false, // change to true if using HTTPS
	})
	if err != nil {
		m.logger.Fatal().Err(err).Msgf("❌ Failed to open connection to Minio: %v", err)
		return nil
	}

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, m.bucketName)
	if err != nil {
		m.logger.Fatal().Err(err).Msgf("❌ Failed to open connection to Minio and check bucket: %v", err)
		return nil
	}
	if !exists {
		if err := client.MakeBucket(ctx, m.bucketName, minio.MakeBucketOptions{}); err != nil {
			m.logger.Fatal().Err(err).Msgf("❌ Minio bucket not found: %v", err)
			return nil
		}
	}
	//minioClient := MinioClient{Client: client}

	return &MinioClient{Client: client, bucketName: m.bucketName}
}

func (m *MinioClient) UploadFile(ctx context.Context, file multipart.File, fileName string, fileSize int64, contentType string) (string, error) {
	uploadInfo, err := m.Client.PutObject(ctx, m.bucketName, fileName, file, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		m.logger.Error().Err(err).Msgf("❌ Failed uploading file to Minio: %v", err)
		return "", err
	}
	m.logger.Info().Msgf("✅ Uploaded %s to MinIO (%d bytes)", fileName, uploadInfo.Size)

	fileURL := fmt.Sprintf("%s/%s/%s", os.Getenv("MINIO_ENDPOINT"), m.bucketName, fileName)
	return fileURL, nil
}

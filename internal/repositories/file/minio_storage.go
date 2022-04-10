package file

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/minio/minio-go/v6"

	"golang_pet_project_1/internal/config"
	"golang_pet_project_1/internal/core/domain"
)

type MinioStorage struct {
	minioClient *minio.Client
	userBucket  string
}

func NewMinioStorage(c *config.Config) (*MinioStorage, error) {
	minioClient, err := minio.New(
		c.MinioHost+":"+c.MinioPort,
		c.MinioRootUser,
		c.MinioRootPassword,
		false,
	)
	if err != nil {
		return &MinioStorage{}, err
	}

	return &MinioStorage{
		minioClient: minioClient,
		userBucket:  c.UserBucketName,
	}, nil
}

func (m *MinioStorage) Save(user *domain.User) (*domain.User, error) {
	b, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		return &domain.User{}, err
	}

	reader := bytes.NewReader(b)
	if _, err = m.minioClient.PutObject(
		m.userBucket,
		strconv.Itoa(user.ID),
		reader,
		reader.Size(),
		minio.PutObjectOptions{},
	); err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (m *MinioStorage) GetByID(id string) (*domain.User, error) {
	reader, err := m.minioClient.GetObject(
		m.userBucket,
		id,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return &domain.User{}, err
	}
	defer reader.Close()

	b := new(bytes.Buffer)
	if _, err := b.ReadFrom(reader); err != nil {
		return &domain.User{}, err
	}

	user := &domain.User{}
	if err = json.Unmarshal(b.Bytes(), user); err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (m *MinioStorage) Update(user *domain.User) (*domain.User, error) {
	err := m.Delete(strconv.Itoa(user.ID))
	if err != nil {
		return &domain.User{}, err
	}

	savedUser, err := m.Save(user)
	if err != nil {
		return &domain.User{}, err
	}

	return savedUser, nil
}

func (m *MinioStorage) Delete(userID string) error {
	err := m.minioClient.RemoveObject(
		m.userBucket,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

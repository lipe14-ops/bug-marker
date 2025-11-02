package repositories

import (
	"bytes"
	"context"
	"io"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/minio/minio-go/v7"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"

	"mime/multipart"
)
 
type ImageRepository interface {
	UploadImage(image *multipart.FileHeader) (*entities.Image, error)
  GetPreSignedImageUrl(image *presenters.Image) (*presenters.Image, error)
}

type imageRepository struct {
	collection *minio.Client
}

func NewImageRepository(collection *minio.Client) ImageRepository {
	return &imageRepository {
		collection: collection,
	}
}

func (r *imageRepository) UploadImage(image *multipart.FileHeader) (*entities.Image, error){
	fileContent, err := image.Open()

	if err != nil {
		return nil, err
	}
	
	defer fileContent.Close()

	buffer := bytes.NewBuffer(nil)

	if _, err := io.Copy(buffer, fileContent); err != nil {
		return nil, err
	}

	fileBytes := buffer.Bytes()
	fileReader := bytes.NewReader(fileBytes)

	responseImage := entities.Image {
		ID:  primitive.NewObjectID(),
		Filename:  image.Filename,
		CreatedAt: time.Now(),
	}

	filename := responseImage.ID.Hex() + path.Ext(responseImage.Filename)

	_, err = r.collection.PutObject(context.Background(), "bugmarker", filename, fileReader, int64(len(fileBytes)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	if err != nil {
		return nil, err
	}

	return 	&responseImage, nil
}

func (r *imageRepository) GetPreSignedImageUrl(image *presenters.Image) (*presenters.Image, error) {
  expireTime := time.Second * 24 * 60 * 60
	filename := image.ID + path.Ext(image.Filename)
  presingedURL, err := r.collection.PresignedGetObject(context.Background(), "bugmarker", filename, expireTime, nil)

	if err != nil {
		return nil, err
	}

  image.Url = presingedURL.String()

  return image, nil
}

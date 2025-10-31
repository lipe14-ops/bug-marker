package repositories

import (
	"context"
	"time"
	"io"
	"bytes"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/minio/minio-go/v7"

	"backend/internal/pkg/entities"

	"mime/multipart"
)
 
type ImageRepository interface {
	UploadImage(image *multipart.FileHeader) (*entities.Image, error)
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

	filename := responseImage.ID.Hex()
	responseImage.Url = "/image/" + filename

	_, err = r.collection.PutObject(context.Background(), "bugmarker", filename, fileReader, int64(len(fileBytes)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	if err != nil {
		return nil, err
	}

	return 	&responseImage, nil
}

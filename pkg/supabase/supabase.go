package supabase

import (
	"mime/multipart"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type Interface interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(link []string) error
}

type supabaseStorage struct {
	client *storage_go.Client
}

// Delete implements Interface3.
func (s *supabaseStorage) Delete(link []string) error {
	_, err := s.client.RemoveFile("user-profile", link)
	if err != nil {
		return err
	}

	return nil
}


func (s *supabaseStorage) Upload(file *multipart.FileHeader) (string, error) {
	fileBody, err := file.Open() 
	if err != nil {
		return "", err
	}

	bucket := os.Getenv("SUPABASE_BUCKET")

	fileName := file.Filename
	contentType := file.Header.Get("Content-Type")
	_, err = s.client.UploadFile(bucket, fileName, fileBody, storage_go.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}

	url := s.client.GetPublicUrl(bucket, fileName).SignedURL

	return url, nil
}

func Init() Interface {
	storageClient := storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_TOKEN"), nil)
	return &supabaseStorage{
		client: storageClient,
	}
}

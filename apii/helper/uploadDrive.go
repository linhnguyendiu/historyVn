package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Use Service account
func ServiceAccount(secretFile string) *http.Client {
	b, err := os.ReadFile(secretFile)
	if err != nil {
		log.Fatal("error while reading the credential file", err)
	}
	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

// func createFolder(service *drive.Service, name string, parentId string) (*drive.File, error) {
// 	d := &drive.File{
// 		Name:     name,
// 		MimeType: "application/vnd.google-apps.folder",
// 		Parents:  []string{parentId},
// 	}

// 	file, err := service.Files.Create(d).Do()

// 	if err != nil {
// 		log.Println("Could not create dir: " + err.Error())
// 		return nil, err
// 	}

// 	return file, nil
// }

func CreateFile(name string, size int64, content []byte) (string, error) {

	client := ServiceAccount("helper/client_secret.json")

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	folderId := "1JJw_fX_6lrqY1pDHoxwnT1WvNGBCwmc5"

	f := &drive.File{
		// MimeType: mimeType,
		Name:    name,
		Parents: []string{folderId},
	}
	file, err := srv.Files.Create(f).Media(bytes.NewReader(content)).Do()

	if err != nil {
		log.Fatalf("not create %v", err)
	}

	log.Printf("File '%s' successfully uploaded", file.Name)
	log.Printf("\nFile Id: '%s' ", file.Id)

	return file.Id, nil
}

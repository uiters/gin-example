package firebase

import (
	storage2 "cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

type Message struct {
	Username string
	Message  string
	CreateAt string
}

func initContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx
}

func InitFirebaseClient() (*db.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(err)
	}

	database := os.Getenv("FIREBASE_DATABASE")

	ctx := initContext()
	config := &firebase.Config{
		DatabaseURL: database,
	}
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	client, err := app.Database(ctx)
	return client, nil
}

func InitFirebaseStorage() (*storage2.Client, error) {
	ctx := initContext()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	client, err := storage2.NewClient(ctx, opt)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return client, nil
}

func PushNotification(username string) error {
	client, err := InitFirebaseClient()
	if err != nil {
		logrus.Print(err)
	}

	newMessage := Message{
		Username: username,
		Message:  "hello",
		CreateAt: time.Now().Format(time.RFC3339),
	}
	_, err = client.NewRef("/"+username).Push(context.Background(), newMessage)
	return err
}

func UploadFile(f multipart.FileHeader) (string, int, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(err)
	}

	bucketName := os.Getenv("FIREBASE_STORAGE")
	ctx := context.Background()
	client, err := InitFirebaseStorage()
	if err != nil {
		logrus.Print(err)
	}

	file, err := f.Open()
	if err != nil {
		logrus.Print(err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	var errs error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wc := client.Bucket(bucketName).Object(f.Filename).NewWriter(ctx)
		if _, err = io.Copy(wc, file); err != nil {
			errs = err
		}
		if err := wc.Close(); err != nil {
			errs = err
		}
		wg.Done()
	}()
	wg.Wait()
	if errs != nil {
		return "", getHTTPCode(errs), errs
	}
	return "gs://" + bucketName + "/" + f.Filename, http.StatusOK, nil
}

/*
func DownloadFile(filename string) ([]byte,error){
	ctx := context.Background()
	client,err:=InitFirebaseStorage()
	if err!=nil{
		logrus.Print(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	rc,err := client.Bucket("uit-lib.appspot.com").Object(filename).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	permissions := os.FileMode(0644).Perm()
	err = ioutil.WriteFile("file.pdf", data, permissions)
	if err != nil {
		// handle error
	}

	return data, nil
}*/

func getHTTPCode(err error) int {
	if err == gorm.ErrRecordNotFound {
		return http.StatusNotFound
	} else if err == gorm.ErrUnaddressable || err == gorm.ErrCantStartTransaction || err == gorm.ErrInvalidSQL || err == gorm.ErrInvalidTransaction {
		return http.StatusInternalServerError
	}
	return http.StatusBadRequest
}
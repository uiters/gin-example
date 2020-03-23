package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"log"
	"time"
)

type Message struct {
	Username string
	Message string
	CreateAt string
}

func initContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx
}


func InitFirebaseClient() (*db.Client,error)  {
	ctx :=initContext()
	config := &firebase.Config{
		DatabaseURL: "https://uit-lib.firebaseio.com/",
	}
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	client, err := app.Database(ctx)
	return client,nil
}

func PushNotification(username string) error  {
	client,err:=InitFirebaseClient()
	if err!=nil{
		logrus.Print(err)
	}

	newMessage :=Message{
		Username: username,
		Message:  "hello",
		CreateAt: time.Now().Format(time.RFC3339),
	}
	_,err = client.NewRef("/"+username).Push(context.Background(),newMessage)
	return err
}
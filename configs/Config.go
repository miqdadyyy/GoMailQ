package configs

import (
	"context"
	"crypto/tls"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"time"
)

func GetDialer() *gomail.Dialer {
	var (
		HOST     = os.Getenv("SMTP_HOST")
		PORT, _  = strconv.Atoi(os.Getenv("SMTP_PORT"))
		USERNAME = os.Getenv("SMTP_USERNAME")
		PASSWORD = os.Getenv("SMTP_PASSWORD")
	)
	dialer := gomail.NewDialer(HOST, PORT, USERNAME, PASSWORD)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer
}

func GetClientCollection() (*mongo.Client, *mongo.Collection, error){
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")));
	if err != nil {
		return nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 60 * time.Second)
	err = client.Connect(ctx);
	if err != nil {
		return nil, nil, err
	}
	err = client.Ping(ctx, readpref.Primary());
	if err != nil {
		return nil, nil, err
	}
	coll := client.Database("tasks").Collection("task")
	return client, coll, err
}

package db

import (
	"Bus/pkg/setting"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Session struct {
	DB      *mongo.Database
	Client  *mongo.Client
	Context context.Context
}

func Connect() (*Session, error) {
	var (
		err      error
		dbname   string
		port     string
		user     string
		password string
		host     string
		hasauth  string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		return nil, err
	}

	dbname = sec.Key("DBNAME").String()
	port = sec.Key("PORT").MustString("27017")
	user = sec.Key("USER").MustString("admin")
	password = sec.Key("PASSWORD").MustString("123456")//c431ca0640930cc4b8b6737787173785
	host = sec.Key("HOST").MustString("mongo")

	hasauth = sec.Key("HASAUTH").MustString("false")
	url := fmt.Sprintf("mongodb://%s:%s", host, port)
	if hasauth == "true" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			user, password, host, port)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		return nil, err
	}
	db := client.Database(dbname)
	return &Session{
		DB:      db,
		Client:  client,
		Context: ctx,
	}, nil
}

func (s *Session) Close() {
	_ = s.Client.Disconnect(s.Context)
}

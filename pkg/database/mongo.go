package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type IMongoDB interface {
	ConnectDatabase() error
	PingDatabase() error
	DisconnectDatabase()
	Database() *mongo.Database
}

type mongoDB struct {
	client *mongo.Client
	db     *mongo.Database
	url    string
	dbName string
}

func (m *mongoDB) ConnectDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.url))
	if err != nil {
		return err
	}
	m.client = client
	m.db = m.client.Database(m.dbName)
	return nil
}

func (m *mongoDB) PingDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return m.client.Ping(ctx, readpref.Primary())
}

func (m *mongoDB) DisconnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := m.client.Disconnect(ctx); err != nil {
		log.Panic(err)
	}
}

func (m *mongoDB) Database() *mongo.Database {
	return m.db
}

func NewMongoDB(url string, dbName string) IMongoDB {
	return &mongoDB{url: url, dbName: dbName}
}

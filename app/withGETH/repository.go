package withGETH

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const collectionName = "tx_monitor"

type IRepository interface {
	InsertTx(transaction map[string]interface{}) error
}

type repository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func (r *repository) InsertTx(transaction map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, transaction)
	return err
}

func NewRepository(db *mongo.Database) IRepository {
	return &repository{db: db, collection: db.Collection(collectionName)}
}

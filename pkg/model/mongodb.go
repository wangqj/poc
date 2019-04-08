package model

import (
	"context"
	"fmt"
	"github.com/go-mesh/openlogging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

const (
	DB             = "rokie"
	CollectionKV   = "kv"
	DefaultTimeout = 5 * time.Second
)

type MongodbService struct {
	c *mongo.Client
}

func (s *MongodbService) CreateOrUpdate(kv *KV) (string, error) {
	if kv.Domain == "" {
		return "", ErrMissingDomain
	}
	ctx, _ := context.WithTimeout(context.Background(), DefaultTimeout)
	collection := s.c.Database(DB).Collection(CollectionKV)
	oid, err := s.Exist(kv.Key, kv.Domain, kv.Labels)
	if err != nil {
		if err != ErrNotExists {
			return "", err
		}
	}
	if oid != "" {
		hex, err := primitive.ObjectIDFromHex(oid)
		if err != nil {
			openlogging.Error(fmt.Sprintf("convert %s ,err:%s", oid, err))
			return "", err
		}
		log.Println(hex)
		ur, err := collection.UpdateOne(ctx, bson.M{"_id": hex}, bson.D{
			{"$set", bson.D{
				{"value", kv.Value},
			}},
		})
		if err != nil {
			return "", err
		}
		openlogging.Debug(fmt.Sprintf("update %s with labels %s value [%s] %d ", kv.Key, kv.Labels, kv.Value, ur.ModifiedCount))
		return oid, nil
	}

	res, err := collection.InsertOne(ctx, kv)
	if err != nil {
		return "", err
	}
	objectID, _ := res.InsertedID.(primitive.ObjectID)
	return objectID.Hex(), nil
}
func (s *MongodbService) Exist(key, domain string, labels Labels) (string, error) {
	if key == "" {
		return "", nil
	}
	if domain == "" {
		return "", ErrMissingDomain
	}
	collection := s.c.Database(DB).Collection(CollectionKV)
	ctx, _ := context.WithTimeout(context.Background(), DefaultTimeout)
	filter := bson.M{"key": key, "domain": domain}
	for k, v := range labels {
		filter["labels."+k] = v
	}
	openlogging.Debug(fmt.Sprintf("find [%s] with lables [%s] in [%s]", key, labels, domain))
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return "", err
	}
	curKV := &KV{} //reuse this pointer to reduce GC, only clear label
	//check label length to get the exact match
	for cur.Next(ctx) { //although complexity is O(n), but there won't be so much labels for one key
		curKV.Labels = nil
		err := cur.Decode(curKV)
		if err != nil {
			openlogging.Error("decode error: " + err.Error())
			return "", err
		}
		openlogging.Debug(fmt.Sprintf("current: %s", curKV))
		if len(curKV.Labels) == len(labels) {
			openlogging.Debug("hit")
			return curKV.ID.Hex(), nil
		}

	}

	return "", ErrNotExists

}
func (s *MongodbService) DeleteByID(id string) error {
	collection := s.c.Database(DB).Collection(CollectionKV)
	ctx, _ := context.WithTimeout(context.Background(), DefaultTimeout)
	collection.DeleteOne(ctx, bson.M{"_id": id})
	return nil
}

func (s *MongodbService) Delete(key, domain string, labels Labels) error {
	return nil
}
func NewMongoService(opts Options) (KVService, error) {
	c, err := getClient(opts)
	if err != nil {
		return nil, err
	}
	m := &MongodbService{
		c: c,
	}
	return m, nil
}
func getClient(opts Options) (*mongo.Client, error) {
	if client == nil {
		var err error
		client, err = mongo.NewClient(options.Client().ApplyURI(opts.URI))
		if err != nil {
			return nil, err
		}
		//ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
		//openlogging.Info("ping to " + opts.URI)
		//err = client.Ping(ctx, readpref.Primary())
		//if err != nil {
		//
		//	return nil, err
		//}
		openlogging.Info("connecting to " + opts.URI)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			return nil, err
		}
		openlogging.Info("connected to " + opts.URI)
	}
	return client, nil
}

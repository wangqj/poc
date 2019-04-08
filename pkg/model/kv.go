//model is abstraction for persistence storage
package model

import (
	"crypto/tls"

	"errors"
	"github.com/apache/servicecomb-rokie/config"
)

var ErrMissingDomain = errors.New("domain info missing, illegal access")
var ErrNotExists = errors.New("key with labels does not exits")

type KVService interface {
	CreateOrUpdate(kv *KV) (string, error)
	//do not use primitive.ObjectID as return to decouple with mongodb, we can afford perf lost
	Exist(key, domain string, labels Labels) (string, error)
	DeleteByID(id string) error
	Delete(key, domain string, labels Labels) error
	//Get(key string, labels Labels) ([]*KV, error)
	//SaveVersion(kv *KV) error
	//GetVersionList(*KV) error
	//RollBack(kv *KV, version string) error
}

type CallOptions struct {
	DocID  string
	Labels map[string]string
}

type Options struct {
	URI      string
	PoolSize int
	SSL      bool
	TLS      *tls.Config
}

func NewKVService() (KVService, error) {
	opts := Options{
		URI:      config.GetDB().URI,
		PoolSize: config.GetDB().PoolSize,
		SSL:      config.GetDB().SSL,
	}
	if opts.SSL {

	}
	return NewMongoService(opts)
}

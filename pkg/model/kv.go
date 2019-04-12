//model is abstraction for persistence storage
package model

import (
	"crypto/tls"

	"errors"
	"github.com/apache/servicecomb-rokie/config"
	"time"
)

var ErrMissingDomain = errors.New("domain info missing, illegal access")
var ErrNotExists = errors.New("key with labels does not exits")
var ErrTooMany = errors.New("key with labels should be only one")
var ErrKeyMustNotEmpty = errors.New("must supply key if you want to get exact one result")

type KVService interface {
	CreateOrUpdate(kv *KV) (*KV, error)
	//do not use primitive.ObjectID as return to decouple with mongodb, we can afford perf lost
	Exist(key, domain string, labels Labels) (string, error)
	DeleteByID(id string) error
	Delete(key, domain string, labels Labels) error
	Find(domain string, options ...CallOption) ([]*KV, error)
	AddHistory(kv *KV) error
	//RollBack(kv *KV, version string) error
}

type Options struct {
	URI      string
	PoolSize int
	SSL      bool
	TLS      *tls.Config
	Timeout  time.Duration
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

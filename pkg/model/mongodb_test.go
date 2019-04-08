package model_test

import (
	"github.com/apache/servicecomb-rokie/pkg/model"
	"github.com/go-chassis/paas-lager"
	"github.com/go-mesh/openlogging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	log.Init(log.Config{
		Writers:     []string{"stdout"},
		LoggerLevel: "DEBUG",
	})

	logger := log.NewLogger("ut")
	openlogging.SetLogger(logger)
}
func TestMongodbService_Exist(t *testing.T) {
	t.Log("connect")
	s, err := model.NewMongoService(model.Options{
		URI: "mongodb://127.0.0.1:27017",
	})
	assert.NoError(t, err)

	t.Log("create a service level config")
	serviceLevelID, err := s.CreateOrUpdate(&model.KV{
		Key:    "timeout",
		Value:  "2s",
		Domain: "default",
		Labels: map[string]string{
			"app":     "mall",
			"service": "cart",
		},
	})
	assert.NoError(t, err)
	t.Log("new id is", serviceLevelID)

}
func TestMongodbService_CreateOrUpdate(t *testing.T) {
	t.Log("connect")
	s, err := model.NewMongoService(model.Options{
		URI: "mongodb://127.0.0.1:27017",
	})
	assert.NoError(t, err)

	t.Log("create a version value config")
	versionLevelID, err := s.CreateOrUpdate(&model.KV{
		Key:    "timeout",
		Value:  "2s",
		Domain: "default",
		Labels: map[string]string{
			"app":     "mall",
			"service": "cart",
			"version": "1.0.0",
		},
	})
	assert.NoError(t, err)
	t.Log("new id is ", versionLevelID)

	t.Log("version value config shoud exists")
	oid, err := s.Exist("timeout", "default", map[string]string{
		"app":     "mall",
		"service": "cart",
		"version": "1.0.0",
	})
	assert.NoError(t, err)
	assert.Equal(t, versionLevelID, oid)
	t.Log(oid)

	t.Log("insert app level config")
	idBefore, err := s.CreateOrUpdate(&model.KV{
		Key:    "timeout",
		Value:  "1s",
		Domain: "default",
		Labels: map[string]string{
			"app": "mall",
		},
	})
	assert.NoError(t, err)
	kvs, err := s.Find("default", model.WithKey("timeout"), model.WithLabels(map[string]string{
		"app": "mall",
	}), model.WithExactOne())
	assert.Equal(t, "1s", kvs[0].Value)

	t.Log("update app level config")
	idAfter, err := s.CreateOrUpdate(&model.KV{
		Key:    "timeout",
		Value:  "3s",
		Domain: "default",
		Labels: map[string]string{
			"app": "mall",
		},
	})
	assert.NoError(t, err)
	t.Log("update id", idAfter)
	assert.Equal(t, idBefore, idAfter)
	kvs, err = s.Find("default", model.WithKey("timeout"), model.WithLabels(map[string]string{
		"app": "mall",
	}), model.WithExactOne())
	assert.Equal(t, "3s", kvs[0].Value)

}

func TestLength(t *testing.T) {
	s, err := model.NewMongoService(model.Options{
		URI: "mongodb://127.0.0.1:27017",
	})
	kvs, err := s.Find("default", model.WithKey("timeout"), model.WithLabels(map[string]string{
		"app": "mall",
	}))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(kvs))
}

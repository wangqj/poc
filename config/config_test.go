package config_test

import (
	"github.com/apache/servicecomb-rokie/config"
	"github.com/go-chassis/go-archaius"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	err := archaius.Init()
	assert.NoError(t, err)
	b := []byte(`
db:
  uri: mongodb://admin:123@127.0.0.1:27017/rokie
  type: mongodb
  poolSize: 10
  ssl: false
  sslCA:
  sslCert:

`)
	defer os.Remove("test.yaml")
	f1, err := os.Create("test.yaml")
	assert.NoError(t, err)
	_, err = io.WriteString(f1, string(b))
	assert.NoError(t, err)
	err = config.Init("test.yaml")
	assert.NoError(t, err)
	assert.Equal(t, 10, config.GetDB().PoolSize)
}

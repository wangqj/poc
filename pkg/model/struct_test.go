package model_test

import (
	"encoding/json"
	"github.com/apache/servicecomb-rokie/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKV_UnmarshalJSON(t *testing.T) {
	kv := &model.KV{
		Value: "test",
		Labels: map[string]string{
			"test": "env",
		},
	}
	b, _ := json.Marshal(kv)
	t.Log(string(b))

	var kv2 model.KV
	err := json.Unmarshal([]byte(` 
        {"value": "1","labels":{"test":"env"}}
    `), &kv2)
	assert.NoError(t, err)
	assert.Equal(t, "env", kv2.Labels["test"])
	assert.Equal(t, "1", kv2.Value)

}

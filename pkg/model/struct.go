package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Labels map[string]string

func (m Labels) ToString() string {
	sb := strings.Builder{}
	for k, v := range m {
		sb.WriteString(k + "=" + v + ",")
	}
	return sb.String()
}
func (m Labels) Encode() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type ID struct {
	ID primitive.ObjectID `json:"_id,omitempty"`
}
type KV struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key       string              `json:"key"`
	Value     string              `json:"value"`
	ValueType string              `json:"valueType"` //ini,json,text,yaml,properties
	Domain    string              `json:"domain"`    //tenant info
	Labels    Labels              `json:"labels"`    //key has labels

	Checker string `json:"check"` //python script
}

func (kv *KV) UnmarshalJSON(bs []byte) (err error) {

	m := make(map[string]string)

	if err = json.Unmarshal(bs, &m); err != nil {
		return err
	}
	delete(m, "a")
	delete(m, "b")
	return err
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Labels map[string]string

//func (m Labels) ToString() string {
//	sb := strings.Builder{}
//	for k, v := range m {
//		sb.WriteString(k + "=" + v + ",")
//	}
//	return sb.String()
//}

type KV struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key       string             `json:"key"`
	Value     string             `json:"value"`
	ValueType string             `json:"valueType"`        //ini,json,text,yaml,properties
	Domain    string             `json:"domain"`           //tenant info
	Labels    map[string]string  `json:"labels,omitempty"` //key has labels
	Checker   string             `json:"check,omitempty"`  //python script
	Revision  int                `json:"revision"`
}
type KVHistory struct {
	KID      string `json:"id,omitempty" bson:"kvID"`
	Value    string `json:"value"`
	Checker  string `json:"check,omitempty"` //python script
	Revision int    `json:"revision"`
}

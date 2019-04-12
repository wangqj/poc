package resource

import (
	"encoding/json"
	"fmt"
	"github.com/apache/servicecomb-rokie/pkg/model"
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-mesh/openlogging"
)

const (
	FindExact               = "exact"
	FindMany                = "greedy"
	MsgDomainMustNotBeEmpty = "domain must not be empty"
	MsgIllegalFindPolicy    = "value of header X-Find can be greedy or exact"
	MsgIllegalLabels        = "label's value can not be empty, " +
		"label can not be duplicated, please check your query parameters"
)

func ReadDomain(context *restful.Context) interface{} {
	return context.ReadRestfulRequest().Attribute("domain")
}
func ReadFindPolicy(context *restful.Context) string {
	policy := context.ReadRestfulRequest().HeaderParameter("X-Find")
	if policy == "" {
		return FindMany
	}
	return policy
}
func WriteErrResponse(context *restful.Context, status int, msg string) {
	context.WriteHeader(status)
	b, _ := json.MarshalIndent(&ErrorMsg{Msg: msg}, "", " ")
	context.Write(b)
}

func ErrLog(action string, kv *model.KV, err error) {
	openlogging.Error(fmt.Sprintf("[%s] [%v] err:%s", action, kv, err.Error()))
}

func InfoLog(action string, kv *model.KV) {
	openlogging.Info(
		fmt.Sprintf("[%s] [%s:%s] in [%s] success", action, kv.Key, kv.Value, kv.Domain))
}

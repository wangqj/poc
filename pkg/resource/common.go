package resource

import (
	"encoding/json"
	"fmt"
	"github.com/apache/servicecomb-rokie/pkg/model"
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-mesh/openlogging"
	"net/http"
)

func WriteErrResponse(context *restful.Context, status int, msg string) {
	context.WriteHeader(http.StatusInternalServerError)
	b, _ := json.MarshalIndent(&ErrorMsg{Msg: msg}, "", " ")
	context.Write(b)
}

func ErrLog(action string, kv *model.KV, err error) {
	openlogging.Error(fmt.Sprintf("[%s] [%s] err:%s", action, kv, err.Error()))
}

func InfoLog(action string, kv *model.KV) {
	openlogging.Info(
		fmt.Sprintf("[%s] [%s:%s] in [%s] success", action, kv.Key, kv.Value, kv.Domain))
}

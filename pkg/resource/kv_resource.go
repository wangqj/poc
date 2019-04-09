//resource package hold http rest API
package resource

import (
	"encoding/json"
	"github.com/apache/servicecomb-rokie/pkg/model"
	goRestful "github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-mesh/openlogging"
	"net/http"
)

type KVResource struct {
}

func (r *KVResource) Put(context *restful.Context) {
	var err error
	key := context.ReadPathParameter("key")
	kv := new(model.KV)
	decoder := json.NewDecoder(context.ReadRequest().Body)
	if err = decoder.Decode(kv); err != nil {
		WriteErrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	domain := ReadDomain(context)
	if domain == nil {
		WriteErrResponse(context, http.StatusInternalServerError, MsgDomainMustNotBeEmpty)
	}
	kv.Key = key
	kv.Domain = domain.(string)
	s, err := model.NewKVService()
	if err != nil {
		WriteErrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = s.CreateOrUpdate(kv)
	if err != nil {
		ErrLog("put", kv, err)
		WriteErrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	InfoLog("put", kv)
	context.WriteHeader(http.StatusOK)
	context.Write([]byte(`true`))

}
func (r *KVResource) Find(context *restful.Context) {
	var err error
	key := context.ReadPathParameter("key")
	if key == "" {
		WriteErrResponse(context, http.StatusForbidden, "key must not be empty")
		return
	}
	values := context.ReadRequest().URL.Query()
	labels := make(map[string]string, len(values))
	for k, v := range values {
		if len(v) != 1 {
			WriteErrResponse(context, http.StatusBadRequest, MsgIllegalLabels)
			return
		}
		labels[k] = v[0]
	}
	s, err := model.NewKVService()
	if err != nil {
		WriteErrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	domain := ReadDomain(context)
	if domain == nil {
		WriteErrResponse(context, http.StatusInternalServerError, MsgDomainMustNotBeEmpty)
		return
	}
	policy := ReadFindPolicy(context)
	var kvs []*model.KV
	switch policy {
	case FindMany:
		kvs, err = s.Find(domain.(string), model.WithKey(key), model.WithLabels(labels))
	case FindExactOne:
		kvs, err = s.Find(domain.(string), model.WithKey(key), model.WithLabels(labels),
			model.WithExactOne())
	default:
		WriteErrResponse(context, http.StatusBadRequest, MsgIllegalFindPolicy)
		return
	}
	if err != nil {
		WriteErrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	err = context.WriteHeaderAndJSON(http.StatusOK, kvs, goRestful.MIME_JSON)
	if err != nil {
		openlogging.Error(err.Error())
	}

}
func (r *KVResource) Delete(context *restful.Context) {

}

//URLPatterns defined config operations
func (r *KVResource) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method:           http.MethodPut,
			Path:             "/v1/kv/{key}",
			ResourceFuncName: "Put",
			FuncDesc:         "create or update key value",
			Parameters: []*restful.Parameters{
				{
					DataType:  "string",
					Name:      "key",
					ParamType: goRestful.PathParameterKind,
				}, {
					DataType:  "string",
					Name:      "X-Domain-Name",
					ParamType: goRestful.HeaderParameterKind,
					Desc:      "set kv to other tenant",
				}, {
					DataType:  "string",
					Name:      "X-Realm",
					ParamType: goRestful.HeaderParameterKind,
					Desc:      "set kv to heterogeneous config server",
				},
			},
			Returns: []*restful.Returns{
				{
					Code:    http.StatusOK,
					Message: "set key value success",
				},
			},
			Consumes: []string{"application/json"},
			Produces: []string{"application/json"},
			Read:     &KVBody{},
		}, {
			Method:           http.MethodGet,
			Path:             "/v1/kv/{key}",
			ResourceFuncName: "Find",
			FuncDesc:         "get key values",
			Parameters: []*restful.Parameters{
				{
					DataType:  "string",
					Name:      "key",
					ParamType: goRestful.PathParameterKind,
				}, {
					DataType:  "string",
					Name:      "X-Domain-Name",
					ParamType: goRestful.HeaderParameterKind,
				}, {
					DataType:  "string",
					Name:      "X-Find",
					ParamType: goRestful.HeaderParameterKind,
					Desc:      "many or exactOne",
				},
			},
			Returns: []*restful.Returns{
				{
					Code:    http.StatusOK,
					Message: "get key value success",
					Model:   []*KVBody{},
				},
			},
			Consumes: []string{"application/json"},
			Produces: []string{"application/json"},
			Read:     &KVBody{},
		},
	}
}

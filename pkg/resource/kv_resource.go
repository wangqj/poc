//resource package hold http rest API
package resource

import (
	"encoding/json"
	"github.com/apache/servicecomb-rokie/pkg/model"
	goRestful "github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/server/restful"
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
	domain := context.ReadRestfulRequest().Attribute("domain")
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
func (r *KVResource) Get(context *restful.Context) {

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
					Desc:      "pull kv from other tenant",
				}, {
					DataType:  "string",
					Name:      "X-Realm",
					ParamType: goRestful.HeaderParameterKind,
					Desc:      "pull kv from heterogeneous config server",
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
			Path:             "/v1//kv/{key}",
			ResourceFuncName: "Put",
			FuncDesc:         "get key values",
			Parameters: []*restful.Parameters{
				{
					DataType:  "string",
					Name:      "project",
					ParamType: goRestful.PathParameterKind,
				}, {
					DataType:  "string",
					Name:      "key",
					ParamType: goRestful.PathParameterKind,
				}, {
					DataType:  "string",
					Name:      "X-Domain-Name",
					ParamType: goRestful.HeaderParameterKind,
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
		},
	}
}

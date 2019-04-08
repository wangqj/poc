package heterogeneous

import (
	"crypto/tls"
	"github.com/apache/servicecomb-rokie/pkg/model"
)

//key is {domain}:{config_server_name}
var configServers = make(map[string]model.KVService, 0)

//NewConfigServer create a new new config server instance
func NewConfigServer(endpoint EndpointOptions) (model.KVService, error) {

}
func RegisterEmbbededConfigServer(name string, s model.KVService) {

}
func RegisterDataCenter(name string, s model.KVService) {

}

//EndpointOptions represents a heterogeneous config server, name "default" represents mongodb
type EndpointOptions struct {
	Name         string
	Address      string
	EndpointType string
	TLS          *tls.Config
}

func Init() error {
}

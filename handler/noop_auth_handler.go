package handler

import (
	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
)

type NoopAuthHandler struct{}

func (bk *NoopAuthHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	inv.SetMetadata("domain", "default")
	chain.Next(inv, cb)
}

func newDomainResolver() handler.Handler {
	return &NoopAuthHandler{}
}

func (bk *NoopAuthHandler) Name() string {
	return "auth-handler"
}
func init() {
	handler.RegisterHandler("auth-handler", newDomainResolver)
}

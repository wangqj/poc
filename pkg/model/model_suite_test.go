package model_test

import (
	"testing"

	"github.com/go-chassis/paas-lager"
	"github.com/go-mesh/openlogging"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Model Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	log.Init(log.Config{
		Writers:     []string{"stdout"},
		LoggerLevel: "DEBUG",
	})

	logger := log.NewLogger("ut")
	openlogging.SetLogger(logger)
})

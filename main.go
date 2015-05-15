package main

import (
	"net/http"

	"github.com/pivotal-cf-experimental/cf-mysql-broker/broker"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-golang/lager"
)

func main() {
	serviceBroker := &broker.Broker{}
	logger := lager.NewLogger("my-service-broker")

	brokerAPI := brokerapi.New(serviceBroker, logger, brokerapi.BrokerCredentials{})
	http.Handle("/", brokerAPI)
	http.ListenAndServe("0.0.0.0:3000", nil)
}

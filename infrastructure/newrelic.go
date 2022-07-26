package infrastructure

import (
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApm(conf *config.Config) (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("kumparan-backend-test"),
		newrelic.ConfigLicense("595b7efbadaa629bc399ac71f6355ceacd4bNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	return app, err
}

package infrastructure

import (
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApm(conf *config.Config) (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(conf.AppName),
		newrelic.ConfigLicense(conf.NewRelic_License),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	return app, err
}

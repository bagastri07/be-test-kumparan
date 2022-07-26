package infrastructure

import (
	"os"

	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
)

func NewRelicApm(conf *config.Config) (*newrelic.Application, error) {
	if !conf.NewRelicOn {
		var Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

		Logger.Warn().Msg("New Relic APM is off")
		return nil, nil
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(conf.AppName),
		newrelic.ConfigLicense(conf.NewRelic_License),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	return app, err
}

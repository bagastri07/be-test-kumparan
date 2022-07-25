package utils

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

func init() {
	zerolog.TimestampFieldName = "timestamp"
}

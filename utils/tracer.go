package utils

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func StartTracer(ctx context.Context, tag, trackerName string) *newrelic.Segment {
	trx := newrelic.FromContext(ctx)
	segment := newrelic.StartSegment(trx, fmt.Sprintf("%s %s", tag, trackerName))
	return segment
}

func StartControllerTracer(eCtx echo.Context, tag, trackerName string) *newrelic.Segment {
	txn := nrecho.FromContext(eCtx)
	segment := newrelic.StartSegment(txn, fmt.Sprintf("%s %s", tag, trackerName))
	return segment
}

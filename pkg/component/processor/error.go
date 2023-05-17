package processor

import (
	"github.com/nehal119/benthos-119/pkg/message"
	"github.com/nehal119/benthos-119/pkg/tracing"
)

// MarkErr marks a message part as having failed. This includes modifying
// metadata to contain this error as well as adding the error to a tracing span
// if the message has one.
func MarkErr(part *message.Part, span *tracing.Span, err error) {
	if err == nil {
		return
	}
	if part != nil {
		part.ErrorSet(err)
	}
	if span == nil && part != nil {
		span = tracing.GetActiveSpan(part)
	}
	if span != nil {
		span.SetTag("error", "true")
		span.LogKV(
			"event", "error",
			"type", err.Error(),
		)
	}
}

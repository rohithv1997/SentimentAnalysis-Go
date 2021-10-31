package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

type loggerBuilder struct {
	timeFormat, module, class, method string
}

func NewLoggerBuilder() *loggerBuilder {
	return &loggerBuilder{
		timeFormat: time.RFC3339Nano,
	}
}

func (builder *loggerBuilder) SetModule(module string) {
	builder.module = module
}

func (builder *loggerBuilder) SetClass(class string) {
	builder.class = class
}

func (builder *loggerBuilder) SetMethod(method string) {
	builder.method = method
}

func (builder *loggerBuilder) Build() *zerolog.Event {
	zerolog.TimeFieldFormat = builder.timeFormat
	return log.Log().
		Str(MetadataModule, builder.module).
		Str(MetadataClass, builder.class).
		Str(MetadataMethod, builder.method)
}

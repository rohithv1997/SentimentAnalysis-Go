package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

type builder struct {
	timeFormat, module, class, method string
}

func NewBuilder() *builder {
	return &builder{
		timeFormat: time.RFC3339Nano,
	}
}

func (builder *builder) SetModule(module string) {
	builder.module = module
}

func (builder *builder) SetClass(class string) {
	builder.class = class
}

func (builder *builder) SetMethod(method string) {
	builder.method = method
}

func (builder *builder) Build() *zerolog.Event {
	zerolog.TimeFieldFormat = builder.timeFormat
	return log.Log().
		Str(MetadataModule, builder.module).
		Str(MetadataClass, builder.class).
		Str(MetadataMethod, builder.method)
}

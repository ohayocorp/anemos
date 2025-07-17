package collectcrds

import (
	"github.com/ohayocorp/anemos/pkg/core"
)

func Add(builder *core.Builder) *core.Component {
	return AddWithOptions(builder, nil)
}

func AddWithOptions(builder *core.Builder, options *Options) *core.Component {
	component := NewComponent(options)
	builder.AddComponent(component)

	return component
}

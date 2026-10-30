package core

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/ohayocorp/anemos/pkg/js"
)

// GenerateUUID returns a randomly generated version 4 UUID.
func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func registerUUID(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(GenerateUUID)).JsModule("uuid").JsName("generateUUID")
}

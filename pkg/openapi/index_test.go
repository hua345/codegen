package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"testing"
)

const testOpenApiFile = "book.yaml"

func TestOpenApi(t *testing.T) {
	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err := loader.LoadSwaggerFromFile(testOpenApiFile)
	if err != nil {
		t.Error(err)
	}
	t.Log(swagger.Info.Title)
	data, err := swagger.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}

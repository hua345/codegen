package openapi

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"testing"
)

const testOpenApiFile = "book.yaml"

func handleApi(api *openapi3.Operation, t *testing.T) {
	t.Log("Description", api.Description)
	t.Log("tags", api.Tags)
	for index, value := range api.Parameters {
		t.Log(index, value)
	}
	t.Log(api.Parameters)
	t.Log(api.RequestBody)
	t.Log(api.Responses)
}
func TestOpenApi(t *testing.T) {
	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err := loader.LoadSwaggerFromFile(testOpenApiFile)
	if err != nil {
		t.Error(err)
	}
	err = swagger.Validate(context.TODO())
	if err != nil {
		t.Error(err)
	} else {
		t.Log("接口文档有效")
	}
	t.Log(swagger.Info.Title)
	for key, value := range swagger.Paths {
		t.Log(key)
		if value.Get != nil {
			handleApi(value.Get, t)
		}
		if value.Post != nil {
			handleApi(value.Post, t)
		}
		if value.Put != nil {
			handleApi(value.Put, t)
		}
		if value.Delete != nil {
			handleApi(value.Delete, t)
		}
	}
}

package openapi

import (
	"context"
	"fmt"
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
func handle(schema *openapi3.SchemaRef)(string,error) {
	schemaType := schema.Value.Type
	schemaFormat := schema.Value.Format
	if schemaType != "" {
		var result string
		switch schemaType {
		case "array":
			// For arrays, we'll get the type of the Items and throw a
			// [] in front of it.
			arrayType, err := schemaToGoType(schema.Items, true)
			if err != nil {
				return "", fmt.Errorf("error generating type for array: %s", err)
			}
			result = "[]" + arrayType
			// Arrays are nullable, so we return our result here, whether or
			// not this field is required
			return result, nil
		case "integer":
			// We default to int32 if format doesn't ask for something else.
			if schemaFormat == "int64" {
				result = "int64"
			} else if schemaFormat == "int32" || schemaFormat == "" {
				result = "int32"
			} else {
				return "", fmt.Errorf("invalid integer format: %s", schemaFormat)
			}
		case "number":
			// We default to float for "number"
			if schemaFormat == "double" {
				result = "float64"
			} else if schemaFormat == "float" || schemaFormat == "" {
				result = "float32"
			} else {
				return "", fmt.Errorf("invalid number format: %s", schemaFormat)
			}
		case "boolean":
			if schemaFormat != "" {
				return "", fmt.Errorf("invalid format (%s) for boolean", f)
			}
			result = "bool"
		case "string":
			switch schemaFormat {
			case "", "password":
				result = "string"
			case "date-time", "date":
				result = "time.Time"
			default:
				return "", fmt.Errorf("invalid string format: %s", f)
			}
		default:
			return "", fmt.Errorf("unhandled Schema type: %s", t)
		}
	}
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
	for key, value := range swagger.Components.Schemas{
		t.Log(key)
		t.Log(value.Value.Type)
		aa:= value.Value.Properties
		t.Log(aa)
	}
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

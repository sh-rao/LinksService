package validation

import (
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"

	model "../link"
)

// TODO Read this from CONFIG
const (
	jsonClassicSchema = "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema\",\n  \"type\": \"object\",\n  \"description\": \"Classic link schema\",\n  \"default\": {},\n  \"additionalProperties\": true,\n  \"required\": [\n    \"title\",\n    \"url\"\n  ],\n  \"properties\": {\n    \"title\": {\n      \"type\": \"string\",\n      \"description\": \"Title of the link\",\n      \"examples\": [\n        \"Test Title\"\n      ],\n      \"maxLength\": 140,\n      \"minLength\": 1\n    },\n    \"url\": {\n      \"type\": \"string\",\n      \"format\": \"uri\",\n      \"description\": \"The url to the link.\",\n      \"pattern\": \"https?://(www\\\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\\\\.[a-zA-Z0-9()]{1,6}\\\\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)\",\n      \"examples\": [\n        \"https://example.com\"\n      ]\n    }\n  }\n}"
)

func ValidateClassicLink(classicLink model.Classic) error {
	schemaLoader := gojsonschema.NewStringLoader(jsonClassicSchema)
	documentLoader := gojsonschema.NewGoLoader(classicLink)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if !result.Valid() {
		fmt.Printf("The document is not valid. see errors :\n")
		var errorMsg string
		for _, err := range result.Errors() {
			errorMsg += err.Description() + "\n"
		}
		return errors.New(errorMsg)
	}

	return nil
}

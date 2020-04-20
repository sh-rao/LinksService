package validation

import (
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"

	"../link"
)

// TODO Read this from CONFIG
const (
	jsonShowListSchema = "{\n\"$schema\": \"http://json-schema.org/draft-07/schema#\",\n\n  \"type\": \"object\",\n  \"description\": \"Shows list link schema\",\n  \"default\": {},\n  \"additionalProperties\": true,\n  \"required\": [\n    \"sold_out\",\n    \"not_on_sale\",\n    \"on_sale\"\n  ],\n\n  \"properties\": {\n    \"sold_out\" : {\n      \"type\": \"object\",\n      \"$ref\": \"#/definitions/show\"\n    },\n    \"not_on_sale\" : {\n      \"type\": \"object\",\n      \"$ref\": \"#/definitions/show\"\n    },\n    \"on_sale\" : {\n      \"type\": \"array\",\n      \"$ref\": \"#/definitions/show\",\n      \"default\": []\n    }\n  },\n\n  \"definitions\": {\n      \"show\": {\n      \"type\": \"object\",\n      \"required\": [\"title\", \"date\", \"venue\"],\n      \"properties\": {\n        \"title\": {\n          \"type\": \"string\",\n          \"description\": \"Title of the show.\",\n          \"examples\": [\n            \"Van Halen\"\n          ],\n          \"maxLength\": 140,\n          \"minLength\": 1\n        },\n        \"date\": {\n          \"type\": \"string\",\n          \"format\": \"date\",\n          \"description\": \"Date of the show.\",\n          \"examples\": [\n            \"Apr 01 2019\"\n          ]\n        },\n        \"venue\": {\n          \"type\": \"string\",\n          \"description\": \"Venue of the show.\",\n          \"examples\": [\n            \"Melbourne, AU\"\n          ],\n          \"maxLength\": 100,\n          \"minLength\": 10\n        }\n      }\n    }\n  }\n}\n"
)

func ValidateShowsList(showsList link.ShowsList) error {
	schemaLoader := gojsonschema.NewStringLoader(jsonShowListSchema)
	documentLoader := gojsonschema.NewGoLoader(showsList)

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

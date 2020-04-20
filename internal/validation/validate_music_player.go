package validation

import (
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"

	"../link"
)

// TODO Read this from CONFIG
const (
	jsonMusicPlayerSchema = "{\n  \"$schema\": \"http://json-schema.org/draft-07/schema\",\n  \"type\": \"object\",\n  \"description\": \"Music player link schema.\",\n  \"default\": {},\n  \"additionalProperties\": true,\n  \"required\": [\n    \"platform\",\n    \"audio_player\"\n  ],\n  \"properties\": {\n    \"platform\": {\n      \"type\": \"string\",\n      \"format\": \"uri\",\n      \"description\": \"Link to the music platform.\",\n      \"pattern\": \"https?://(www\\\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\\\\.[a-zA-Z0-9()]{1,6}\\\\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)\",\n      \"examples\": [\n        \"Spotify\",\n        \"Apple Music\"\n      ]\n    },\n    \"audio_player\": {\n      \"type\": \"string\",\n      \"format\": \"uri\",\n      \"description\": \"The url to the embedded audio player.\",\n      \"pattern\": \"https?://(www\\\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\\\\.[a-zA-Z0-9()]{1,6}\\\\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)\"\n    }\n  }\n}"
)

func ValidateMusicPlayer(musicPlayer link.MusicPlayer) error {
	schemaLoader := gojsonschema.NewStringLoader(jsonMusicPlayerSchema)
	documentLoader := gojsonschema.NewGoLoader(musicPlayer)

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

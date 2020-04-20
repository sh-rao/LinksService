package validation_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"../link"
	v "../validation"
)

func TestValidateClassicLink(t *testing.T) {
	t.Run("good json body", func(t *testing.T) {
		classicLinkJson := "{\n  \"title\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZAB\",\n  \"url\": \"http://example.com?type=what\"\n}"
		var classicLink link.Classic
		json.Unmarshal([]byte(classicLinkJson), &classicLink)
		err := v.ValidateClassicLink(classicLink)
		assert.Nil(t, err)
	})
	t.Run("bad json body", func(t *testing.T) {
		classicLinkJson := "{\n  \"bad_title\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZAB\",\n  \"url\": \"://example.com?type=what\"\n}"
		var classicLink link.Classic
		json.Unmarshal([]byte(classicLinkJson), &classicLink)
		err := v.ValidateClassicLink(classicLink)
		assert.NotNil(t, err)
	})
}

package validation_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"../link"
	v "../validation"
)

func TestValidateShowsList(t *testing.T) {
	t.Run("good json body", func(t *testing.T) {
		showsListLinkJson := "{\n\"sold_out\": {\n\"title\": \"ABCDEFGHIJ\",\n\"date\": \"ABCDEF\",\n\"venue\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZABCD\"\n},\n\"not_on_sale\": {\n\"title\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZAB\",\n\"date\": \"ABCDEFGHIJKLMNOPQ\",\n\"venue\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZABCD\"\n},\n\"on_sale\": [{\n\"title\": \"ABCDEFGHIJKL\",\n\"date\": \"ABCDEFGHIJKLMNOPQRSTUVWX\",\n\"venue\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZABCD\"\n}]\n}"
		var showsListLink link.ShowsList
		json.Unmarshal([]byte(showsListLinkJson), &showsListLink)
		err := v.ValidateShowsList(showsListLink)
		assert.Nil(t, err)
	})
	t.Run("bad json body", func(t *testing.T) {
		showsListLinkJson := "{\n  \"sold_out\": {\n    \"title\": \"ABCDEFGHIJ\",\n    \"venue\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZABCD\"\n  },\n  \"not_on_sale\": {\n    \"title\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZAB\",\n    \"date\": \"ABCDEFGHIJKLMNOPQ\",\n    \"venue\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZABCD\"\n  },\n  \"on_sale\": {\n    \"title\": \"ABCDEFGHIJKL\",\n    \"date\": \"ABCDEFGHIJKLMNOPQRSTUVWX\",\n    \"venue\": \"ABCDEFGHIJKLMNOPQRSTUVWXYZABCD\"\n  }\n}"
		var showsListLink link.ShowsList
		json.Unmarshal([]byte(showsListLinkJson), &showsListLink)
		err := v.ValidateShowsList(showsListLink)
		assert.NotNil(t, err)
	})
}

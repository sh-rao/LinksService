package validation_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"../link"
	v "../validation"
)

func TestValidateMusicPlayer(t *testing.T) {
	t.Run("good json body", func(t *testing.T) {
		musicPlayerLinkJson := "{\n\"platform\": \"https://spotify.com\",\n\"audio_player\": \"https://audioplayer.com?song=Time&artist=PinkFloyd\"\n}"
		var musicPlayerLink link.MusicPlayer
		json.Unmarshal([]byte(musicPlayerLinkJson), &musicPlayerLink)
		err := v.ValidateMusicPlayer(musicPlayerLink)
		assert.Nil(t, err)
	})
	t.Run("bad json body", func(t *testing.T) {
		musicPlayerLinkJson := "{\n  \"audio_player\": \"http://audioplayer.com?song=Time&artist=PinkFloyd\"\n}"
		var musicPlayerLink link.MusicPlayer
		json.Unmarshal([]byte(musicPlayerLinkJson), &musicPlayerLink)
		err := v.ValidateMusicPlayer(musicPlayerLink)
		assert.NotNil(t, err)
	})
}

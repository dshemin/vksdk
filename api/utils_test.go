package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_UtilsCheckLink(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.UtilsCheckLink(map[string]string{
		"url": "http://google.ru",
	})
	assert.NoError(t, err)
}

func TestVK_UtilsGetShortLink(t *testing.T) {
	needUserToken(t)

	shortLink, err := vkUser.UtilsGetShortLink(map[string]string{
		"url":     "http://google.ru",
		"private": "1",
	})
	assert.NoError(t, err)

	_, err = vkUser.UtilsGetLastShortenedLinks(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.UtilsDeleteFromLastShortened(map[string]string{
		"key": shortLink.Key,
	})
	assert.NoError(t, err)
}

func TestVK_UtilsGetLinkStats(t *testing.T) {
	needGroupToken(t)

	params := map[string]string{
		"key":             "8TDuIz",
		"interval":        "month",
		"intervals_count": "12",
	}

	_, err := vkGroup.UtilsGetLinkStats(params)
	assert.NoError(t, err)

	_, err = vkGroup.UtilsGetLinkStatsExtended(params)
	assert.NoError(t, err)
}

func TestVK_UtilsGetServerTime(t *testing.T) {
	needGroupToken(t)

	_, err := vkGroup.UtilsGetServerTime(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_UtilsResolveScreenName(t *testing.T) {
	needGroupToken(t)

	f := func(name string) {
		t.Helper()
		_, err := vkGroup.UtilsResolveScreenName(map[string]string{
			"screen_name": name,
		})
		if err != nil {
			t.Errorf("VK.UtilsResolveScreenName() err = %v", err)
		}
	}

	f("durov")
	f("api")
	f("app6991405")
	f("z")
}

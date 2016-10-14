package slack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareMesssage(t *testing.T) {
	slack := &Slack{
		Webhook:  "webhook",
		Channel:  "test-channel",
		Emoji:    "test-emoji",
		Username: "test-name",
	}
	// data = appendOption(data, "color", color)
	// data = appendOption(data, "channel", slack.Channel)
	// data = appendOption(data, "icon_emoji", slack.Emoji)
	// data = appendOption(data, "username", slack.Username)
	msg := slack.prepareMsg("test", danger)
	data := `"mrkdwn_in": ["text"], "text": "test", "color": "danger", "channel": "test-channel", "icon_emoji": "test-emoji", "username": "test-name"`
	expected := []byte(fmt.Sprintf(`{"attachments": [{%s}]}`, data))

	assert.Equal(t, expected, msg, "it appends only chosen options")

	// with no emoji, channel & username provided

	slack.Emoji = ""
	slack.Username = ""
	slack.Channel = ""
	msg = slack.prepareMsg("test", danger)
	data = `"mrkdwn_in": ["text"], "text": "test", "color": "danger"`
	expected = []byte(fmt.Sprintf(`{"attachments": [{%s}]}`, data))

	assert.Equal(t, expected, msg, "it appends all options")
}

func TestappendOption(t *testing.T) {
	data := `"mrkdwn_in": ["text"]`
	data = appendOption(data, "color", danger)
	expected := `"mrkdwn_in": ["text"], "color": "danger"`

	assert.Equal(t, expected, data, "it appends options in proper format")
}

func TestInit(t *testing.T) {
	slack := Init("webhook", "channel", "emoji", "name")
	assert.Equal(t, "webhook", slack.Webhook, "it sets webhook")
	assert.Equal(t, "channel", slack.Channel, "it sets channel")
	assert.Equal(t, "emoji", slack.Emoji, "it sets emoji")
	assert.Equal(t, "name", slack.Username, "it sets username")
}

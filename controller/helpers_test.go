package controller

import "testing"

func TestBotToken(t *testing.T) {
	expected := "Bot deadbeef"
	token, err := BotToken("test.token")
	if err != nil {
		t.Error(err)
	}
	if token != expected {
		t.Errorf("token mismatch: %s", token)
	}
}

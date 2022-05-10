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

func TestMatches(t *testing.T) {
	sheSaidPass := []string{
		"its so big",
		"it's so hard!",
		"its so huge ",
	}
	sheSaidFail := []string{
		"",
		"it is sooo big",
		"it's not hard",
		"deadbeef",
	}

	for _, pattern := range sheSaidPass {
		if matches(sheSaidRegex, pattern) != true {
			t.Errorf("TestMatch Pass mismatch: %s, %s", sheSaidRegex, pattern)
		}
	}
	for _, pattern := range sheSaidFail {
		if matches(sheSaidRegex, pattern) != false {
			t.Errorf("TestMatch Fail mismatch: %s, %s", sheSaidRegex, pattern)
		}
	}
}

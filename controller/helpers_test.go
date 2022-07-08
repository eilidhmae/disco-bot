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

func mustPass(regex string, patterns []string) (string, bool) {
	for _, pat := range patterns {
		if matches(regex, pat) != true {
			return pat, false
		}
	}
	return "", true
}

func mustFail(regex string, patterns []string) (string, bool) {
	for _, pat := range patterns {
		if matches(regex, pat) != false {
			return pat, false
		}
	}
	return "", true
}

func TestMatches(t *testing.T) {
	var resp string
	var ok bool

	// test sheSaidRegex
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
	resp, ok = mustPass(sheSaidRegex, sheSaidPass)
	if !ok {
		t.Errorf("sheSaidPass mismatch: %s", resp)
	}
	resp, ok = mustFail(sheSaidRegex, sheSaidFail)
	if !ok {
		t.Errorf("sheSaidFail mismatch: %s", resp)
	}
}

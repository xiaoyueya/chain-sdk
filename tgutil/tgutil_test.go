package tgutil

import "testing"

func TestSendTg(t *testing.T) {
	SendTgMsg("1944447539:AAGIYbhMpJXr-tTinsrzOkWNYbJLZmXmw34", "-524307389", "test anything", "nothing")
	select {}
}

package selenium

import (
	"github.com/tebeka/selenium"
	"github.com/viant/endly"
)

// Session represents a selenium session
type Session struct {
	ID      string
	Browser string
	driver  selenium.WebDriver
}

// SeleniumSessions reprents selenium sessions.
type sessions struct {
	Sessions map[string]*Session
}

var sessionKey = (*sessions)(nil)

func Sessions(context *endly.Context) map[string]*Session {
	var result *sessions
	if !context.Contains(sessionKey) {
		result = &sessions{
			Sessions: make(map[string]*Session),
		}
		context.Put(sessionKey, result)
	}
	context.GetInto(sessionKey, &result)
	return result.Sessions
}

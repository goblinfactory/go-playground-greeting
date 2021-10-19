package dependencyinjection

import "errors"

// -------------------------
// 		BUSINESS LOGIC
// -------------------------

//DataStore ...
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Logger ...
type Logger interface {
	Log(message string)
}

// LoggerAdapter adapts the existing log function
type LoggerAdapter func(string)

// Log ...
func (l LoggerAdapter) Log(message string) {
	l(message)
}

// SimpleLogic ...
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// SayHello ..
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name + ". Available balance is [ R 0000,000.00]. \nAmount outstanding : [ R 2,6 Million ] Please settle your Foshini Jeanpants account within 30 days.", nil
}

// SayGoodbye ..
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

// NewSimpleLogic ..
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{l, ds}
}

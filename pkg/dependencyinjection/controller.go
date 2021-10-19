package dependencyinjection

import "net/http"

// Logic ..
type Logic interface {
	SayHello(userID string) (string, error)
}

// Controller ..
type Controller struct {
	l     Logger
	logic Logic
}

// SayHello ..
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// accept interfaces ...return structs

// NewController ..
func NewController(l Logger, logic Logic) Controller {
	return Controller{l, logic}
}

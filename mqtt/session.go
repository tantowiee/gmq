package mqtt

// Session represents a Session which is a stateful interaction between a Client and a Server.
type Session struct {
	CleanSession bool
	ClientID     string
}

// NewSession creates and returns a Session.
func NewSession() *Session {
	return &Session{}
}

package sessionmgmt

import "time"

type (
	TokenT string

	SessionConfig struct {
		SessionIdleTimeoutMin uint32
		SessionHardTimeoutMin uint32
	}

	Session struct {
		Token TokenT

		CreatedAt      time.Time
		LastAccessedAt time.Time

		config     *SessionConfig
		sessionObj *SessionObj
	}

	SessionObj interface{}
)

func DefaultSessionConfig() (sc *SessionConfig) {
	sc = &SessionConfig{
		SessionIdleTimeoutMin: 120,
		SessionHardTimeoutMin: 12000,
	}
	return
}

func NewSession(sCfg *SessionConfig) (session *Session, err error) {
	session = &Session{
		config: sCfg,
	}
	if session.config == nil {
		session.config = DefaultSessionConfig()
	}
	return
}

func (session *Session) IsActive() (isActive bool) {
	currTime := time.Now().UTC()
	// Check if hard time out has exceeded
	if uint32(currTime.Sub(session.CreatedAt).Minutes()) > session.config.SessionHardTimeoutMin {
		isActive = false
		return
	}
	// Check if idle time out has exceeded
	if uint32(currTime.Sub(session.LastAccessedAt).Minutes()) > session.config.SessionIdleTimeoutMin {
		isActive = false
		return
	}
	isActive = true
	return
}

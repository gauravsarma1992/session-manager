package sessionmgmt

type (
	User struct {
		ID       string     `json:"id"`
		Username string     `json:"username"`
		Password string     `json:"-"`
		Sessions []*Session `json:"sessions"`
	}
)

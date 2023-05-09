package middleware

type Session struct {
	Id        string
	UserID    string
	CompanyID string
	Role      Role
}

type Role string

var (
	AdminRole  Role = "admin"
	ClientRole Role = "client"
	MemberRole Role = "member"
)

func (r Role) IsAdmin() bool {
	return r == AdminRole
}

func getSessionInfo(token string) (Session, error) {
	//mock session
	return Session{
		Id:        "1",
		UserID:    "2",
		CompanyID: "1",
		Role:      ClientRole,
	}, nil
}

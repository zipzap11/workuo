package invitation

type InvitationCore struct {
	ID          uint
	RecruiterID uint
	UserID      uint
	JobID       uint
}

type Service interface {
}

type Repository interface {
}

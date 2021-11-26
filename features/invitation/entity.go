package invitation

type InvitationCore struct {
	ID          uint
	RecruiterID uint
	UserID      uint
	JobID       uint
}

type Service interface {
	InviteUser(InvitationCore) error
}

type Repository interface {
	InviteUser(InvitationCore) error
}

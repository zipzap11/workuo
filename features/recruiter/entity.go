package recruiter

type RecruiterCore struct {
	ID      uint
	Company string
	Address string
	Bio     string
}

type Service interface {
	RegisterRecruiter(data RecruiterCore) error
}

type Repository interface {
	CreateRecruiter(data RecruiterCore) error
}

package recruiter

type RecruiterCore struct {
	ID       uint
	Company  string
	Address  string
	Bio      string
	Email    string
	Password string
}

type Service interface {
	RegisterRecruiter(data RecruiterCore) error
}

type Repository interface {
	CreateRecruiter(data RecruiterCore) error
}

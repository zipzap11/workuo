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
	LoginRecruiter(data RecruiterCore) (RecruiterCore, error)
}

type Repository interface {
	CreateRecruiter(data RecruiterCore) error
	CheckRecruiter(data RecruiterCore) (RecruiterCore, error)
}

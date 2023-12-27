package usecases

type UserUseCase interface {
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	Store(u *User) error
}

type CompanyUseCase struct {
	GetByEmail(email string) (*Company, error)
	Store(c *Company) error
}

type userUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) UserUseCase {
	return &userUseCase{repo}
}

type companyUseCase struct {
	repo CompanyRepository
}

func NewCompanyUseCase(repo CompanyRepository) CompanyUseCase {
	return &companyUseCase{repo}
}
package user

type GetUserByParamType string

const (
	ByEmail GetUserByParamType = "email"
	ById    GetUserByParamType = "id"
)

type GetUserParams struct {
	Id   *int
	Type GetUserByParamType
}

type IUserRepository interface {
	Create(newUser UserModel) error
	GetUser(params GetUserParams) (*UserModel, error)
	Delete(id int) error
	Update(user UserModel) error
	GetAll() ([]UserModel, error)
}

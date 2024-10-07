package usecases

import (
	"errors"

	"github.com/Edu4rdoNeves/EasyStrock/internal/core/repository"
	"github.com/Edu4rdoNeves/EasyStrock/internal/domain/model"
	"github.com/Edu4rdoNeves/EasyStrock/internal/tools"
)

type ILoginUseCases interface {
	Login(login *model.Login) (string, error)
}

type LoginUseCases struct {
	repository repository.ILoginRepository
}

func NewLoginBusiness(iRepository repository.ILoginRepository) ILoginUseCases {
	return &LoginUseCases{
		iRepository,
	}
}

func (b *LoginUseCases) Login(login *model.Login) (string, error) {
	var user model.Users

	err := b.repository.Login(login, &user)
	if err != nil {
		return "", err
	}

	if tools.SHA256Enconder(user.Password) != tools.SHA256Enconder(login.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := tools.NewJWTService().GenerateToken(uint(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}

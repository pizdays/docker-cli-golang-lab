package usecases

import (
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/domains"
)

type UseCase struct {
	Repo domains.Repository
}

// init usecase
func NewUseCase(repo domains.Repository) domains.UseCase {
	return &UseCase{
		Repo: repo,
	}
}
func (u *UseCase) CreateService() {

	return
}

func (u *UseCase) UpdateService() {

	return
}

func (u *UseCase) DeleteService() {

	return
}

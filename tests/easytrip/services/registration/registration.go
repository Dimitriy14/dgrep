package registration

import (
	"github.com/oreuta/easytrip/models"
	"github.com/oreuta/easytrip/repository"
)

type RegService interface {
	CanRegistr(data models.User) (err error)
	CanLogIN(data models.User) (user models.User, err error)
}

type RegServiceStruct struct{}

func (a *RegServiceStruct) CanRegistr(data models.User) (err error) {
	err = repository.InsertInto(data)
	return
}

func (a *RegServiceStruct) CanLogIN(data models.User) (user models.User, err error) {
	user, err = repository.CheckUser(data)
	return
}

func New() RegService {
	return &RegServiceStruct{}
}

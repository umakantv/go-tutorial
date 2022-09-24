package service

import "auth/domain"

type UserService struct {
	repo UserRepo
}

func (s UserService) Create(user domain.User) (*domain.User, error) {

}

func (s UserService) Find(username string) (*domain.User, error) {

}

func (s UserService) Update(user domain.User) (*domain.User, error) {

}

func (s UserService) Delete(user domain.User) error {

}

func NewUserService(reop)

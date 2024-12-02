package service

import "github.com/AntonZatsepilin/My-JWT/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

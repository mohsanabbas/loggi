package repository

import "fmt"

type Repository interface {
	Print(string) (string, error)
}

func NewRepository() Repository {
	return &repository{}
}

type repository struct{}

func (p *repository) Print(s string) (string, error) {
	result := fmt.Sprintf("I am from Repository: %s", s)
	return result, nil
}

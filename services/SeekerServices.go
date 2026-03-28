package services

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"www.jobRadar.com/entities"
	"www.jobRadar.com/repository"
)

type SeekerServices struct {
	repo repository.SeekerRepository
}

func NewSeekerServices(repo repository.SeekerRepository) *SeekerServices {
	return &SeekerServices{
		repo: repo,
	}
}

func (seeker_service *SeekerServices) RegisterService(ctx context.Context, seeker entities.Seeker) error {

	password, err := bcrypt.GenerateFromPassword([]byte(seeker.Password), 10)
	if err != nil {
		log.Println(err)
	}

	seeker.Password = string(password)
	return seeker_service.repo.Create(ctx, seeker)
}

func (s *SeekerServices) LoginSeekerService(
	ctx context.Context,
	email string,
	password string,
) (*entities.Seeker, error) {

	seeker, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword(
		[]byte(seeker.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return seeker, nil
}

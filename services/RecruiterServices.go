package services

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"www.jobRadar.com/entities"
	"www.jobRadar.com/repository"
)

type RecruiterServices struct {
	repo repository.RecruiterRepository
}

func NewRecruiterServices(repo repository.RecruiterRepository) *RecruiterServices {
	return &RecruiterServices{
		repo: repo,
	}
}

func (recruiter_service *RecruiterServices) RegisterService(ctx context.Context, recruiter entities.Recruiter) error {

	password, err := bcrypt.GenerateFromPassword([]byte(recruiter.Password), 10)
	if err != nil {
		log.Println(err)
		return err
	}

	recruiter.Password = string(password)
	return recruiter_service.repo.Create(ctx, recruiter)
}

func (recruiter_service *RecruiterServices) LoginService(ctx context.Context, email, password string) (*entities.Recruiter, error) {

	recruiter, err := recruiter_service.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid email")
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword(
		[]byte(recruiter.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, fmt.Errorf("invalid  password")
	}

	return recruiter, nil
}

func (recruiter_service *RecruiterServices) CreateJob(
	ctx context.Context,
	recruiterID uuid.UUID,
	job entities.JobPost,
) (uuid.UUID, error) {

	// 🔐 Basic validations
	if recruiterID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("company id is required")
	}

	if job.Name == "" {
		return uuid.Nil, fmt.Errorf("job name is required")
	}

	if job.Amount <= 0 {
		return uuid.Nil, fmt.Errorf("amount must be greater than 0")
	}

	// 🧠 Optional: more validations
	if job.Time == "" {
		return uuid.Nil, fmt.Errorf("time is required")
	}

	// 👉 Call repository
	jobID, err := recruiter_service.repo.CreatePost(ctx, recruiterID, job)
	if err != nil {
		return uuid.Nil, err
	}

	return jobID, nil
}

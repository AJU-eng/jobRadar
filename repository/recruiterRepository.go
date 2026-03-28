package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"www.jobRadar.com/entities"
)

type PostgresRecruiterRepo struct {
	db *sql.DB
}

func NewRecruiterRepo(db *sql.DB) *PostgresRecruiterRepo {
	return &PostgresRecruiterRepo{
		db: db,
	}
}

type RecruiterRepository interface {
	Create(ctx context.Context, recruiter entities.Recruiter) error
	GetByEmail(
		ctx context.Context,
		email string,
	) (*entities.Recruiter, error)
	CreatePost(
		ctx context.Context,
		recruiterID uuid.UUID,
		job entities.JobPost,
	) (uuid.UUID, error)
}

func (recruiterDB *PostgresRecruiterRepo) Create(ctx context.Context, recruiter entities.Recruiter) error {

	var recruiter_id uuid.UUID
	query := `
	INSERT INTO recruiters (
		name,
		email,
		license_no,
		password,
		location,
		phone_no
	) VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id
	`

	err := recruiterDB.db.QueryRowContext(
		ctx,
		query,
		recruiter.Name,
		recruiter.Email,
		recruiter.License_no,
		recruiter.Password,
		recruiter.Location,
		recruiter.Phone_no,
	).Scan(&recruiter_id)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to insert seeker: %w", err)
	}

	return nil
}

func (recruiterDB *PostgresRecruiterRepo) GetByEmail(
	ctx context.Context,
	email string,
) (*entities.Recruiter, error) {

	var recruiter entities.Recruiter

	log.Println(email)

	query := `
	SELECT  name, email, password, license_no, location, phone_no
	FROM recruiters
	WHERE email = $1
	`

	err := recruiterDB.db.QueryRowContext(ctx, query, email).Scan(
		&recruiter.Name,
		&recruiter.Email,
		&recruiter.Password,
		&recruiter.License_no,
		&recruiter.Location,
		&recruiter.Phone_no,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("recruiter not found")
		}

		log.Println(err)
		return nil, err
	}

	return &recruiter, nil
}

func (recruiterDB *PostgresRecruiterRepo) CreatePost(
	ctx context.Context,
	recruiterID uuid.UUID,
	job entities.JobPost,
) (uuid.UUID, error) {

	var jobID uuid.UUID

	query := `
	INSERT INTO job_posts (
		comp_id,
		name,
		description,
		amount,
		time,
		time_range,
		period
	) VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id
	`

	err := recruiterDB.db.QueryRowContext(
		ctx,
		query,
		recruiterID,
		job.Name,
		job.Description,
		job.Amount,
		job.Time,
		job.TimeRange,
		job.Period,
	).Scan(&jobID)

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create job post: %w", err)
	}

	return jobID, nil
}

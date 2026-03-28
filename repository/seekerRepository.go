package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"www.jobRadar.com/entities"
)

type PostgresSeekerRepo struct {
	db *sql.DB
}

func NewSeekerRepo(db *sql.DB) *PostgresSeekerRepo {
	return &PostgresSeekerRepo{
		db: db,
	}
}

type SeekerRepository interface {
	Create(ctx context.Context, seeker entities.Seeker) error
	GetByEmail(
		ctx context.Context,
		email string,
	) (*entities.Seeker, error)
}

func (seekerDB *PostgresSeekerRepo) Create(ctx context.Context, seeker entities.Seeker) error {

	var seeker_id uuid.UUID
	query := `
	INSERT INTO seekers (
		name,
		email,
		gender,
		password,
		age,
		qualification,
		adhar_no,
		phone_no,
		location
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id
	`

	err := seekerDB.db.QueryRowContext(
		ctx,
		query,
		seeker.Name,
		seeker.Email,
		seeker.Gender,
		seeker.Password,
		seeker.Age,
		seeker.Qualification,
		seeker.Adhar_no,
		seeker.Phone_no,
		seeker.Location,
	).Scan(&seeker_id)

	if err != nil {
		return fmt.Errorf("failed to insert seeker: %w", err)
	}

	return nil
}

func (seekerDB *PostgresSeekerRepo) GetByEmail(
	ctx context.Context,
	email string,
) (*entities.Seeker, error) {

	var seeker entities.Seeker

	log.Println(email, "repo")

	query := `
	SELECT  name, email, password, gender, age,
	       qualification, adhar_no, phone_no, location
	FROM seekers
	WHERE email = $1
	`

	err := seekerDB.db.QueryRowContext(ctx, query, email).Scan(
		&seeker.Name,
		&seeker.Email,
		&seeker.Password,
		&seeker.Gender,
		&seeker.Age,
		&seeker.Qualification,
		&seeker.Adhar_no,
		&seeker.Phone_no,
		&seeker.Location,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid email or password")
		}
		return nil, err
	}

	return &seeker, nil
}

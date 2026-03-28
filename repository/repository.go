package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func New(addr string, maxOpenConnecs, maxIdleConnecs int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	duration, err := time.ParseDuration(maxIdleTime)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	context, cancel := context.WithTimeout(context.Background(), duration)

	defer cancel()

	err = db.PingContext(context)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

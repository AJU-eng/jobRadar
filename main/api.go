package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"www.jobRadar.com/handler"
	"www.jobRadar.com/repository"
	"www.jobRadar.com/services"
)

type application struct {
	config Config
}

type Config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	addr           string
	maxOpenConnecs int
	maxIdleConnecs int
	maxIdleTime    string
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := repository.New(
		app.config.dbConfig.addr,
		app.config.dbConfig.maxOpenConnecs,
		app.config.dbConfig.maxIdleConnecs,
		app.config.dbConfig.maxIdleTime)

	if err != nil {
		log.Println(err)
		return nil
	}

	seeker_repo := repository.NewSeekerRepo(db)
	seeker_service := services.NewSeekerServices(seeker_repo)
	seeker_handler := handler.NewSeekerHandler(*seeker_service)
	recuiter_repo := repository.NewRecruiterRepo(db)
	recruiter_service := services.NewRecruiterServices(recuiter_repo)
	recruiter_handler := handler.NewRecruiteHandler(*recruiter_service)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/seekerRegister", seeker_handler.Register)
		r.Post("/recruiterRegister", recruiter_handler.Register)
		r.Post("/recruiterLogin", recruiter_handler.LoginRecruiter)
		r.Post("/seekerLogin", seeker_handler.LoginSeeker)
		r.Post("/createJob", recruiter_handler.CreateJob)
		r.Get("/jobs", recruiter_handler.GetJobsByLocation)
	})

	return r
}

func (app *application) Run(mux http.Handler) error {
	server := http.Server{
		Addr:           app.config.addr,
		Handler:        mux,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 10,
	}

	return server.ListenAndServe()
}

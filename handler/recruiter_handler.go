package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"www.jobRadar.com/entities"
	"www.jobRadar.com/services"
)

type RecruiterHandler struct {
	service services.RecruiterServices
}

func NewRecruiteHandler(service services.RecruiterServices) *RecruiterHandler {
	return &RecruiterHandler{
		service: service,
	}
}

func (recruite *RecruiterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var RecruiterRequest entities.Recruiter

	err := json.NewDecoder(r.Body).Decode(&RecruiterRequest)

	if err != nil {
		log.Println(err)
		return
	}

	errr := recruite.service.RegisterService(r.Context(), RecruiterRequest)

	if errr != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "recruiter registered",
	})
}

func (recruite *RecruiterHandler) LoginRecruiter(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(req, "handler")

	_, err = recruite.service.LoginService(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "recruiter logined",
	})
}

func (recruite *RecruiterHandler) CreateJob(w http.ResponseWriter, r *http.Request) {

	var req entities.JobPost

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// 🔐 Get recruiter ID from context (set by middleware)
	recruiterID, err := uuid.Parse("75406089-2f44-4061-99a4-9b85bee7db9c")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Map request → entity
	job := entities.JobPost{
		Name:        req.Name,
		Description: req.Description,
		Amount:      req.Amount,
		Time:        req.Time,
		TimeRange:   req.TimeRange,
		Period:      req.Period,
	}

	// Call service
	jobID, err := recruite.service.CreateJob(r.Context(), recruiterID, job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Response
	resp := map[string]interface{}{
		"message": "job created successfully",
		"job_id":  jobID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

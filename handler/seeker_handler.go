package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"www.jobRadar.com/entities"
	"www.jobRadar.com/services"
)

type SeekerHandler struct {
	service services.SeekerServices
}

func NewSeekerHandler(service services.SeekerServices) *SeekerHandler {
	return &SeekerHandler{
		service: service,
	}
}

func (s *SeekerHandler) Register(w http.ResponseWriter, r *http.Request) {
	var SeekerRequest entities.Seeker

	err := json.NewDecoder(r.Body).Decode(&SeekerRequest)

	if err != nil {
		log.Println(err)
		return
	}

	errr := s.service.RegisterService(r.Context(), SeekerRequest)

	if errr != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "seeker registered",
	})
}

func (seeker *SeekerHandler) LoginSeeker(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, err = seeker.service.LoginSeekerService(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode("seeker logined")
}

package httpVacancy

import (
	"github.com/gorilla/mux"
	"joblessness/haha/middleware"
	"joblessness/haha/vacancy"
)

func RegisterHTTPEndpoints(router *mux.Router, m *middleware.AuthMiddleware, uc vacancy.VacancyUseCase) {
	h := NewHandler(uc)

	router.HandleFunc("/vacancies", m.CheckAuth(h.CreateVacancy)).Methods("POST")
	router.HandleFunc("/vacancies", h.GetVacancies).Methods("GET")
	router.HandleFunc("/vacancies/{vacancy_id}", h.GetVacancy).Methods("GET")
	router.HandleFunc("/vacancies/{vacancy_id}", m.CheckAuth(h.ChangeVacancy)).Methods("PUT")
	router.HandleFunc("/vacancies/{vacancy_id}", m.CheckAuth(h.DeleteVacancy) ).Methods("DELETE")
}
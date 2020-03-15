package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"joblessness/haha/models"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

type VacancyHandler struct {
	Vacancies map[uint]*models.Vacancy
	Mu        sync.RWMutex
	VacancyId uint64
}

func (api *VacancyHandler) getNewVacancyId() uint64 {
	return atomic.AddUint64(&api.VacancyId, 1)
}

func NewVacancyHandler() *VacancyHandler {
	return &VacancyHandler {
		Vacancies: map[uint]*models.Vacancy {
			1: {1, "name", "description", "skills", 100500, "address", "phone number"},
		},
		Mu:        sync.RWMutex{},
		VacancyId: 1,
	}
}

func (api *VacancyHandler) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /vacancies")

	var vacancy models.Vacancy
	json.NewDecoder(r.Body).Decode(&vacancy)

	if vacancy.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newId := api.getNewVacancyId()
	vacancy.ID = newId

	api.Mu.Lock()
	api.Vacancies[uint(newId)] = &vacancy
	api.Mu.Unlock()

	type Response struct {
		ID uint64 `json:"id"`
	}

	jsonData, err := json.Marshal(Response{newId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (api *VacancyHandler) GetVacancies(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /vacancies")

	var vacancies []models.Vacancy
	api.Mu.RLock()
	for _, vacancy := range api.Vacancies {
		vacancies = append(vacancies, *vacancy)
	}
	api.Mu.RUnlock()

	if len(vacancies) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	jsonData, err := json.Marshal(vacancies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *VacancyHandler) GetVacancy(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /vacancies/{vacancy_id}")

	vacancyId, err := strconv.Atoi(mux.Vars(r)["vacancy_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	api.Mu.RLock()
	vacancy, ok := api.Vacancies[uint(vacancyId)]
	api.Mu.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(vacancy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (api *VacancyHandler) ChangeVacancy(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT /vacancies/{vacancy_id}")

	vacancyId, err := strconv.ParseUint(mux.Vars(r)["vacancy_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := api.Vacancies[uint(vacancyId)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)

	name := data["name"]
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := data["description"]
	skills := data["skills"]
	salary, _ := strconv.Atoi(data["salary"])
	address := data["address"]
	phoneNumber := data["phone-number"]

	api.Mu.Lock()
	api.Vacancies[uint(vacancyId)] = &models.Vacancy{vacancyId, name, description, skills, salary, address, phoneNumber}
	api.Mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func (api *VacancyHandler) DeleteVacancy(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE /vacancies/{vacancy_id}")

	vacancyId, err := strconv.Atoi(mux.Vars(r)["vacancy_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	api.Mu.Lock()
	if _, ok := api.Vacancies[uint(vacancyId)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		api.Mu.Unlock()
		return
	}

	delete(api.Vacancies, uint(vacancyId))
	api.Mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

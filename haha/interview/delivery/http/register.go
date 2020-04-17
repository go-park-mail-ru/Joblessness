package interviewHttp

import (
	"github.com/gorilla/mux"
	interviewInterfaces "joblessness/haha/interview/interfaces"
	"joblessness/haha/middleware"
)

func RegisterHTTPEndpoints(router *mux.Router, m *middleware.SessionHandler, uc interviewInterfaces.InterviewUseCase) {
	h := NewHandler(uc)
	chatRouter := router.PathPrefix("/chat").Subrouter()

	router.HandleFunc("/{summary_id:[0-9]+}/response", m.OrganizationRequired(h.ResponseSummary)).Methods("PUT")
	chatRouter.HandleFunc("", m.UserRequired(h.EnterChat)).Methods("PUT")
	chatRouter.HandleFunc("/conversation/{user_id:[0-9]+}", m.UserRequired(h.History)).Methods("GET")
	chatRouter.HandleFunc("/conversation/", m.UserRequired(h.GetConversations)).Methods("GET")
}
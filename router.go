//package main
//
//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"github.com/kabukky/httpscerts"
//	"log"
//	"net/http"
//)
//
//func echoFunc(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("POST /users/logout")
//
//	Cors.PrivateApi(&w, r)
//
//	params := mux.Vars(r)
//	message := params["message"]
//	fmt.Fprintf(w, "Hello %s!", message)
//}
//
//
//func contentTipeMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Add("Content-Type", "application/json")
//		next.ServeHTTP(w, r)
//	})
//}
//
//func StartRouter() {
//	router := mux.NewRouter().PathPrefix("/api").Subrouter()//.StrictSlash(true)
//
//	router.Use(contentTipeMiddleware)
//	router.HandleFunc("/echo/{message}", echoFunc)
//	router.Methods("OPTIONS").HandlerFunc(Cors.Preflight)
//
//	// users
//	authApi := NewAuthHandler()
//
//	router.HandleFunc("/users/login", authApi.Login).Methods("POST")
//	router.HandleFunc("/users/logout", authApi.Logout).Methods("POST")
//	router.HandleFunc("/users", authApi.Register).Methods("POST")
//
//	router.HandleFunc("/user/{user_id}", authApi.GetUserPage).Methods("GET")
//	router.HandleFunc("/users/{user_id}/avatar", authApi.SetAvatar).Methods("POST")
//	router.HandleFunc("/user/{user_id}", authApi.ChangeUserInfo).Methods("POST")
//
//	// vacancies
//	vacancyApi := NewVacancyHandler()
//
//	router.HandleFunc("/vacancies", vacancyApi.CreateVacancy).Methods("POST")
//	router.HandleFunc("/vacancies", vacancyApi.GetVacancies).Methods("GET")
//	router.HandleFunc("/vacancies/{vacancy_id}", vacancyApi.GetVacancy).Methods("GET")
//	router.HandleFunc("/vacancies/{vacancy_id}", vacancyApi.ChangeVacancy).Methods("PUT")
//	router.HandleFunc("/vacancies/{vacancy_id}", vacancyApi.DeleteVacancy).Methods("DELETE")
//
//	// summaries
//	summaryApi := NewSummaryHandler()
//
//	router.HandleFunc("/summaries", summaryApi.CreateSummary).Methods("POST")
//	router.HandleFunc("/summaries", summaryApi.GetSummaries).Methods("GET")
//	router.HandleFunc("/summaries/{summary_id}", summaryApi.GetSummary).Methods("GET")
//	router.HandleFunc("/summaries/{summary_id}", summaryApi.ChangeSummary).Methods("PUT")
//	router.HandleFunc("/summaries/{summary_id}", summaryApi.DeleteSummary).Methods("DELETE")
//
//	//Миша
//	//apiRouter.HandleFunc("/user/{id}", HandleSetPrivateInfo).Methods("Put")
//	//apiRouter.HandleFunc("/user/{id}", HandleGetPrivateInfo).Methods("Get")
//	//apiRouter.HandleFunc("/avatar/{id}", HandleSetAvatar).Methods("Put")
//	//apiRouter.HandleFunc("/avatar/{id}", HandleGetAvatar).Methods("Get")
//
//	//Huvalk
//	//apiRouter.HandleFunc("/resume", HadleCreateResume).Methods("Post")
//	//apiRouter.HandleFunc("/resume/{id}", HandleChangeResume).Methods("Put")
//	//apiRouter.HandleFunc("/resume/{id}", HandleGetResume).Methods("Get")
//	//apiRouter.HandleFunc("/resume/{id}", HandleRemoveResume).Methods("Delete")
//
//	//Huvalk
//	//apiRouter.HandleFunc("/vacancy", HandleCreateVacancy).Methods("Post")
//	//apiRouter.HandleFunc("/vacancy/{id}", HandleChangeVacancy).Methods("Put")
//	//apiRouter.HandleFunc("/vacancy/{id}", HandleGetVacancy).Methods("Get")
//	//apiRouter.HandleFunc("/vacancy/{id}", HandleRemoveVacancy).Methods("Delete")
//
//	//Сережа М
//	//apiRouter.HandleFunc("/resume/short/{id}", HandleGetShortResume).Methods("Get")
//	//apiRouter.HandleFunc("/vacancy/short/{id}", HandleGetShortVacancy).Methods("Get")
//	//apiRouter.HandleFunc("/vacancys/{from}/{to}", HandleGetShortVacancy).Methods("Get")
//	log.Println("** Trying to start Service Started on Port 8000 **")
//
//	err := httpscerts.Check("cert.pem", "key.pem")
//	if err != nil {
//		err = httpscerts.Generate("cert.pem", "key.pem", "91.210.170.6:8000")
//		if err != nil {
//			log.Fatal("Ошибка: Не можем сгенерировать https сертификат.")
//		}
//	}
//	log.Println("** Service Started on Port 8000 **")
//	http.Handle("/", router)
//	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", nil)
//}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func echoFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /users/logout")

	//Cors.PrivateApi(&w, r)

	params := mux.Vars(r)
	message := params["message"]
	fmt.Fprintf(w, "Hello %s!", message)
}


func contentTipeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func StartRouter() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()//.StrictSlash(true)

	router.Use(contentTipeMiddleware)
	router.HandleFunc("/echo/{message}", echoFunc)
	router.Methods("OPTIONS").HandlerFunc(Cors.Preflight)

	// users
	authApi := NewAuthHandler()

	router.HandleFunc("/users/login", authApi.Login).Methods("POST")
	router.HandleFunc("/users/logout", authApi.Logout).Methods("POST")
	router.HandleFunc("/users", authApi.Register).Methods("POST")

	router.HandleFunc("/user/{user_id}", authApi.GetUserPage).Methods("GET")
	router.HandleFunc("/users/{user_id}/avatar", authApi.SetAvatar).Methods("POST")
	router.HandleFunc("/user/{user_id}", authApi.ChangeUserInfo).Methods("POST")

	// vacancies
	vacancyApi := NewVacancyHandler()

	router.HandleFunc("/vacancies", vacancyApi.CreateVacancy).Methods("POST")
	router.HandleFunc("/vacancies", vacancyApi.GetVacancies).Methods("GET")
	router.HandleFunc("/vacancies/{vacancy_id}", vacancyApi.GetVacancy).Methods("GET")
	router.HandleFunc("/vacancies/{vacancy_id}", vacancyApi.ChangeVacancy).Methods("PUT")
	router.HandleFunc("/vacancies/{vacancy_id}", vacancyApi.DeleteVacancy).Methods("DELETE")

	// summaries
	summaryApi := NewSummaryHandler()

	router.HandleFunc("/summaries", summaryApi.CreateSummary).Methods("POST")
	router.HandleFunc("/summaries", summaryApi.GetSummaries).Methods("GET")
	router.HandleFunc("/summaries/{summary_id}", summaryApi.GetSummary).Methods("GET")
	router.HandleFunc("/summaries/{summary_id}", summaryApi.ChangeSummary).Methods("PUT")
	router.HandleFunc("/summaries/{summary_id}", summaryApi.DeleteSummary).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Server started")
	http.ListenAndServe(":8001", router)
}
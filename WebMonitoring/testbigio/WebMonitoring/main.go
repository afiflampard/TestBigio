package main

import (
	"fmt"
	"net/http"
	"onboarding/config"
	"onboarding/controller"
	"onboarding/middleware"
	"onboarding/migrate"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	db := config.Init()

	migrate.Migrate(db)
	controller.InitiateDb(db)

	router := mux.NewRouter()

	router.HandleFunc("/", index)

	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println(port)
	subRouter := router.PathPrefix("/user").Subrouter()
	subProtectedRouter := router.PathPrefix("/user").Subrouter()

	subRouter.HandleFunc("/v1/test", HelloWorld).Methods("GET")

	subRouter.HandleFunc("/v1/login", controller.Authenticate).Methods("POST")

	subRouter.HandleFunc("/users", controller.GetUsers).Methods("GET")
	subProtectedRouter.Use(middleware.JwtVerifyToken)
	subProtectedRouter.HandleFunc("/v1/signup", controller.CreateAccount).Methods("POST")
	subProtectedRouter.HandleFunc("/v1/isiraport", controller.CreateRaport).Methods("POST")
	subProtectedRouter.HandleFunc("/v1/getraport", controller.GetRaport).Methods("GET")
	subProtectedRouter.HandleFunc("/v1/user/{id}", controller.GetUserById).Methods("GET")
	subProtectedRouter.HandleFunc("/v1/user/{id}", controller.UpdateUsers).Methods("PUT")
	subProtectedRouter.HandleFunc("/v1/user/photo/{id}", controller.UpdatePhoto).Methods("PUT")
	subProtectedRouter.HandleFunc("/v1/user/{id}", controller.DeleteUser).Methods("DELETE")

	//router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//log.Fatal(http.ListenAndServe(":"+port, router))
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}

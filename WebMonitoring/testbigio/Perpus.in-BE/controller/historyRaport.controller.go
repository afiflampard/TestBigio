package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onboarding/helpers"
	"onboarding/models"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateRaport = func(w http.ResponseWriter, r *http.Request) {
	newDataRaport := &models.Raport{}
	params := mux.Vars(r)
	idGuru := r.Header.Get("user_id")
	i, err := strconv.Atoi(idGuru)
	if err != nil {
		fmt.Println(err)
	}
	idSiswa, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewDecoder(r.Body).Decode(newDataRaport)
	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, "Invalid Request")
	}
	dataRaport := &models.Raport{
		IDSiswa: uint(idSiswa),
		Nama:    newDataRaport.Nama,
		Mapel:   newDataRaport.Mapel,
		Nilai:   newDataRaport.Nilai,
	}
	_, isValid := dataRaport.IsiRaport(GetDb(), idSiswa, i)
	if isValid {
		helpers.ResponseWithJson(w, 200, map[string]string{
			"success": "data sudah masuk db",
		})
	}
}

var GetRaport = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idSiswa, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	dataRaport := &models.Raport{}
	isValid, resp := dataRaport.GetRaport(GetDb(), idSiswa)
	if isValid {
		helpers.ResponseWithJson(w, 200, resp)
		return
	}
	helpers.MessageResponses(false, http.StatusBadRequest, "Data tidak ada")

}

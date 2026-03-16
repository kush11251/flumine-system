package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "flumine-system/src/models"
    "flumine-system/src/utils"
)

func metricsHandler(writer http.ResponseWriter, request *http.Request) {
    data := &models.Data{}
    err := json.NewDecoder(request.Body).Decode(data)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
n    }
    err = utils.ValidateData(data)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
    }
    writer.WriteHeader(http.StatusOK)
    fmt.Fprint(writer, "Metrics endpoint reached successfully.")
}
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/service"
	"log"
	"net/http"
	"strconv"
)

type CustomerHandler struct{
	custService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{custService: customerService}
}

func (s *CustomerHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *CustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	customer, err := s.custService.GetCustomer(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Get Customer error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, customer)
}

func (s *CustomerHandler) Post(w http.ResponseWriter, r *http.Request) {

	var cust = &model.Customer{}
	err := json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.CreateCustomer(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Create customer error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		s.responseBuilder(w, errMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, "customer created")
}

func (s *CustomerHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}
	var cust = &model.Customer{}
	err = json.NewDecoder(r.Body).Decode(cust)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.UpdateCustomer(custID, cust)
	if err != nil {
		errMsg := fmt.Sprintf("Update customer error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	s.responseBuilder(w, "customer updated")
}

func (s *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custID, err := strconv.Atoi(vars["id"])
	if err != nil {
		errMsg := fmt.Sprintf("Response builder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	err = s.custService.DeleteCustomer(custID)
	if err != nil {
		errMsg := fmt.Sprintf("Delete customer error : %v", err)

		w.WriteHeader(http.StatusNotFound)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, "customer deleted")
}

func (s *CustomerHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}

func (s *CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request){
	limit, err := strconv.Atoi(r.FormValue("limit"))
	page, err := strconv.Atoi(r.FormValue("page"))
	offset := (page - 1) * limit

	customers, err := s.custService.GetAllCustomer(offset, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Get Customer error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	s.responseBuilder(w, customers)
}

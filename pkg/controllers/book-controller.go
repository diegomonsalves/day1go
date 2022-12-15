package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/ahkil/go-bookstore/pkg/utils"
	"github.com/ahkil/go-bookstore/pkg/models"
	
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	NewBooks:=models.GetAllBooks()
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r);
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	//pendiente en el 1:53:50
}

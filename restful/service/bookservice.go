package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gpmgo/gopm/modules/log"
	"learn-go/restful/entity"
	"net/http"
)

func GetBookById(w http.ResponseWriter, r *http.Request){
	book := entity.Book{"2019-08-05-01","xieshang","math","1995","beijing"}

	params := mux.Vars(r)
	for _,item := range params{
		fmt.Printf("item %s",item)
	}

	resp,err := json.Marshal(book)
	if err != nil {
		log.Error("GetBookById error %s ",err)
	}
	w.Write(resp)
}
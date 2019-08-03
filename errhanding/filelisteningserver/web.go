package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"learn-go/errhanding/filelisteningserver/filelistening"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func errWrapper(
		handler appHandler) func(
			writer http.ResponseWriter, r *http.Request){
		return func(writer http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				log.Error("Panic error: %s",err)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}()

			err := handler(writer,r)
			if userErr, ok := err.(userError); ok{
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}

			if err != nil{
				log.Error("Error handling request: %s",err.Error())
				code := http.StatusOK
				switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError
				}
				http.Error(
					writer,
					http.StatusText(code),
					code)
			}
		}
}

func main(){
	http.HandleFunc("/",errWrapper(filelistening.HandleListFile))
	err := http.ListenAndServe(":8888",nil)
	if err != nil{
		panic(err)
	}
}
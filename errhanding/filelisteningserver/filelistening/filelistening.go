package filelistening

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefixPath  = "/list/"

type userError string

func (err userError) Message() string{
	return string(err)
}

func (err userError) Error() string{
	return err.Message()
}


func HandleListFile(writer http.ResponseWriter, request *http.Request) error{

	if strings.Index(request.URL.Path,prefixPath) != 0 {
		return userError("Path must start with "+prefixPath)
	}
	path := request.URL.Path[len(prefixPath):]
	file,err := os.Open(path)
	if err != nil{
		return err
	}

	all,err := ioutil.ReadAll(file)
	if err != nil{
		return err
	}
	writer.Write(all)
	return nil
}

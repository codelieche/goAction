package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	// 判断url是否含有/list/
	if strings.Index(request.URL.Path, prefix) != 0 {
		//return errors.New("path must start with" + prefix)
		return userError("Path must start with " + prefix)
	}

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer, err.Error(),
		//	http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	// 正常执行完毕，无错误就返回nil
	return nil
}

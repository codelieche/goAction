package objects

import (
	"fmt"
	"goAction/books/src/objects-storage/settings"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/**
上传对象
*/

// 上传对象的处理函数
// Method： PUT
// PATH： /objects/:name 【这里不考虑子目录】
func handlerPutObject(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.EscapedPath(), "/")[2]
	//log.Println(filename)
	filePath := fmt.Sprintf("%s/objects/%s", settings.STORATE_ROOT, filename)
	if file, err := os.Create(filePath); err != nil {
		log.Println("创建文件失败！", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		defer file.Close()
		// 请先自己创建好需要储存文件的目录
		// 复制文件
		io.Copy(file, r.Body)
		w.Write([]byte(filePath))
		return
	}

}

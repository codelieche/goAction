package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/codelieche/goAction/tutorial/books/objects-storage/settings"
)

/**
获取object
*/

// 获取对象处理函数
// Method：GET
// PATH：/objects/:name
// 示例：/objects/abc.jpg
func handlerGetObject(w http.ResponseWriter, r *http.Request) {
	// 根据URL.Path得到文件名
	filename := strings.Split(r.URL.Path, "/")[2]
	filePath := fmt.Sprintf("%s/objects/%s", settings.STORATE_ROOT, filename)

	// 打开文件
	if file, err := os.Open(filePath); err != nil {
		log.Println("获取文件失败：", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		defer file.Close()
		io.Copy(w, file)
		return
	}
}

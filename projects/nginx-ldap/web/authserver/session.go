package authserver

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"time"

	"goAction/projects/nginx-ldap/web/settings"

	"github.com/gorilla/sessions"
)

// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
// 1. 准备32位字节
var store = sessions.NewCookieStore([]byte(settings.Config.WebSecretKey))

// 生成唯一的session id字符串
// 如果量大了，需要重新调整
func GenerateSessionId() string {
	// 1. 准备32位字节
	b := make([]byte, 32)
	hash := md5.New()

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Println(err)
		now := time.Now().Format(time.RFC3339Nano)
		hash.Write([]byte(now))
	} else {
		//log.Print(n)
		hash.Write(b)
	}

	//2. md5
	return hex.EncodeToString(hash.Sum(nil))
}

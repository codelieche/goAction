package main

import (
	"fmt"
	"os"
)

func readAndPrintFile(filename string) {
	// 读取文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("出现错误：%s\n", err.Error())
		return
	}
	defer file.Close()
	// 读取文件内容
	data := make([]byte, 1024)

	count, _ := file.Read(data)
	fmt.Println("Read %d bytes.\n%q", count, data[:count])

	// 输出读取的内容
	if count > 0 {
		fmt.Println(string(data[:count-1]))
	}
}

func createFile(filename string) {

	// 如果想写入内容，得加入os.O_RDWR
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	count, _ := file.Write([]byte("我的测试文件\n"))

	fmt.Println("写入数据的长度:", count)

	fileinfo, _ := file.Stat()
	file.WriteString("写入点字符\n")

	fmt.Println("\nfileinfo:\n\tfile name: ", fileinfo.Name(),
		"\n\tis dir:", fileinfo.IsDir(),
		"\n\tFileMode:", fileinfo.Mode())

	fmt.Println("create file done")

}

func main() {
	fmt.Println("=== main ===")
	fmt.Print("os.Getwd():")
	fmt.Println(os.Getwd())
	fmt.Print("os.Hostname():")
	fmt.Println(os.Hostname())

	//readAndPrintFile("tmp/log.txt")
	//readAndPrintFile("tmp/file01.txt")

	fmt.Println("\n=== createFile() ===")
	createFile("./tmp/file01.txt")

	fmt.Println("\n=== Done ===")

}

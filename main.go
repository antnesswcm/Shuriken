package Shuriken

import (
	"fmt"
)

func main() {
	test1()
}
func test1() {
	var testPath string = `D:\\测试文件夹\`
	dirs, files, err := GetAllFiles(testPath, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("目录:")
	for i, v := range dirs {
		fmt.Printf("%d\t%s\n", i, v)
	}
	fmt.Println("文件:")
	for i, v := range files {
		fmt.Printf("%d\t%s\n", i, v)
	}
}

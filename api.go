package Shuriken

// GetAllFiles 递归返回路径下的文件与目录
// 参数 path string 指定遍历的目录
// 参数 bottomUp bool 是否自底向上
// 返回值 dirs []string 返回得到的目录切片
// 返回值 files []string 返回得到的文件切片
func GetAllFiles(path string, bottomUp bool) (dirs []string, files []string, err error) {
	return travelDirectory(path, true, bottomUp)
}

// GetFiles 返回路径下的文件与目录，不递归
// 参数 path string 指定遍历的目录
// 返回值 dirs []string 返回得到的目录切片
// 返回值 files []string 返回得到的文件切片
func GetFiles(path string) (dirs []string, files []string, err error) {
	return travelDirectory(path, false, false)
}

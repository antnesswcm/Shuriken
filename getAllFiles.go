package Shuriken

import (
	"io/fs"
	"path/filepath"
	"sort"
)

// travelDirectory 返回路径下的文件与目录
// 参数 path string 指定遍历的目录
// 参数 recursive bool 是否遍历子目录
// 参数 bottomUp bool 是否自底向上
// 返回值 dirs []string 返回得到的目录切片
// 返回值 files []string 返回得到的文件切片
func travelDirectory(path string, recursive bool, bottomUp bool) (dirs []string, files []string, err error) {

	// 使用一个切片来存储所有文件和文件夹的信息。
	var entries []struct {
		path string
		info fs.DirEntry
	}

	// 使用 filepath.WalkDir 函数遍历指定路径下的所有文件和文件夹。
	err = filepath.WalkDir(path, func(filePath string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filePath != path {
			// 将文件或文件夹的信息添加到切片中。
			entries = append(entries, struct {
				path string
				info fs.DirEntry
			}{filePath, dirEntry})
			if !recursive && dirEntry.IsDir() {
				// 如果递归标志为 false，则跳过该目录。
				return fs.SkipDir
			}
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	if bottomUp {
		// 对切片进行反向排序，使得最深层次的文件或文件夹排在前面。
		sort.Slice(entries, func(i, j int) bool {
			return len(entries[i].path) > len(entries[j].path)
		})
	}

	// 遍历切片，将文件和文件夹分别添加到对应的列表中。
	for _, entry := range entries {
		if entry.info.IsDir() {
			dirs = append(dirs, entry.path)
		} else {
			files = append(files, entry.path)
		}
	}

	return dirs, files, nil
}

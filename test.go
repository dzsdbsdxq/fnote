//package main
//
//import (
//	"fmt"
//	"os"
//	"path/filepath"
//)
//
//// 分页遍历文件夹
//func ListDirPaginated(dir string, page, pageSize int) ([]string, error) {
//	if page <= 0 {
//		return nil, fmt.Errorf("invalid page number: %d", page)
//	}
//	if pageSize <= 0 {
//		return nil, fmt.Errorf("invalid page size: %d", pageSize)
//	}
//
//	// 计算起始索引和结束索引
//	start := (page - 1) * pageSize
//	end := start + pageSize
//
//	var files []string
//	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		if !info.IsDir() {
//			files = append(files, path)
//		}
//		return nil
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	// 分页处理
//	if start >= len(files) {
//		return []string{}, nil
//	}
//	if end > len(files) {
//		end = len(files)
//	}
//	return files[start:end], nil
//}
//
//func main() {
//	dir := "./"     // 替换为你需要遍历的目录
//	page := 1       // 页码
//	pageSize := 100 // 每页大小
//
//	files, err := ListDirPaginated(dir, page, pageSize)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	for _, file := range files {
//		fmt.Println(file)
//	}
//}

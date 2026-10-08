package commit

import (
	"fmt"
	"strconv"
	"strings"
)

// 将用户选择的 number Path 加入
func getFilePath(nums string) ([]string, error) {
	n := strings.Fields(nums)
	stat, err := getModifier()
	if err != nil {
		return nil, err
	}
	var filePath []string

	for _, v := range n {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		if i < 0 || i >= len(stat) {
			return nil, fmt.Errorf("编号 %d 超出范围（0-%d）", i, len(stat)-1)
		}
		filePath = append(filePath, stat[i].FilePath)
	}

	return filePath, nil
}

func GetFilePath(nums string) ([]string, error) {
	filePath, err := getFilePath(nums)
	if err != nil {
		return nil, err
	}

	return filePath, nil
}

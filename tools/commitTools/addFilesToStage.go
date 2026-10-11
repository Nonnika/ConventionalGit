package commitTools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nonnika/ConventionalGit/model/colorsModel"
)

var rangeRe = regexp.MustCompile(`^(-?\d+)-(-?\d+)$`)

// 将用户选择的 number Path 加入
func getFilePathByNumber(nums string) ([]string, error) {
	n := strings.Fields(nums)
	if len(n) == 0 {
		return nil, fmt.Errorf("文件编号不能为空")
	}
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
		filePath = append(filePath, stat[i].FileNewPath)
	}

	return filePath, nil
}

// 范围选择函数
func getFilePathByRange(start, end int) ([]string, error) {
	stat, err := getModifier()
	if err != nil {
		return nil, err
	}

	if start < 0 || start >= len(stat) {
		return nil, fmt.Errorf("%s左范围不合法%s", colorsModel.RED, colorsModel.END)
	}
	if end < start || end >= len(stat) {
		return nil, fmt.Errorf("%s右范围不合法%s", colorsModel.RED, colorsModel.END)
	}

	var filePath []string

	for _, v := range stat[start : end+1] {
		filePath = append(filePath, v.FileNewPath)
	}

	return filePath, nil
}

func GetFilePath(nums string) ([]string, error) {
	nums = strings.TrimSpace(nums)
	// 如果有范围字符进入范围函数
	n := rangeRe.FindStringSubmatch(nums)
	if n != nil {
		n1, err := strconv.Atoi(n[1])
		if err != nil {
			return nil, err
		}

		n2, err := strconv.Atoi(n[2])
		if err != nil {
			return nil, err
		}
		return getFilePathByRange(n1, n2)
	}

	return getFilePathByNumber(nums)
}

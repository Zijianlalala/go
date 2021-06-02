// Package keyboard reads user input from the keyboard
package keyboard

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

// 从键盘输入一个数字，返回浮点数格式以及错误状态
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil{
		return 0, err
	}
	return grade, nil
}
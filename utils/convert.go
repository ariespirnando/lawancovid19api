package utils

import (
	"strings"
	"strconv"
)

func ConvertToInt(str string) int64{
	str = strings.ReplaceAll(str, ",", "")
	x, _ := strconv.Atoi(str)
	m := int64(x)
	return m
}
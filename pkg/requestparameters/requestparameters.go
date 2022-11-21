package requestparameters

import (
	"errors"
	"regexp"
	"strconv"
)

func GetRequestParameters(str string) (string, float32, float32, error) {
	//регулярное выражение чтобы найти группы ip адрес сервера первое и второе число
	reg := regexp.MustCompile(`(\d+.\d+.\d+.\d+:\d+)\s(\d+.*\d*)\s(\d+.*\d*)\s*\r*\n$`)
	strGroup := reg.FindSubmatch([]byte(str))
	if strGroup == nil {
		return "", 0, 0, errors.New("failed to read the parameters, need <ip-address>:<port> <float number 1> <float number 2>")
	}

	//конвертируем числа в float32
	value1, err := strconv.ParseFloat(string(strGroup[2]), 32)
	if err != nil {
		return "", 0, 0, errors.New("failed to read first number")
	}

	value2, err := strconv.ParseFloat(string(strGroup[3]), 32)
	if err != nil {
		return "", 0, 0, errors.New("failed to read second number")
	}

	return string(strGroup[1]), float32(value1), float32(value2), nil
}

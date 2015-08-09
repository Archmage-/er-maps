package util

import (
	"strconv"
)

func ConvertToInts(items []string) ([]int, error) {
	intItems := make([]int, len(items))
	for i, str := range items {
		res, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intItems[i] = res
	}
	return intItems, nil
}

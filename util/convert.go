package util

import (
	"strconv"
	"strings"
)

// ArrayStrToUint64 converts a slice of string to a slice of uint64.
//
// strconv.ParseUint is called for each element. If any error occurs, this
// function panics with the error.
func ArrayStrToUint64(strs []string) []uint64 {
	nums := make([]uint64, len(strs))
	for i, str := range strs {
		nums[i] = Must(strconv.ParseUint(strings.TrimSpace(str), 10, 64))
	}
	return nums
}

// ArrayStrToInt64 converts a slice of string to a slice of int64.
//
// strconv.ParseInt is called for each element. If any error occurs, this
// function panics with the error.
func ArrayStrToInt64(strs []string) []int64 {
	nums := make([]int64, len(strs))
	for i, str := range strs {
		nums[i] = Must(strconv.ParseInt(strings.TrimSpace(str), 10, 64))
	}
	return nums
}

// ArrayStrToFloat64 converts a slice of string to a slice of float64.
//
// strconv.ParseFloat is called for each element. If any error occurs, this
// function panics with the error.
func ArrayStrToFloat64(strs []string) []float64 {
	nums := make([]float64, len(strs))
	for i, str := range strs {
		nums[i] = Must(strconv.ParseFloat(strings.TrimSpace(str), 64))
	}
	return nums
}

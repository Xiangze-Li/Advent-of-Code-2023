package util_test

import (
	"advent2023/util"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestAssociative(t *testing.T) {
	tests := []struct {
		name string
		exec func() any
		want any
	}{
		{
			name: "Float to rounded int",
			exec: func() any {
				return util.Associative([]float64{1.1, 2.2, 3.3}, func(f float64) (float64, int64) {
					return f, int64(math.Round(f))
				})
			},
			want: map[float64]int64{1.1: 1, 2.2: 2, 3.3: 3},
		},
		{
			name: "Int to string",
			exec: func() any {
				return util.Associative([]int{1, 2, 3}, func(i int) (int, string) {
					return i, strconv.Itoa(i)
				})
			},
			want: map[int]string{1: "1", 2: "2", 3: "3"},
		},
		{
			name: "Empty slice",
			exec: func() any {
				return util.Associative([]int{}, func(i int) (int, string) {
					return i, strconv.Itoa(i)
				})
			},
			want: map[int]string{},
		},
		{
			name: "Duplicate keys",
			exec: func() any {
				count := 0
				return util.Associative([]int{1, 2, 3, 2}, func(i int) (int, int) {
					count++
					return i, count
				})
			},
			want: map[int]int{1: 1, 2: 4, 3: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Util.Associative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToVis(t *testing.T) {
	tests := []struct {
		name string
		exec func() any
		want any
	}{
		{
			name: "Int array",
			exec: func() any {
				return util.ToVis([][2]int{
					{1, 2},
					{3, 4},
					{1, 2},
					{5, 6},
				})
			},
			want: map[[2]int]bool{
				{1, 2}: true,
				{3, 4}: true,
				{5, 6}: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToVis() = %v, want %v", got, tt.want)
			}
		})
	}
}

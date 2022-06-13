package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{2, 1},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, topKFrequent(tt.args.nums, tt.args.k), "topKFrequent(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}

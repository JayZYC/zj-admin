package util

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{[]string{"172.19.0.186:1883", "172.19.0.187:1883", "172.19.0.188:1883"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shuffle(tt.args.arr)
			fmt.Println(got)
		})
	}
}

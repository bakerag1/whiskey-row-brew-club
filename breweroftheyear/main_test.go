package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWinners(t *testing.T) {
	type args struct {
		file string
	}
	arg := args{file: "./test-comps.yaml"}
	tests := []struct {
		name string
		args args
		want []Standing
	}{
		{
			name: "check scores",
			args: arg,
			want: []Standing{
				{Place: "1st", Name: "Susan", Points: 12},
				{Place: "2nd", Name: "Andrew", Points: 11},
				{Place: "3rd", Name: "andrew", Points: 5},
				{Place: "4th", Name: "Wes", Points: 4},
				{Place: "5th", Name: "Tom", Points: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				got := GetWinners(tt.args.file)
				assert.Equal(t, tt.want, got)

			})
	}
}

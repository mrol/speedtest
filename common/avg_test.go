package common

import "testing"

func Test_avgFloat64(t *testing.T) {
	type args struct {
		in []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1 value",
			args: args{in: []float64{12}},
			want: 12,
		},
		{
			name: "0 values",
			args: args{in: []float64{0}},
			want: 0,
		},
		{
			name: "a lot of values",
			args: args{in: []float64{1, 3, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AvgFloat64(tt.args.in); got != tt.want {
				t.Errorf("AvgFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

package convert

import "testing"

func TestConvertColumnToNumber(t *testing.T) {
	type args struct {
		column string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{column: "A"},
			want: 0,
		},
		{
			name: "B",
			args: args{column: "B"},
			want: 1,
		},
		{
			name: "AA",
			args: args{column: "AA"},
			want: 26,
		},
		{
			name: "AB",
			args: args{column: "AB"},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertColumnToNumber(tt.args.column); got != tt.want {
				t.Errorf("ConvertColumnToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

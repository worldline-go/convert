package convert

import "testing"

func TestRows_Include(t *testing.T) {
	type arg struct {
		v    int
		want bool
	}
	tests := []struct {
		name    string
		rows    []string
		args    []arg
		wantErr bool
	}{
		{
			name: "single range",
			rows: []string{"30"},
			args: []arg{
				{v: 30, want: true},
				{v: 31, want: false},
			},
			wantErr: false,
		},
		{
			name: "multiple ranges",
			rows: []string{"29-55", "1-10,20-30"},
			args: []arg{
				{v: 30, want: true},
				{v: 31, want: true},
				{v: 56, want: false},
				{v: 10, want: true},
				{v: 11, want: false},
				{v: 20, want: true},
				{v: 31, want: true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewRows(tt.rows)
			if err != nil {
				if tt.wantErr {
					return
				}

				t.Errorf("NewRows() error = %v", err)
			}

			for _, a := range tt.args {
				if got := r.IsInclude(a.v); got != a.want {
					t.Errorf("Rows.IsInclude(%v) = %v, want %v", a.v, got, a.want)
				}
			}
		})
	}
}

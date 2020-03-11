package luhnalg

import "testing"

func TestLuhnAlgorithm_Validate(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		la   *luhnAlgorithm
		args args
		want bool
	}{
		{
			name: "Comply to Luhn Algorithm checksum: True",
			la:   &luhnAlgorithm{},
			args: args{"4111111111111111"},
			want: true,
		},
		{
			name: "Comply to Luhn Algorithm checksum: False",
			la:   &luhnAlgorithm{},
			args: args{"4111111111111"},
			want: false,
		},
		{
			name: "Comply to Luhn Algorithm checksum (with spaces): True",
			la:   &luhnAlgorithm{},
			args: args{"5105 1051 0510 5100"},
			want: true,
		},
		{
			name: "Comply to Luhn Algorithm checksum (with spaces): False",
			la:   &luhnAlgorithm{},
			args: args{"5105 1051 0510 5106"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.la.Validate(tt.args.data); got != tt.want {
				t.Errorf("luhnAlgorithm.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLuhnAlgorithm(b *testing.B) {
	luhnManager := New()
	for i := 0; i < b.N; i++ {
		luhnManager.Validate("5105 1051 0510 5100")
	}
}

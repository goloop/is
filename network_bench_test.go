package is

import "testing"

var ipResult bool

func BenchmarkIPValidation(b *testing.B) {
	for _, addr := range []string{"127.0.0.1", "2001:db8::1", "invalid"} {
		b.Run(addr, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ipResult = IP(addr)
			}
		})
	}
}

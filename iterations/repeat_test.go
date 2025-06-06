package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expect %q got %q", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

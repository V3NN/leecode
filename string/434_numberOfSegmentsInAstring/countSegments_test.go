package __numberOfSegmentsInAstring

import "testing"

func TestCountSegments(t *testing.T) {
	t.Log(countSegments("   "))
	t.Log(countSegments("Hello, my name is John"))
}

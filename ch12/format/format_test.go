package format

import "testing"

func TestAny(t *testing.T) {
	tests := []struct {
		input interface{}
		want  string
	}{
		{"", ""},
		{17, "17"},
		{uint64(13), "13"},
		{false, "false"},
	}
	for _, test := range tests {
		if got := Any(test.input); got != test.want {
			t.Errorf("Any(%q) = %v", test.input, got)
		}
	}
}

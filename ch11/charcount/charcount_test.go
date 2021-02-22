package charcount

import (
	"bytes"
	"testing"
)

func TestEmptyCharcount(t *testing.T) {
	buf := bytes.NewBufferString("")
	got, err := Charcount(buf)
	if err != nil {
		t.Fatalf("Charcount(%q): %v", "", err)
	}
	if len(*got.counts) > 0 {
		t.Errorf("Charcount(%q) len(counts) = %v, want %v", "", len(*got.counts), 0)
	}
}

func TestAbc(t *testing.T) {
	buf := bytes.NewBufferString("Abc")
	got, err := Charcount(buf)
	if err != nil {
		t.Fatalf("Charcount(%q): %v", "Abc", err)
	}
	if len(*got.counts) != 3 {
		t.Errorf("Charcount(%q) len(counts) = %v, want %v", "Abc", len(*got.counts), 0)
	}
	if (*got.counts)[rune('A')] != 1 {
		t.Errorf("Charcount(%q) counts['A'] = %v, want %v", "Abc", (*got.counts)[rune('A')], 1)
	}
	if (*got.counts)[rune('b')] != 1 {
		t.Errorf("Charcount(%q) counts['b'] = %v, want %v", "Abc", (*got.counts)[rune('b')], 1)
	}
	if (*got.counts)[rune('c')] != 1 {
		t.Errorf("Charcount(%q) counts['c'] = %v, want %v", "Abc", (*got.counts)[rune('c')], 1)
	}
	if (*got).utflen[0] != 0 {
		t.Errorf("Charcount(%q) utflen[0] = %v, want %v", "Abc", (*got).utflen[0], 0)
	}
	if (*got).utflen[1] != 3 {
		t.Errorf("Charcount(%q) utflen[1] = %v, want %v", "Abc", (*got).utflen[1], 3)
	}
	if (*got).utflen[2] != 0 {
		t.Errorf("Charcount(%q) utflen[2] = %v, want %v", "Abc", (*got).utflen[2], 0)
	}
	if (*got).utflen[3] != 0 {
		t.Errorf("Charcount(%q) utflen[3] = %v, want %v", "Abc", (*got).utflen[3], 0)
	}
	if (*got).utflen[4] != 0 {
		t.Errorf("Charcount(%q) utflen[4] = %v, want %v", "Abc", (*got).utflen[4], 0)
	}
	if (*got).invalid != 0 {
		t.Errorf("Charcount(%q) invalid = %v, want 0", "Abc", (*got).utflen)
	}
}

package bot

import (
	"os"
	"testing"
)

func SetupTesting(t *testing.T) {
	os.Setenv("RAIN_TESTING", "TEST")
}

func TestNewBot(t *testing.T) {
	SetupTesting(t)

	b, err := NewBot("Niladon", "string")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	ds := b.ds.(*DST)
	if ds.id != 0 {
		t.Fatalf("DST set incorrect ID, got %d, expecting 0\n", ds.id)
	}

	if b.name != "Niladon" {
		t.Errorf("NewBot failed to set the correct name! Got %s instead", b.name)
	}
}

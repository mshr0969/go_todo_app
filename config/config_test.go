package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("want %d, but %d", wantPort, got.Port)
	}
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}

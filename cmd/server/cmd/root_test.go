package cmd

import "testing"

func TestFlags(t *testing.T) {
	valid := [][]string{
		{"--port", "0"},
		{"--port", "8080"},
		{"--port", "65535"},
	}

	invalid := [][]string{
		{"--port", "-1"},
		{"--port", "65536"},
	}

	for _, flag := range invalid {
		err := rootCmd.ParseFlags(flag)

		if err == nil {
			t.Fatalf("Expected error, got nil on flag '%s' with value '%s'", flag[0], flag[1])
		}
	}

	for _, flag := range valid {
		err := rootCmd.ParseFlags(flag)

		if err != nil {
			t.Fatalf("Expected nil, got error on flag '%s' with value '%s'", flag[0], flag[1])
		}
	}
}

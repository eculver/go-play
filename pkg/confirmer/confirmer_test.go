package confirmer_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/eculver/go-play/pkg/confirmer"
)

func TestDefaultConfirmer(t *testing.T) {
	tests := []struct {
		scenario string
		msg      string
		in       []string
		expected bool
	}{
		{
			scenario: "Default Confirmer denies for negative values",
			in:       []string{"n"},
			expected: false,
		},
		{
			scenario: "Default Confirmer accepts for affirmative values",
			in:       []string{"Y"},
			expected: true,
		},
		{
			scenario: "Default Confirmer denies after 3 retries",
			in:       []string{"f", "o", "o"},
			expected: false,
		},
		{
			scenario: "Default Confirmer accepts after retries",
			in:       []string{"", "o", "Y"},
			expected: true,
		},
		{
			scenario: "Default Confirmer denies when it can't read input",
			in:       []string{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			buffer := bytes.Buffer{}
			for _, i := range test.in {
				buffer.Write([]byte(fmt.Sprintf("%s\n", i)))
			}

			confirmed := confirmer.Confirm("confirmation", &buffer)

			if confirmed != test.expected {
				t.Errorf("expected %t but got %t", test.expected, confirmed)
			}
		})
	}
}

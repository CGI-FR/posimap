package config_test

import (
	"testing"

	"github.com/cgi-fr/posimap/internal/appli/config"
)

//nolint:funlen
func TestPicture_Compile(t *testing.T) {
	t.Parallel()

	//nolint:exhaustruct
	tests := []struct {
		name     string
		pic      config.Picture
		expected string
		wantErr  bool
	}{
		{
			name:     "valid picture",
			pic:      "X(10)",
			expected: "X(10)",
		},
		{
			name:     "valid numeric picture",
			pic:      "9(5)V(2)",
			expected: "9(5)V(2)",
		},
		{
			name:     "signed numeric picture",
			pic:      "S9(2)V99",
			expected: "S9(2)V(2)",
		},
		{
			name:     "no decimals numeric picture",
			pic:      "S9",
			expected: "S9",
		},
		{
			name:     "alpha picture",
			pic:      "AAA",
			expected: "A(3)",
		},
		{
			name:     "numeric picture",
			pic:      "9V99",
			expected: "9V(2)",
		},
		{
			name:    "invalid picture type",
			pic:     "Z(10)",
			wantErr: true,
		},
		{
			name:    "negative length",
			pic:     "X(-5)",
			wantErr: true,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			picture, err := testcase.pic.Compile()
			if (err != nil) != testcase.wantErr {
				t.Errorf("Compile() error = %v, wantErr %v", err, testcase.wantErr)
			}

			if err == nil && picture.String() != testcase.expected {
				t.Errorf("Compile() = %v, want %v", picture.String(), testcase.expected)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestFindTagNode(t *testing.T) {
	// TODO
}

func TestFindLinks(t *testing.T) {
	// TODO
}

func TestNumberFormatter(t *testing.T) {
	testCases := []struct {
		name     string
		issues   []string
		expected string
	}{
		{
			name:     "Only single issue linked",
			issues:   []string{"g/foo/bar/issues/1"},
			expected: "1",
		},
		{
			name:     "Multiple issues linked",
			issues:   []string{"g/foo/bar/issues/1", "g/foo/bar/issues/2"},
			expected: "1 2",
		},
		{
			name:     "No linked issue",
			issues:   nil,
			expected: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			f := &NumberFormatter{}
			actual := f.Format(tt.issues)

			if actual != tt.expected {
				t.Errorf("Expected: %v, actual: %v", tt.expected, actual)
			}
		})
	}
}

func Test_URLFormatter(t *testing.T) {
	testCases := []struct {
		name     string
		issues   []string
		expected string
	}{
		{
			name:     "Only single issue linked",
			issues:   []string{"g/foo/bar/issues/1"},
			expected: "g/foo/bar/issues/1",
		},
		{
			name:     "Multiple issues linked",
			issues:   []string{"g/foo/bar/issues/1", "g/foo/bar/issues/2"},
			expected: "g/foo/bar/issues/1 g/foo/bar/issues/2",
		},
		{
			name:     "No linked issue",
			issues:   nil,
			expected: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			f := &URLFormatter{}
			actual := f.Format(tt.issues)

			if actual != tt.expected {
				t.Errorf("Expected: %v, actual: %v", tt.expected, actual)
			}
		})
	}
}

func Test_ExternalIssueRefFormatter(t *testing.T) {
	testCases := []struct {
		name     string
		issues   []string
		expected string
	}{
		{
			name:     "Only single issue linked",
			issues:   []string{"g/foo/bar/issues/1"},
			expected: "foo/bar#1",
		},
		{
			name:     "Multiple issues linked",
			issues:   []string{"g/foo/bar/issues/1", "g/foo/bar/issues/2"},
			expected: "foo/bar#1 foo/bar#2",
		},
		{
			name:     "No linked issue",
			issues:   nil,
			expected: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			f := &ExternalIssueRefFormatter{}
			actual := f.Format(tt.issues)

			if actual != tt.expected {
				t.Errorf("Expected: %v, actual: %v", tt.expected, actual)
			}
		})
	}
}

package version

import "testing"

func TestNewVersion(t *testing.T) {
	testCases := []struct {
		s  string
		ok bool
	}{
		{"1", true},
		{"1.2", true},
		{"0.3", true},
		{"00.1", true},
		{"0.", false},
		{"1.2.", false},
		{"1.2.3", true},
		{"..", false},
		{".", false},
		{"", false},
	}

	for _, tc := range testCases {
		_, err := NewVersion(tc.s)
		if err == nil && !tc.ok {
			t.Fatalf("Expected error, but success: %s", tc.s)
		}
		if err != nil && tc.ok {
			t.Fatalf("Unexpected error %v: %s", err, tc.s)
		}
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		s string
	}{
		{"1.2.3"},
	}

	for _, tc := range testCases {
		v, err := NewVersion(tc.s)
		if err != nil {
			t.Fatalf("Unexpected error %v: %s", err, tc.s)
		}
		if v.String() != tc.s {
			t.Fatalf("Not equals (%s!=%s)", v.String(), tc.s)
		}
	}
}

func TestCmp(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected int
	}{
		{"1", "1", 0},
		{"0", "0", 0},
		{"1", "2", -1},
		{"4", "3", 1},
		{"1.2.3", "1.2.3", 0},
		{"2.3.3", "2.3.4", -1},
		{"4.3", "4.1", 1},
		{"1.0", "1.0.0", 0},
		{"1.2.3", "1.2.3.4", -1},
		{"78.90.12", "78.90", 1},
	}

	for _, tc := range testCases {
		v1, err := NewVersion(tc.v1)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		v2, err := NewVersion(tc.v2)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		actual := v1.Cmp(v2)
		if tc.expected != actual {
			t.Errorf("Got %v, want %v: %s, %s", actual, tc.expected, tc.v1, tc.v2)
		}
	}
}

func TestEquals(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected bool
	}{
		{"1.2", "1.2", true},
		{"2.3.4", "3.4.5", false},
		{"4", "3", false},
	}

	for _, tc := range testCases {
		v1, err := NewVersion(tc.v1)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		v2, err := NewVersion(tc.v2)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		actual := v1.Equals(v2)
		if tc.expected != actual {
			t.Errorf("Got %v, want %v: %s, %s", actual, tc.expected, tc.v1, tc.v2)
		}
	}
}

func TestLessThan(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected bool
	}{
		{"1.2", "1.2", false},
		{"2.3.4", "3.4.5", true},
		{"4", "3", false},
	}

	for _, tc := range testCases {
		v1, err := NewVersion(tc.v1)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		v2, err := NewVersion(tc.v2)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		actual := v1.LessThan(v2)
		if tc.expected != actual {
			t.Errorf("Got %v, want %v: %s, %s", actual, tc.expected, tc.v1, tc.v2)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected bool
	}{
		{"1.2", "1.2", false},
		{"2.3.4", "3.4.5", false},
		{"4", "3", true},
	}

	for _, tc := range testCases {
		v1, err := NewVersion(tc.v1)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		v2, err := NewVersion(tc.v2)
		if err != nil {
			t.Fatalf("Unexepcted error %v", err)
		}

		actual := v1.GreaterThan(v2)
		if tc.expected != actual {
			t.Errorf("Got %v, want %v: %s, %s", actual, tc.expected, tc.v1, tc.v2)
		}
	}
}

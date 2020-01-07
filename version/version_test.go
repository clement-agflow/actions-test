package version

import "testing"

func TestNew(t *testing.T) {
	cases := []struct {
		rawString string
		success   bool
	}{
		{"lol", false},
		{"1.2", false},
		{"1.2.3", true},
		{"12.22.3", true},
	}

	for _, c := range cases {
		v, err := New(c.rawString)
		if (err == nil) != c.success {
			t.Fatalf("wrong result for: %q", c.rawString)
		}
		if err == nil {
			t.Log(v)
		}
	}

}

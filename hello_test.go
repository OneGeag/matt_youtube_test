package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			// for empty input
			result: "Hello, world!",
		},
		{
			// input has one name
			items:  []string{"Matt"},
			result: "Hello, Matt!",
		},
		{
			// input has several names
			items:  []string{"Matt", "John"},
			result: "Hello, Matt, John!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result,
				st.items, s)
		}
	}
}

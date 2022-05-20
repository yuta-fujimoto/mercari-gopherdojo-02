package download

import "testing"


func TestRun(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		url string
		} {
		{
			name: "normal",
			url: "http://example.com",
		},
		{
			name: "large",
			url: "https://releases.ubuntu.com/focal/ubuntu-20.04.4-live-server-amd64.iso",
		},
	}
	for _, td := range cases {
		td := td
		t.Run(td.name, func(t *testing.T) {
			t.Parallel()
			got := Run(td.url, 2)
			if got != nil {
				t.Fatal(got.Error())
			}
		})
	}
}

func TestError(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		url string
	} {
		{
			name: "invalid protocol",
			url: "invalid",
		},
		{
			name: "invalid url",
			url: "https://invalid",
		},
		{
			name: "does not support range request",
			url: "https://google.com",
		},
	}
	for _, td := range cases {
		td := td
		t.Run(td.name, func(t *testing.T) {
			t.Parallel()
			got := Run(td.url, 2)
			if got == nil {
				t.Fatal("no error")
			}
		})
	}
}



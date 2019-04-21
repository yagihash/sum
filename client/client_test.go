package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewClient(t *testing.T) {
	ts := makeServer(t)
	cases := []struct {
		name     string
		input    string
		expected Client
	}{
		{name: fmt.Sprintf("OK(%s)", ts.URL), input: ts.URL, expected: Client{
			URL:    ts.URL,
			Md5sum: "3858f62230ac3c915f300c664312c63f",
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, _ := NewClient(c.input)
			opt := cmpopts.IgnoreUnexported(*actual)
			if diff := cmp.Diff(*actual, c.expected, opt); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func TestFetch(t *testing.T) {
	ts := makeServer(t)
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: fmt.Sprintf("OK(%s)", ts.URL), input: ts.URL, expected: "foobar"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client, _ := NewClient(c.input)
			actual, _, err := client.Fetch()
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(actual, c.expected); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func TestMd5Sum(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "OK(foobar)", input: "foobar", expected: "3858f62230ac3c915f300c664312c63f"},
		{name: "OK(\\xe1\\x4f\\x2e\\xf8\\x65\\x88\\x23\\xf7)", input: "\xe1\x4f\x2e\xf8\x65\x88\x23\xf7", expected: "1ec011d378e50faa62dff81a3b9fd94f"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := makeServer(t)
			client, err := NewClient(ts.URL)
			if err != nil {
				t.Error(err)
			}
			actual := client.md5sum(c.input)

			if diff := cmp.Diff(actual, c.expected); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func makeServer(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("foobar")); err != nil {
				t.Fatal(err)
			}
		},
	))
}

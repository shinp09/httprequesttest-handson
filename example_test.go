package example_test

import (
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client  {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

const VaildToken = "Valid Token"
//
//func client(t *testing.T) *http.Client  {
//	t.Helper()
//}
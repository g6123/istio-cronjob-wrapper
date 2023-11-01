package pkg

import "net/http"

func checkRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

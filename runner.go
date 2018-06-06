package behaviours

import (
	"net/http"

	"github.com/mitchellh/go-testing-interface"
)

// Runner is an interface that "makes" a Request.
// The net/http.Client implementation fills this interface,
// but it can also be filled by something else if necessary.
type Runner interface {
	Do(*http.Request) (*http.Response, error)
}

// Run runs the supplied Behaviours using the supplied Runner,
// passing the Response from each Request into the next
// Behaviour's RequestFunc. This allows a chain of Requests to
// be programmatically created, and each Request can use information
// returned in the previous Response.
func Run(t testing.T, r Runner, bs []Behaviour) {
	var resp *http.Response
	for _, b := range bs {
		t.Logf("Running behaviour %q", b.Description)
		req, err := b.RequestFunc(resp)
		if err != nil {
			t.Fatalf("Error running behaviour %q: %s", b.Description, err)
		}
		resp, err = r.Do(req)
		checkErr := b.CheckFunc(resp, err)
		if checkErr != nil {
			t.Fatalf("Error checking behaviour %q: %s", b.Description, err)
		}
	}
}

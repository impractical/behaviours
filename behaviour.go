package behaviours

import "net/http"

// A Behaviour is a codified expectation that a Request (potentially made
// using the output of a previous Request) will yield a Response/error
// combination that meets a certain criteria. Behaviours should be chained
// together, with each RequestFunc's Request being called, having its
// Response and error checked by CheckFunc, and the Response being passed
// as an argument into the next Behaviour's RequestFunc. This allows for
// programmatic testing of complex behaviours against APIs, when Requests
// need to use values or information from Responses in forming the next
// Request.
type Behaviour struct {
	Description string
	RequestFunc func(*http.Response) (*http.Request, error)
	CheckFunc   func(*http.Response, error) error
}

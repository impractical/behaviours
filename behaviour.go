package behaviours

import "net/http"

// A Behaviour is a codified expectation that a Request (potentially made
// using the output of a previous Request) will yield a Response/error
// combination that meets a certain criteria. Behaviours should be chained
// together, with each Behaviours's GetRequest being called, having its
// Response and error checked by CheckResponse, and the Response being passed
// as an argument into the next Behaviour's GetRequest. This allows for
// programmatic testing of complex behaviours against APIs, when Requests
// need to use values or information from Responses in forming the next
// Request.
type Behaviour interface {
	GetDescription() string
	GetRequest(*http.Response) (*http.Request, error)
	CheckResponse(*http.Response, error) error
}

// StatelessBehaviour is a helper type that implements the Behaviour interface.
// It does not allow for keeping state between API requests, and is thus only
// useful for a stateless request flow. It does minimize the amount of
// boilerplate required when running a stateless request flow, however.
type StatelessBehaviour struct {
	Description       string
	GetRequestFunc    func(*http.Response) (*http.Request, error)
	CheckResponseFunc func(*http.Response, error) error
}

var _ Behaviour = StatelessBehaviour{}

func (s StatelessBehaviour) GetDescription() string {
	return s.Description
}

func (s StatelessBehaviour) GetRequest(r *http.Response) (*http.Request, error) {
	return s.GetRequestFunc(r)
}

func (s StatelessBehaviour) CheckResponse(r *http.Response, err error) error {
	return s.CheckResponseFunc(r, err)
}

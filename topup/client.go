//
//
// File generated from our OpenAPI spec
//
//

// Package topup provides the /topups APIs
package topup

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/form"
)

// Client is used to invoke /topups APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new topup.
func New(params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().New(params)
}

// New creates a new topup.
func (c Client) New(params *stripe.TopupParams) (*stripe.Topup, error) {
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodPost, "/v1/topups", c.Key, params, topup)
	return topup, err
}

// Get returns the details of a topup.
func Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Get(id, params)
}

// Get returns the details of a topup.
func (c Client) Get(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	path := stripe.FormatURLPath("/v1/topups/%s", id)
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, topup)
	return topup, err
}

// Update updates a topup's properties.
func Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Update(id, params)
}

// Update updates a topup's properties.
func (c Client) Update(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	path := stripe.FormatURLPath("/v1/topups/%s", id)
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, topup)
	return topup, err
}

// Cancel is the method for the `POST /v1/topups/{topup}/cancel` API.
func Cancel(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/topups/{topup}/cancel` API.
func (c Client) Cancel(id string, params *stripe.TopupParams) (*stripe.Topup, error) {
	path := stripe.FormatURLPath("/v1/topups/%s/cancel", id)
	topup := &stripe.Topup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, topup)
	return topup, err
}

// List returns a list of topups.
func List(params *stripe.TopupListParams) *Iter {
	return getC().List(params)
}

// List returns a list of topups.
func (c Client) List(listParams *stripe.TopupListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TopupList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/topups", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for topups.
type Iter struct {
	*stripe.Iter
}

// Topup returns the topup which the iterator is currently pointing to.
func (i *Iter) Topup() *stripe.Topup {
	return i.Current().(*stripe.Topup)
}

// TopupList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TopupList() *stripe.TopupList {
	return i.List().(*stripe.TopupList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}

package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockOptions struct {
	APIOptions error
	AppID      string
}

type MockClient struct {
	options MockOptions
}

//func (objectself *Client) accessExternalEndpoint() (string, error) {
//	return "{'externaldata':'29'}", nil
//}

func Test_main(t *testing.T) {
	var expected = "{'finaloutput':{'externaldata':'27'}}"

	/* This tests with the actual client calling an actual endpoint.  We don't want this. */
	var cl = &Client{options: Options{APIOptions: errors.New(""), AppID: "app123"}}
	var observed = useClient(cl)

	/* This tests with a mocked client which merely simulates calling an actual endpoint.  We want this. */
	//var mock_cl = &MockClient{options: MockOptions{APIOptions: errors.New(""), AppID: "app123"}}
	//var observed = useClient(mock_cl)

	assert.Equal(t, expected, observed)
}

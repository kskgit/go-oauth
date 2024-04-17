package main

import (
	"net/url"

	"github.com/stretchr/testify/mock"
)

type MockHttpRequest struct {
	mock.Mock
}

func (m *MockHttpRequest) URL() *url.URL {
	args := m.Called()	
	return args.Get(0).(*url.URL)
}

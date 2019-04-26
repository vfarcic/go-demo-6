package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FunctionalTestSuite struct {
	suite.Suite
	hostIp      string
	servicePath string
}

func TestFunctionalTestSuite(t *testing.T) {
	s := new(FunctionalTestSuite)
	s.hostIp = os.Getenv("ADDRESS")
	s.servicePath = "/demo"
	if len(os.Getenv("SERVICE_PATH")) > 0 {
		s.servicePath = os.Getenv("SERVICE_PATH")
	}
	suite.Run(t, s)
}

func (s *FunctionalTestSuite) SetupTest() {
}

// Functional

func (s FunctionalTestSuite) Test_Hello_ReturnsStatus200() {
	address := fmt.Sprintf("%s%s/hello", s.hostIp, s.servicePath)
	logPrintf("Sending a request to %s\n", address)
	resp, err := http.Get(address)

	s.NoError(err)
	s.Equal(200, resp.StatusCode, "ADDR: ", address)
}

func (s FunctionalTestSuite) Test_Person_ReturnsStatus200() {
	address := fmt.Sprintf("%s%s/person", s.hostIp, s.servicePath)
	logPrintf("Sending a request to %s\n", address)
	resp, err := http.Get(address)

	s.NoError(err)
	s.Equal(200, resp.StatusCode, "ADDR: %s", address)
}

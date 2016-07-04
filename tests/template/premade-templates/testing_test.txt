package RBTtemplate

import (
	"testing"

	"github.com/wolfchase/rainbot/template"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type TestingSuite struct {
	suite.Suite
	// Members
}

func (s *TestingSuite) SetupTest() {
	// Setup
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *TestingSuite) TestSomething() {

}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestTestingSuite(t *testing.T) {
	suite.Run(t, new(TestingSuite))
}

/**************************************************************************************************/

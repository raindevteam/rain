package tmpl

const TestTemplate = `package RBT{{.Lib}}

import (
	"testing"

	"github.com/wolfchase/rainbot/{{.Lib}}"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type {{.SuiteName}}Suite struct {
	suite.Suite
	// Members
}

func (s *{{.SuiteName}}Suite) SetupTest() {
	// Setup
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *{{.SuiteName}}Suite) TestSomething() {

}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func Test{{.SuiteName}}Suite(t *testing.T) {
	suite.Run(t, new({{.SuiteName}}Suite))
}

/**************************************************************************************************/
`

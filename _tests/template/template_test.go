package RBTtemplate

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wolfchase/rainbot/template"
)

/****                                 Suite Configuration                                      ****/

type TemplateSuite struct {
	suite.Suite
	// Members
}

func (s *TemplateSuite) SetupTest() {
	// Setup
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *TemplateSuite) TestCreateTestTemplate() {
	tmpl.CreateTestTemplate("template", "testing")

	file, err := ioutil.ReadFile("premade-templates/testing_test.txt")

	if err != nil {
		s.FailNow("Error reading file", err.Error())
	}

	templ, err := ioutil.ReadFile("testing_test.go")

	if err != nil {
		s.FailNow("Error reading file", err.Error())
	}

	s.Equal(string(file[:]), string(templ[:]))
}

func (s *TemplateSuite) TearDownSuite() {
	err := os.Remove("testing_test.go")

	if err != nil {
		s.FailNow("Could remove file", err.Error())
	}
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestTemplateSuite(t *testing.T) {
	suite.Run(t, new(TemplateSuite))
}

/**************************************************************************************************/

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

func (s *TemplateSuite) CheckTmplAgainst(premade string, created string) (bool, error) {
	var err error

	file, err := ioutil.ReadFile(premade)
	if err != nil {
		return false, err
	}

	templ, err := ioutil.ReadFile(created)
	if err != nil {
		return false, err
	}

	return string(file[:]) == string(templ[:]), nil
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *TemplateSuite) TestCreateTestTemplate() {
	tmpl.CreateTestTemplate("template", "testing")

	check, err := s.CheckTmplAgainst("premade-templates/testing_test.txt", "testing_test.go")

	if err != nil {
		s.FailNow("Error while checking templates", err.Error())
	}

	s.True(check, "Template matched premade template")
}

func (s *TemplateSuite) TestGoModTemplate() {
	tmpl.CreateModTemplate("go", "test")
	check, err := s.CheckTmplAgainst("premade-templates/go_module.txt", "test.go")

	if err != nil {
		s.FailNow("Error while checking templates", err.Error())
	}

	s.True(check, "Template matched premade template")
}

func (s *TemplateSuite) TestJSModTemplate() {
	tmpl.CreateModTemplate("js", "test")
	check, err := s.CheckTmplAgainst("premade-templates/js_module.txt", "test.js")

	if err != nil {
		s.FailNow("Error while checking templates", err.Error())
	}

	s.True(check, "Template matched premade template")
}

func (s *TemplateSuite) TearDownSuite() {
	if err := os.Remove("testing_test.go"); err != nil {
		s.FailNow("Could remove file", err.Error())
	}

	if err := os.Remove("test.go"); err != nil {
		s.FailNow("Could remove file", err.Error())
	}

	if err := os.Remove("test.js"); err != nil {
		s.FailNow("Could remove file", err.Error())
	}
}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestTemplateSuite(t *testing.T) {
	suite.Run(t, new(TemplateSuite))
}

/**************************************************************************************************/

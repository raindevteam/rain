package Tbot

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type HandlerSuite struct {
	suite.Suite
	h *rainbot.Handler
}

func (s *HandlerSuite) SetupTest() {
	s.h = rainbot.NewHandler()
}

/**************************************************************************************************/

/****                                      Tests Go Here                                       ****/

func (s *HandlerSuite) TestAddModule() {

}

/**************************************************************************************************/

/****                                        Run Suite                                         ****/

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

/**************************************************************************************************/

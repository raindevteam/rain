package rbot

import (
	"testing"

	"github.com/raindevteam/rain/rbot"
	"github.com/stretchr/testify/suite"
)

/****                                 Suite Configuration                                      ****/

type HandlerSuite struct {
	suite.Suite
	h *rbot.Handler
}

func (s *HandlerSuite) SetupTest() {
	s.h = rbot.NewHandler()
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

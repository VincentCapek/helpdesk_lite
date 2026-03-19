package models_test

import (
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ModelSuite struct {
	*suite.Model
}

func Test_ModelSuite(t *testing.T) {
	ms := &ModelSuite{Model: suite.NewModel()}
	suite.Run(t, ms)
}
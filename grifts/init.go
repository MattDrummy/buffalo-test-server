package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/mattdrummy/buffalo-test-server/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}

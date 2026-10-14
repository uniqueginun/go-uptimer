package tests

import (
	"github.com/goravel/framework/testing"

	"uptimer/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}

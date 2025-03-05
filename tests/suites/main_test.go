package suites_test

import (
	"path/filepath"
	"testing"
)

var featuresPath string

func TestMain(m *testing.M) {
	featuresPath, _ = filepath.Abs("../features")

	m.Run()
}

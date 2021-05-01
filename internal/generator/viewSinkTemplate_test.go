package generator_test

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func validateViewSink(tmpDir string, t *testing.T) {
	s, err := ioutil.ReadFile(path.Join(tmpDir, "internal", "kafmesh", "details", "testId_test_viewSink.km.go"))
	if err != nil {
		t.Fatal(err)
	}

	err = cupaloy.SnapshotMulti("validateViewSink", s)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}

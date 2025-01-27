//go:build unit || ALL
// +build unit ALL

/*
 * Copyright 2020 VMware, Inc.  All rights reserved.  Licensed under the Apache v2 License.
 */

package govcd

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

// goldenString is a test helper to manage Golden files. It supports `update` parameter which may be
// useful for writing such files (manual or automated way).
func goldenString(t *testing.T, goldenFile string, actual string, update bool) string {
	t.Helper()

	goldenPath := "../test-resources/golden/" + t.Name() + "_" + goldenFile + ".golden"

	f, err := os.OpenFile(filepath.Clean(goldenPath), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		t.Fatalf("unable to find golden file '%s': %s", goldenPath, err)
	}
	defer safeClose(f)

	if update {
		_, err := f.WriteString(actual)
		if err != nil {
			t.Fatalf("error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("error opening file %s: %s", goldenPath, err)
	}
	return string(content)
}

// goldenBytes wraps goldenString and returns []byte
func goldenBytes(t *testing.T, goldenFile string, actual []byte, update bool) []byte {
	return []byte(goldenString(t, goldenFile, string(actual), update))
}

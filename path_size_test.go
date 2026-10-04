package code

import "testing"

func TestGetPathSizeDirectory(t *testing.T) {
    size, err := GetPathSize("testdata", false, false, false)

    if err != nil {
        t.Fatal(err)
    }

    if size != "8B" {
        t.Errorf("expected 8B, got %s", size)
    }
}
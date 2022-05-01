package lib_test

import (
	"net/url"
	"testing"
	"weblinks/lib"
)

type TestSameOrLevel struct {
	pathRoot string
	pathLink string
	expected bool
}

func TestIsSameOrUpper(t *testing.T) {
	scenery := []TestSameOrLevel{
		{"http://test.com/asdas/asde", "http://test.com/asdas/asde", true},
		{"http://test.com/asdas/asde", "http://test.com/asdas/asde/sdfg", false},
		{"http://test.com/asdas/asde/sdfg", "http://test.com/asdas/asde", true},
		{"http://test.com/mhgjm/asde", "http://test.com/asdas/asde", false},
	}
	for _, scene := range scenery {
		urlRoot, err := url.Parse(scene.pathRoot)
		if err != nil {
			panic(err)
		}
		urlLink, err := url.Parse(scene.pathLink)
		if err != nil {
			panic(err)
		}
		if lib.IsSameOrUpper(urlRoot, urlLink) != scene.expected {
			if scene.expected {
				t.Errorf("Path %s should be SameOrUpper of %s", scene.pathLink, scene.pathRoot)
			} else {
				t.Errorf("Path %s should not be SameOrUpper of %s", scene.pathLink, scene.pathRoot)
			}
		}
	}
}

func TestIsSameOrLower(t *testing.T) {
	scenery := []TestSameOrLevel{
		{"http://test.com/asdas/asde", "http://test.com/asdas/asde", true},
		{"http://test.com/asdas/asde", "http://test.com/asdas/asde/sdfg", true},
		{"http://test.com/asdas/asde/sdfg", "http://test.com/asdas/asde", false},
		{"http://test.com/mhgjm/asde", "http://test.com/asdas/asde", false},
	}
	for _, scene := range scenery {
		urlRoot, err := url.Parse(scene.pathRoot)
		if err != nil {
			panic(err)
		}
		urlLink, err := url.Parse(scene.pathLink)
		if err != nil {
			panic(err)
		}
		if lib.IsSameOrLower(urlRoot, urlLink) != scene.expected {
			if scene.expected {
				t.Errorf("Path %s should be SameOrLower of %s", scene.pathLink, scene.pathRoot)
			} else {
				t.Errorf("Path %s should not be SameOrLower of %s", scene.pathLink, scene.pathRoot)
			}
		}
	}
}

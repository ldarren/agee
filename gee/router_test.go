package gee

import (
	"testing"
)

func TestSplitPatternHelloWorld(t *testing.T) {
	pattern := "/hello/world"
	parts := splitPattern(pattern)
	if 2 != len(parts) || "hello" != parts[0] || "world" != parts[1] {
		t.Fatalf(`splitPattern(%q) = %q, want parts eq len 2`, pattern, parts)
	}
}

func TestSplitPatternRedundantSlash(t *testing.T) {
	pattern := "/hello//world"
	parts := splitPattern(pattern)
	if 2 != len(parts) || "hello" != parts[0] || "world" != parts[1] {
		t.Fatalf(`splitPattern(%q) = %q, want parts eq len 2`, pattern, parts)
	}
}

func TestSplitPatternWithWildCard(t *testing.T) {
	pattern := "/hello/*world/foo"
	parts := splitPattern(pattern)
	if 2 != len(parts) || "hello" != parts[0] || "*world" != parts[1] {
		t.Fatalf(`splitPattern(%q) = %q, want parts eq len 2`, pattern, parts)
	}
}

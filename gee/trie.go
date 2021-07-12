package gee

import (
	"strings"
)

type node struct {
	pattern  string
	part     string
	isWild   bool
	children []*node
}

func newNode(part string) *node {
	return &node{
		part:     part,
		isWild:   strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*"),
		children: make([]*node, 0),
	}
}

func (n *node) matchChild(part string) *node {
	var wild *node = nil
	for _, child := range n.children {
		if child.part == part {
			return child
		}
		if child.isWild {
			wild = child
		}
	}
	return wild
}

func (n *node) insert(pattern string, parts []string, i int) {
	if len(parts) == i {
		n.pattern = pattern
		return
	}

	part := parts[i]
	child := n.matchChild(part)
	if nil == child {
		child = newNode(part)
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, i+1)
}

func (n *node) search(parts []string, i int) *node {
	if len(parts) == i {
		return n
	}

	part := parts[i]
	child := n.matchChild(part)
	if nil == child {
		return child
	}
	return child.search(parts, i+1)
}

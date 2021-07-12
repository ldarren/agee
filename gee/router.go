package gee

import (
	"strings"
)

type router struct {
	head     *node
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		head:     newNode(""),
		handlers: make(map[string]HandleFunc),
	}
}

func splitPattern(pattern string) []string {
	p := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, part := range p {
		if "" == part {
			continue
		}
		parts = append(parts, part)
		if '*' == part[0] {
			break
		}
	}
	return parts
}

func (r *router) add(method string, pattern string, handler HandleFunc) {
	parts := splitPattern(pattern)
	head := r.head
	head.insert(pattern, parts, 0)

	r.handlers[method+"-"+pattern] = handler
}

func (r *router) get(method string, path string) (HandleFunc, map[string]string) {
	head := r.head
	if nil == head {
		return nil, nil
	}

	values := splitPattern(path)
	n := head.search(values, 0)
	if nil == n {
		return nil, nil
	}

	parts := splitPattern(n.pattern)
	params := make(map[string]string)

	for i, part := range parts {
		switch c := part[0]; c {
		case ':':
			params[part[1:]] = values[i]
		case '*':
			params[part[1:]] = strings.Join(values[i:], "/")
			break
		}
	}

	handler := r.handlers[method+"-"+n.pattern]
	return handler, params
}

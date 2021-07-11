package gee

import (
	"testing"
	"reflect"
//	"fmt"
)

func TestSplitPattern(t *testing.T) {
	ok := reflect.DeepEqual(splitPattern("/hello/world"), []string{"hello", "world"})
	ok = ok && reflect.DeepEqual(splitPattern("/hello//world"), []string{"hello", "world"})
	ok = ok && reflect.DeepEqual(splitPattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(splitPattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(splitPattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test splitPattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newRouter()
	r.add("GET", "/", nil)
	r.add("GET", "/hello/:name", nil)
/*	r.add("GET", "/hello/b/c", nil)
	r.add("GET", "/hi/:name", nil)
	r.add("GET", "/assets/*filepath", nil)

	_, ps := r.get("GET", "/hello/geektutu")

	if ps == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched , params['name']: %s\n", ps["name"])*/
}

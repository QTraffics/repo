package stringconf

import (
	"iter"
	"slices"
	"strings"
	"sync"
)

type SplitFunc func(raw string) []string

func SplitBy(v string) SplitFunc {
	if len(v) == 0 {
		panic("empty split sep")
	}
	return func(raw string) []string {
		return slices.DeleteFunc(strings.Split(raw, v), func(vs string) bool {
			return len(vs) == 0 || vs == v
		})
	}
}

type Fields struct {
	Raw       string
	SplitFunc SplitFunc

	current  int
	elements []string

	once sync.Once
}

func (f *Fields) lazy() {
	f.once.Do(func() {
		if f.SplitFunc == nil {
			f.SplitFunc = SplitBy(",")
		}
		if len(f.Raw) == 0 {
			f.elements = []string{}
			f.current = 0
			return
		}

		splitResult := f.SplitFunc(f.Raw)
		f.current = 0
		f.elements = splitResult
	})
}

func (f *Fields) Index(n int) string {
	f.lazy()
	return f.elements[n]
}

func (f *Fields) Len() int {
	f.lazy()
	return len(f.elements)
}

func (f *Fields) Pick() string {
	if !f.HasNext() {
		return ""
	}
	return f.elements[f.current]
}

func (f *Fields) PickKV() (k, v string) {
	if !f.HasNext() {
		return "", ""
	}

	pick := f.elements[f.current]
	n := strings.SplitN(pick, "=", 2)
	if len(n) < 2 {
		return n[0], ""
	}
	return n[0], n[1]
}

func (f *Fields) PickNext() string {
	v := f.Pick()
	f.Next()
	return v
}

func (f *Fields) PickKVNext() (k, v string) {
	k, v = f.PickKV()
	f.Next()
	return
}

func (f *Fields) HasNext() bool {
	f.lazy()
	return f.current < len(f.elements)
}

func (f *Fields) Next() bool {
	f.lazy()
	f.current++
	return f.current < len(f.elements)
}

func (f *Fields) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for f.HasNext() {
			if yield(f.PickNext()) {
				continue
			}
			break
		}
	}
}

func (f *Fields) IterKV() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for f.HasNext() {
			if yield(f.PickKVNext()) {
				continue
			}
			break
		}
	}
}

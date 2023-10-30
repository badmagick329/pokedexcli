package pokecache

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type TestCacheCase struct {
	name   string
	key    string
	action string
	val    []byte
	want   []byte
}

func TestCache(t *testing.T) {
	cases := testCacheCases()
	cache := NewCache(time.Minute * 1)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []byte
			if tc.action == "add" {
				got = cache.Add(tc.key, tc.val)
			} else if tc.action == "get" {
				got = cache.Get(tc.key)
			}
			diff := cmp.Diff(string(tc.want), string(got))
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func testCacheCases() []TestCacheCase {
	return []TestCacheCase{
		{
			name:   "add test",
			key:    "test",
			action: "add",
			val:    []byte("test val"),
			want:   []byte("test val"),
		},
		{
			name:   "get test",
			key:    "test",
			action: "get",
			val:    []byte{},
			want:   []byte("test val"),
		},
		{
			name:   "get non existent test",
			key:    "something",
			action: "get",
			val:    []byte{},
			want:   []byte{},
		},
		{
			name:   "add empty",
			key:    "",
			action: "add",
			val:    []byte("empty key"),
			want:   []byte("empty key"),
		},
		{
			name:   "get empty",
			key:    "",
			action: "get",
			val:    []byte{},
			want:   []byte("empty key"),
		},
		{
			name:   "add empty val",
			key:    "no val",
			action: "add",
			val:    []byte{},
			want:   []byte{},
		},
		{
			name:   "get empty val",
			key:    "no val",
			action: "get",
			val:    []byte{},
			want:   []byte{},
		},
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 250
	c := NewCache(interval)
	c.Add("test1", []byte("test1"))
	time.Sleep(time.Millisecond * 10)
	val := c.Get("test1")
	if val == nil {
		t.Errorf("Value was reaped too early")
	}
	time.Sleep(interval)
	val = c.Get("test1")
	if val != nil {
		t.Errorf("Value was not reaped")
	}
}

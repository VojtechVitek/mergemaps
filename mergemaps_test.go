package mergemaps_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/VojtechVitek/mergemaps"
)

func ExampleMerge() {
	m := map[string]int{}
	m1 := map[string]int{"foo": 0}
	m2 := map[string]int{"bar": 1, "baz": 2}

	mergemaps.Merge(m, m1, 0)
	mergemaps.Merge(m, m2, 0)

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
	// Output:
	// foo: 0
	// bar: 1
	// baz: 2
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		dst        interface{}
		src        interface{}
		flags      int
		shouldPass bool
		expected   interface{}
	}{
		{ // Test empty maps
			map[int]int{},
			map[int]int{},
			0,
			true,
			map[int]int{},
		},
		{ // Test dst + src => expected
			map[int]string{1: "foo"},
			map[int]string{2: "bar"},
			0,
			true,
			map[int]string{1: "foo", 2: "bar"},
		},
		{ // Test dst + src => expected, do not overwrite dst
			map[string]string{"foo": "bar"},
			map[string]string{"foo": ""},
			0,
			true,
			map[string]string{"foo": "bar"},
		},
		{ // Test dst + src => expected, overwrite dst
			map[string]string{"foo": "bar"},
			map[string]string{"foo": ""},
			mergemaps.OverwriteExistingDstKey,
			true,
			map[string]string{"foo": ""},
		},
		{ // Test dst + src => expected, error on existing key value
			map[string]string{"foo": "bar"},
			map[string]string{"foo": "bar"},
			mergemaps.ErrorOnExistingDstKey | mergemaps.OverwriteExistingDstKey,
			false,
			map[string]string{"foo": "bar"},
		},
		{ // Test dst + src => expected, do not error on same key value
			map[string]string{"foo": "bar"},
			map[string]string{"foo": "bar"},
			mergemaps.ErrorOnDifferentDstKeyValue | mergemaps.OverwriteExistingDstKey,
			true,
			map[string]string{"foo": "bar"},
		},
		{ // Test dst + src => expected, error on different key value
			map[string]string{"foo": "bar"},
			map[string]string{"foo": ""},
			mergemaps.ErrorOnDifferentDstKeyValue | mergemaps.OverwriteExistingDstKey,
			false,
			map[string]string{"foo": "bar"},
		},
	}

	for i, test := range testCases {
		err := mergemaps.Merge(test.dst, test.src, test.flags)
		if err != nil && test.shouldPass {
			t.Errorf("Unexpected error while merging § on testCase[%v].", i)
		}
		if err == nil && !test.shouldPass {
			t.Errorf("Unexpected non-error while merging maps on testCase[%v].", i)
		}
		if !reflect.DeepEqual(test.dst, test.expected) {
			t.Errorf("Unexpected map on testCase[%v]. Expected: %v, got: %v.", i, test.expected, test.dst)
		}
	}
}

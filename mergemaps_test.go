package mergemaps_test

import (
	"reflect"
	"testing"

	"github.com/VojtechVitek/mergemaps"
)

func TestMergeInto(t *testing.T) {
	var nilMap map[int]int

	testCases := []struct {
		dst      interface{}
		src      interface{}
		flags    int
		err      bool
		expected interface{}
	}{
		{ // Can't merge into nil
			dst:      nil,
			src:      map[int]int{},
			flags:    0,
			err:      true,
			expected: nil,
		},
		{ // Merge nil into an empty map
			dst:      map[int]int{},
			src:      nil,
			flags:    0,
			err:      false,
			expected: map[int]int{},
		},
		{ // Can't merge into nil map
			dst:      nilMap,
			src:      map[int]int{},
			flags:    0,
			err:      true,
			expected: nilMap,
		},
		{ // Can't merge into pointer
			dst:      &nilMap,
			src:      map[int]int{},
			flags:    0,
			err:      true,
			expected: &nilMap,
		},
		{ // Test empty maps
			dst:      map[int]int{},
			src:      map[int]int{},
			flags:    0,
			err:      false,
			expected: map[int]int{},
		},
		{ // Test dst + src => expected
			dst:      map[int]byte{0: 0, 1: 1},
			src:      map[int]byte{2: 2, 3: 3},
			flags:    0,
			err:      false,
			expected: map[int]byte{0: 0, 1: 1, 2: 2, 3: 3},
		},
		{ // Test dst + src => expected, do not overwrite dst
			dst:      map[string]string{"foo": "bar"},
			src:      map[string]string{"foo": ""},
			flags:    0,
			err:      false,
			expected: map[string]string{"foo": "bar"},
		},
		{ // Test dst + src => expected, overwrite dst
			dst:      map[string]string{"foo": "bar"},
			src:      map[string]string{"foo": ""},
			flags:    mergemaps.OverwriteExistingDstKey,
			err:      false,
			expected: map[string]string{"foo": ""},
		},
		{ // Test dst + src => expected, error on existing key value
			dst:      map[string]string{"foo": "bar"},
			src:      map[string]string{"foo": "bar"},
			flags:    mergemaps.ErrorOnExistingDstKey | mergemaps.OverwriteExistingDstKey,
			err:      true,
			expected: map[string]string{"foo": "bar"},
		},
		{ // Test dst + src => expected, do not error on same key value
			dst:      map[string]string{"foo": "bar"},
			src:      map[string]string{"foo": "bar"},
			flags:    mergemaps.ErrorOnDifferentDstKeyValue | mergemaps.OverwriteExistingDstKey,
			err:      false,
			expected: map[string]string{"foo": "bar"},
		},
		{ // Test dst + src => expected, error on different key value
			dst:      map[string]string{"foo": "bar"},
			src:      map[string]string{"foo": ""},
			flags:    mergemaps.ErrorOnDifferentDstKeyValue | mergemaps.OverwriteExistingDstKey,
			err:      true,
			expected: map[string]string{"foo": "bar"},
		},
	}

	for i, test := range testCases {
		err := mergemaps.MergeInto(test.dst, test.src, test.flags)
		if err != nil && !test.err {
			t.Errorf("Unexpected error while merging maps on testCase[%v]: %v.", i, err)
		}
		if err == nil && test.err {
			t.Errorf("Unexpected non-error while merging maps on testCase[%v].", i)
		}

		if !reflect.DeepEqual(test.dst, test.expected) {
			t.Errorf("Unexpected map on testCase[%v]. Expected: %#v, got: %#v.", i, test.expected, test.dst)
		}

	}
}

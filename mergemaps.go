package mergemaps

import (
	"fmt"
	"reflect"
)

// Merge flags
const (
	OverwriteExistingDstKey     = 1 << iota
	ErrorOnExistingDstKey       = 1 << iota
	ErrorOnDifferentDstKeyValue = 1 << iota
)

// MergeInto merges items from a src map into a dst map.
// Returns an error when the maps are not of the same type.
// Flags:
// - ErrorOnExistingDstKey
//     When set: Return an error if any of the dst keys is already set.
// - ErrorOnDifferentDstKeyValue
//     When set: Return an error if any of the dst keys is already set
//               to a different value than src key.
// - OverwriteDstKey
//     When set: Overwrite existing dst key value with src key value.
func MergeInto(dst, src interface{}, flags int) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	if !dstVal.IsValid() {
		return fmt.Errorf("Dst is not a valid value")
	}
	if dstVal.Kind() != reflect.Map {
		return fmt.Errorf("Dst is not a map: %v", dstVal.Kind())
	}

	if !srcVal.IsValid() || srcVal.IsNil() {
		// Nothing to merge
		return nil
	}

	dstTyp := dstVal.Type()
	srcTyp := srcVal.Type()
	if !dstTyp.AssignableTo(srcTyp) {
		return fmt.Errorf("Type mismatch, can't assign '%v' to '%v'", srcTyp, dstTyp)
	}

	if dstVal.IsNil() {
		return fmt.Errorf("Dst value is nil")
	}

	for _, k := range srcVal.MapKeys() {
		if dstVal.MapIndex(k).IsValid() {
			if flags&ErrorOnExistingDstKey != 0 {
				return fmt.Errorf("ErrorOnExistingDstKey flag: Dst key already set to a different value, '%v'='%v'", k, dstVal.MapIndex(k))
			}
			if dstVal.MapIndex(k).String() != srcVal.MapIndex(k).String() {
				if flags&ErrorOnDifferentDstKeyValue != 0 {
					return fmt.Errorf("ErrorOnDifferentDstKeyValue flag: Dst key already set to a different value, '%v'='%v'", k, dstVal.MapIndex(k))
				}
				if flags&OverwriteExistingDstKey != 0 {
					dstVal.SetMapIndex(k, srcVal.MapIndex(k))
				}
			}
		} else {
			dstVal.SetMapIndex(k, srcVal.MapIndex(k))
		}
	}

	return nil
}

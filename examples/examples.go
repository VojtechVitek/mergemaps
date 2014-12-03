package main

import (
	"fmt"

	"github.com/VojtechVitek/mergemaps"
)

func mergeMapStringString() {
	dst := map[string]string{"foo": "bar"}
	src := map[string]string{"baz": ""}

	// Merge src map into dst map
	err := mergemaps.MergeInto(dst, src, 0 /*flags*/)
	if err != nil {
		fmt.Errorf("Can't merge src into dst: %v", err)
	}

	// Prints map[string]string{"baz":"", "foo":"bar"}
	fmt.Printf("%#v\n", dst)
}

func mergeMapBoolInt() {
	dst := map[bool]int{true: 1}
	src := map[bool]int{false: 0}

	// Merge src map into dst map
	err := mergemaps.MergeInto(dst, src, 0 /*flags*/)
	if err != nil {
		fmt.Errorf("Can't merge src into dst: %v", err)
	}

	// Prints map[bool]int{true:1, false:0}
	fmt.Printf("%#v\n", dst)
}

func mergeMapRuneFloat() {
	dst := map[rune]float64{'♥': 99.999, 'x': 0.0}
	src := map[rune]float64{'Ř': -10.10, '⌘': -0.0}

	// Merge src map into dst map
	err := mergemaps.MergeInto(dst, src, 0 /*flags*/)
	if err != nil {
		fmt.Errorf("Can't merge src into dst: %v", err)
	}

	// Prints map[int32]float64{9829:99.999, 120:0, 344:-10.1, 8984:0}
	fmt.Printf("%#v\n", dst)
}

func main() {
	mergeMapStringString()
	mergeMapBoolInt()
	mergeMapRuneFloat()
}

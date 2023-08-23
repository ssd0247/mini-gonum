package elements

import (
	"fmt"
	"strings"
	"testing"
)

func Test_PaddingWithInvalidPrefix(t *testing.T) {
	v := &IntVector{Values: []int{2}}
	actual := v.Padding(10, 10, "pref")
	expected := fmt.Errorf(pad_prefix_type_err, "pref")
	if strings.TrimSpace(actual.Error()) != strings.TrimSpace(expected.Error()) {
		t.Errorf("Expected = %s, Got = %s", expected.Error(), actual.Error())
	}
}

func Test_PaddingWithInvalidLen(t *testing.T) {
	v := &IntVector{Values: []int{4, 5}}
	actual := v.Padding(3, -1, true)
	expected := pad_out_of_domain_err
	if strings.TrimSpace(actual.Error()) != strings.TrimSpace(expected) {
		t.Errorf("Expected = %s, Got = %s", expected, actual.Error())
	}
}

func Test_PaddingWithValidArguments(t *testing.T) {
	v := &IntVector{Values: []int{3, 4, 5}}
	v.Padding(2, 1, true)
	expected := []int{2, 3, 4, 5}
	if len(expected) != len(v.Values) {
		t.Errorf("Expected = %v, Got = %v", expected, v.Values)
	} else {
		for i := range v.Values {
			if v.Values[i] != expected[i] {
				t.Errorf("Different values at array-index %d", i)
			}
		}
	}
	v.Padding(1, 2, false)
	expected = []int{2, 3, 4, 5, 1, 1}
	if len(expected) != len(v.Values) {
		t.Errorf("Expected = %v, Got = %v", expected, v.Values)
	} else {
		for i := range v.Values {
			if v.Values[i] != expected[i] {
				t.Errorf("Different values at array-index %d", i)
			}
		}
	}
}

func Test_CopyFromNilSrc(t *testing.T) {
	v := &IntVector{Values: []int{1, 2, 3, 4, 5}}
	var src []int
	if ans := v.CopyFrom(src); ans {
		t.Errorf("Expected = %v, Got = %v", false, ans)
	}
}

func Test_CopyFromNilValues(t *testing.T) {
	v := &IntVector{Values: nil}
	dest := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if ans := v.CopyFrom(dest); ans {
		for i := range dest {
			if dest[i] != v.Values[i] {
				t.Errorf("Different values at array-index %d", i)
			}
		}
	} else {
		t.Errorf("Expected = %v, Got = %v", true, ans)
	}
}

func Test_CopyFromLongerSrc(t *testing.T) {
	v := &IntVector{Values: []int{1, 2, 3, 4}}
	dest := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if ans := v.CopyFrom(dest); ans {
		t.Errorf("Expected = %v, Got = %v", true, ans)
	}
}

func Test_CopyFromDefault(t *testing.T) {
	v := &IntVector{Values: []int{1, 2, 3, 4}}
	dest := []int{9, 8, 7}
	if ans := v.CopyFrom(dest); !ans {
		t.Errorf("Expected = %v, Got = %v", true, ans)
	}
	fmt.Println(v.Values)
}

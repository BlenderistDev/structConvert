package dry

import (
	"reflect"
	"testing"
)

func TestCheckEqual(t *testing.T, expected, testing interface{}) {
	if !reflect.DeepEqual(expected, testing) {
		t.Errorf("expect %v, got %v", expected, testing)
	}
}

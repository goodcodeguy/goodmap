package goodmap

import (
	"reflect"
	"strings"
	"testing"
)

type TestSource struct {
	Name string
	ID   uint
}

type TestDest struct {
	Name string
	ID   uint
}

type SuperMapper struct {
	Mapper
}

var superMapper = SuperMapper{}
var genericMapper = Mapper{}

func (mapper SuperMapper) MapFromString(src string, dest interface{}) {
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr {
		panic("dest must be a pointer")
	}

	val.Elem().Set(reflect.ValueOf("bar"))
}

func Test_MapToSimple(t *testing.T) {
	src := "test"
	var dest string

	genericMapper.MapFromString(src, &dest)
	if strings.Compare(src, dest) != 0 {
		t.Errorf("Expected %s, got %s", src, dest)
	}
}

func Test_SuperMapper(t *testing.T) {
	var dest string

	superMapper.MapFromString("foo", &dest)
	if strings.Compare("bar", dest) != 0 {
		t.Errorf("Expected %s, got %s", "bar", dest)
	}
}

func Test_MapFromStringNoDestinationPointer(t *testing.T) {
	var dest string

	err := genericMapper.MapFromString("test", dest)
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func Test_MapFromIntWithNoDestinationPointer(t *testing.T) {
	var dest string

	err := genericMapper.MapFromInt(1, dest)
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func Test_MapFromStringToInt(t *testing.T) {
	src := ""
	var dest int

	err := genericMapper.MapFromString(src, &dest)
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func Test_MapFromIntToString(t *testing.T) {
	src := 5
	var dest string

	_ = genericMapper.MapFromInt(src, &dest)
	if strings.Compare("5", dest) != 0 {
		t.Errorf("Expected %s, got %s", "5", dest)
	}
}

func Test_MapFromIntToInt(t *testing.T) {
	src := 5
	var dest int64

	_ = genericMapper.MapFromInt(src, &dest)
	if int64(src) != dest {
		t.Errorf("Expected %d, got %d", src, dest)
	}
}

func Test_MapFromBoolTrueToInt(t *testing.T) {
	src := true
	var dest int64

	_ = genericMapper.MapFromBool(src, &dest)
	if dest != 1 {
		t.Errorf("Expected 1, got %d", dest)
	}
}

func Test_MapFromBoolFalseToInt(t *testing.T) {
	src := false
	var dest int64

	_ = genericMapper.MapFromBool(src, &dest)
	if dest != 0 {
		t.Errorf("Expected 0, got %d", dest)
	}
}

func Test_MapFromBoolWithNoDestinationPointer(t *testing.T) {
	var dest string

	err := genericMapper.MapFromBool(false, dest)
	if err == nil {
		t.Errorf("Expected error.")
	}
}

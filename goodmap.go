package goodmap

import (
	"errors"
	"reflect"
	"strconv"
)

type MapperInterface interface {
	MapFromString(string, interface{}) error
}

type Mapper struct {
	MapperInterface
}

func (mapper Mapper) MapFromString(src string, dest interface{}) error {
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr {
		return errors.New("Destination must be a pointer")
	}

	if val.Elem().Kind() != reflect.String {
		return errors.New("Destination must be a pointer to a String")
	}

	val.Elem().Set(reflect.ValueOf(src))
	return nil
}

func (mapper Mapper) MapFromInt(src int, dest interface{}) error {
	destVal := reflect.ValueOf(dest)
	srcVal := int64(src)

	if destVal.Kind() != reflect.Ptr {
		return errors.New("Destination must be a pointer")
	}

	switch destVal.Elem().Kind() {
	case reflect.String:
		destVal.Elem().Set(reflect.ValueOf(strconv.FormatInt(srcVal, 10)))
	case reflect.Int:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		destVal.Elem().Set(reflect.ValueOf(srcVal))
	}

	return nil
}

func (mapper Mapper) MapFromBool(src bool, dest interface{}) error {
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr {
		return errors.New("Destination must be a pointer.")
	}

	switch destVal.Elem().Kind() {
	case reflect.Int:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		if src {
			destVal.Elem().Set(reflect.ValueOf(int64(1)))
		} else {
			destVal.Elem().Set(reflect.ValueOf(int64(0)))
		}
	}
	return nil
}

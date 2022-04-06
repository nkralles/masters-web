package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Load(data interface{}) error {
	v := reflect.ValueOf(data).Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		key := t.Field(i).Tag.Get("env")
		if key == "" {
			continue
		}
		if !f.CanSet() {
			continue
		}
		s := os.Getenv(key)
		var err error
		switch f.Kind() {
		case reflect.String:
			f.SetString(s)
		case reflect.Bool:
			var b bool
			if s == "" {
				f.SetBool(false)
			} else if b, err = strconv.ParseBool(s); err == nil {
				f.SetBool(b)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			if s == "" {
				f.SetInt(0)
			} else if i, err = strconv.ParseInt(s, 10, 64); err == nil {
				f.SetInt(i)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			if s == "" {
				f.SetUint(0)
			} else if i, err = strconv.ParseUint(s, 10, 64); err == nil {
				f.SetUint(i)
			}
		case reflect.Float32, reflect.Float64:
			var n float64
			if s == "" {
				f.SetFloat(0.0)
			} else if n, err = strconv.ParseFloat(s, 64); err == nil {
				f.SetFloat(n)
			}
		default:
			continue
		}
		if err != nil {
			return fmt.Errorf("failed to load environment variable '%s': %v", key, err)
		}
	}
	return nil
}
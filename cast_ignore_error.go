// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast provides easy and safe casting in Go.
package cast

import "time"

// ToBoolIgnoreError casts an interface to a bool type.
func ToBoolIgnoreError(i interface{}) bool {
	v, _ := ToBool(i)
	return v
}

// ToTimeIgnoreError casts an interface to a time.Time type.
func ToTimeIgnoreError(i interface{}) time.Time {
	v, _ := ToTime(i)
	return v
}

// ToDurationIgnoreError casts an interface to a time.Duration type.
func ToDurationIgnoreError(i interface{}) time.Duration {
	v, _ := ToDuration(i)
	return v
}

// ToFloat64IgnoreError casts an interface to a float64 type.
func ToFloat64IgnoreError(i interface{}) float64 {
	v, _ := ToFloat64(i)
	return v
}

// ToFloat32IgnoreError casts an interface to a float32 type.
func ToFloat32IgnoreError(i interface{}) float32 {
	v, _ := ToFloat32(i)
	return v
}

// ToInt64IgnoreError casts an interface to an int64 type.
func ToInt64IgnoreError(i interface{}) int64 {
	v, _ := ToInt64(i)
	return v
}

// ToInt32IgnoreError casts an interface to an int32 type.
func ToInt32IgnoreError(i interface{}) int32 {
	v, _ := ToInt32(i)
	return v
}

// ToInt16IgnoreError casts an interface to an int16 type.
func ToInt16IgnoreError(i interface{}) int16 {
	v, _ := ToInt16(i)
	return v
}

// ToInt8IgnoreError casts an interface to an int8 type.
func ToInt8IgnoreError(i interface{}) int8 {
	v, _ := ToInt8(i)
	return v
}

// ToIntIgnoreError casts an interface to an int type.
func ToIntIgnoreError(i interface{}) int {
	v, _ := ToInt(i)
	return v
}

// ToUintIgnoreError casts an interface to a uint type.
func ToUintIgnoreError(i interface{}) uint {
	v, _ := ToUint(i)
	return v
}

// ToUint64IgnoreError casts an interface to a uint64 type.
func ToUint64IgnoreError(i interface{}) uint64 {
	v, _ := ToUint64(i)
	return v
}

// ToUint32IgnoreError casts an interface to a uint32 type.
func ToUint32IgnoreError(i interface{}) uint32 {
	v, _ := ToUint32(i)
	return v
}

// ToUint16IgnoreError casts an interface to a uint16 type.
func ToUint16IgnoreError(i interface{}) uint16 {
	v, _ := ToUint16(i)
	return v
}

// ToUint8IgnoreError casts an interface to a uint8 type.
func ToUint8IgnoreError(i interface{}) uint8 {
	v, _ := ToUint8(i)
	return v
}

// ToStringIgnoreError casts an interface to a string type.
func ToStringIgnoreError(i interface{}) string {
	v, _ := ToString(i)
	return v
}

// ToStringMapStringIgnoreError casts an interface to a map[string]string type.
func ToStringMapStringIgnoreError(i interface{}) map[string]string {
	v, _ := ToStringMapString(i)
	return v
}

// ToStringMapStringSliceIgnoreError casts an interface to a map[string][]string type.
func ToStringMapStringSliceIgnoreError(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSlice(i)
	return v
}

// ToStringMapBoolIgnoreError casts an interface to a map[string]bool type.
func ToStringMapBoolIgnoreError(i interface{}) map[string]bool {
	v, _ := ToStringMapBool(i)
	return v
}

// ToStringMapIntIgnoreError casts an interface to a map[string]int type.
func ToStringMapIntIgnoreError(i interface{}) map[string]int {
	v, _ := ToStringMapInt(i)
	return v
}

// ToStringMapInt64IgnoreError casts an interface to a map[string]int64 type.
func ToStringMapInt64IgnoreError(i interface{}) map[string]int64 {
	v, _ := ToStringMapInt64(i)
	return v
}

// ToStringMapIgnoreError casts an interface to a map[string]interface{} type.
func ToStringMapIgnoreError(i interface{}) map[string]interface{} {
	v, _ := ToStringMap(i)
	return v
}

// ToSliceIgnoreError casts an interface to a []interface{} type.
func ToSliceIgnoreError(i interface{}) []interface{} {
	v, _ := ToSlice(i)
	return v
}

// ToBoolSliceIgnoreError casts an interface to a []bool type.
func ToBoolSliceIgnoreError(i interface{}) []bool {
	v, _ := ToBoolSlice(i)
	return v
}

// ToStringSliceIgnoreError casts an interface to a []string type.
func ToStringSliceIgnoreError(i interface{}) []string {
	v, _ := ToStringSlice(i)
	return v
}

// ToIntSliceIgnoreError casts an interface to a []int type.
func ToIntSliceIgnoreError(i interface{}) []int {
	v, _ := ToIntSlice(i)
	return v
}

// ToDurationSliceIgnoreError casts an interface to a []time.Duration type.
func ToDurationSliceIgnoreError(i interface{}) []time.Duration {
	v, _ := ToDurationSlice(i)
	return v
}

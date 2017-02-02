package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

var (
	// reflectMapAny     = reflectCast(map[string]interface{}(nil), visitMapInterface)
	reFoldMapBool    = liftFold(map[string]bool(nil), foldMapBool)
	reFoldMapInt     = liftFold(map[string]int(nil), foldMapInt)
	reFoldMapInt8    = liftFold(map[string]int8(nil), foldMapInt8)
	reFoldMapInt16   = liftFold(map[string]int16(nil), foldMapInt16)
	reFoldMapInt32   = liftFold(map[string]int32(nil), foldMapInt32)
	reFoldMapInt64   = liftFold(map[string]int64(nil), foldMapInt64)
	reFoldMapUint    = liftFold(map[string]uint(nil), foldMapUint)
	reFoldMapUint8   = liftFold(map[string]uint8(nil), foldMapUint8)
	reFoldMapUint16  = liftFold(map[string]uint16(nil), foldMapUint16)
	reFoldMapUint32  = liftFold(map[string]uint32(nil), foldMapUint32)
	reFoldMapUint64  = liftFold(map[string]uint64(nil), foldMapUint64)
	reFoldMapFloat32 = liftFold(map[string]float32(nil), foldMapFloat32)
	reFoldMapFloat64 = liftFold(map[string]float64(nil), foldMapFloat64)
	reFoldMapString  = liftFold(map[string]string(nil), foldMapString)
)

var tMapAny = reflect.TypeOf(map[string]interface{}(nil))

func reflectMapAny(C *foldContext, v reflect.Value) error {
	if v.Type().Name() != "" {
		v = v.Convert(tArrayAny)
	}
	return foldMapInterface(C, v.Interface())
}

func foldMapInterface(C *foldContext, v interface{}) error {
	m := v.(map[string]interface{})
	if err := C.OnObjectStart(len(m), structform.AnyType); err != nil {
		return err
	}

	for k, v := range m {
		if err := C.OnKey(k); err != nil {
			return err
		}
		if err := foldInterfaceValue(C, v); err != nil {
			return err
		}
	}
	return C.OnObjectFinished()
}

func foldMapBool(C *foldContext, v interface{}) error {
	return C.OnBoolObject(v.(map[string]bool))
}

func foldMapString(C *foldContext, v interface{}) error {
	return C.OnStringObject(v.(map[string]string))
}

func foldMapInt8(C *foldContext, v interface{}) error {
	return C.OnInt8Object(v.(map[string]int8))
}

func foldMapInt16(C *foldContext, v interface{}) error {
	return C.OnInt16Object(v.(map[string]int16))
}

func foldMapInt32(C *foldContext, v interface{}) error {
	return C.OnInt32Object(v.(map[string]int32))
}

func foldMapInt64(C *foldContext, v interface{}) error {
	return C.OnInt64Object(v.(map[string]int64))
}

func foldMapInt(C *foldContext, v interface{}) error {
	return C.OnIntObject(v.(map[string]int))
}

func foldMapUint8(C *foldContext, v interface{}) error {
	return C.OnUint8Object(v.(map[string]uint8))
}

func foldMapUint16(C *foldContext, v interface{}) error {
	return C.OnUint16Object(v.(map[string]uint16))
}

func foldMapUint32(C *foldContext, v interface{}) error {
	return C.OnUint32Object(v.(map[string]uint32))
}

func foldMapUint64(C *foldContext, v interface{}) error {
	return C.OnUint64Object(v.(map[string]uint64))
}

func foldMapUint(C *foldContext, v interface{}) error {
	return C.OnUintObject(v.(map[string]uint))
}

func foldMapFloat32(C *foldContext, v interface{}) error {
	return C.OnFloat32Object(v.(map[string]float32))
}

func foldMapFloat64(C *foldContext, v interface{}) error {
	return C.OnFloat64Object(v.(map[string]float64))
}
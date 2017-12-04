// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package bar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"math"
	"strconv"
	"strings"
)

type BarException struct {
	StringField string `json:"stringField,required"`
}

// ToWire translates a BarException struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *BarException) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a BarException struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BarException struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BarException
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BarException) FromWire(w wire.Value) error {
	var err error

	stringFieldIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		}
	}

	if !stringFieldIsSet {
		return errors.New("field StringField of BarException is required")
	}

	return nil
}

// String returns a readable string representation of a BarException
// struct.
func (v *BarException) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++

	return fmt.Sprintf("BarException{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this BarException match the
// provided BarException.
//
// This function performs a deep comparison.
func (v *BarException) Equals(rhs *BarException) bool {
	if !(v.StringField == rhs.StringField) {
		return false
	}

	return true
}

func (v *BarException) Error() string {
	return v.String()
}

type BarRequest struct {
	StringField string    `json:"stringField,required"`
	BoolField   bool      `json:"boolField,required"`
	BinaryField []byte    `json:"binaryField,required"`
	Timestamp   Timestamp `json:"timestamp,required"`
	EnumField   Fruit     `json:"enumField,required"`
	LongField   Long      `json:"longField,required"`
}

// ToWire translates a BarRequest struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *BarRequest) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueBool(v.BoolField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.BinaryField == nil {
		return w, errors.New("field BinaryField of BarRequest is required")
	}
	w, err = wire.NewValueBinary(v.BinaryField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++

	w, err = v.Timestamp.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 4, Value: w}
	i++

	w, err = v.EnumField.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++

	w, err = v.LongField.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 6, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Timestamp_Read(w wire.Value) (Timestamp, error) {
	var x Timestamp
	err := x.FromWire(w)
	return x, err
}

func _Fruit_Read(w wire.Value) (Fruit, error) {
	var v Fruit
	err := v.FromWire(w)
	return v, err
}

func _Long_Read(w wire.Value) (Long, error) {
	var x Long
	err := x.FromWire(w)
	return x, err
}

// FromWire deserializes a BarRequest struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BarRequest struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BarRequest
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BarRequest) FromWire(w wire.Value) error {
	var err error

	stringFieldIsSet := false
	boolFieldIsSet := false
	binaryFieldIsSet := false
	timestampIsSet := false
	enumFieldIsSet := false
	longFieldIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBool {
				v.BoolField, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				boolFieldIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				v.BinaryField, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
				binaryFieldIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TI64 {
				v.Timestamp, err = _Timestamp_Read(field.Value)
				if err != nil {
					return err
				}
				timestampIsSet = true
			}
		case 5:
			if field.Value.Type() == wire.TI32 {
				v.EnumField, err = _Fruit_Read(field.Value)
				if err != nil {
					return err
				}
				enumFieldIsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TI64 {
				v.LongField, err = _Long_Read(field.Value)
				if err != nil {
					return err
				}
				longFieldIsSet = true
			}
		}
	}

	if !stringFieldIsSet {
		return errors.New("field StringField of BarRequest is required")
	}

	if !boolFieldIsSet {
		return errors.New("field BoolField of BarRequest is required")
	}

	if !binaryFieldIsSet {
		return errors.New("field BinaryField of BarRequest is required")
	}

	if !timestampIsSet {
		return errors.New("field Timestamp of BarRequest is required")
	}

	if !enumFieldIsSet {
		return errors.New("field EnumField of BarRequest is required")
	}

	if !longFieldIsSet {
		return errors.New("field LongField of BarRequest is required")
	}

	return nil
}

// String returns a readable string representation of a BarRequest
// struct.
func (v *BarRequest) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++
	fields[i] = fmt.Sprintf("BoolField: %v", v.BoolField)
	i++
	fields[i] = fmt.Sprintf("BinaryField: %v", v.BinaryField)
	i++
	fields[i] = fmt.Sprintf("Timestamp: %v", v.Timestamp)
	i++
	fields[i] = fmt.Sprintf("EnumField: %v", v.EnumField)
	i++
	fields[i] = fmt.Sprintf("LongField: %v", v.LongField)
	i++

	return fmt.Sprintf("BarRequest{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this BarRequest match the
// provided BarRequest.
//
// This function performs a deep comparison.
func (v *BarRequest) Equals(rhs *BarRequest) bool {
	if !(v.StringField == rhs.StringField) {
		return false
	}
	if !(v.BoolField == rhs.BoolField) {
		return false
	}
	if !bytes.Equal(v.BinaryField, rhs.BinaryField) {
		return false
	}
	if !(v.Timestamp == rhs.Timestamp) {
		return false
	}
	if !v.EnumField.Equals(rhs.EnumField) {
		return false
	}
	if !(v.LongField == rhs.LongField) {
		return false
	}

	return true
}

type BarRequestRecur struct {
	Name  string           `json:"name,required"`
	Recur *BarRequestRecur `json:"recur,omitempty"`
}

// ToWire translates a BarRequestRecur struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *BarRequestRecur) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Recur != nil {
		w, err = v.Recur.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarRequestRecur_Read(w wire.Value) (*BarRequestRecur, error) {
	var v BarRequestRecur
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a BarRequestRecur struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BarRequestRecur struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BarRequestRecur
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BarRequestRecur) FromWire(w wire.Value) error {
	var err error

	nameIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Recur, err = _BarRequestRecur_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !nameIsSet {
		return errors.New("field Name of BarRequestRecur is required")
	}

	return nil
}

// String returns a readable string representation of a BarRequestRecur
// struct.
func (v *BarRequestRecur) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	if v.Recur != nil {
		fields[i] = fmt.Sprintf("Recur: %v", v.Recur)
		i++
	}

	return fmt.Sprintf("BarRequestRecur{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this BarRequestRecur match the
// provided BarRequestRecur.
//
// This function performs a deep comparison.
func (v *BarRequestRecur) Equals(rhs *BarRequestRecur) bool {
	if !(v.Name == rhs.Name) {
		return false
	}
	if !((v.Recur == nil && rhs.Recur == nil) || (v.Recur != nil && rhs.Recur != nil && v.Recur.Equals(rhs.Recur))) {
		return false
	}

	return true
}

type BarResponse struct {
	StringField        string           `json:"stringField,required"`
	IntWithRange       int32            `json:"intWithRange,required"`
	IntWithoutRange    int32            `json:"intWithoutRange,required"`
	MapIntWithRange    map[UUID]int32   `json:"mapIntWithRange,required"`
	MapIntWithoutRange map[string]int32 `json:"mapIntWithoutRange,required"`
	BinaryField        []byte           `json:"binaryField,required"`
}

type _Map_UUID_I32_MapItemList map[UUID]int32

func (m _Map_UUID_I32_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := k.ToWire()
		if err != nil {
			return err
		}

		vw, err := wire.NewValueI32(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_UUID_I32_MapItemList) Size() int {
	return len(m)
}

func (_Map_UUID_I32_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_UUID_I32_MapItemList) ValueType() wire.Type {
	return wire.TI32
}

func (_Map_UUID_I32_MapItemList) Close() {}

type _Map_String_I32_MapItemList map[string]int32

func (m _Map_String_I32_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := wire.NewValueI32(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_I32_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_I32_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_I32_MapItemList) ValueType() wire.Type {
	return wire.TI32
}

func (_Map_String_I32_MapItemList) Close() {}

// ToWire translates a BarResponse struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *BarResponse) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueI32(v.IntWithRange), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	w, err = wire.NewValueI32(v.IntWithoutRange), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.MapIntWithRange == nil {
		return w, errors.New("field MapIntWithRange of BarResponse is required")
	}
	w, err = wire.NewValueMap(_Map_UUID_I32_MapItemList(v.MapIntWithRange)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 4, Value: w}
	i++
	if v.MapIntWithoutRange == nil {
		return w, errors.New("field MapIntWithoutRange of BarResponse is required")
	}
	w, err = wire.NewValueMap(_Map_String_I32_MapItemList(v.MapIntWithoutRange)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++
	if v.BinaryField == nil {
		return w, errors.New("field BinaryField of BarResponse is required")
	}
	w, err = wire.NewValueBinary(v.BinaryField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 6, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UUID_Read(w wire.Value) (UUID, error) {
	var x UUID
	err := x.FromWire(w)
	return x, err
}

func _Map_UUID_I32_Read(m wire.MapItemList) (map[UUID]int32, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TI32 {
		return nil, nil
	}

	o := make(map[UUID]int32, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _UUID_Read(x.Key)
		if err != nil {
			return err
		}

		v, err := x.Value.GetI32(), error(nil)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func _Map_String_I32_Read(m wire.MapItemList) (map[string]int32, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TI32 {
		return nil, nil
	}

	o := make(map[string]int32, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}

		v, err := x.Value.GetI32(), error(nil)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a BarResponse struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a BarResponse struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v BarResponse
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *BarResponse) FromWire(w wire.Value) error {
	var err error

	stringFieldIsSet := false
	intWithRangeIsSet := false
	intWithoutRangeIsSet := false
	mapIntWithRangeIsSet := false
	mapIntWithoutRangeIsSet := false
	binaryFieldIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				v.IntWithRange, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				intWithRangeIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TI32 {
				v.IntWithoutRange, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				intWithoutRangeIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TMap {
				v.MapIntWithRange, err = _Map_UUID_I32_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				mapIntWithRangeIsSet = true
			}
		case 5:
			if field.Value.Type() == wire.TMap {
				v.MapIntWithoutRange, err = _Map_String_I32_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				mapIntWithoutRangeIsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TBinary {
				v.BinaryField, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
				binaryFieldIsSet = true
			}
		}
	}

	if !stringFieldIsSet {
		return errors.New("field StringField of BarResponse is required")
	}

	if !intWithRangeIsSet {
		return errors.New("field IntWithRange of BarResponse is required")
	}

	if !intWithoutRangeIsSet {
		return errors.New("field IntWithoutRange of BarResponse is required")
	}

	if !mapIntWithRangeIsSet {
		return errors.New("field MapIntWithRange of BarResponse is required")
	}

	if !mapIntWithoutRangeIsSet {
		return errors.New("field MapIntWithoutRange of BarResponse is required")
	}

	if !binaryFieldIsSet {
		return errors.New("field BinaryField of BarResponse is required")
	}

	return nil
}

// String returns a readable string representation of a BarResponse
// struct.
func (v *BarResponse) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++
	fields[i] = fmt.Sprintf("IntWithRange: %v", v.IntWithRange)
	i++
	fields[i] = fmt.Sprintf("IntWithoutRange: %v", v.IntWithoutRange)
	i++
	fields[i] = fmt.Sprintf("MapIntWithRange: %v", v.MapIntWithRange)
	i++
	fields[i] = fmt.Sprintf("MapIntWithoutRange: %v", v.MapIntWithoutRange)
	i++
	fields[i] = fmt.Sprintf("BinaryField: %v", v.BinaryField)
	i++

	return fmt.Sprintf("BarResponse{%v}", strings.Join(fields[:i], ", "))
}

func _Map_UUID_I32_Equals(lhs, rhs map[UUID]int32) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

func _Map_String_I32_Equals(lhs, rhs map[string]int32) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this BarResponse match the
// provided BarResponse.
//
// This function performs a deep comparison.
func (v *BarResponse) Equals(rhs *BarResponse) bool {
	if !(v.StringField == rhs.StringField) {
		return false
	}
	if !(v.IntWithRange == rhs.IntWithRange) {
		return false
	}
	if !(v.IntWithoutRange == rhs.IntWithoutRange) {
		return false
	}
	if !_Map_UUID_I32_Equals(v.MapIntWithRange, rhs.MapIntWithRange) {
		return false
	}
	if !_Map_String_I32_Equals(v.MapIntWithoutRange, rhs.MapIntWithoutRange) {
		return false
	}
	if !bytes.Equal(v.BinaryField, rhs.BinaryField) {
		return false
	}

	return true
}

type Fruit int32

const (
	FruitApple  Fruit = 0
	FruitBanana Fruit = 1
)

// Fruit_Values returns all recognized values of Fruit.
func Fruit_Values() []Fruit {
	return []Fruit{
		FruitApple,
		FruitBanana,
	}
}

// UnmarshalText tries to decode Fruit from a byte slice
// containing its name.
//
//   var v Fruit
//   err := v.UnmarshalText([]byte("APPLE"))
func (v *Fruit) UnmarshalText(value []byte) error {
	switch string(value) {
	case "APPLE":
		*v = FruitApple
		return nil
	case "BANANA":
		*v = FruitBanana
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "Fruit")
	}
}

// ToWire translates Fruit into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v Fruit) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes Fruit from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return Fruit(0), err
//   }
//
//   var v Fruit
//   if err := v.FromWire(x); err != nil {
//     return Fruit(0), err
//   }
//   return v, nil
func (v *Fruit) FromWire(w wire.Value) error {
	*v = (Fruit)(w.GetI32())
	return nil
}

// String returns a readable string representation of Fruit.
func (v Fruit) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "APPLE"
	case 1:
		return "BANANA"
	}
	return fmt.Sprintf("Fruit(%d)", w)
}

// Equals returns true if this Fruit value matches the provided
// value.
func (v Fruit) Equals(rhs Fruit) bool {
	return v == rhs
}

// MarshalJSON serializes Fruit into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v Fruit) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"APPLE\""), nil
	case 1:
		return ([]byte)("\"BANANA\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode Fruit from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *Fruit) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "Fruit")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "Fruit")
		}
		*v = (Fruit)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "Fruit")
	}
}

type Long int64

// ToWire translates Long into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v Long) ToWire() (wire.Value, error) {
	x := (int64)(v)
	return wire.NewValueI64(x), error(nil)
}

// String returns a readable string representation of Long.
func (v Long) String() string {
	x := (int64)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes Long from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *Long) FromWire(w wire.Value) error {
	x, err := w.GetI64(), error(nil)
	*v = (Long)(x)
	return err
}

// Equals returns true if this Long is equal to the provided
// Long.
func (lhs Long) Equals(rhs Long) bool {
	return (lhs == rhs)
}

type ParamsStruct struct {
	UserUUID string `json:"-"`
}

// ToWire translates a ParamsStruct struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ParamsStruct) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.UserUUID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ParamsStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ParamsStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ParamsStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ParamsStruct) FromWire(w wire.Value) error {
	var err error

	userUUIDIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.UserUUID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				userUUIDIsSet = true
			}
		}
	}

	if !userUUIDIsSet {
		return errors.New("field UserUUID of ParamsStruct is required")
	}

	return nil
}

// String returns a readable string representation of a ParamsStruct
// struct.
func (v *ParamsStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("UserUUID: %v", v.UserUUID)
	i++

	return fmt.Sprintf("ParamsStruct{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ParamsStruct match the
// provided ParamsStruct.
//
// This function performs a deep comparison.
func (v *ParamsStruct) Equals(rhs *ParamsStruct) bool {
	if !(v.UserUUID == rhs.UserUUID) {
		return false
	}

	return true
}

type QueryParamsOptsStruct struct {
	Name      string  `json:"name,required"`
	UserUUID  *string `json:"userUUID,omitempty"`
	AuthUUID  *string `json:"authUUID,omitempty"`
	AuthUUID2 *string `json:"authUUID2,omitempty"`
}

// ToWire translates a QueryParamsOptsStruct struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *QueryParamsOptsStruct) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.UserUUID != nil {
		w, err = wire.NewValueString(*(v.UserUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.AuthUUID != nil {
		w, err = wire.NewValueString(*(v.AuthUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.AuthUUID2 != nil {
		w, err = wire.NewValueString(*(v.AuthUUID2)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a QueryParamsOptsStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a QueryParamsOptsStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v QueryParamsOptsStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *QueryParamsOptsStruct) FromWire(w wire.Value) error {
	var err error

	nameIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UserUUID = &x
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.AuthUUID = &x
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.AuthUUID2 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !nameIsSet {
		return errors.New("field Name of QueryParamsOptsStruct is required")
	}

	return nil
}

// String returns a readable string representation of a QueryParamsOptsStruct
// struct.
func (v *QueryParamsOptsStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	if v.UserUUID != nil {
		fields[i] = fmt.Sprintf("UserUUID: %v", *(v.UserUUID))
		i++
	}
	if v.AuthUUID != nil {
		fields[i] = fmt.Sprintf("AuthUUID: %v", *(v.AuthUUID))
		i++
	}
	if v.AuthUUID2 != nil {
		fields[i] = fmt.Sprintf("AuthUUID2: %v", *(v.AuthUUID2))
		i++
	}

	return fmt.Sprintf("QueryParamsOptsStruct{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this QueryParamsOptsStruct match the
// provided QueryParamsOptsStruct.
//
// This function performs a deep comparison.
func (v *QueryParamsOptsStruct) Equals(rhs *QueryParamsOptsStruct) bool {
	if !(v.Name == rhs.Name) {
		return false
	}
	if !_String_EqualsPtr(v.UserUUID, rhs.UserUUID) {
		return false
	}
	if !_String_EqualsPtr(v.AuthUUID, rhs.AuthUUID) {
		return false
	}
	if !_String_EqualsPtr(v.AuthUUID2, rhs.AuthUUID2) {
		return false
	}

	return true
}

// GetUserUUID returns the value of UserUUID if it is set or its
// zero value if it is unset.
func (v *QueryParamsOptsStruct) GetUserUUID() (o string) {
	if v.UserUUID != nil {
		return *v.UserUUID
	}

	return
}

// GetAuthUUID returns the value of AuthUUID if it is set or its
// zero value if it is unset.
func (v *QueryParamsOptsStruct) GetAuthUUID() (o string) {
	if v.AuthUUID != nil {
		return *v.AuthUUID
	}

	return
}

// GetAuthUUID2 returns the value of AuthUUID2 if it is set or its
// zero value if it is unset.
func (v *QueryParamsOptsStruct) GetAuthUUID2() (o string) {
	if v.AuthUUID2 != nil {
		return *v.AuthUUID2
	}

	return
}

type QueryParamsStruct struct {
	Name      string  `json:"name,required"`
	UserUUID  *string `json:"userUUID,omitempty"`
	AuthUUID  *string `json:"authUUID,omitempty"`
	AuthUUID2 *string `json:"authUUID2,omitempty"`
}

// ToWire translates a QueryParamsStruct struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *QueryParamsStruct) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.UserUUID != nil {
		w, err = wire.NewValueString(*(v.UserUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.AuthUUID != nil {
		w, err = wire.NewValueString(*(v.AuthUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.AuthUUID2 != nil {
		w, err = wire.NewValueString(*(v.AuthUUID2)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a QueryParamsStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a QueryParamsStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v QueryParamsStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *QueryParamsStruct) FromWire(w wire.Value) error {
	var err error

	nameIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UserUUID = &x
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.AuthUUID = &x
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.AuthUUID2 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !nameIsSet {
		return errors.New("field Name of QueryParamsStruct is required")
	}

	return nil
}

// String returns a readable string representation of a QueryParamsStruct
// struct.
func (v *QueryParamsStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	if v.UserUUID != nil {
		fields[i] = fmt.Sprintf("UserUUID: %v", *(v.UserUUID))
		i++
	}
	if v.AuthUUID != nil {
		fields[i] = fmt.Sprintf("AuthUUID: %v", *(v.AuthUUID))
		i++
	}
	if v.AuthUUID2 != nil {
		fields[i] = fmt.Sprintf("AuthUUID2: %v", *(v.AuthUUID2))
		i++
	}

	return fmt.Sprintf("QueryParamsStruct{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this QueryParamsStruct match the
// provided QueryParamsStruct.
//
// This function performs a deep comparison.
func (v *QueryParamsStruct) Equals(rhs *QueryParamsStruct) bool {
	if !(v.Name == rhs.Name) {
		return false
	}
	if !_String_EqualsPtr(v.UserUUID, rhs.UserUUID) {
		return false
	}
	if !_String_EqualsPtr(v.AuthUUID, rhs.AuthUUID) {
		return false
	}
	if !_String_EqualsPtr(v.AuthUUID2, rhs.AuthUUID2) {
		return false
	}

	return true
}

// GetUserUUID returns the value of UserUUID if it is set or its
// zero value if it is unset.
func (v *QueryParamsStruct) GetUserUUID() (o string) {
	if v.UserUUID != nil {
		return *v.UserUUID
	}

	return
}

// GetAuthUUID returns the value of AuthUUID if it is set or its
// zero value if it is unset.
func (v *QueryParamsStruct) GetAuthUUID() (o string) {
	if v.AuthUUID != nil {
		return *v.AuthUUID
	}

	return
}

// GetAuthUUID2 returns the value of AuthUUID2 if it is set or its
// zero value if it is unset.
func (v *QueryParamsStruct) GetAuthUUID2() (o string) {
	if v.AuthUUID2 != nil {
		return *v.AuthUUID2
	}

	return
}

type Timestamp int64

// ToWire translates Timestamp into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v Timestamp) ToWire() (wire.Value, error) {
	x := (int64)(v)
	return wire.NewValueI64(x), error(nil)
}

// String returns a readable string representation of Timestamp.
func (v Timestamp) String() string {
	x := (int64)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes Timestamp from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *Timestamp) FromWire(w wire.Value) error {
	x, err := w.GetI64(), error(nil)
	*v = (Timestamp)(x)
	return err
}

// Equals returns true if this Timestamp is equal to the provided
// Timestamp.
func (lhs Timestamp) Equals(rhs Timestamp) bool {
	return (lhs == rhs)
}

type UUID string

// ToWire translates UUID into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v UUID) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

// String returns a readable string representation of UUID.
func (v UUID) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes UUID from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *UUID) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (UUID)(x)
	return err
}

// Equals returns true if this UUID is equal to the provided
// UUID.
func (lhs UUID) Equals(rhs UUID) bool {
	return (lhs == rhs)
}

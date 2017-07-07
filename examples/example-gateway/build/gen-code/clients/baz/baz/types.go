// Code generated by thriftrw v1.3.0
// @generated

package baz

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

type AuthErr struct {
	Message string `json:"message,required"`
}

func (v *AuthErr) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *AuthErr) FromWire(w wire.Value) error {
	var err error
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !messageIsSet {
		return errors.New("field Message of AuthErr is required")
	}
	return nil
}

func (v *AuthErr) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("AuthErr{%v}", strings.Join(fields[:i], ", "))
}

func (v *AuthErr) Equals(rhs *AuthErr) bool {
	if !(v.Message == rhs.Message) {
		return false
	}
	return true
}

func (v *AuthErr) Error() string {
	return v.String()
}

type BazRequest struct {
	B1 bool   `json:"b1,required"`
	S2 string `json:"s2,required"`
	I3 int32  `json:"i3,required"`
}

func (v *BazRequest) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueBool(v.B1), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.S2), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	w, err = wire.NewValueI32(v.I3), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *BazRequest) FromWire(w wire.Value) error {
	var err error
	b1IsSet := false
	s2IsSet := false
	i3IsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				v.B1, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				b1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.S2, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				s2IsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TI32 {
				v.I3, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				i3IsSet = true
			}
		}
	}
	if !b1IsSet {
		return errors.New("field B1 of BazRequest is required")
	}
	if !s2IsSet {
		return errors.New("field S2 of BazRequest is required")
	}
	if !i3IsSet {
		return errors.New("field I3 of BazRequest is required")
	}
	return nil
}

func (v *BazRequest) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("B1: %v", v.B1)
	i++
	fields[i] = fmt.Sprintf("S2: %v", v.S2)
	i++
	fields[i] = fmt.Sprintf("I3: %v", v.I3)
	i++
	return fmt.Sprintf("BazRequest{%v}", strings.Join(fields[:i], ", "))
}

func (v *BazRequest) Equals(rhs *BazRequest) bool {
	if !(v.B1 == rhs.B1) {
		return false
	}
	if !(v.S2 == rhs.S2) {
		return false
	}
	if !(v.I3 == rhs.I3) {
		return false
	}
	return true
}

type Fruit int32

const (
	FruitApple  Fruit = 0
	FruitBanana Fruit = 1
)

func (v Fruit) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *Fruit) FromWire(w wire.Value) error {
	*v = (Fruit)(w.GetI32())
	return nil
}

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

func (v Fruit) Equals(rhs Fruit) bool {
	return v == rhs
}

func (v Fruit) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"APPLE\""), nil
	case 1:
		return ([]byte)("\"BANANA\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

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
		switch w {
		case "APPLE":
			*v = FruitApple
			return nil
		case "BANANA":
			*v = FruitBanana
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "Fruit")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "Fruit")
	}
}

type OtherAuthErr struct {
	Message string `json:"message,required"`
}

func (v *OtherAuthErr) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *OtherAuthErr) FromWire(w wire.Value) error {
	var err error
	messageIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		}
	}
	if !messageIsSet {
		return errors.New("field Message of OtherAuthErr is required")
	}
	return nil
}

func (v *OtherAuthErr) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	return fmt.Sprintf("OtherAuthErr{%v}", strings.Join(fields[:i], ", "))
}

func (v *OtherAuthErr) Equals(rhs *OtherAuthErr) bool {
	if !(v.Message == rhs.Message) {
		return false
	}
	return true
}

func (v *OtherAuthErr) Error() string {
	return v.String()
}

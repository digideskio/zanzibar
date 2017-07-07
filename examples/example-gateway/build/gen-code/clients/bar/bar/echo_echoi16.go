// Code generated by thriftrw v1.3.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Echo_EchoI16_Args struct {
	Arg int16 `json:"arg,required"`
}

func (v *Echo_EchoI16_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueI16(v.Arg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Echo_EchoI16_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI16 {
				v.Arg, err = field.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of Echo_EchoI16_Args is required")
	}
	return nil
}

func (v *Echo_EchoI16_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("Echo_EchoI16_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_EchoI16_Args) Equals(rhs *Echo_EchoI16_Args) bool {
	if !(v.Arg == rhs.Arg) {
		return false
	}
	return true
}

func (v *Echo_EchoI16_Args) MethodName() string {
	return "echoI16"
}

func (v *Echo_EchoI16_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Echo_EchoI16_Helper = struct {
	Args           func(arg int16) *Echo_EchoI16_Args
	IsException    func(error) bool
	WrapResponse   func(int16, error) (*Echo_EchoI16_Result, error)
	UnwrapResponse func(*Echo_EchoI16_Result) (int16, error)
}{}

func init() {
	Echo_EchoI16_Helper.Args = func(arg int16) *Echo_EchoI16_Args {
		return &Echo_EchoI16_Args{Arg: arg}
	}
	Echo_EchoI16_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Echo_EchoI16_Helper.WrapResponse = func(success int16, err error) (*Echo_EchoI16_Result, error) {
		if err == nil {
			return &Echo_EchoI16_Result{Success: &success}, nil
		}
		return nil, err
	}
	Echo_EchoI16_Helper.UnwrapResponse = func(result *Echo_EchoI16_Result) (success int16, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Echo_EchoI16_Result struct {
	Success *int16 `json:"success,omitempty"`
}

func (v *Echo_EchoI16_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueI16(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_EchoI16_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Echo_EchoI16_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI16 {
				var x int16
				x, err = field.Value.GetI16(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Echo_EchoI16_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Echo_EchoI16_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("Echo_EchoI16_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_EchoI16_Result) Equals(rhs *Echo_EchoI16_Result) bool {
	if !_I16_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	return true
}

func (v *Echo_EchoI16_Result) MethodName() string {
	return "echoI16"
}

func (v *Echo_EchoI16_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

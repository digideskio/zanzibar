// Code generated by zanzibar
// @generated
// Checksum : G8LkrERLNd6dIPSd4SCAtQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package foo

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/base/base"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(in *jlexer.Lexer, out *FooStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var FooStringSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fooString":
			out.FooString = string(in.String())
			FooStringSet = true
		case "fooI32":
			if in.IsNull() {
				in.Skip()
				out.FooI32 = nil
			} else {
				if out.FooI32 == nil {
					out.FooI32 = new(int32)
				}
				*out.FooI32 = int32(in.Int32())
			}
		case "fooI16":
			if in.IsNull() {
				in.Skip()
				out.FooI16 = nil
			} else {
				if out.FooI16 == nil {
					out.FooI16 = new(int16)
				}
				*out.FooI16 = int16(in.Int16())
			}
		case "fooDouble":
			if in.IsNull() {
				in.Skip()
				out.FooDouble = nil
			} else {
				if out.FooDouble == nil {
					out.FooDouble = new(float64)
				}
				*out.FooDouble = float64(in.Float64())
			}
		case "fooBool":
			if in.IsNull() {
				in.Skip()
				out.FooBool = nil
			} else {
				if out.FooBool == nil {
					out.FooBool = new(bool)
				}
				*out.FooBool = bool(in.Bool())
			}
		case "fooMap":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.FooMap = make(map[string]string)
				} else {
					out.FooMap = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.FooMap)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(base.Message)
				}
				(*out.Message).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !FooStringSet {
		in.AddError(fmt.Errorf("key 'fooString' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(out *jwriter.Writer, in FooStruct) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fooString\":")
	out.String(string(in.FooString))
	if in.FooI32 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooI32\":")
		if in.FooI32 == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.FooI32))
		}
	}
	if in.FooI16 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooI16\":")
		if in.FooI16 == nil {
			out.RawString("null")
		} else {
			out.Int16(int16(*in.FooI16))
		}
	}
	if in.FooDouble != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooDouble\":")
		if in.FooDouble == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.FooDouble))
		}
	}
	if in.FooBool != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooBool\":")
		if in.FooBool == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.FooBool))
		}
	}
	if len(in.FooMap) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"fooMap\":")
		if in.FooMap == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.FooMap {
				if !v2First {
					out.RawByte(',')
				}
				v2First = false
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	if in.Message != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"message\":")
		if in.Message == nil {
			out.RawString("null")
		} else {
			(*in.Message).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FooStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FooStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FooStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FooStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo1(in *jlexer.Lexer, out *FooName) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo1(out *jwriter.Writer, in FooName) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FooName) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FooName) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FooName) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FooName) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo1(l, v)
}
func easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo2(in *jlexer.Lexer, out *FooException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var TeapotSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "teapot":
			out.Teapot = string(in.String())
			TeapotSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !TeapotSet {
		in.AddError(fmt.Errorf("key 'teapot' is required"))
	}
}
func easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo2(out *jwriter.Writer, in FooException) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"teapot\":")
	out.String(string(in.Teapot))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FooException) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FooException) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FooException) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FooException) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsFooFoo2(l, v)
}

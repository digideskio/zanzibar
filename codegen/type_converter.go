// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package codegen

import (
	"fmt"
	"strings"

	"bytes"
	"github.com/pkg/errors"
	"go.uber.org/thriftrw/compile"
)

// LineBuilder struct just appends/builds lines.
type LineBuilder struct {
	lines []string
}

// append helper will add a line to TypeConverter
func (l *LineBuilder) append(parts ...string) {
	line := strings.Join(parts, "")
	l.lines = append(l.lines, line)
}

// appendf helper will add a formatted line to TypeConverter
func (l *LineBuilder) appendf(format string, parts ...interface{}) {
	line := fmt.Sprintf(format, parts...)
	l.lines = append(l.lines, line)
}

// GetLines returns the lines in the line builder
func (l *LineBuilder) GetLines() []string {
	return l.lines
}

type FieldAssignment struct {
	From           *compile.FieldSpec
	FromIdentifier string
	To             *compile.FieldSpec
	ToIdentifier   string
	TypeCast       string
	ExtraNilCheck  bool
	buf            *bytes.Buffer
}

func (f *FieldAssignment) IsTransform() bool {
	if f.FromIdentifier[:3] != "in." || f.ToIdentifier[:4] != "out." {
		panic("isTransform check called on unexpected input  fromIdentifer " + f.FromIdentifier + " toIdentifer " + f.ToIdentifier)
	}
	return f.FromIdentifier[3:] != f.ToIdentifier[4:]
}

func (f *FieldAssignment) append(items ...string) {
	for _, item := range items {
		_, _ = f.buf.WriteString(item)
	}
	f.buf.WriteString("\n")
}

func (f *FieldAssignment) Generate(indent string) string {
	f.buf = new(bytes.Buffer)

	if !f.IsTransform() {
		//  do an extra nil check for Overrides
		if !f.From.Required && f.To.Required || f.ExtraNilCheck {
			// need to nil check otherwise we could be cast or deref nil
			f.append(indent, "if ", f.FromIdentifier, " != nil {")
			f.append(indent, "\t", f.ToIdentifier, " = ", f.TypeCast, "(", f.FromIdentifier, ")")
			f.append(indent, "}")
		} else {
			f.append(indent, f.ToIdentifier, " = ", f.TypeCast, "(", f.FromIdentifier, ")")
		}
		return f.buf.String()
	}

	//  generates nil checks for intermediate objects on the f.FromIdentifier path
	checks := convertIdentifiersToNilChecks(getMiddleIdentifiers(f.FromIdentifier))

	if !f.ExtraNilCheck {
		// if from is required or from/to are both optional   then we don't need the final outer check
		if f.From.Required || !f.From.Required && !f.To.Required {
			checks = checks[:len(checks)-1]
		}
		if len(checks) == 0 {
			f.append(indent, f.ToIdentifier, " = ", f.TypeCast, "(", f.FromIdentifier, ")")
			return f.buf.String()
		}
	}
	f.append(indent, "if ", strings.Join(checks, " && "), " {")
	f.append(indent, "\t", f.ToIdentifier, " = ", f.TypeCast, "(", f.FromIdentifier, ")")
	f.append(indent, "}")
	// TODO else?  log? should this silently eat intermediate nils as none-assignment,  should set nil?
	return f.buf.String()
}

// PackageNameResolver interface allows for resolving what the
// package name for a thrift file is. This depends on where the
// thrift-based structs are generated.
type PackageNameResolver interface {
	TypePackageName(thriftFile string) (string, error)
}

// TypeConverter can generate a function body that converts two thriftrw
// FieldGroups from one to another. It's assumed that the converted code
// operates on two variables, "in" and "out" and that both are a go struct.
type TypeConverter struct {
	LineBuilder
	Helper PackageNameResolver
}

func (c *TypeConverter) getGoTypeName(valueType compile.TypeSpec) (string, error) {
	return GoType(c.Helper, valueType)
}

//   input of "A.B.C.D" returns ["A","A.B", "A.B.C", "A.B.C.D"]
func getMiddleIdentifiers(identifier string) []string {
	subIds := strings.Split(identifier, ".")

	middleIds := make([]string, 0, len(subIds))

	middleIds = append(middleIds, subIds[0])
	for i := 1; i < len(subIds); i++ {
		middleIds = append(middleIds, fmt.Sprintf("%s.%s", middleIds[i-1], subIds[i]))
	}

	return middleIds
}

//  converts a list of identifier paths into boolean nil check expressions on those paths
func convertIdentifiersToNilChecks(identifiers []string) []string {
	checks := make([]string, 0, len(identifiers)-1)

	for i, ident := range identifiers {
		if i == 0 {
			continue // discard the check on "in" object
		}

		checks = append(checks, ident+" != nil")
	}
	return checks
}

//  tranforms will have different paths,   to->from proxy mappings will have the same identifier path aside from prefix
func isTransform(from string, to string) bool {
	if from[:3] != "in." || to[:4] != "out." {
		panic("isTransform check called on unexpected input")
	}
	return from[3:] != to[4:]
}

// helper func for assign WithOverride
func (c *TypeConverter) assignWithChecks(
	indent string,
	toIdentifier string,
	toTypeName string,
	toRequired bool,
	fromIdentifier string,
	fromRequired bool,
	extraNilCheck bool,
) {

	if !isTransform(fromIdentifier, toIdentifier) {

		if !fromRequired && toRequired || extraNilCheck {
			// need to nil check otherwise we could be cast or deref nil
			c.append(indent, "if ", fromIdentifier, " != nil {")
			c.append(indent, "\t", toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
			c.append(indent, "}")
		} else {
			c.append(indent, toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
		}
		return
	}

	//  generates nil checks for intermediate objects on the fromIdentifier path
	checks := convertIdentifiersToNilChecks(getMiddleIdentifiers(fromIdentifier))

	if !extraNilCheck {
		// if from is required or from/to are both optional   then we don't need the final outer check
		if fromRequired || !fromRequired && !toRequired {
			checks = checks[:len(checks)-1]
		}
		if len(checks) == 0 {
			c.append(indent, toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
			return
		}
	}
	c.append(indent, "if ", strings.Join(checks, " && "), " {")
	c.append(indent, "\t", toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
	c.append(indent, "}")
	// TODO else?  log/stat it?  should this siliently eat intermediate nils as none-assignement,  should set nil?
}

// helper func for assign WithOverride
func (c *TypeConverter) assignWithChecks2(
	indent string,

	toIdentifier string,
	toTypeName string,
	toRequired bool,
	fromIdentifier string,
	fromRequired bool,
	extraNilCheck bool,
) {

	if !isTransform(fromIdentifier, toIdentifier) {

		if !fromRequired && toRequired || extraNilCheck {
			// need to nil check otherwise we could be cast or deref nil
			c.append(indent, "if ", fromIdentifier, " != nil {")
			c.append(indent, "\t", toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
			c.append(indent, "}")
		} else {
			c.append(indent, toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
		}
		return
	}

	//  generates nil checks for intermediate objects on the fromIdentifier path
	checks := convertIdentifiersToNilChecks(getMiddleIdentifiers(fromIdentifier))

	if !extraNilCheck {
		// if from is required or from/to are both optional   then we don't need the final outer check
		if fromRequired || !fromRequired && !toRequired {
			checks = checks[:len(checks)-1]
		}
		if len(checks) == 0 {
			c.append(indent, toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
			return
		}
	}
	c.append(indent, "if ", strings.Join(checks, " && "), " {")
	c.append(indent, "\t", toIdentifier, " = ", toTypeName, "(", fromIdentifier, ")")
	c.append(indent, "}")
	// TODO else?  log/stat it?  should this siliently eat intermediate nils as none-assignement,  should set nil?
}

func (c *TypeConverter) assignWithOverride2(
	indent string,
	defaultAssign *FieldAssignment,
	overrideAssign *FieldAssignment,
) {
	if defaultAssign.FromIdentifier != "" {
		c.append(defaultAssign.Generate(indent))
		if overrideAssign != nil {
			overrideAssign.ExtraNilCheck = true
		}
	}

	if overrideAssign != nil && overrideAssign.FromIdentifier != "" {
		c.append(overrideAssign.Generate(indent))
	}
}

func (c *TypeConverter) assignWithOverride(
	indent string,
	toIdentifier string,
	toTypeCast string,
	toRequired bool,
	fromIdentifier string,
	fromRequired bool,
	overriddenTypeCast string,
	overriddenIdentifier string,
	overrideRequired bool,
) {
	if overriddenIdentifier != "" {
		c.assignWithChecks(indent, toIdentifier, overriddenTypeCast, toRequired, overriddenIdentifier, overrideRequired, false)
		// second assignment uses extraNilCheck so that we do not re-assign to a nil value
		c.assignWithChecks(indent, toIdentifier, toTypeCast, toRequired, fromIdentifier, fromRequired, true)
		return
	}
	c.assignWithChecks(indent, toIdentifier, toTypeCast, toRequired, fromIdentifier, fromRequired, false)
}

func (c *TypeConverter) getIdentifierName(fieldType compile.TypeSpec) (string, error) {
	t, err := goCustomType(c.Helper, fieldType)
	if err != nil {
		return "", errors.Wrapf(
			err,
			"could not lookup fieldType when building converter for %s",
			fieldType.ThriftName(),
		)
	}
	return t, nil
}

func (c *TypeConverter) genConverterForStruct(
	toFieldName string,
	toFieldType *compile.StructSpec,
	fromFieldType compile.TypeSpec,
	fromIdentifier string,
	keyPrefix string,
	fromPrefix string,
	indent string,
	fieldMap map[string]FieldMapperEntry,
) error {
	toIdentifier := "out." + keyPrefix

	typeName, err := c.getIdentifierName(toFieldType)
	if err != nil {
		return err
	}
	subToFields := toFieldType.Fields

	// if no fromFieldType assume we're constructing from transform fieldMap  TODO: make this less subtle
	if fromFieldType == nil {
		//  in the direct assignment we do a nil check on fromField here. its hard for unknown number of transforms
		//  initialize the toField with an empty struct
		c.append(indent, toIdentifier, " = &", typeName, "{}")

		// recursive call
		err = c.genStructConverter(
			keyPrefix+".",
			fromPrefix+".",
			indent,
			nil,
			subToFields,
			fieldMap,
		)
		if err != nil {
			return err
		}
		return nil
	}

	fromFieldStruct, ok := fromFieldType.(*compile.StructSpec)
	if !ok {
		return errors.Errorf(
			"could not convert struct fields, "+
				"incompatible type for %s :",
			toFieldName,
		)
	}

	c.append(indent, "if ", fromIdentifier, " != nil {")

	c.append(indent, "\t", toIdentifier, " = &", typeName, "{}")

	subFromFields := fromFieldStruct.Fields
	err = c.genStructConverter(
		keyPrefix+".",
		fromPrefix+".",
		indent+"\t",
		subFromFields,
		subToFields,
		fieldMap,
	)
	if err != nil {
		return err
	}

	c.append(indent, "} else {")
	c.append(indent, "\t", toIdentifier, " = nil")
	c.append(indent, "}")

	return nil
}

func (c *TypeConverter) genConverterForPrimitiveOrTypedef(
	indent string,
	toTypeCast string,
	defaultAssignment *FieldAssignment,
	overrideAssignment *FieldAssignment,
) {

	//func (c *TypeConverter) genConverterForPrimitive(
	//	toField *compile.FieldSpec,
	//	toIdentifier string,
	//	fromField *compile.FieldSpec,
	//	fromIdentifier string,
	//	defaultFromField *compile.FieldSpec,
	//	defaultFromIdentifier string,
	//	indent string,
	//) error {
	//	toTypeCast, err := c.getGoTypeName(toField.Type)
	//	if err != nil {
	//		return err
	//	}
	//
	//	defaultAssignment := &FieldAssignment{
	//		From:           defaultFromField,
	//		FromIdentifier: defaultFromIdentifier,
	//		To:             toField,
	//		ToIdentifier:   toIdentifier,
	//		ExtraNilCheck:  false,
	//		TypeCast:       toTypeCast,
	//	}
	//
	//	overrideAssignment := &FieldAssignment{
	//		From:           fromField,
	//		FromIdentifier: fromIdentifier,
	//		To:             toField,
	//		ToIdentifier:   toIdentifier,
	//		ExtraNilCheck:  false,
	//		TypeCast:       toTypeCast,
	//	}

	if defaultAssignment.To.Required {
		if overrideAssignment != nil && !overrideAssignment.From.Required {
			// Dereference the pointer
			overrideAssignment.TypeCast = "*"
		}
		if !defaultAssignment.From.Required {
			// Dereference the pointer
			defaultAssignment.TypeCast = "*"
		}
		c.assignWithOverride2(indent, defaultAssignment, overrideAssignment)

	} else {
		if overrideAssignment != nil {
			if !overrideAssignment.From.Required {
				overrideAssignment.TypeCast = fmt.Sprintf("(*%s)", toTypeCast)
			} else {
				overrideAssignment.TypeCast = fmt.Sprintf("(*%s)&", toTypeCast)
			}
		}
		defaultAssignment.TypeCast = fmt.Sprintf("(*%s)", toTypeCast)
		if defaultAssignment.From.Required {
			defaultAssignment.TypeCast = fmt.Sprintf("(*%s)&", toTypeCast)
		}
		c.assignWithOverride2(indent, defaultAssignment, overrideAssignment)
	}
}

func (c *TypeConverter) genConverterForList2(
	toFieldType *compile.ListSpec,
	toField *compile.FieldSpec,
	toIdentifier string,
	defaultFromField *compile.FieldSpec,
	defaultFromIdentifier string,
	overrideFromField *compile.FieldSpec,
	overrideFromIdentifier string,
	keyPrefix string,
	indent string,
) error {
	typeName, err := c.getGoTypeName(toFieldType.ValueSpec)
	if err != nil {
		return err
	}

	valueStruct, isStruct := toFieldType.ValueSpec.(*compile.StructSpec)
	sourceIdentifier := defaultFromIdentifier
	checkOverride := false
	if overrideFromField != nil {
		// Determine which map (from or overrride) to use
		c.appendf("sourceList := %s", defaultFromIdentifier)
		if isStruct {
			c.appendf("isOverridden := false")
		}
		// TODO(sindelar): Verify how optional thrift lists are defined.
		c.appendf("if %s != nil {", overrideFromIdentifier)

		c.appendf("\tsourceList = %s", overrideFromIdentifier)
		if isStruct {
			c.append("\tisOverridden = true")
		}
		c.append("}")

		sourceIdentifier = "sourceList"
		checkOverride = true
	}

	if isStruct {
		c.appendf(
			"%s = make([]*%s, len(%s))",
			toIdentifier, typeName, sourceIdentifier,
		)
	} else {
		c.appendf(
			"%s = make([]%s, len(%s))",
			toIdentifier, typeName, sourceIdentifier,
		)
	}

	c.append("for index, value := range ", sourceIdentifier, " {")

	if isStruct {
		nestedIndent := "\t" + indent

		if checkOverride {
			nestedIndent = "\t" + nestedIndent
			c.append("\t", "if isOverridden {")

			overrideFromFieldType, ok := overrideFromField.Type.(*compile.ListSpec)
			if !ok {
				return errors.Errorf(
					"Could not convert field (%s): type is not list",
					defaultFromField.Name,
				)
			}

			err = c.genConverterForStruct(
				toField.Name,
				valueStruct,
				overrideFromFieldType.ValueSpec,
				"value",
				keyPrefix+pascalCase(toField.Name)+"[index]",
				strings.TrimPrefix(overrideFromIdentifier, "in.")+"[index]",
				nestedIndent,
				nil,
			)
			if err != nil {
				return err
			}
			c.append("\t", "} else {")
		}

			defaultFieldType, ok := defaultFromField.Type.(*compile.ListSpec)
			if !ok {
				return errors.Errorf(
					"Could not convert field (%s): type is not list",
					defaultFromField.Name,
				)
			}

			err = c.genConverterForStruct(
				toField.Name,
				valueStruct,
				defaultFieldType.ValueSpec,
				"value",
				keyPrefix+pascalCase(toField.Name)+"[index]",
				strings.TrimPrefix(defaultFromIdentifier, "in.")+"[index]",
				nestedIndent,
				nil,
			)
			if err != nil {
				return err
			}
		if checkOverride {
			c.append("\t", "}")
		}
	} else {
		c.append("\t", toIdentifier, "[index] = ", typeName, "(value)")
	}

	c.append("}")
	return nil
}

//func (c *TypeConverter) genConverterForList(
//	toFieldType *compile.ListSpec,
//	toField *compile.FieldSpec,
//	fromField *compile.FieldSpec,
//	overriddenField *compile.FieldSpec,
//	toIdentifier string,
//	fromIdentifier string,
//	overriddenIdentifier string,
//	keyPrefix string,
//	indent string,
//) error {
//	typeName, err := c.getGoTypeName(toFieldType.ValueSpec)
//	if err != nil {
//		return err
//	}
//
//	valueStruct, isStruct := toFieldType.ValueSpec.(*compile.StructSpec)
//	sourceIdentifier := fromIdentifier
//	checkOverride := false
//	if overriddenIdentifier != "" {
//		// Determine which map (from or overrride) to use
//		c.appendf("sourceList := %s", overriddenIdentifier)
//		if isStruct {
//			c.appendf("isOverridden := false")
//		}
//		// TODO(sindelar): Verify how optional thrift lists are defined.
//		c.appendf("if %s != nil {", fromIdentifier)
//
//		c.appendf("\tsourceList = %s", fromIdentifier)
//		if isStruct {
//			c.append("\tisOverridden = true")
//		}
//		c.append("}")
//
//		sourceIdentifier = "sourceList"
//		checkOverride = true
//	}
//
//	if isStruct {
//		c.appendf(
//			"%s = make([]*%s, len(%s))",
//			toIdentifier, typeName, sourceIdentifier,
//		)
//	} else {
//		c.appendf(
//			"%s = make([]%s, len(%s))",
//			toIdentifier, typeName, sourceIdentifier,
//		)
//	}
//
//	c.append("for index, value := range ", sourceIdentifier, " {")
//
//	if isStruct {
//		nestedIndent := "\t" + indent
//
//		if checkOverride {
//			nestedIndent = "\t" + nestedIndent
//			c.append("\t", "if isOverridden {")
//		}
//		fromFieldType, ok := fromField.Type.(*compile.ListSpec)
//		if !ok {
//			return errors.Errorf(
//				"Could not convert field (%s): type is not list",
//				fromField.Name,
//			)
//		}
//
//		err = c.genConverterForStruct(
//			toField.Name,
//			valueStruct,
//			fromFieldType.ValueSpec,
//			"value",
//			keyPrefix+pascalCase(toField.Name)+"[index]",
//			strings.TrimPrefix(fromIdentifier, "in.")+"[index]",
//			nestedIndent,
//			nil,
//		)
//		if err != nil {
//			return err
//		}
//		if checkOverride {
//			c.append("\t", "} else {")
//
//			overriddenFieldType, ok := overriddenField.Type.(*compile.ListSpec)
//			if !ok {
//				return errors.Errorf(
//					"Could not convert field (%s): type is not list",
//					overriddenField.Name,
//				)
//			}
//
//			err = c.genConverterForStruct(
//				toField.Name,
//				valueStruct,
//				overriddenFieldType.ValueSpec,
//				"value",
//				keyPrefix+pascalCase(toField.Name)+"[index]",
//				strings.TrimPrefix(overriddenIdentifier, "in.")+"[index]",
//				nestedIndent,
//				nil,
//			)
//			if err != nil {
//				return err
//			}
//			c.append("\t", "}")
//		}
//	} else {
//		c.append("\t", toIdentifier, "[index] = ", typeName, "(value)")
//	}
//
//	c.append("}")
//	return nil
//}

func (c *TypeConverter) genConverterForMap2(
	toFieldType *compile.MapSpec,
	toField *compile.FieldSpec,
	toIdentifier string,
	defaultFromField *compile.FieldSpec,
	defaultFromIdentifier string,
	overrideFromField *compile.FieldSpec,
	overrideFromIdentifier string,
	keyPrefix string,
	indent string,
) error {
	typeName, err := c.getGoTypeName(toFieldType.ValueSpec)
	if err != nil {
		return err
	}

	valueStruct, isStruct := toFieldType.ValueSpec.(*compile.StructSpec)
	sourceIdentifier := defaultFromIdentifier
	_, isStringKey := toFieldType.KeySpec.(*compile.StringSpec)

	if !isStringKey {
		realType := compile.RootTypeSpec(toFieldType.KeySpec)
		switch realType.(type) {
		case *compile.StringSpec:
			keyType, _ := c.getGoTypeName(toFieldType.KeySpec)
			c.appendf(
				"%s = make(map[%s]%s, len(%s))",
				toIdentifier, keyType, typeName, sourceIdentifier,
			)
			c.appendf("for key, value := range %s {", sourceIdentifier)
			c.append("\t", toIdentifier, "[ ", keyType, "(key)] = ", typeName, "(value)")
			c.append("}")
			return nil
		default:
			return errors.Errorf(
				"could not convert key (%s), map is not string-keyed.", toField.Name)
		}
	}

	checkOverride := false
	if overrideFromIdentifier != "" {
		// Determine which map (from or overrride) to use
		c.appendf("sourceList := %s", defaultFromIdentifier)

		if isStruct {
			c.appendf("isOverridden := false")
		}
		// TODO(sindelar): Verify how optional thrift map are defined.
		c.appendf("if %s != nil {", overrideFromIdentifier)

		c.appendf("\tsourceList = %s", overrideFromIdentifier)
		if isStruct {
			c.append("\tisOverridden = true")
		}
		c.append("}")

		sourceIdentifier = "sourceList"
		checkOverride = true
	}

	if isStruct {
		c.appendf(
			"%s = make(map[string]*%s, len(%s))",
			toIdentifier, typeName, sourceIdentifier,
		)
	} else {
		c.appendf(
			"%s = make(map[string]%s, len(%s))",
			toIdentifier, typeName, sourceIdentifier,
		)
	}

	c.appendf("for key, value := range %s {", sourceIdentifier)

	if isStruct {
		nestedIndent := "\t" + indent

		if checkOverride {
			nestedIndent = "\t" + nestedIndent
			c.append("\t", "if isOverridden {")

			overrideFieldType, ok := overrideFromField.Type.(*compile.MapSpec)
			if !ok {
				return errors.Errorf(
					"Could not convert field (%s): type is not map",
					overrideFromField.Name,
				)
			}

			err = c.genConverterForStruct(
				toField.Name,
				valueStruct,
				overrideFieldType.ValueSpec,
				"value",
				keyPrefix+pascalCase(toField.Name)+"[key]",
				strings.TrimPrefix(overrideFromIdentifier, "in.")+"[key]",
				nestedIndent,
				nil,
			)
			if err != nil {
				return err
			}
			c.append("\t", "} else {")
		}

		fromFieldType, ok := defaultFromField.Type.(*compile.MapSpec)
		if !ok {
			return errors.Errorf(
				"Could not convert field (%s): type is not map",
				defaultFromField.Name,
			)
		}

		err = c.genConverterForStruct(
			toField.Name,
			valueStruct,
			fromFieldType.ValueSpec,
			"value",
			keyPrefix+pascalCase(toField.Name)+"[key]",
			strings.TrimPrefix(defaultFromIdentifier, "in.")+"[key]",
			nestedIndent,
			nil,
		)
		if err != nil {
			return err
		}
		if checkOverride {
			c.append("\t", "}")
		}
	} else {
		c.append("\t", toIdentifier, "[key] = ", typeName, "(value)")
	}

	c.append("}")
	return nil
}

//func (c *TypeConverter) genConverterForMap(
//	toFieldType *compile.MapSpec,
//	toField *compile.FieldSpec,
//	fromField *compile.FieldSpec,
//	overriddenField *compile.FieldSpec,
//	toIdentifier string,
//	fromIdentifier string,
//	overriddenIdentifier string,
//	keyPrefix string,
//	indent string,
//) error {
//	typeName, err := c.getGoTypeName(toFieldType.ValueSpec)
//	if err != nil {
//		return err
//	}
//
//	valueStruct, isStruct := toFieldType.ValueSpec.(*compile.StructSpec)
//	sourceIdentifier := fromIdentifier
//	_, isStringKey := toFieldType.KeySpec.(*compile.StringSpec)
//
//	if !isStringKey {
//		realType := compile.RootTypeSpec(toFieldType.KeySpec)
//		switch realType.(type) {
//		case *compile.StringSpec:
//			keyType, _ := c.getGoTypeName(toFieldType.KeySpec)
//			c.appendf(
//				"%s = make(map[%s]%s, len(%s))",
//				toIdentifier, keyType, typeName, sourceIdentifier,
//			)
//			c.appendf("for key, value := range %s {", sourceIdentifier)
//			c.append("\t", toIdentifier, "[ ", keyType, "(key)] = ", typeName, "(value)")
//			c.append("}")
//			return nil
//		default:
//			return errors.Errorf(
//				"could not convert key (%s), map is not string-keyed.", toField.Name)
//		}
//	}
//
//	checkOverride := false
//	if overriddenIdentifier != "" {
//		// Determine which map (from or overrride) to use
//		c.appendf("sourceList := %s", overriddenIdentifier)
//
//		if isStruct {
//			c.appendf("isOverridden := false")
//		}
//		// TODO(sindelar): Verify how optional thrift map are defined.
//		c.appendf("if %s != nil {", fromIdentifier)
//
//		c.appendf("\tsourceList = %s", fromIdentifier)
//		if isStruct {
//			c.append("\tisOverridden = true")
//		}
//		c.append("}")
//
//		sourceIdentifier = "sourceList"
//		checkOverride = true
//	}
//
//	if isStruct {
//		c.appendf(
//			"%s = make(map[string]*%s, len(%s))",
//			toIdentifier, typeName, sourceIdentifier,
//		)
//	} else {
//		c.appendf(
//			"%s = make(map[string]%s, len(%s))",
//			toIdentifier, typeName, sourceIdentifier,
//		)
//	}
//
//	c.appendf("for key, value := range %s {", sourceIdentifier)
//
//	if isStruct {
//		nestedIndent := "\t" + indent
//
//		if checkOverride {
//			nestedIndent = "\t" + nestedIndent
//			c.append("\t", "if isOverridden {")
//		}
//
//		fromFieldType, ok := fromField.Type.(*compile.MapSpec)
//		if !ok {
//			return errors.Errorf(
//				"Could not convert field (%s): type is not map",
//				fromField.Name,
//			)
//		}
//
//		err = c.genConverterForStruct(
//			toField.Name,
//			valueStruct,
//			fromFieldType.ValueSpec,
//			"value",
//			keyPrefix+pascalCase(toField.Name)+"[key]",
//			strings.TrimPrefix(fromIdentifier, "in.")+"[key]",
//			nestedIndent,
//			nil,
//		)
//		if err != nil {
//			return err
//		}
//
//		if checkOverride {
//			c.append("\t", "} else {")
//
//			overriddenFieldType, ok := overriddenField.Type.(*compile.MapSpec)
//			if !ok {
//				return errors.Errorf(
//					"Could not convert field (%s): type is not map",
//					overriddenField.Name,
//				)
//			}
//
//			err = c.genConverterForStruct(
//				toField.Name,
//				valueStruct,
//				overriddenFieldType.ValueSpec,
//				"value",
//				keyPrefix+pascalCase(toField.Name)+"[key]",
//				strings.TrimPrefix(overriddenIdentifier, "in.")+"[key]",
//				nestedIndent,
//				nil,
//			)
//			if err != nil {
//				return err
//			}
//			c.append("\t", "}")
//		}
//	} else {
//		c.append("\t", toIdentifier, "[key] = ", typeName, "(value)")
//	}
//
//	c.append("}")
//	return nil
//}

// recursive function to walk a DFS on toFields and try to assign fromFields or fieldMap tranforms
//  generated code is appended as we traverse the toFields thrift type structure
//  keyPrefix - the identifier (path) of the current position in the "to" struct
//  fromPrefix - the identifier (path) of the corresponding position in the "from" struct
//  indent - a string of tabs for current block scope
//  fromFields - fields in the current from struct,  can be nil  if only fieldMap transforms are applicable in the path
//  toFields - fields in the current to struct
//  fieldMap - a data structure specifying configured transforms   Map[toIdentifier ] ->  fromField FieldMapperEntry
func (c *TypeConverter) genStructConverter(
	keyPrefix string,
	fromPrefix string,
	indent string,
	fromFields []*compile.FieldSpec,
	toFields []*compile.FieldSpec,
	fieldMap map[string]FieldMapperEntry,
) error {

	for i := 0; i < len(toFields); i++ {
		toField := toFields[i]

		toSubIdentifier := keyPrefix + pascalCase(toField.Name)
		toIdentifier := "out." + toSubIdentifier

		var defaultFromField *compile.FieldSpec
		defaultFromIdentifier := ""

		var overrideFromField *compile.FieldSpec
		overrideFromIdentifier := ""

		for j := 0; j < len(fromFields); j++ {
			if fromFields[j].Name == toField.Name {
				defaultFromField = fromFields[j]
				defaultFromIdentifier = "in." + fromPrefix + pascalCase(defaultFromField.Name)
				break
			}
		}

		// Check for mapped field

		// check if this toField satisfies a fieldMap transform
		transformFrom, ok := fieldMap[toSubIdentifier]
		if ok {
			transformFromIdentifier := "in." + transformFrom.QualifiedName
			// no existing direct fromField,  just assign the transform
			if defaultFromField == nil {
				defaultFromField = transformFrom.Field
				defaultFromIdentifier = transformFromIdentifier
				// else there is a conflicting direct fromField
			} else {
				//  depending on Override flag either the direct fromField or transformFrom is the OverrideField
				if transformFrom.Override {
					// if transform is optional its the overrideFrom
					if !transformFrom.Field.Required {
						overrideFromField = transformFrom.Field
						overrideFromIdentifier = transformFromIdentifier
						// If override is true and the new field is required,
						// there's a default instantiation value and will always
						// overwrite.
					} else {
						defaultFromField = transformFrom.Field
						defaultFromIdentifier = transformFromIdentifier
					}
					// transform not an override
				} else {
					// transform is only relevant if the defaultFrom is optional
					if !defaultFromField.Required {

						// somewhat counterintuitive, but def of override:false means
						// only assign the transform if initial field is not nil
						// to achieve this we swap; make transform default,
						// assign the default as an override if its not nil
						overrideFromField = defaultFromField
						overrideFromIdentifier = defaultFromIdentifier
						defaultFromField = transformFrom.Field
						defaultFromIdentifier = transformFromIdentifier
					}
				}
			}
		}

		// neither direct or transform fromField was found
		if defaultFromField == nil {
			// search the fieldMap toField identifiers for matching identifier prefix
			// e.g.  the current toField is a struct and something within it has a transform
			//   a full match identifiers for transform non-struct types would have been caught above
			hasStructFieldMapping := false
			for toID := range fieldMap {
				if strings.HasPrefix(toID, toSubIdentifier) {
					hasStructFieldMapping = true
				}
			}
			//  if there's no fromField and no fieldMap transform that could be applied
			if !hasStructFieldMapping {
				if toField.Required {
					// unrecoverable error
					return errors.Errorf(
						"required toField %s does not have a valid fromField mapping",
						toField.Name,
					)
				}

				// the toField is optional and there's nothing that maps to it or its sub-fields so we skip it
				continue
			}
		}

		//if fromIdentifier == "" && fromField != nil {
		//	// should we set this if no fromField ??
		//	fromIdentifier = "in." + fromPrefix + pascalCase(fromField.Name)
		//}

		// Override thrift type names to avoid naming collisions between endpoint
		// and client types.
		switch toFieldType := toField.Type.(type) {
		case
			*compile.BoolSpec,
			*compile.I8Spec,
			*compile.I16Spec,
			*compile.I32Spec,
			*compile.EnumSpec,
			*compile.I64Spec,
			*compile.DoubleSpec,
			*compile.StringSpec,
			*compile.TypedefSpec:

			var (
				toTypeCast string
				err        error
			)
			if _, ok := toField.Type.(*compile.TypedefSpec); ok {
				toTypeCast, err = c.getIdentifierName(toField.Type)
			} else {
				toTypeCast, err = c.getGoTypeName(toField.Type)
			}
			if err != nil {
				return err
			}

			defaultAssignment := &FieldAssignment{
				From:           defaultFromField,
				FromIdentifier: defaultFromIdentifier,
				To:             toField,
				ToIdentifier:   toIdentifier,
				ExtraNilCheck:  false,
				TypeCast:       toTypeCast,
			}

			var overrideAssignment *FieldAssignment
			if overrideFromField != nil {
				overrideAssignment = &FieldAssignment{
					From:           overrideFromField,
					FromIdentifier: overrideFromIdentifier,
					To:             toField,
					ToIdentifier:   toIdentifier,
					ExtraNilCheck:  false,
					TypeCast:       toTypeCast,
				}
			}

			c.genConverterForPrimitiveOrTypedef(
				indent, toTypeCast, defaultAssignment, overrideAssignment)

		case *compile.BinarySpec:
			// TODO: handle override. Check if binarySpec can be optional.
			c.append(toIdentifier, " = []byte(", defaultFromIdentifier, ")")
		//case *compile.TypedefSpec:
		//	typeName, err := c.getIdentifierName(toField.Type)
		//	if err != nil {
		//		return err
		//	}
		//	defaultFieldRequired := false
		//
		//	// TODO: typedef for struct is invalid here ...
		//	if toField.Required {
		//		defaultType := typeName
		//		if !overrideFromField.Required {
		//			// Dereference the pointer
		//			typeName = "*"
		//		}
		//		if defaultFromField != nil {
		//			if !defaultFromField.Required {
		//				// Dereference the pointer
		//				defaultType = "*"
		//			} else {
		//				defaultFieldRequired = true
		//			}
		//		}
		//		c.assignWithOverride(
		//			indent,
		//			toIdentifier,
		//			typeName,
		//			toField.Required,
		//			overrideFromIdentifier,
		//			overrideFromField.Required,
		//			defaultType,
		//			defaultFromIdentifier,
		//			defaultFieldRequired,
		//		)
		//	} else {
		//		fromType := fmt.Sprintf("(*%s)", typeName)
		//		if overrideFromField.Required {
		//			fromType = fmt.Sprintf("(*%s)&", typeName)
		//		}
		//		defaultType := fmt.Sprintf("(*%s)", typeName)
		//		if defaultFromField != nil && defaultFromField.Required {
		//			defaultType = fmt.Sprintf("(*%s)&", typeName)
		//			defaultFieldRequired = true
		//		}
		//
		//		c.assignWithOverride(
		//			indent,
		//			toIdentifier,
		//			fromType,
		//			toField.Required,
		//			overrideFromIdentifier,
		//			overrideFromField.Required,
		//			defaultType,
		//			defaultFromIdentifier,
		//			defaultFieldRequired,
		//		)
		//	}

		case *compile.StructSpec:
			var stFromType   compile.TypeSpec
			stFromType = nil
			stFromPrefix := keyPrefix
			stFromIdentifier := ""

			if defaultFromField != nil {
				if isTransform(defaultFromIdentifier, toIdentifier) && overrideFromField != nil {
					stFromType = overrideFromField.Type
					stFromPrefix = keyPrefix + pascalCase(overrideFromField.Name)
					stFromIdentifier = overrideFromIdentifier
				} else {
					stFromType = defaultFromField.Type
					stFromPrefix = keyPrefix + pascalCase(defaultFromField.Name)
					stFromIdentifier = defaultFromIdentifier
				}
			}

			err := c.genConverterForStruct(
				toField.Name,
				toFieldType,
				stFromType,
				stFromIdentifier,
				keyPrefix+pascalCase(toField.Name),
				stFromPrefix,
				indent,
				fieldMap,
			)
			if err != nil {
				return err
			}
		case *compile.ListSpec:
			err := c.genConverterForList2(
				toFieldType,
				toField,
				toIdentifier,
				defaultFromField,
				defaultFromIdentifier,
				overrideFromField,
				overrideFromIdentifier,
				keyPrefix,
				indent,
			)

			if err != nil {
				return err
			}
		case *compile.MapSpec:
			err := c.genConverterForMap2(
				toFieldType,
				toField,
				toIdentifier,
				defaultFromField,
				defaultFromIdentifier,
				overrideFromField,
				overrideFromIdentifier,
				keyPrefix,
				indent,
			)
			if err != nil {
				return err
			}
		default:
			// fmt.Printf("Unknown type %s for field %s \n",
			// 	toField.Type.TypeCode().String(), toField.Name,
			// )

			// pkgName, err := h.TypePackageName(toField.Type.ThriftFile())
			// if err != nil {
			// 	return nil, err
			// }
			// typeName := pkgName + "." + toField.Type.ThriftName()
			// line := toIdentifier + "(*" + typeName + ")" + postfix
			// c.Lines = append(c.Lines, line)
		}
	}

	return nil
}

// GenStructConverter will add lines to the TypeConverter for mapping
// from one go struct to another based on two thriftrw.FieldGroups.
// fieldMap is a may from keys that are the qualified field path names for
// destination fields (sent to the downstream client) and the entries are source
// fields (from the incoming request)
func (c *TypeConverter) GenStructConverter(
	fromFields []*compile.FieldSpec,
	toFields []*compile.FieldSpec,
	fieldMap map[string]FieldMapperEntry,
) error {
	// Add compiled FieldSpecs to the FieldMapperEntry
	fieldMap = addSpecToMap(fieldMap, fromFields, "")
	// Check for vlaues not populated recursively by addSpecToMap
	for k, v := range fieldMap {
		if fieldMap[k].Field == nil {
			return errors.Errorf(
				"Failed to find field ( %s ) for transform.",
				v.QualifiedName,
			)
		}
	}

	err := c.genStructConverter("", "", "", fromFields, toFields, fieldMap)
	if err != nil {
		return err
	}

	return nil
}

// FieldMapperEntry defines a source field and optional arguments
// converting and overriding fields.
type FieldMapperEntry struct {
	QualifiedName string //
	Field         *compile.FieldSpec

	Override      bool
	typeConverter string // TODO: implement. i.e string(int) etc
	transform     string // TODO: implement. i.e. camelCasing, Title, etc
}

func addSpecToMap(
	overrideMap map[string]FieldMapperEntry,
	fields compile.FieldGroup,
	prefix string,
) map[string]FieldMapperEntry {
	for k, v := range overrideMap {
		for _, spec := range fields {
			fieldQualName := prefix + pascalCase(spec.Name)
			if v.QualifiedName == fieldQualName {
				v.Field = spec
				overrideMap[k] = v
				break
			} else if strings.HasPrefix(v.QualifiedName, fieldQualName) {
				structSpec := spec.Type.(*compile.StructSpec)
				overrideMap = addSpecToMap(overrideMap, structSpec.Fields, fieldQualName+".")
			}
		}
	}
	return overrideMap
}

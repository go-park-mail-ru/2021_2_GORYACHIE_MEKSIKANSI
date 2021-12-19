// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package order

import (
	profile "2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson120d1ca2Decode20212GORYACHIEMEKSIKANSIInternalsOrder(in *jlexer.Lexer, out *CreateOrder) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "methodPay":
			out.MethodPay = string(in.String())
		case "address":
			easyjson120d1ca2Decode20212GORYACHIEMEKSIKANSIInternalsProfile(in, &out.Address)
		case "comment":
			out.Comment = string(in.String())
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
func easyjson120d1ca2Encode20212GORYACHIEMEKSIKANSIInternalsOrder(out *jwriter.Writer, in CreateOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"methodPay\":"
		out.RawString(prefix[1:])
		out.String(string(in.MethodPay))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		easyjson120d1ca2Encode20212GORYACHIEMEKSIKANSIInternalsProfile(out, in.Address)
	}
	{
		const prefix string = ",\"comment\":"
		out.RawString(prefix)
		out.String(string(in.Comment))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2Encode20212GORYACHIEMEKSIKANSIInternalsOrder(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2Decode20212GORYACHIEMEKSIKANSIInternalsOrder(l, v)
}
func easyjson120d1ca2Decode20212GORYACHIEMEKSIKANSIInternalsProfile(in *jlexer.Lexer, out *profile.AddressCoordinates) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "coordinates":
			easyjson120d1ca2Decode20212GORYACHIEMEKSIKANSIInternalsProfile1(in, &out.Coordinates)
		case "alias":
			out.Alias = string(in.String())
		case "city":
			out.City = string(in.String())
		case "street":
			out.Street = string(in.String())
		case "house":
			out.House = string(in.String())
		case "flat":
			out.Flat = string(in.String())
		case "porch":
			out.Porch = int(in.Int())
		case "floor":
			out.Floor = int(in.Int())
		case "intercom":
			out.Intercom = string(in.String())
		case "comment":
			out.Comment = string(in.String())
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
func easyjson120d1ca2Encode20212GORYACHIEMEKSIKANSIInternalsProfile(out *jwriter.Writer, in profile.AddressCoordinates) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"coordinates\":"
		out.RawString(prefix[1:])
		easyjson120d1ca2Encode20212GORYACHIEMEKSIKANSIInternalsProfile1(out, in.Coordinates)
	}
	if in.Alias != "" {
		const prefix string = ",\"alias\":"
		out.RawString(prefix)
		out.String(string(in.Alias))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	if in.Street != "" {
		const prefix string = ",\"street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	if in.House != "" {
		const prefix string = ",\"house\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	if in.Flat != "" {
		const prefix string = ",\"flat\":"
		out.RawString(prefix)
		out.String(string(in.Flat))
	}
	if in.Porch != 0 {
		const prefix string = ",\"porch\":"
		out.RawString(prefix)
		out.Int(int(in.Porch))
	}
	if in.Floor != 0 {
		const prefix string = ",\"floor\":"
		out.RawString(prefix)
		out.Int(int(in.Floor))
	}
	if in.Intercom != "" {
		const prefix string = ",\"intercom\":"
		out.RawString(prefix)
		out.String(string(in.Intercom))
	}
	if in.Comment != "" {
		const prefix string = ",\"comment\":"
		out.RawString(prefix)
		out.String(string(in.Comment))
	}
	out.RawByte('}')
}
func easyjson120d1ca2Decode20212GORYACHIEMEKSIKANSIInternalsProfile1(in *jlexer.Lexer, out *profile.Coordinates) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "latitude":
			out.Latitude = float32(in.Float32())
		case "longitude":
			out.Longitude = float32(in.Float32())
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
func easyjson120d1ca2Encode20212GORYACHIEMEKSIKANSIInternalsProfile1(out *jwriter.Writer, in profile.Coordinates) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"latitude\":"
		out.RawString(prefix[1:])
		out.Float32(float32(in.Latitude))
	}
	{
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.Float32(float32(in.Longitude))
	}
	out.RawByte('}')
}
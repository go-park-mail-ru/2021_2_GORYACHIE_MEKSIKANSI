// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package util

import (
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

func easyjsonA1ea4324Decode20212GORYACHIEMEKSIKANSIInternalsUtil(in *jlexer.Lexer, out *ResponseStatus) {
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
		case "status":
			out.StatusHTTP = int(in.Int())
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
func easyjsonA1ea4324Encode20212GORYACHIEMEKSIKANSIInternalsUtil(out *jwriter.Writer, in ResponseStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.StatusHTTP))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1ea4324Encode20212GORYACHIEMEKSIKANSIInternalsUtil(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1ea4324Decode20212GORYACHIEMEKSIKANSIInternalsUtil(l, v)
}

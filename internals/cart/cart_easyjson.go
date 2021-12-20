// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cart

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

func easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart(in *jlexer.Lexer, out *CartRequest) {
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
		case "cart":
			easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart1(in, &out.Cart)
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
func easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart(out *jwriter.Writer, in CartRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cart\":"
		out.RawString(prefix[1:])
		easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart1(out, in.Cart)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CartRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CartRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart(l, v)
}
func easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart1(in *jlexer.Lexer, out *RequestCartDefault) {
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
		case "restaurant":
			easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart2(in, &out.Restaurant)
		case "dishes":
			if in.IsNull() {
				in.Skip()
				out.Dishes = nil
			} else {
				in.Delim('[')
				if out.Dishes == nil {
					if !in.IsDelim(']') {
						out.Dishes = make([]DishesRequest, 0, 0)
					} else {
						out.Dishes = []DishesRequest{}
					}
				} else {
					out.Dishes = (out.Dishes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 DishesRequest
					easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart3(in, &v1)
					out.Dishes = append(out.Dishes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "promo_code":
			out.PromoCode = string(in.String())
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
func easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart1(out *jwriter.Writer, in RequestCartDefault) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"restaurant\":"
		out.RawString(prefix[1:])
		easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart2(out, in.Restaurant)
	}
	{
		const prefix string = ",\"dishes\":"
		out.RawString(prefix)
		if in.Dishes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Dishes {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart3(out, v3)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"promo_code\":"
		out.RawString(prefix)
		out.String(string(in.PromoCode))
	}
	out.RawByte('}')
}
func easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart3(in *jlexer.Lexer, out *DishesRequest) {
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
		case "id":
			out.Id = int(in.Int())
		case "itNum":
			out.ItemNumber = int(in.Int())
		case "count":
			out.Count = int(in.Int())
		case "radios":
			if in.IsNull() {
				in.Skip()
				out.Radios = nil
			} else {
				in.Delim('[')
				if out.Radios == nil {
					if !in.IsDelim(']') {
						out.Radios = make([]RadiosCartRequest, 0, 4)
					} else {
						out.Radios = []RadiosCartRequest{}
					}
				} else {
					out.Radios = (out.Radios)[:0]
				}
				for !in.IsDelim(']') {
					var v4 RadiosCartRequest
					easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart4(in, &v4)
					out.Radios = append(out.Radios, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ingredients":
			if in.IsNull() {
				in.Skip()
				out.Ingredients = nil
			} else {
				in.Delim('[')
				if out.Ingredients == nil {
					if !in.IsDelim(']') {
						out.Ingredients = make([]IngredientsCartRequest, 0, 8)
					} else {
						out.Ingredients = []IngredientsCartRequest{}
					}
				} else {
					out.Ingredients = (out.Ingredients)[:0]
				}
				for !in.IsDelim(']') {
					var v5 IngredientsCartRequest
					easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart5(in, &v5)
					out.Ingredients = append(out.Ingredients, v5)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart3(out *jwriter.Writer, in DishesRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"itNum\":"
		out.RawString(prefix)
		out.Int(int(in.ItemNumber))
	}
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix)
		out.Int(int(in.Count))
	}
	{
		const prefix string = ",\"radios\":"
		out.RawString(prefix)
		if in.Radios == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Radios {
				if v6 > 0 {
					out.RawByte(',')
				}
				easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart4(out, v7)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"ingredients\":"
		out.RawString(prefix)
		if in.Ingredients == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Ingredients {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart5(out, v9)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart5(in *jlexer.Lexer, out *IngredientsCartRequest) {
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
		case "id":
			out.Id = int(in.Int())
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
func easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart5(out *jwriter.Writer, in IngredientsCartRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	out.RawByte('}')
}
func easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart4(in *jlexer.Lexer, out *RadiosCartRequest) {
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
		case "rId":
			out.RadiosId = int(in.Int())
		case "id":
			out.Id = int(in.Int())
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
func easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart4(out *jwriter.Writer, in RadiosCartRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.RadiosId))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int(int(in.Id))
	}
	out.RawByte('}')
}
func easyjsonDdb0949aDecode20212GORYACHIEMEKSIKANSIInternalsCart2(in *jlexer.Lexer, out *RestaurantRequest) {
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
		case "id":
			out.Id = int(in.Int())
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
func easyjsonDdb0949aEncode20212GORYACHIEMEKSIKANSIInternalsCart2(out *jwriter.Writer, in RestaurantRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	out.RawByte('}')
}

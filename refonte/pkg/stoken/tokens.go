package stoken

import "github.com/cgi-fr/posimap/refonte/api"

const (
	RecordStart api.StructToken = '('
	RecordEnd   api.StructToken = ')'
	ObjectStart api.StructToken = '{'
	ObjectEnd   api.StructToken = '}'
	ArrayStart  api.StructToken = '['
	ArrayEnd    api.StructToken = ']'
	Separator   api.StructToken = ','
	Key         api.StructToken = ':'
	String      api.StructToken = '"'
	Number      api.StructToken = '0'
	Boolean     api.StructToken = 't'
	Null        api.StructToken = 'n'
	EOF         api.StructToken = 0
)

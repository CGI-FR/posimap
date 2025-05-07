package codec

const (
	VisibleBlankRunes = " \u00A0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205F\u3000" //nolint:lll
	ControlBlankRunes = "\t\n\v\f\r\u0085\u2028\u2029"
	BlankRunes        = VisibleBlankRunes + ControlBlankRunes
)

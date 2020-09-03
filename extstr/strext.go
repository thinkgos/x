package extstr

import (
	"strings"
	"unicode"
)

// Recombine 转换驼峰字符串为用sep分隔的字符串
// example: sep = '_'
// HelloWorld -> hello_world
// Hello_World -> hello_world
// HiHello_World -> hi_hello_world
// IDCom -> id_com
// IDcom -> i_dcom
// nameIDCom -> name_id_com
// nameIDcom -> name_i_dcom
func Recombine(str string, sep byte) string {
	if str == "" {
		return str
	}

	var isLastCaseUpper bool
	var isCurrCaseUpper bool
	var isNextCaseUpper bool
	var isNextNumberUpper bool
	var buf = strings.Builder{}

	for i, v := range str[:len(str)-1] {
		isNextCaseUpper = str[i+1] >= 'A' && str[i+1] <= 'Z'
		isNextNumberUpper = str[i+1] >= '0' && str[i+1] <= '9'

		if i > 0 {
			if isCurrCaseUpper {
				if isLastCaseUpper && (isNextCaseUpper || isNextNumberUpper) {
					buf.WriteRune(v)
				} else {
					if str[i-1] != sep && str[i+1] != sep {
						buf.WriteRune(rune(sep))
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(str)-2 && (isNextCaseUpper && !isNextNumberUpper) {
					buf.WriteRune(rune(sep))
				}
			}
		} else {
			isCurrCaseUpper = true
			buf.WriteRune(v)
		}
		isLastCaseUpper = isCurrCaseUpper
		isCurrCaseUpper = isNextCaseUpper
	}

	buf.WriteByte(str[len(str)-1])

	return strings.ToLower(buf.String())
}

// UnRecombine 转换sep分隔的字符串为驼峰字符串
// example: sep = '_'
// hello_world -> HelloWorld
func UnRecombine(str string, sep byte) string {
	if str == "" {
		return ""
	}

	bStr := strings.Builder{}
	for _, s := range strings.Split(str, string(sep)) {
		bStr.WriteString(strings.Title(s))
	}
	return bStr.String()
}

// Recode 重组转换一些特殊的字符
type Recode struct {
	replacer *strings.Replacer
}

// NewRecode 创建一个Recode,以initialisms为自定义的Replacer
// example:
// API -> Api
// ID -> id
func NewRecode(initialisms []string) *Recode {
	initialismsForReplacer := make([]string, 0, len(initialisms)*2)
	for _, s := range initialisms {
		initialismsForReplacer = append(initialismsForReplacer, s, strings.Title(strings.ToLower(s)))
	}

	return &Recode{replacer: strings.NewReplacer(initialismsForReplacer...)}
}

// Recombine 转换驼峰字符串为用sep分隔的字符串,特殊字符由initialisms决定取代
// example1: sep = '_'
// HelloWorld -> hello_world
// Hello_World -> hello_world
// HiHello_World -> hi_hello_world
// example2: sep = '_' initialisms = [ID]
// IDCom -> id_com
// IDcom -> idcom
// nameIDCom -> name_id_com
// nameIDcom -> name_idcom
func (sf Recode) Recombine(str string, sep byte) string {
	return Recombine(sf.replacer.Replace(str), sep)
}

var (
	// DefaultInitialisms default initialism for snake case
	DefaultInitialisms = []string{
		"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP",
		"HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA",
		"SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8",
		"VM", "XML", "XSRF", "XSS",
	}
	defaultReplacer *strings.Replacer
)

func init() {
	initialismsForReplacer := make([]string, 0, len(DefaultInitialisms)*2)
	for _, s := range DefaultInitialisms {
		initialismsForReplacer = append(initialismsForReplacer, s, strings.Title(strings.ToLower(s)))
	}

	defaultReplacer = strings.NewReplacer(initialismsForReplacer...)
}

// SnakeCase 转换驼峰字符串为用'_'分隔的字符串,特殊字符由DefaultInitialisms决定取代
// example2: sep = '_' initialisms = DefaultInitialisms
// IDCom -> id_com
// IDcom -> idcom
// nameIDCom -> name_id_com
// nameIDcom -> name_idcom
func SnakeCase(str string) string {
	return Recombine(defaultReplacer.Replace(str), '_')
}

// CamelCase to camel case string
// id_com -> IdCom
// idcom -> Idcom
// name_id_com -> NameIdCom
// name_idcom -> NameIdcom
func CamelCase(str string) string {
	return UnRecombine(str, '_')
}

// isSeparator reports whether the rune could mark a word boundary.
// TODO: update when package unicode captures more of the properties.
// see strings isSeparator
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case r >= '0' && r <= '9':
			return false
		case r >= 'a' && r <= 'z':
			return false
		case r >= 'A' && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}

	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

// LowTitle 首字母小写
// see strings.Title
func LowTitle(str string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(func(r rune) rune {
		if isSeparator(prev) {
			prev = r
			return unicode.ToLower(r)
		}
		prev = r
		return r
	}, str)
}

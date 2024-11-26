package is

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/goloop/g"
)

var (
	// selectorStrictRegex is the regular expression for a valid CSS
	// selector in strict mode.
	selectorStrictRegex = regexp.MustCompile(`^[a-zA-Z_-][a-zA-Z\d_-]*$`)

	// selectorRegex is the regular expression for a valid CSS
	// selector in non-strict mode.
	selectorRegex = regexp.MustCompile(`^(#|\.)?[a-zA-Z_-][a-zA-Z\d_-]*$`)

	// ErrLanguageNotSupported indicates that the specified
	// programming language is not supported.
	ErrLanguageNotSupported = func(lang string) error {
		return fmt.Errorf("programming language '%s' is not supported", lang)
	}
)

// reservedWords contains reserved keywords for various programming languages.
var reservedWords = map[string]map[string]struct{}{
	"python": {
		"and": {}, "as": {}, "assert": {}, "async": {}, "await": {},
		"break": {}, "class": {}, "continue": {}, "def": {}, "del": {},
		"elif": {}, "else": {}, "except": {}, "False": {}, "finally": {},
		"for": {}, "from": {}, "global": {}, "if": {}, "import": {},
		"in": {}, "is": {}, "lambda": {}, "None": {}, "nonlocal": {},
		"not": {}, "or": {}, "pass": {}, "raise": {}, "return": {},
		"True": {}, "try": {}, "while": {}, "with": {}, "yield": {},
		"match": {}, "case": {}, "type": {}, "print": {},
	},

	"javascript": {
		"abstract": {}, "arguments": {}, "await": {}, "boolean": {},
		"break": {}, "byte": {}, "case": {}, "catch": {}, "char": {},
		"class": {}, "const": {}, "continue": {}, "debugger": {},
		"default": {}, "delete": {}, "do": {}, "double": {}, "else": {},
		"enum": {}, "eval": {}, "export": {}, "extends": {}, "false": {},
		"final": {}, "finally": {}, "float": {}, "for": {}, "function": {},
		"goto": {}, "if": {}, "implements": {}, "import": {}, "in": {},
		"instanceof": {}, "int": {}, "interface": {}, "let": {}, "long": {},
		"native": {}, "new": {}, "null": {}, "package": {}, "private": {},
		"protected": {}, "public": {}, "return": {}, "short": {},
		"static": {}, "super": {}, "switch": {}, "synchronized": {},
		"this": {}, "throw": {}, "throws": {}, "transient": {}, "true": {},
		"try": {}, "typeof": {}, "var": {}, "void": {}, "volatile": {},
		"while": {}, "with": {}, "yield": {}, "async": {},
	},

	"typescript": {
		"any": {}, "as": {}, "async": {}, "await": {}, "boolean": {},
		"break": {}, "case": {}, "catch": {}, "class": {}, "const": {},
		"constructor": {}, "continue": {}, "debugger": {}, "declare": {},
		"default": {}, "delete": {}, "do": {}, "else": {}, "enum": {},
		"export": {}, "extends": {}, "false": {}, "finally": {}, "for": {},
		"from": {}, "function": {}, "get": {}, "if": {}, "implements": {},
		"import": {}, "in": {}, "infer": {}, "instanceof": {}, "interface": {},
		"is": {}, "keyof": {}, "let": {}, "module": {}, "namespace": {},
		"never": {}, "new": {}, "null": {}, "number": {}, "object": {},
		"package": {}, "private": {}, "protected": {}, "public": {},
		"readonly": {}, "require": {}, "return": {}, "set": {}, "static": {},
		"string": {}, "super": {}, "switch": {}, "this": {}, "throw": {},
		"true": {}, "try": {}, "type": {}, "typeof": {}, "undefined": {},
		"unique": {}, "unknown": {}, "var": {}, "void": {}, "while": {},
		"with": {}, "yield": {},
	},

	"go": {
		"break": {}, "case": {}, "chan": {}, "const": {}, "continue": {},
		"default": {}, "defer": {}, "else": {}, "fallthrough": {}, "for": {},
		"func": {}, "go": {}, "goto": {}, "if": {}, "import": {},
		"interface": {}, "map": {}, "package": {}, "range": {}, "return": {},
		"select": {}, "struct": {}, "switch": {}, "type": {}, "var": {},
	},

	"rust": {
		"as": {}, "break": {}, "const": {}, "continue": {}, "crate": {},
		"dyn": {}, "else": {}, "enum": {}, "extern": {}, "false": {},
		"fn": {}, "for": {}, "if": {}, "impl": {}, "in": {}, "let": {},
		"loop": {}, "match": {}, "mod": {}, "move": {}, "mut": {}, "pub": {},
		"ref": {}, "return": {}, "self": {}, "Self": {}, "static": {},
		"struct": {}, "super": {}, "trait": {}, "true": {}, "type": {},
		"unsafe": {}, "use": {}, "where": {}, "while": {}, "async": {},
		"await": {}, "try": {}, "union": {},
	},

	"cpp": {
		"alignas": {}, "alignof": {}, "and": {}, "and_eq": {}, "asm": {},
		"auto": {}, "bitand": {}, "bitor": {}, "bool": {}, "break": {},
		"case": {}, "catch": {}, "char": {}, "char8_t": {}, "char16_t": {},
		"char32_t": {}, "class": {}, "compl": {}, "concept": {}, "const": {},
		"consteval": {}, "constexpr": {}, "constinit": {}, "const_cast": {},
		"continue": {}, "co_await": {}, "co_return": {}, "co_yield": {},
		"decltype": {}, "default": {}, "delete": {}, "do": {}, "double": {},
		"dynamic_cast": {}, "else": {}, "enum": {}, "explicit": {},
		"export": {}, "extern": {}, "false": {}, "float": {}, "for": {},
		"friend": {}, "goto": {}, "if": {}, "inline": {}, "int": {},
		"long": {}, "mutable": {}, "namespace": {}, "new": {}, "noexcept": {},
		"not": {}, "not_eq": {}, "nullptr": {}, "operator": {}, "or": {},
		"or_eq": {}, "private": {}, "protected": {}, "public": {},
		"register": {}, "reinterpret_cast": {}, "requires": {}, "return": {},
		"short": {}, "signed": {}, "sizeof": {}, "static": {},
		"static_assert": {}, "static_cast": {}, "struct": {}, "switch": {},
		"template": {}, "this": {}, "thread_local": {}, "throw": {},
		"true": {}, "try": {}, "typedef": {}, "typeid": {}, "typename": {},
		"union": {}, "unsigned": {}, "using": {}, "virtual": {}, "void": {},
		"volatile": {}, "wchar_t": {}, "while": {}, "xor": {}, "xor_eq": {},
	},

	"c": {
		"auto": {}, "break": {}, "case": {}, "char": {}, "const": {},
		"continue": {}, "default": {}, "do": {}, "double": {}, "else": {},
		"enum": {}, "extern": {}, "float": {}, "for": {}, "goto": {},
		"if": {}, "inline": {}, "int": {}, "long": {}, "register": {},
		"restrict": {}, "return": {}, "short": {}, "signed": {}, "sizeof": {},
		"static": {}, "struct": {}, "switch": {}, "typedef": {}, "union": {},
		"unsigned": {}, "void": {}, "volatile": {}, "while": {}, "_Bool": {},
		"_Complex": {}, "_Imaginary": {},
		"_Alignas": {}, "_Alignof": {}, "_Atomic": {}, "_Generic": {},
		"_Noreturn": {}, "_Static_assert": {}, "_Thread_local": {},
		"#define": {}, "#include": {}, "#undef": {}, "#ifdef": {},
		"#ifndef": {}, "#if": {}, "#else": {}, "#elif": {}, "#endif": {},
		"#error": {}, "#pragma": {}, "#line": {},
		"size_t": {}, "ptrdiff_t": {}, "NULL": {}, "true": {}, "false": {},
		"stdin": {}, "stdout": {}, "stderr": {}, "errno": {}, "EOF": {},
		"int8_t": {}, "uint8_t": {}, "int16_t": {}, "uint16_t": {},
		"int32_t": {}, "uint32_t": {}, "int64_t": {}, "uint64_t": {},
		"intptr_t": {}, "uintptr_t": {}, "wchar_t": {}, "time_t": {},
		"INFINITY": {}, "NAN": {}, "INT_MAX": {}, "INT_MIN": {},
		"UINT_MAX": {}, "offsetof": {}, "va_list": {},
		"va_start": {}, "va_end": {}, "va_arg": {},
	},

	"java": {
		"abstract": {}, "assert": {}, "boolean": {}, "break": {},
		"byte": {}, "case": {}, "catch": {}, "char": {}, "class": {},
		"const": {}, "continue": {}, "default": {}, "do": {}, "double": {},
		"else": {}, "enum": {}, "extends": {}, "final": {}, "finally": {},
		"float": {}, "for": {}, "goto": {}, "if": {}, "implements": {},
		"import": {}, "instanceof": {}, "int": {}, "interface": {},
		"long": {}, "native": {}, "new": {}, "package": {}, "private": {},
		"protected": {}, "public": {}, "return": {}, "short": {},
		"static": {}, "strictfp": {}, "super": {}, "switch": {},
		"synchronized": {}, "this": {}, "throw": {}, "throws": {},
		"transient": {}, "try": {}, "void": {}, "volatile": {}, "while": {},
		"true": {}, "false": {}, "null": {}, "var": {}, "sealed": {},
		"permits": {}, "record": {}, "yield": {},
	},

	"kotlin": {
		"abstract": {}, "actual": {}, "annotation": {}, "as": {}, "break": {},
		"by": {}, "catch": {}, "class": {}, "companion": {}, "const": {},
		"constructor": {}, "continue": {}, "crossinline": {}, "data": {},
		"delegate": {}, "do": {}, "dynamic": {}, "else": {}, "enum": {},
		"expect": {}, "external": {}, "false": {}, "field": {}, "final": {},
		"finally": {}, "for": {}, "fun": {}, "get": {}, "if": {}, "import": {},
		"infix": {}, "init": {}, "inline": {}, "inner": {}, "interface": {},
		"internal": {}, "is": {}, "lateinit": {}, "noinline": {}, "null": {},
		"object": {}, "open": {}, "operator": {}, "out": {}, "override": {},
		"package": {}, "private": {}, "protected": {}, "public": {},
		"reified": {}, "return": {}, "sealed": {}, "set": {}, "super": {},
		"suspend": {}, "tailrec": {}, "this": {}, "throw": {}, "true": {},
		"try": {}, "typealias": {}, "val": {}, "var": {}, "vararg": {},
		"when": {}, "while": {}, "in": {},
	},

	"swift": {
		"any": {}, "as": {}, "associatedtype": {}, "associativity": {},
		"case": {}, "catch": {}, "class": {}, "continue": {},
		"default": {}, "defer": {}, "deinit": {}, "didSet": {}, "do": {},
		"dynamic": {}, "else": {}, "enum": {}, "extension": {},
		"false": {}, "fileprivate": {}, "final": {}, "for": {}, "func": {},
		"get": {}, "guard": {}, "if": {}, "import": {}, "in": {}, "infix": {},
		"init": {}, "inout": {}, "internal": {}, "is": {}, "lazy": {},
		"let": {}, "mutating": {}, "nil": {}, "none": {}, "nonmutating": {},
		"open": {}, "operator": {}, "optional": {}, "override": {},
		"precedence": {}, "prefix": {}, "private": {}, "protocol": {},
		"repeat": {}, "required": {}, "return": {}, "right": {}, "self": {},
		"set": {}, "static": {}, "struct": {}, "subscript": {}, "super": {},
		"switch": {}, "throw": {}, "throws": {}, "true": {}, "try": {},
		"typealias": {}, "unowned": {}, "var": {}, "weak": {}, "where": {},
		"while": {}, "willSet": {}, "break": {}, "fallthrough": {}, "left": {},
		"public": {}, "type": {}, "postfix": {}, "convenience": {},
	},

	"sql": {
		"add": {}, "all": {}, "alter": {}, "and": {}, "any": {}, "as": {},
		"asc": {}, "backup": {}, "between": {}, "by": {}, "case": {},
		"column": {}, "constraint": {}, "create": {}, "database": {},
		"delete": {}, "desc": {}, "distinct": {}, "drop": {}, "exec": {},
		"exists": {}, "foreign": {}, "from": {}, "full": {}, "group": {},
		"in": {}, "index": {}, "inner": {}, "insert": {}, "into": {}, "is": {},
		"join": {}, "key": {}, "left": {}, "like": {}, "limit": {}, "not": {},
		"null": {}, "or": {}, "order": {}, "outer": {}, "primary": {},
		"procedure": {}, "right": {}, "rownum": {}, "select": {}, "set": {},
		"table": {}, "top": {}, "truncate": {}, "union": {}, "unique": {},
		"update": {}, "values": {}, "view": {}, "where": {}, "with": {},
		"check": {}, "default": {}, "having": {},
	},

	"php": {}, // all variables has $ prefix in PHP

	// "php": {
	// 	"abstract": {}, "and": {}, "array": {}, "as": {}, "break": {},
	// 	"case": {}, "catch": {}, "class": {}, "clone": {}, "const": {},
	// 	"declare": {}, "default": {}, "die": {}, "do": {}, "echo": {},
	// 	"elseif": {}, "empty": {}, "enddeclare": {}, "endfor": {},
	// 	"endif": {}, "endswitch": {}, "endwhile": {}, "eval": {}, "exit": {},
	// 	"extends": {}, "final": {}, "finally": {}, "fn": {}, "for": {},
	// 	"function": {}, "global": {}, "goto": {}, "if": {}, "implements": {},
	// 	"include": {}, "include_once": {}, "instanceof": {}, "insteadof": {},
	// 	"interface": {}, "isset": {}, "list": {}, "match": {}, "namespace": {},
	// 	"new": {}, "or": {}, "print": {}, "private": {}, "protected": {},
	// 	"require": {}, "require_once": {}, "return": {}, "static": {},
	// 	"throw": {}, "trait": {}, "try": {}, "unset": {}, "use": {},
	// 	"while": {}, "xor": {}, "yield": {}, "yield from": {},
	// 	"callable": {}, "continue": {}, "else": {}, "endforeach": {},
	// 	"foreach": {}, "public": {}, "switch": {}, "var": {}, "true": {},
	// 	"false": {}, "null": {},
	// },

	"ruby": {
		"alias": {}, "and": {}, "begin": {}, "break": {}, "case": {},
		"def": {}, "defined?": {}, "do": {}, "else": {}, "elsif": {}, "end": {},
		"ensure": {}, "false": {}, "for": {}, "if": {}, "in": {}, "module": {},
		"next": {}, "nil": {}, "not": {}, "or": {}, "redo": {}, "rescue": {},
		"retry": {}, "return": {}, "self": {}, "super": {}, "then": {},
		"undef": {}, "unless": {}, "until": {}, "when": {}, "while": {},
		"class": {}, "true": {}, "yield": {},
	},

	"csharp": {
		"abstract": {}, "as": {}, "base": {}, "bool": {}, "break": {},
		"case": {}, "catch": {}, "char": {}, "checked": {}, "class": {},
		"continue": {}, "decimal": {}, "default": {}, "delegate": {}, "do": {},
		"double": {}, "else": {}, "enum": {}, "event": {}, "explicit": {},
		"extern": {}, "false": {}, "finally": {}, "fixed": {}, "float": {},
		"for": {}, "foreach": {}, "goto": {}, "if": {}, "implicit": {},
		"in": {}, "int": {}, "interface": {}, "internal": {}, "is": {},
		"long": {}, "namespace": {}, "new": {}, "null": {}, "object": {},
		"operator": {}, "out": {}, "override": {}, "params": {}, "private": {},
		"protected": {}, "public": {}, "readonly": {}, "ref": {}, "return": {},
		"sbyte": {}, "sealed": {}, "short": {}, "sizeof": {}, "stackalloc": {},
		"static": {}, "string": {}, "struct": {}, "switch": {}, "this": {},
		"throw": {}, "true": {}, "try": {}, "typeof": {}, "uint": {},
		"unchecked": {}, "unsafe": {}, "ushort": {}, "using": {}, "virtual": {},
		"void": {}, "volatile": {}, "while": {}, "add": {}, "alias": {},
		"ascending": {}, "async": {}, "await": {}, "by": {}, "descending": {},
		"dynamic": {}, "equals": {}, "from": {}, "get": {}, "global": {},
		"group": {}, "into": {}, "join": {}, "let": {}, "nameof": {},
		"notnull": {}, "on": {}, "orderby": {}, "partial": {}, "remove": {},
		"select": {}, "set": {}, "unmanaged": {}, "value": {}, "var": {},
		"when": {}, "where": {}, "with": {}, "yield": {}, "byte": {},
		"const": {}, "lock": {}, "ulong": {},
	},

	"html": {
		"!doctype": {}, "a": {}, "abbr": {}, "address": {}, "area": {},
		"article": {}, "aside": {}, "audio": {}, "b": {}, "base": {},
		"bdi": {}, "bdo": {}, "blockquote": {}, "body": {}, "br": {},
		"button": {}, "canvas": {}, "caption": {}, "cite": {}, "code": {},
		"col": {}, "colgroup": {}, "data": {}, "datalist": {}, "dd": {},
		"del": {}, "details": {}, "dfn": {}, "dialog": {}, "div": {},
		"dl": {}, "dt": {}, "em": {}, "embed": {}, "fieldset": {},
		"figure": {}, "footer": {}, "form": {}, "h1": {}, "h2": {}, "h3": {},
		"h4": {}, "h5": {}, "h6": {}, "head": {}, "header": {}, "hr": {},
		"html": {}, "i": {}, "iframe": {}, "img": {}, "input": {}, "ins": {},
		"kbd": {}, "label": {}, "legend": {}, "li": {}, "link": {}, "main": {},
		"map": {}, "mark": {}, "meta": {}, "meter": {}, "nav": {},
		"object": {}, "ol": {}, "optgroup": {}, "option": {}, "output": {},
		"p": {}, "param": {}, "picture": {}, "pre": {}, "progress": {}, "q": {},
		"rp": {}, "rt": {}, "ruby": {}, "s": {}, "samp": {}, "script": {},
		"section": {}, "select": {}, "small": {}, "source": {}, "span": {},
		"strong": {}, "style": {}, "sub": {}, "summary": {}, "sup": {},
		"table": {}, "tbody": {}, "td": {}, "template": {}, "textarea": {},
		"tfoot": {}, "th": {}, "thead": {}, "time": {}, "title": {}, "tr": {},
		"track": {}, "u": {}, "ul": {}, "var": {}, "video": {}, "wbr": {},
		"figcaption": {}, "noscript": {},
	},

	"css": {
		"@charset": {}, "@import": {}, "@namespace": {}, "@media": {},
		"@supports": {}, "@document": {}, "@page": {}, "@font-face": {},
		"@keyframes": {}, "@viewport": {}, "@counter-style": {},
		"@layer": {}, "@property": {}, "active": {}, "after": {}, "align": {},
		"align-content": {}, "align-items": {}, "align-self": {}, "all": {},
		"animation": {}, "appearance": {}, "attr": {}, "auto": {},
		"background": {}, "before": {}, "border": {}, "bottom": {}, "box": {},
		"break": {}, "calc": {}, "caret": {}, "checked": {}, "clear": {},
		"clip": {}, "color": {}, "column": {}, "content": {}, "counter": {},
		"cursor": {}, "direction": {}, "display": {}, "element": {},
		"enabled": {}, "filter": {}, "first": {}, "flex": {}, "float": {},
		"focus": {}, "font": {}, "gap": {}, "grid": {}, "height": {},
		"hover": {}, "hsl": {}, "hsla": {}, "if": {}, "image": {}, "import": {},
		"important": {}, "in": {}, "inherit": {}, "initial": {}, "inset": {},
		"justify": {}, "lang": {}, "last": {}, "left": {}, "line": {},
		"link": {}, "list": {}, "margin": {}, "mask": {}, "matrix": {},
		"max": {}, "media": {}, "min": {}, "mix": {}, "nth": {}, "only": {},
		"opacity": {}, "order": {}, "outline": {}, "overflow": {},
		"page": {}, "perspective": {}, "placeholder": {}, "position": {},
		"property": {}, "rgb": {}, "rgba": {}, "right": {}, "root": {},
		"rotate": {}, "scale": {}, "scroll": {}, "selection": {}, "shadow": {},
		"shape": {}, "size": {}, "skew": {}, "space": {}, "src": {},
		"step": {}, "style": {}, "supports": {}, "table": {}, "target": {},
		"text": {}, "top": {}, "transform": {}, "transition": {},
		"type": {}, "url": {}, "user": {}, "valid": {}, "var": {},
		"vertical": {}, "visibility": {}, "visited": {}, "white": {},
		"z-index": {}, "@font-feature-values": {}, "backdrop-filter": {},
		"translate": {}, "width": {}, "empty": {}, "padding": {},
	},

	"markdown": {
		"#": {}, "##": {}, "###": {}, "####": {}, "#####": {}, "######": {},
		"*": {}, "**": {}, "_": {}, "__": {}, "___": {},
		"`": {}, "``": {}, "```": {}, "~~~~": {}, ">": {}, ">>": {},
		"-": {}, "+": {}, "1.": {}, "[": {}, "]": {}, "(": {}, ")": {},
		"{": {}, "}": {}, "!": {}, "|": {}, "\\": {}, "---": {},
		"===": {}, "***": {}, "<!--": {}, "-->": {},
	},

	"regex": {
		"^": {}, "$": {}, ".": {}, "*": {}, "+": {}, "?": {}, "|": {},
		"\\": {}, "(": {}, ")": {}, "[": {}, "]": {}, "{": {}, "}": {},
		"\\d": {}, "\\D": {}, "\\w": {}, "\\W": {}, "\\s": {}, "\\S": {},
		"\\b": {}, "\\B": {}, "\\n": {}, "\\r": {}, "\\t": {}, "\\f": {},
		"\\v": {}, "\\0": {}, "\\x": {}, "\\u": {}, "\\c": {}, "\\p": {},
	},

	"bash": {
		"alias": {}, "bg": {}, "bind": {}, "break": {}, "builtin": {},
		"caller": {}, "cd": {}, "command": {}, "compgen": {}, "complete": {},
		"compopt": {}, "continue": {}, "declare": {}, "dirs": {}, "disown": {},
		"echo": {}, "enable": {}, "eval": {}, "exec": {}, "exit": {},
		"export": {}, "false": {}, "fc": {}, "fg": {}, "for": {},
		"function": {}, "getopts": {}, "hash": {}, "help": {}, "history": {},
		"if": {}, "jobs": {}, "kill": {}, "let": {}, "local": {},
		"logout": {}, "mapfile": {}, "popd": {}, "printf": {}, "pushd": {},
		"pwd": {}, "read": {}, "readarray": {}, "readonly": {}, "return": {},
		"set": {}, "shift": {}, "shopt": {}, "source": {}, "suspend": {},
		"test": {}, "times": {}, "trap": {}, "true": {}, "type": {},
		"typeset": {}, "ulimit": {}, "umask": {}, "unalias": {}, "unset": {},
		"until": {}, "wait": {}, "while": {}, "do": {}, "done": {},
		"elif": {}, "else": {}, "esac": {}, "fi": {}, "in": {},
		"then": {}, "case": {},
	},

	"assembly": {
		"mov": {}, "add": {}, "sub": {}, "mul": {}, "div": {}, "jmp": {},
		"cmp": {}, "je": {}, "jne": {}, "call": {}, "ret": {}, "push": {},
		"pop": {}, "nop": {}, "and": {}, "or": {}, "xor": {}, "shl": {},
		"shr": {}, "inc": {}, "dec": {}, "eax": {}, "ebx": {}, "ecx": {},
		"edx": {}, "esi": {}, "edi": {}, "esp": {}, "ebp": {}, "ax": {},
		"bx": {}, "cx": {}, "dx": {}, "data": {}, "text": {}, "byte": {},
		"word": {}, "dword": {}, "qword": {}, "ptr": {}, "proc": {},
		"segment": {}, "ends": {}, "assume": {}, "org": {}, "offset": {},
		"endp": {},
	},

	"perl": {
		"__data__": {}, "__end__": {}, "__file__": {}, "__line__": {},
		"begin": {}, "check": {}, "end": {}, "init": {}, "unitcheck": {},
		"and": {}, "cmp": {}, "continue": {}, "do": {}, "else": {},
		"elsif": {}, "eq": {}, "for": {}, "foreach": {}, "ge": {},
		"gt": {}, "if": {}, "le": {}, "lt": {}, "ne": {}, "not": {},
		"or": {}, "package": {}, "sub": {}, "unless": {}, "until": {},
		"while": {}, "xor": {}, "my": {}, "our": {}, "use": {},
	},

	"lua": {
		"and": {}, "break": {}, "do": {}, "else": {}, "elseif": {},
		"end": {}, "false": {}, "for": {}, "function": {}, "goto": {},
		"if": {}, "in": {}, "local": {}, "nil": {}, "not": {}, "or": {},
		"repeat": {}, "return": {}, "then": {}, "true": {}, "until": {},
		"while": {},
	},

	"haskell": {
		"case": {}, "class": {}, "data": {}, "default": {}, "deriving": {},
		"do": {}, "else": {}, "if": {}, "import": {}, "in": {}, "infix": {},
		"infixl": {}, "infixr": {}, "instance": {}, "let": {}, "module": {},
		"newtype": {}, "of": {}, "then": {}, "type": {}, "where": {},
		"foreign": {}, "forall": {}, "mdo": {}, "family": {}, "role": {},
		"pattern": {}, "static": {}, "group": {}, "by": {}, "using": {},
		"qualified": {},
	},

	"scala": {
		"abstract": {}, "case": {}, "catch": {}, "class": {}, "def": {},
		"do": {}, "else": {}, "extends": {}, "false": {}, "final": {},
		"finally": {}, "for": {}, "forsome": {}, "if": {}, "implicit": {},
		"import": {}, "lazy": {}, "match": {}, "new": {}, "null": {},
		"object": {}, "override": {}, "package": {}, "private": {},
		"protected": {}, "return": {}, "sealed": {}, "super": {}, "this": {},
		"throw": {}, "trait": {}, "try": {}, "true": {}, "type": {},
		"val": {}, "var": {}, "while": {}, "with": {}, "yield": {},
	},

	"objectivec": {
		"@interface": {}, "@implementation": {}, "@end": {}, "@protocol": {},
		"@class": {}, "@public": {}, "@protected": {}, "@private": {},
		"@property": {}, "@synthesize": {}, "@dynamic": {}, "@selector": {},
		"@optional": {}, "@required": {}, "@try": {}, "@catch": {},
		"@finally": {}, "@throw": {}, "@synchronized": {},
		"@autoreleasepool": {},
	},

	"dart": {
		"abstract": {}, "as": {}, "assert": {}, "async": {}, "await": {},
		"break": {}, "case": {}, "catch": {}, "class": {}, "const": {},
		"continue": {}, "covariant": {}, "default": {}, "deferred": {},
		"do": {}, "dynamic": {}, "else": {}, "enum": {}, "export": {},
		"extends": {}, "extension": {}, "external": {}, "factory": {},
		"false": {}, "final": {}, "finally": {}, "for": {}, "function": {},
		"get": {}, "hide": {}, "if": {}, "implements": {}, "import": {},
		"in": {}, "interface": {}, "is": {}, "late": {}, "library": {},
		"mixin": {}, "new": {}, "null": {}, "on": {}, "operator": {},
		"part": {}, "required": {}, "rethrow": {}, "return": {}, "set": {},
		"show": {}, "static": {}, "super": {}, "switch": {}, "sync": {},
		"this": {}, "throw": {}, "true": {}, "try": {}, "typedef": {},
		"var": {}, "void": {}, "while": {}, "with": {}, "yield": {},
	},

	"elixir": {
		"after": {}, "alias": {}, "and": {}, "case": {}, "cond": {},
		"def": {}, "defimpl": {}, "defmacro": {}, "defmodule": {},
		"defp": {}, "defprotocol": {}, "do": {}, "else": {}, "end": {},
		"fn": {}, "for": {}, "if": {}, "import": {}, "in": {}, "nil": {},
		"not": {}, "or": {}, "quote": {}, "raise": {}, "receive": {},
		"require": {}, "rescue": {}, "try": {}, "unless": {}, "unquote": {},
		"use": {}, "when": {},
	},

	"erlang": {
		"after": {}, "begin": {}, "case": {}, "catch": {}, "cond": {},
		"end": {}, "fun": {}, "if": {}, "let": {}, "of": {}, "query": {},
		"receive": {}, "try": {}, "when": {},
	},

	"r": {
		"if": {}, "else": {}, "repeat": {}, "while": {}, "function": {},
		"for": {}, "in": {}, "next": {}, "break": {}, "true": {},
		"false": {}, "null": {}, "inf": {}, "nan": {}, "na": {},
		"na_integer_": {}, "na_real_": {}, "na_complex_": {},
		"na_character_": {},
	},

	"matlab": {
		"break": {}, "case": {}, "catch": {}, "classdef": {}, "continue": {},
		"else": {}, "elseif": {}, "end": {}, "for": {}, "function": {},
		"global": {}, "if": {}, "otherwise": {}, "parfor": {}, "persistent": {},
		"return": {}, "spmd": {}, "switch": {}, "try": {}, "while": {},
	},

	"vbnet": {
		"addhandler": {}, "addressof": {}, "alias": {}, "and": {},
		"andalso": {}, "as": {}, "boolean": {}, "byref": {}, "byte": {},
		"byval": {}, "call": {}, "case": {}, "catch": {}, "cbool": {},
		"cbyte": {}, "cchar": {}, "cdate": {}, "cdec": {}, "cdbl": {},
		"char": {}, "cint": {}, "class": {}, "clng": {}, "cobj": {},
		"const": {}, "continue": {}, "csbyte": {}, "cshort": {}, "csng": {},
		"cstr": {}, "ctype": {}, "cuint": {}, "culng": {}, "cushort": {},
		"date": {}, "decimal": {}, "declare": {}, "default": {}, "delegate": {},
		"dim": {}, "directcast": {}, "do": {}, "double": {}, "each": {},
		"else": {}, "elseif": {}, "end": {}, "enum": {}, "erase": {},
		"error": {}, "event": {}, "exit": {}, "false": {}, "finally": {},
		"for": {}, "friend": {}, "function": {}, "get": {}, "gettype": {},
		"global": {}, "gosub": {}, "goto": {}, "handles": {}, "if": {},
		"implements": {}, "imports": {}, "in": {}, "inherits": {},
		"interface": {}, "is": {}, "isnot": {}, "let": {}, "lib": {},
		"like": {}, "long": {}, "loop": {}, "me": {}, "mod": {}, "module": {},
		"mustinherit": {}, "mustoverride": {}, "mybase": {}, "myclass": {},
		"namespace": {}, "narrowing": {}, "new": {}, "next": {}, "not": {},
		"nothing": {}, "notinheritable": {}, "notoverridable": {}, "object": {},
		"of": {}, "on": {}, "operator": {}, "option": {}, "optional": {},
		"or": {}, "orelse": {}, "out": {}, "overloads": {}, "overridable": {},
		"overrides": {}, "paramarray": {}, "partial": {}, "private": {},
		"protected": {}, "public": {}, "raiseevent": {}, "readonly": {},
		"redim": {}, "rem": {}, "removehandler": {}, "resume": {}, "return": {},
		"sbyte": {}, "select": {}, "set": {}, "shadows": {}, "shared": {},
		"short": {}, "single": {}, "static": {}, "step": {}, "stop": {},
		"string": {}, "structure": {}, "sub": {}, "synclock": {}, "then": {},
		"throw": {}, "to": {}, "true": {}, "try": {}, "trycast": {},
		"uinteger": {}, "ulong": {}, "ushort": {}, "using": {}, "variant": {},
		"wend": {}, "when": {}, "while": {}, "widening": {}, "with": {},
		"withevents": {}, "writeonly": {}, "xor": {}, "integer": {},
		"typeof": {},
	},

	"cobol": {
		"accept": {}, "add": {}, "call": {}, "cancel": {}, "close": {},
		"compute": {}, "configuration": {}, "copy": {}, "data": {},
		"delete": {}, "display": {}, "divide": {}, "else": {}, "end-if": {},
		"exit": {}, "fd": {}, "file": {}, "if": {}, "initialize": {},
		"inspect": {}, "move": {}, "multiply": {}, "open": {}, "perform": {},
		"read": {}, "rewrite": {}, "search": {}, "stop": {}, "subtract": {},
		"write": {}, "program-id": {}, "author": {}, "date-written": {},
		"environment": {}, "division": {}, "section": {}, "working-storage": {},
	},

	"fortran": {
		"allocate": {}, "assign": {}, "backspace": {}, "block": {},
		"data": {}, "call": {}, "close": {}, "common": {}, "continue": {},
		"deallocate": {}, "do": {}, "else": {}, "else if": {},
		"end": {}, "end file": {}, "endfile": {}, "endif": {}, "entry": {},
		"equivalence": {}, "external": {}, "format": {}, "function": {},
		"go to": {}, "goto": {}, "if": {}, "implicit": {}, "inquire": {},
		"integer": {}, "intrinsic": {}, "logical": {}, "open": {},
		"parameter": {}, "pause": {}, "print": {}, "program": {}, "read": {},
		"return": {}, "rewind": {}, "stop": {}, "subroutine": {}, "then": {},
		"type": {}, "write": {},
	},

	"julia": {
		"begin": {}, "while": {}, "for": {}, "return": {}, "break": {},
		"continue": {}, "function": {}, "macro": {}, "quote": {}, "try": {},
		"catch": {}, "finally": {}, "let": {}, "if": {}, "else": {},
		"elseif": {}, "struct": {}, "module": {}, "baremodule": {}, "using": {},
		"import": {}, "export": {}, "const": {}, "global": {}, "local": {},
		"mutable": {}, "abstract": {}, "primitive": {}, "type": {},
		"immutable": {}, "typealias": {}, "bitstype": {}, "do": {}, "end": {},
		"where": {},
	},

	"prolog": {
		":-": {}, "--": {}, "?-": {}, "is": {}, "assert": {}, "retract": {},
		"clause": {}, "true": {}, "fail": {}, "not": {}, "repeat": {},
		"call": {}, "if": {}, "then": {}, "else": {},
	},

	"eiffel": {
		"across": {}, "agent": {}, "alias": {}, "all": {}, "and": {},
		"as": {}, "assign": {}, "check": {}, "class": {}, "convert": {},
		"create": {}, "debug": {}, "deferred": {}, "do": {}, "else": {},
		"elseif": {}, "end": {}, "ensure": {}, "expanded": {}, "export": {},
		"external": {}, "feature": {}, "from": {}, "frozen": {}, "if": {},
		"implies": {}, "inherit": {}, "inspect": {}, "invariant": {},
		"like": {}, "local": {}, "loop": {}, "not": {}, "note": {},
		"obsolete": {}, "old": {}, "once": {}, "or": {}, "prefix": {},
		"redefine": {}, "rename": {}, "require": {}, "rescue": {}, "retry": {},
		"select": {}, "separate": {}, "then": {}, "undefine": {}, "until": {},
		"variant": {}, "when": {},
	},
}

// languageAliases maps different variations of language
// names to their canonical form.
var languageAliases = map[string]string{
	"js":      "javascript",
	"jsx":     "javascript",
	"ts":      "typescript",
	"py":      "python",
	"rb":      "ruby",
	"golang":  "go",
	"c++":     "cpp",
	"cxx":     "cpp",
	"h":       "cpp",
	"hpp":     "cpp",
	"arduino": "cpp",
}

// languageConfig holds language-specific validation rules.
type languageConfig struct {
	language      string          // the language name
	caseSensitive bool            // is the language case-sensitive
	allowUnicode  bool            // allow unicode characters in identifiers
	checkFirst    func(rune) bool // check the first character
	validChars    func(rune) bool // check if a character is valid
}

var languageConfigs = map[string]languageConfig{
	"python": {
		language:      "python",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"javascript": {
		language:      "javascript",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' || r == '$'
		},
	},
	"typescript": {
		language:      "typescript",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' || r == '$'
		},
	},
	"php": {
		language:      "php",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"ruby": {
		language:      "ruby",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) ||
				r == '_' || r == '@' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"perl": {
		language:      "perl",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"lua": {
		language:      "lua",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"go": {
		language:      "go",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"rust": {
		language:      "rust",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"cpp": {
		language:      "cpp",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"c": {
		language:      "c",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"java": {
		language:      "java",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '$'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' || r == '$'
		},
	},
	"kotlin": {
		language:      "kotlin",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"swift": {
		language:      "swift",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"sql": {
		language:      "sql",
		caseSensitive: false,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '@' || r == '#'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' ||
				r == '@' || r == '#' || r == '$'
		},
	},
	"html": {
		language:      "html",
		caseSensitive: false,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '-'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' || r == '-'
		},
	},
	"css": {
		language:      "css",
		caseSensitive: false,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_' || r == '-' || r == '@'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' || r == '-' || r == '@'
		},
	},
	"r": {
		language:      "r",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '.' || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '.' || r == '_'
		},
	},
	"matlab": {
		language:      "matlab",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r)
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
	"haskell": {
		language:      "haskell",
		caseSensitive: true,
		allowUnicode:  false,
		checkFirst: func(r rune) bool {
			return unicode.IsLower(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) ||
				unicode.IsNumber(r) || r == '_' || r == '\''
		},
	},
	"scala": {
		language:      "scala",
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	},
}

// isLanguageSupported checks if the language is supported.
func isLanguageSupported(language string) bool {
	language = strings.ToLower(language)
	if alias, ok := languageAliases[language]; ok {
		language = alias
	}

	_, exists := reservedWords[language]
	return exists
}

// isReservedInAnyLanguage checks if the given word is reserved
// in any programming language.
func isReservedInAnyLanguage(word string) bool {
	for _, langWords := range reservedWords {
		if _, reserved := langWords[word]; reserved {
			return true
		}
	}

	return false
}

// isValidIdentifier checks if the variable name conforms to the basic rules:
// starts with a valid character and contains only valid characters.
func isValidIdentifier(v string, config languageConfig) bool {
	r := []rune(v)

	// Check for empty string.
	if len(r) == 0 {
		return false
	}

	// Check for allowed unicode characters.
	if !config.allowUnicode {
		for _, c := range r {
			if c > unicode.MaxASCII {
				return false
			}
		}
	}

	// Check the first character.
	if !config.checkFirst(r[0]) {
		return false
	}

	// Special case for Ruby.
	if config.language == "ruby" {
		// Remove prefixs @, @@, $.
		if strings.HasPrefix(v, "@@") {
			r = r[2:]
		} else if strings.HasPrefix(v, "@") || strings.HasPrefix(v, "$") {
			r = r[1:]
		}

		// Check first character after prefix.
		if !unicode.IsLetter(r[0]) && r[0] != '_' {
			return false
		}

		// Check other characters.
		for i, c := range r {
			if !config.validChars(c) {
				// Symbol ? or ! is allowed only at the end
				// of the variable name.
				if (c != '?' && c != '!') || i != len(r)-1 {
					return false
				}
			}
		}

		return true
	}

	// Check other characters.
	for _, c := range r[1:] {
		// Special case for @.
		if c == '@' && !config.checkFirst('@') {
			return false
		}

		if !config.validChars(c) {
			return false
		}
	}

	return true
}

// VariableNameFor validates if the given string can be used as a
// variable name in the specified programming language.
//
// Parameters:
//   - v: string to validate;
//   - language: programming language to check against (case-insensitive);
//
// Returns:
//   - bool: true if the given name is valid for the specified language;
//   - error: ErrLanguageNotSupported if the language is not supported.
func VariableNameFor(v string, language string) (bool, error) {
	// Empty string is not a valid variable name regardless of language.
	runes := []rune(v)
	if len(runes) == 0 {
		return false, nil
	}

	// Save original language name for error messages.
	originalLanguage := language

	// Normalize language name.
	language = strings.TrimSpace(language)

	// Empty language name is not valid.
	if language == "" {
		return false, ErrLanguageNotSupported(originalLanguage)
	}

	// Convert to lowercase for case-insensitive comparison.
	language = strings.ToLower(language)
	if alias, ok := languageAliases[language]; ok {
		language = alias
	}

	// Check if language is supported.
	if !isLanguageSupported(language) {
		return false, ErrLanguageNotSupported(originalLanguage)
	}

	// Get language config
	config, _ := languageConfigs[language]

	if langWords, exists := reservedWords[language]; exists {
		checkWord := v

		// Remove prefixes and suffixes for checking reserved words.
		switch language {
		case "php":
			if strings.HasPrefix(checkWord, "$") {
				checkWord = checkWord[1:]
			}
		case "ruby":
			if strings.HasPrefix(checkWord, "@@") {
				checkWord = checkWord[2:]
			} else if strings.HasPrefix(checkWord, "@") {
				checkWord = checkWord[1:]
			} else if strings.HasPrefix(checkWord, "$") {
				checkWord = checkWord[1:]
			}
			if strings.HasSuffix(checkWord, "?") {
				checkWord = checkWord[:len(checkWord)-1]
			} else if strings.HasSuffix(checkWord, "!") {
				checkWord = checkWord[:len(checkWord)-1]
			}
		}

		// Check for reserved words considering case sensitivity.
		if config.caseSensitive {
			if _, isReserved := langWords[checkWord]; isReserved {
				return false, nil
			}
		} else {
			checkWord = strings.ToLower(checkWord)
			for reserved := range langWords {
				if strings.ToLower(reserved) == checkWord {
					return false, nil
				}
			}
		}
	}

	// Special cases for Ruby.
	if language == "ruby" {
		// Check for class variables (@@).
		if strings.HasPrefix(v, "@@") {
			runes = runes[2:] // skip @@ for further checks
			if len(runes) == 0 {
				return false, nil
			}
			if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
				return false, nil
			}
		} else if strings.HasPrefix(v, "@") { // instance variables
			runes = runes[1:] // skip @ for further checks
			if len(runes) == 0 {
				return false, nil
			}
			if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
				return false, nil
			}
		} else if strings.HasPrefix(v, "$") { // global variables
			runes = runes[1:] // skip $ for further checks
			if len(runes) == 0 {
				return false, nil
			}
			if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
				return false, nil
			}
		}

		// If single ?, ! or only prefix without variable name.
		if len(runes) == 1 && (runes[0] == '?' || runes[0] == '!' ||
			runes[0] == '@' || runes[0] == '$') {
			return false, nil
		}

		// Check ending character ? or !
		runeCount := len(runes)
		for i, r := range runes {
			if r == '?' || r == '!' {
				// If ? or ! is the first character or not at the end.
				if i == 0 || i != runeCount-1 {
					return false, nil
				}
				// Check if there is at least one valid character before ? or !.
				prev := runes[i-1]
				if !unicode.IsLetter(prev) &&
					!unicode.IsNumber(prev) && prev != '_' {
					return false, nil
				}
			}
			if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' &&
				!(i == runeCount-1 && (r == '?' || r == '!')) {
				return false, nil
			}
		}

		return true, nil
	}

	// Check first character.
	if !config.checkFirst(runes[0]) {
		return false, nil
	}

	// Special case for PHP: if starts with $, check the next character.
	if language == "php" && runes[0] == '$' {
		if len(runes) < 2 {
			return false, nil
		}
		runes = runes[1:] // skip $ for further checks.
		if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
			return false, nil
		}
	}

	// Check remaining characters.
	for _, r := range runes {
		// For non-Unicode languages, check if character is ASCII.
		if !config.allowUnicode && r > unicode.MaxASCII {
			return false, nil
		}

		if !config.validChars(r) {
			return false, nil
		}
	}

	return true, nil
}

// VarFor is synonym for VariableNameFor function.
func VarFor(v string, language string) (bool, error) {
	return VariableNameFor(v, language)
}

// VariableName validates if the given string can be used as a
// variable name without considering specific programming language.
//
// Parameters:
//   - v: string to validate
//   - strict: if true, checks if the word is reserved in
//     any programming language.
//
// Returns true if the given name is valid.
func VariableName(v string, strict ...bool) bool {
	// Empty string is not a valid variable name.
	if v == "" {
		return false
	}

	// Default configuration
	config := languageConfig{
		caseSensitive: true,
		allowUnicode:  true,
		checkFirst: func(r rune) bool {
			return unicode.IsLetter(r) || r == '_'
		},
		validChars: func(r rune) bool {
			return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
		},
	}

	// Convert to runes for proper Unicode handling.
	runes := []rune(v)

	// Check first character
	if !config.checkFirst(runes[0]) {
		return false
	}

	// Check all characters explicitly
	for _, r := range runes {
		if !config.validChars(r) {
			return false
		}
	}

	// Check if word is reserved in any language when strict mode is enabled
	if g.All(strict...) && isReservedInAnyLanguage(v) {
		return false
	}

	return true
}

// Var is synonym for VariableName function.
func Var(v string, strict ...bool) bool {
	return VariableName(v, strict...)
}

// SelectorName checks if the given string is a valid CSS/HTML selector name.
// It supports simple selectors including elements (e.g., "div"), classes
// (e.g., ".myClass"), and IDs (e.g., "#myID").
//
// Parameters:
//   - v: string to validate;
//   - strict: if true, the leading characters # and . are inadmissible.
//
// Returns true if the given name is a valid selector.
func SelectorName(v string, strict ...bool) bool {
	if len(v) == 0 {
		return false
	}

	// Determine if strict mode is enabled.
	isStrict := g.All(strict...)

	// Select the appropriate regex based on the mode.
	rgex := g.If(isStrict, selectorStrictRegex, selectorRegex)
	isValidSelector := rgex.MatchString(v)

	// If valid selector, also check if it's not a reserved word in CSS.
	if isValidSelector {
		// Remove leading # or . for checking.
		checkWord := v
		if len(checkWord) > 0 && (checkWord[0] == '#' || checkWord[0] == '.') {
			checkWord = checkWord[1:]
		}

		// Check in CSS reserved words if strict mode is enabled.
		if isStrict {
			if cssWords, exists := reservedWords["css"]; exists {
				if _, reserved := cssWords[checkWord]; reserved {
					return false
				}
			}
		}
	}

	return isValidSelector
}

// Sel is synonym for SelectorName function.
func Sel(v string, strict ...bool) bool {
	return SelectorName(v, strict...)
}

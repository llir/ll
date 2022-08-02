// generated by Textmapper; DO NOT EDIT

package ll

import (
	"strings"
	"unicode/utf8"
)

// Lexer uses a generated DFA to scan through a utf-8 encoded input string. If
// the string starts with a BOM character, it gets skipped.
type Lexer struct {
	source string

	ch          rune // current character, -1 means EOI
	offset      int  // character offset
	tokenOffset int  // last token offset
	line        int  // current line number (1-based)
	tokenLine   int  // last token line
	scanOffset  int  // scanning offset
	value       interface{}

	State int // lexer state, modifiable
}

var bomSeq = "\xef\xbb\xbf"

// Init prepares the lexer l to tokenize source by performing the full reset
// of the internal state.
func (l *Lexer) Init(source string) {
	l.source = source

	l.ch = 0
	l.offset = 0
	l.tokenOffset = 0
	l.line = 1
	l.tokenLine = 1
	l.State = 0

	if strings.HasPrefix(source, bomSeq) {
		l.offset += len(bomSeq)
	}

	l.rewind(l.offset)
}

// Next finds and returns the next token in l.source. The source end is
// indicated by Token.EOI.
//
// The token text can be retrieved later by calling the Text() method.
func (l *Lexer) Next() Token {
restart:
	l.tokenLine = l.line
	l.tokenOffset = l.offset

	state := tmStateMap[l.State]
	hash := uint32(0)
	backupToken := -1
	var backupOffset int
	backupHash := hash
	for state >= 0 {
		var ch int
		if uint(l.ch) < tmRuneClassLen {
			ch = int(tmRuneClass[l.ch])
		} else if l.ch < 0 {
			state = int(tmLexerAction[state*tmNumClasses])
			continue
		} else {
			ch = 1
		}
		state = int(tmLexerAction[state*tmNumClasses+ch])
		if state > tmFirstRule {
			if state < 0 {
				state = (-1 - state) * 2
				backupToken = tmBacktracking[state]
				backupOffset = l.offset
				backupHash = hash
				state = tmBacktracking[state+1]
			}
			hash = hash*uint32(31) + uint32(l.ch)

			if l.ch == '\n' {
				l.line++
			}

			// Scan the next character.
			// Note: the following code is inlined to avoid performance implications.
			l.offset = l.scanOffset
			if l.offset < len(l.source) {
				r, w := rune(l.source[l.offset]), 1
				if r >= 0x80 {
					// not ASCII
					r, w = utf8.DecodeRuneInString(l.source[l.offset:])
				}
				l.scanOffset += w
				l.ch = r
			} else {
				l.ch = -1 // EOI
			}
		}
	}

	token := Token(tmFirstRule - state)
recovered:
	switch token {
	case LABEL_IDENT_TOK:
		hh := hash & 127
		switch hh {
		case 0:
			if hash == 0x5bd73a00 && "declaration:" == l.source[l.tokenOffset:l.offset] {
				token = DECLARATIONCOLON
				break
			}
		case 1:
			if hash == 0x8fd67c81 && "stride:" == l.source[l.tokenOffset:l.offset] {
				token = STRIDECOLON
				break
			}
		case 7:
			if hash == 0xa3b52e07 && "offset:" == l.source[l.tokenOffset:l.offset] {
				token = OFFSETCOLON
				break
			}
		case 10:
			if hash == 0xdfb6b50a && "globals:" == l.source[l.tokenOffset:l.offset] {
				token = GLOBALSCOLON
				break
			}
		case 11:
			if hash == 0xaf42d10b && "count:" == l.source[l.tokenOffset:l.offset] {
				token = COUNTCOLON
				break
			}
			if hash == 0xfcc78e8b && "splitDebugInlining:" == l.source[l.tokenOffset:l.offset] {
				token = SPLITDEBUGINLININGCOLON
				break
			}
		case 12:
			if hash == 0xe6d32c8c && "imports:" == l.source[l.tokenOffset:l.offset] {
				token = IMPORTSCOLON
				break
			}
		case 14:
			if hash == 0x674398e && "rank:" == l.source[l.tokenOffset:l.offset] {
				token = RANKCOLON
				break
			}
		case 15:
			if hash == 0xfb94388f && "getter:" == l.source[l.tokenOffset:l.offset] {
				token = GETTERCOLON
				break
			}
			if hash == 0x63bd70f && "name:" == l.source[l.tokenOffset:l.offset] {
				token = NAMECOLON
				break
			}
		case 19:
			if hash == 0x3cf10193 && "exportSymbols:" == l.source[l.tokenOffset:l.offset] {
				token = EXPORTSYMBOLSCOLON
				break
			}
			if hash == 0xb255ef13 && "filename:" == l.source[l.tokenOffset:l.offset] {
				token = FILENAMECOLON
				break
			}
			if hash == 0xb42de293 && "flags:" == l.source[l.tokenOffset:l.offset] {
				token = FLAGSCOLON
				break
			}
			if hash == 0xae94893 && "virtualIndex:" == l.source[l.tokenOffset:l.offset] {
				token = VIRTUALINDEXCOLON
				break
			}
		case 20:
			if hash == 0x724b1314 && "runtimeLang:" == l.source[l.tokenOffset:l.offset] {
				token = RUNTIMELANGCOLON
				break
			}
		case 22:
			if hash == 0xf01e996 && "annotations:" == l.source[l.tokenOffset:l.offset] {
				token = ANNOTATIONSCOLON
				break
			}
		case 26:
			if hash == 0x28737d9a && "defaulted:" == l.source[l.tokenOffset:l.offset] {
				token = DEFAULTEDCOLON
				break
			}
			if hash == 0x7e83e59a && "templateParams:" == l.source[l.tokenOffset:l.offset] {
				token = TEMPLATEPARAMSCOLON
				break
			}
		case 27:
			if hash == 0x4d09869b && "dataLocation:" == l.source[l.tokenOffset:l.offset] {
				token = DATALOCATIONCOLON
				break
			}
			if hash == 0x765eff9b && "setter:" == l.source[l.tokenOffset:l.offset] {
				token = SETTERCOLON
				break
			}
		case 30:
			if hash == 0x5ef58d1e && "discriminator:" == l.source[l.tokenOffset:l.offset] {
				token = DISCRIMINATORCOLON
				break
			}
		case 32:
			if hash == 0x2b0e8ba0 && "extraData:" == l.source[l.tokenOffset:l.offset] {
				token = EXTRADATACOLON
				break
			}
		case 35:
			if hash == 0xb1651ea3 && "dwoId:" == l.source[l.tokenOffset:l.offset] {
				token = DWOIDCOLON
				break
			}
			if hash == 0xf0975e23 && "elements:" == l.source[l.tokenOffset:l.offset] {
				token = ELEMENTSCOLON
				break
			}
		case 36:
			if hash == 0x2dd0a4 && "arg:" == l.source[l.tokenOffset:l.offset] {
				token = ARGCOLON
				break
			}
			if hash == 0x389b97a4 && "column:" == l.source[l.tokenOffset:l.offset] {
				token = COLUMNCOLON
				break
			}
		case 41:
			if hash == 0xc20043a9 && "nodes:" == l.source[l.tokenOffset:l.offset] {
				token = NODESCOLON
				break
			}
		case 42:
			if hash == 0x75a2f2aa && "debugBaseAddress:" == l.source[l.tokenOffset:l.offset] {
				token = DEBUGBASEADDRESSCOLON
				break
			}
		case 45:
			if hash == 0xcc4d9ead && "rangesBaseAddress:" == l.source[l.tokenOffset:l.offset] {
				token = RANGESBASEADDRESSCOLON
				break
			}
			if hash == 0xae4e252d && "retainedNodes:" == l.source[l.tokenOffset:l.offset] {
				token = RETAINEDNODESCOLON
				break
			}
		case 48:
			if hash == 0x631808b0 && "operands:" == l.source[l.tokenOffset:l.offset] {
				token = OPERANDSCOLON
				break
			}
		case 49:
			if hash == 0x518ce8b1 && "identifier:" == l.source[l.tokenOffset:l.offset] {
				token = IDENTIFIERCOLON
				break
			}
		case 51:
			if hash == 0x14e78833 && "apinotes:" == l.source[l.tokenOffset:l.offset] {
				token = APINOTESCOLON
				break
			}
		case 54:
			if hash == 0x6a45736 && "unit:" == l.source[l.tokenOffset:l.offset] {
				token = UNITCOLON
				break
			}
		case 55:
			if hash == 0xa11e98b7 && "entity:" == l.source[l.tokenOffset:l.offset] {
				token = ENTITYCOLON
				break
			}
		case 58:
			if hash == 0x31bd57ba && "splitDebugFilename:" == l.source[l.tokenOffset:l.offset] {
				token = SPLITDEBUGFILENAMECOLON
				break
			}
		case 61:
			if hash == 0x599211bd && "isDefinition:" == l.source[l.tokenOffset:l.offset] {
				token = ISDEFINITIONCOLON
				break
			}
			if hash == 0xcfab18bd && "isOptimized:" == l.source[l.tokenOffset:l.offset] {
				token = ISOPTIMIZEDCOLON
				break
			}
		case 62:
			if hash == 0x5ceba3e && "file:" == l.source[l.tokenOffset:l.offset] {
				token = FILECOLON
				break
			}
			if hash == 0xcbe7b9be && "upperBound:" == l.source[l.tokenOffset:l.offset] {
				token = UPPERBOUNDCOLON
				break
			}
		case 64:
			if hash == 0x35cb40 && "sdk:" == l.source[l.tokenOffset:l.offset] {
				token = SDKCOLON
				break
			}
			if hash == 0x69b5840 && "type:" == l.source[l.tokenOffset:l.offset] {
				token = TYPECOLON
				break
			}
		case 67:
			if hash == 0xed8781c3 && "attributes:" == l.source[l.tokenOffset:l.offset] {
				token = ATTRIBUTESCOLON
				break
			}
			if hash == 0x23880cc3 && "checksumkind:" == l.source[l.tokenOffset:l.offset] {
				token = CHECKSUMKINDCOLON
				break
			}
			if hash == 0x94454c3 && "stringLength:" == l.source[l.tokenOffset:l.offset] {
				token = STRINGLENGTHCOLON
				break
			}
		case 70:
			if hash == 0xc58c5f46 && "dwarfAddressSpace:" == l.source[l.tokenOffset:l.offset] {
				token = DWARFADDRESSSPACECOLON
				break
			}
			if hash == 0x7a9945c6 && "isDecl:" == l.source[l.tokenOffset:l.offset] {
				token = ISDECLCOLON
				break
			}
			if hash == 0x6234ec6 && "line:" == l.source[l.tokenOffset:l.offset] {
				token = LINECOLON
				break
			}
		case 71:
			if hash == 0x59e67d47 && "encoding:" == l.source[l.tokenOffset:l.offset] {
				token = ENCODINGCOLON
				break
			}
		case 73:
			if hash == 0xa6960349 && "thrownTypes:" == l.source[l.tokenOffset:l.offset] {
				token = THROWNTYPESCOLON
				break
			}
		case 74:
			if hash == 0x96ae4b4a && "linkageName:" == l.source[l.tokenOffset:l.offset] {
				token = LINKAGENAMECOLON
				break
			}
		case 75:
			if hash == 0x6b2f65cb && "stringLengthExpression:" == l.source[l.tokenOffset:l.offset] {
				token = STRINGLENGTHEXPRESSIONCOLON
				break
			}
		case 77:
			if hash == 0x2f676f4d && "header:" == l.source[l.tokenOffset:l.offset] {
				token = HEADERCOLON
				break
			}
		case 80:
			if hash == 0x4c73e350 && "spFlags:" == l.source[l.tokenOffset:l.offset] {
				token = SPFLAGSCOLON
				break
			}
		case 81:
			if hash == 0x12532d51 && "configMacros:" == l.source[l.tokenOffset:l.offset] {
				token = CONFIGMACROSCOLON
				break
			}
		case 82:
			if hash == 0x4e37e52 && "containingType:" == l.source[l.tokenOffset:l.offset] {
				token = CONTAININGTYPECOLON
				break
			}
			if hash == 0xf38af4d2 && "scopeLine:" == l.source[l.tokenOffset:l.offset] {
				token = SCOPELINECOLON
				break
			}
		case 83:
			if hash == 0x8e024053 && "allocated:" == l.source[l.tokenOffset:l.offset] {
				token = ALLOCATEDCOLON
				break
			}
		case 85:
			if hash == 0xaba949d5 && "align:" == l.source[l.tokenOffset:l.offset] {
				token = ALIGNCOLON
				break
			}
		case 86:
			if hash == 0x88a9ffd6 && "vtableHolder:" == l.source[l.tokenOffset:l.offset] {
				token = VTABLEHOLDERCOLON
				break
			}
		case 87:
			if hash == 0x17d02c57 && "checksum:" == l.source[l.tokenOffset:l.offset] {
				token = CHECKSUMCOLON
				break
			}
			if hash == 0xc4ee1ed7 && "virtuality:" == l.source[l.tokenOffset:l.offset] {
				token = VIRTUALITYCOLON
				break
			}
		case 89:
			if hash == 0x6862059 && "size:" == l.source[l.tokenOffset:l.offset] {
				token = SIZECOLON
				break
			}
		case 90:
			if hash == 0x17fda && "cc:" == l.source[l.tokenOffset:l.offset] {
				token = CCCOLON
				break
			}
			if hash == 0x58795a5a && "runtimeVersion:" == l.source[l.tokenOffset:l.offset] {
				token = RUNTIMEVERSIONCOLON
				break
			}
		case 91:
			if hash == 0x28bdb9db && "isUnsigned:" == l.source[l.tokenOffset:l.offset] {
				token = ISUNSIGNEDCOLON
				break
			}
		case 92:
			if hash == 0x9bf65c && "debugInfoForProfiling:" == l.source[l.tokenOffset:l.offset] {
				token = DEBUGINFOFORPROFILINGCOLON
				break
			}
			if hash == 0xdb62f9dc && "stringLocationExpression:" == l.source[l.tokenOffset:l.offset] {
				token = STRINGLOCATIONEXPRESSIONCOLON
				break
			}
		case 93:
			if hash == 0xa97308dd && "lowerBound:" == l.source[l.tokenOffset:l.offset] {
				token = LOWERBOUNDCOLON
				break
			}
		case 95:
			if hash == 0x15146ddf && "emissionKind:" == l.source[l.tokenOffset:l.offset] {
				token = EMISSIONKINDCOLON
				break
			}
		case 96:
			if hash == 0x3633e0 && "tag:" == l.source[l.tokenOffset:l.offset] {
				token = TAGCOLON
				break
			}
		case 97:
			if hash == 0xcccfb6e1 && "types:" == l.source[l.tokenOffset:l.offset] {
				token = TYPESCOLON
				break
			}
		case 98:
			if hash == 0x5a7fd7e2 && "language:" == l.source[l.tokenOffset:l.offset] {
				token = LANGUAGECOLON
				break
			}
		case 99:
			if hash == 0x8dd720e3 && "nameTableKind:" == l.source[l.tokenOffset:l.offset] {
				token = NAMETABLEKINDCOLON
				break
			}
		case 101:
			if hash == 0x5c784e5 && "expr:" == l.source[l.tokenOffset:l.offset] {
				token = EXPRCOLON
				break
			}
			if hash == 0xb91d9865 && "retainedTypes:" == l.source[l.tokenOffset:l.offset] {
				token = RETAINEDTYPESCOLON
				break
			}
		case 102:
			if hash == 0xc9e48c66 && "scope:" == l.source[l.tokenOffset:l.offset] {
				token = SCOPECOLON
				break
			}
		case 104:
			if hash == 0xb29e66e8 && "enums:" == l.source[l.tokenOffset:l.offset] {
				token = ENUMSCOLON
				break
			}
			if hash == 0xc14e22e8 && "producer:" == l.source[l.tokenOffset:l.offset] {
				token = PRODUCERCOLON
				break
			}
		case 105:
			if hash == 0xcee59ce9 && "value:" == l.source[l.tokenOffset:l.offset] {
				token = VALUECOLON
				break
			}
		case 107:
			if hash == 0x75b114eb && "sysroot:" == l.source[l.tokenOffset:l.offset] {
				token = SYSROOTCOLON
				break
			}
		case 109:
			if hash == 0xd63af6d && "directory:" == l.source[l.tokenOffset:l.offset] {
				token = DIRECTORYCOLON
				break
			}
			if hash == 0xfc2e84ed && "includePath:" == l.source[l.tokenOffset:l.offset] {
				token = INCLUDEPATHCOLON
				break
			}
		case 111:
			if hash == 0x9322fbef && "baseType:" == l.source[l.tokenOffset:l.offset] {
				token = BASETYPECOLON
				break
			}
			if hash == 0x6aa06aef && "thisAdjustment:" == l.source[l.tokenOffset:l.offset] {
				token = THISADJUSTMENTCOLON
				break
			}
		case 115:
			if hash == 0x3135e6f3 && "macros:" == l.source[l.tokenOffset:l.offset] {
				token = MACROSCOLON
				break
			}
			if hash == 0x371df3 && "var:" == l.source[l.tokenOffset:l.offset] {
				token = VARCOLON
				break
			}
		case 116:
			if hash == 0xfac8cbf4 && "associated:" == l.source[l.tokenOffset:l.offset] {
				token = ASSOCIATEDCOLON
				break
			}
		case 121:
			if hash == 0xe6c2fff9 && "isLocal:" == l.source[l.tokenOffset:l.offset] {
				token = ISLOCALCOLON
				break
			}
		case 124:
			if hash == 0x2e67257c && "inlinedAt:" == l.source[l.tokenOffset:l.offset] {
				token = INLINEDATCOLON
				break
			}
		case 126:
			if hash == 0x4714cffe && "isImplicitCode:" == l.source[l.tokenOffset:l.offset] {
				token = ISIMPLICITCODECOLON
				break
			}
		case 127:
			if hash == 0x877c9b7f && "source:" == l.source[l.tokenOffset:l.offset] {
				token = SOURCECOLON
				break
			}
		}
	case METADATA_NAME_TOK:
		hh := hash & 31
		switch hh {
		case 1:
			if hash == 0xe9c3fc41 && "!DILocalVariable" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDILOCALVARIABLE
				break
			}
		case 2:
			if hash == 0x45c23ec2 && "!DIBasicType" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIBASICTYPE
				break
			}
			if hash == 0x49dd4b02 && "!DIFile" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIFILE
				break
			}
			if hash == 0x664d6da2 && "!DIMacroFile" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIMACROFILE
				break
			}
		case 3:
			if hash == 0xe56a8583 && "!DISubrange" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDISUBRANGE
				break
			}
		case 4:
			if hash == 0xa1dae1e4 && "!DISubroutineType" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDISUBROUTINETYPE
				break
			}
		case 5:
			if hash == 0x5fa48d45 && "!DIGlobalVariable" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIGLOBALVARIABLE
				break
			}
		case 6:
			if hash == 0xf22af706 && "!DIMacro" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIMACRO
				break
			}
		case 7:
			if hash == 0xb3e7b7a7 && "!DIObjCProperty" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIOBJCPROPERTY
				break
			}
		case 10:
			if hash == 0x7f3c70ca && "!DIEnumerator" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIENUMERATOR
				break
			}
			if hash == 0xe823714a && "!DISubprogram" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDISUBPROGRAM
				break
			}
		case 11:
			if hash == 0xf77c972b && "!DILexicalBlockFile" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDILEXICALBLOCKFILE
				break
			}
		case 13:
			if hash == 0xec9f4e0d && "!DIImportedEntity" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIIMPORTEDENTITY
				break
			}
		case 14:
			if hash == 0xb5a7f8ce && "!DIArgList" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIARGLIST
				break
			}
			if hash == 0xf21cda2e && "!DILabel" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDILABEL
				break
			}
		case 15:
			if hash == 0xe4764d0f && "!DILexicalBlock" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDILEXICALBLOCK
				break
			}
			if hash == 0x4eb4666f && "!DITemplateTypeParameter" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDITEMPLATETYPEPARAMETER
				break
			}
		case 17:
			if hash == 0x91d7de11 && "!DICompileUnit" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDICOMPILEUNIT
				break
			}
			if hash == 0x4dbccab1 && "!DIDerivedType" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIDERIVEDTYPE
				break
			}
			if hash == 0x58598bf1 && "!DIStringType" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDISTRINGTYPE
				break
			}
		case 18:
			if hash == 0x53f9b272 && "!DIModule" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIMODULE
				break
			}
		case 21:
			if hash == 0xf5ac8395 && "!DINamespace" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDINAMESPACE
				break
			}
		case 24:
			if hash == 0xd990cc58 && "!DITemplateValueParameter" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDITEMPLATEVALUEPARAMETER
				break
			}
		case 27:
			if hash == 0xd4125b9b && "!DICompositeType" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDICOMPOSITETYPE
				break
			}
			if hash == 0xd1aac23b && "!DILocation" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDILOCATION
				break
			}
		case 28:
			if hash == 0x9920ad1c && "!DICommonBlock" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDICOMMONBLOCK
				break
			}
		case 29:
			if hash == 0x4689e9bd && "!DIGlobalVariableExpression" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIGLOBALVARIABLEEXPRESSION
				break
			}
			if hash == 0x37539a1d && "!GenericDINode" == l.source[l.tokenOffset:l.offset] {
				token = EXCLGENERICDINODE
				break
			}
		case 30:
			if hash == 0x4b182b9e && "!DIExpression" == l.source[l.tokenOffset:l.offset] {
				token = EXCLDIEXPRESSION
				break
			}
		}
	}
	switch token {
	case INVALID_TOKEN:
		if backupToken >= 0 {
			token = Token(backupToken)
			hash = backupHash
			l.rewind(backupOffset)
		} else if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
		if token != INVALID_TOKEN {
			goto recovered
		}
	case 2, 3:
		goto restart
	}
	return token
}

// Pos returns the start and end positions of the last token returned by Next().
func (l *Lexer) Pos() (start, end int) {
	start = l.tokenOffset
	end = l.offset
	return
}

// Line returns the line number of the last token returned by Next() (1-based).
func (l *Lexer) Line() int {
	return l.tokenLine
}

// Text returns the substring of the input corresponding to the last token.
func (l *Lexer) Text() string {
	return l.source[l.tokenOffset:l.offset]
}

// Value returns the value associated with the last returned token.
func (l *Lexer) Value() interface{} {
	return l.value
}

// rewind can be used in lexer actions to accept a portion of a scanned token, or to include
// more text into it.
func (l *Lexer) rewind(offset int) {
	if offset < l.offset {
		l.line -= strings.Count(l.source[offset:l.offset], "\n")
	} else {
		if offset > len(l.source) {
			offset = len(l.source)
		}
		l.line += strings.Count(l.source[l.offset:offset], "\n")
	}

	// Scan the next character.
	l.scanOffset = offset
	l.offset = offset
	if l.offset < len(l.source) {
		r, w := rune(l.source[l.offset]), 1
		if r >= 0x80 {
			// not ASCII
			r, w = utf8.DecodeRuneInString(l.source[l.offset:])
		}
		l.scanOffset += w
		l.ch = r
	} else {
		l.ch = -1 // EOI
	}
}

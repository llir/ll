// generated by Textmapper; DO NOT EDIT

package ll

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the llvm language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
	EOI
	INVALID_TOKEN
	COMMENT
	WHITESPACE
	GLOBAL_IDENT_TOK
	LOCAL_IDENT_TOK
	LABEL_IDENT_TOK
	ATTR_GROUP_ID_TOK
	COMDAT_NAME_TOK
	METADATA_NAME_TOK
	METADATA_ID_TOK
	DWARF_TAG_TOK
	DWARF_ATT_ENCODING_TOK
	DI_FLAG_TOK
	DISP_FLAG_TOK
	DWARF_LANG_TOK
	DWARF_CC_TOK
	CHECKSUM_KIND_TOK
	DWARF_VIRTUALITY_TOK
	DWARF_MACINFO_TOK
	DWARF_OP_TOK
	EMISSION_KIND_TOK
	NAME_TABLE_KIND_TOK
	INT_LIT_TOK
	FLOAT_LIT_TOK
	STRING_LIT_TOK
	INT_TYPE_TOK
	AARCH64_VECTOR_PCS             // aarch64_vector_pcs
	ACQ_REL                        // acq_rel
	ACQUIRE                        // acquire
	ADD                            // add
	ADDRSPACE                      // addrspace
	ADDRSPACECAST                  // addrspacecast
	AFN                            // afn
	ALIAS                          // alias
	ALIGN                          // align
	ALIGNSTACK                     // alignstack
	ALLOCA                         // alloca
	ALLOCSIZE                      // allocsize
	ALWAYSINLINE                   // alwaysinline
	AMDGPU_CS                      // amdgpu_cs
	AMDGPU_ES                      // amdgpu_es
	AMDGPU_GS                      // amdgpu_gs
	AMDGPU_HS                      // amdgpu_hs
	AMDGPU_KERNEL                  // amdgpu_kernel
	AMDGPU_LS                      // amdgpu_ls
	AMDGPU_PS                      // amdgpu_ps
	AMDGPU_VS                      // amdgpu_vs
	AND                            // and
	ANY                            // any
	ANYREGCC                       // anyregcc
	APPENDING                      // appending
	ARCP                           // arcp
	ARGMEMONLY                     // argmemonly
	ARM_AAPCS_VFPCC                // arm_aapcs_vfpcc
	ARM_AAPCSCC                    // arm_aapcscc
	ARM_APCSCC                     // arm_apcscc
	ASHR                           // ashr
	ASM                            // asm
	ATOMIC                         // atomic
	ATOMICRMW                      // atomicrmw
	ATTRIBUTES                     // attributes
	AVAILABLE_EXTERNALLY           // available_externally
	AVR_INTRCC                     // avr_intrcc
	AVR_SIGNALCC                   // avr_signalcc
	BITCAST                        // bitcast
	BLOCKADDRESS                   // blockaddress
	BR                             // br
	BUILTIN                        // builtin
	BYVAL                          // byval
	CHAR_C                         // c
	CALL                           // call
	CALLBR                         // callbr
	CALLER                         // caller
	CATCH                          // catch
	CATCHPAD                       // catchpad
	CATCHRET                       // catchret
	CATCHSWITCH                    // catchswitch
	CC                             // cc
	CCC                            // ccc
	CLEANUP                        // cleanup
	CLEANUPPAD                     // cleanuppad
	CLEANUPRET                     // cleanupret
	CMPXCHG                        // cmpxchg
	COLD                           // cold
	COLDCC                         // coldcc
	COMDAT                         // comdat
	COMMON                         // common
	CONSTANT                       // constant
	CONTRACT                       // contract
	CONVERGENT                     // convergent
	CXX_FAST_TLSCC                 // cxx_fast_tlscc
	DATALAYOUT                     // datalayout
	DECLARE                        // declare
	DEFAULT                        // default
	DEFINE                         // define
	DEREFERENCEABLE                // dereferenceable
	DEREFERENCEABLE_OR_NULL        // dereferenceable_or_null
	DISTINCT                       // distinct
	DLLEXPORT                      // dllexport
	DLLIMPORT                      // dllimport
	DOUBLE                         // double
	DSO_LOCAL                      // dso_local
	DSO_PREEMPTABLE                // dso_preemptable
	EQ                             // eq
	EXACT                          // exact
	EXACTMATCH                     // exactmatch
	EXTERN_WEAK                    // extern_weak
	EXTERNAL                       // external
	EXTERNALLY_INITIALIZED         // externally_initialized
	EXTRACTELEMENT                 // extractelement
	EXTRACTVALUE                   // extractvalue
	FADD                           // fadd
	FALSE                          // false
	FAST                           // fast
	FASTCC                         // fastcc
	FCMP                           // fcmp
	FDIV                           // fdiv
	FENCE                          // fence
	FILTER                         // filter
	FLOAT                          // float
	FMUL                           // fmul
	FNEG                           // fneg
	FP128                          // fp128
	FPEXT                          // fpext
	FPTOSI                         // fptosi
	FPTOUI                         // fptoui
	FPTRUNC                        // fptrunc
	FREM                           // frem
	FROM                           // from
	FSUB                           // fsub
	GC                             // gc
	GETELEMENTPTR                  // getelementptr
	GHCCC                          // ghccc
	GLOBAL                         // global
	HALF                           // half
	HHVM_CCC                       // hhvm_ccc
	HHVMCC                         // hhvmcc
	HIDDEN                         // hidden
	ICMP                           // icmp
	IFUNC                          // ifunc
	IMMARG                         // immarg
	INACCESSIBLEMEM_OR_ARGMEMONLY  // inaccessiblemem_or_argmemonly
	INACCESSIBLEMEMONLY            // inaccessiblememonly
	INALLOCA                       // inalloca
	INBOUNDS                       // inbounds
	INDIRECTBR                     // indirectbr
	INITIALEXEC                    // initialexec
	INLINEHINT                     // inlinehint
	INRANGE                        // inrange
	INREG                          // inreg
	INSERTELEMENT                  // insertelement
	INSERTVALUE                    // insertvalue
	INTEL_OCL_BICC                 // intel_ocl_bicc
	INTELDIALECT                   // inteldialect
	INTERNAL                       // internal
	INTTOPTR                       // inttoptr
	INVOKE                         // invoke
	JUMPTABLE                      // jumptable
	LABEL                          // label
	LANDINGPAD                     // landingpad
	LARGEST                        // largest
	LINKONCE                       // linkonce
	LINKONCE_ODR                   // linkonce_odr
	LOAD                           // load
	LOCAL_UNNAMED_ADDR             // local_unnamed_addr
	LOCALDYNAMIC                   // localdynamic
	LOCALEXEC                      // localexec
	LSHR                           // lshr
	MAX                            // max
	METADATA                       // metadata
	MIN                            // min
	MINSIZE                        // minsize
	MODULE                         // module
	MONOTONIC                      // monotonic
	MSP430_INTRCC                  // msp430_intrcc
	MUL                            // mul
	MUSTTAIL                       // musttail
	NAKED                          // naked
	NAND                           // nand
	NE                             // ne
	NEST                           // nest
	NINF                           // ninf
	NNAN                           // nnan
	NOALIAS                        // noalias
	NOBUILTIN                      // nobuiltin
	NOCAPTURE                      // nocapture
	NOCF_CHECK                     // nocf_check
	NODUPLICATE                    // noduplicate
	NODUPLICATES                   // noduplicates
	NOFREE                         // nofree
	NOIMPLICITFLOAT                // noimplicitfloat
	NOINLINE                       // noinline
	NONE                           // none
	NONLAZYBIND                    // nonlazybind
	NONNULL                        // nonnull
	NORECURSE                      // norecurse
	NOREDZONE                      // noredzone
	NORETURN                       // noreturn
	NOSYNC                         // nosync
	NOTAIL                         // notail
	NOUNWIND                       // nounwind
	NSW                            // nsw
	NSZ                            // nsz
	NULL                           // null
	NUW                            // nuw
	OEQ                            // oeq
	OGE                            // oge
	OGT                            // ogt
	OLE                            // ole
	OLT                            // olt
	ONE                            // one
	OPAQUE                         // opaque
	OPTFORFUZZING                  // optforfuzzing
	OPTNONE                        // optnone
	OPTSIZE                        // optsize
	OR                             // or
	ORD                            // ord
	PARTITION                      // partition
	PERSONALITY                    // personality
	PHI                            // phi
	PPC_FP128                      // ppc_fp128
	PREFIX                         // prefix
	PRESERVE_ALLCC                 // preserve_allcc
	PRESERVE_MOSTCC                // preserve_mostcc
	PRIVATE                        // private
	PROLOGUE                       // prologue
	PROTECTED                      // protected
	PTRTOINT                       // ptrtoint
	PTX_DEVICE                     // ptx_device
	PTX_KERNEL                     // ptx_kernel
	READNONE                       // readnone
	READONLY                       // readonly
	REASSOC                        // reassoc
	RELEASE                        // release
	RESUME                         // resume
	RET                            // ret
	RETURNED                       // returned
	RETURNS_TWICE                  // returns_twice
	SAFESTACK                      // safestack
	SAMESIZE                       // samesize
	SANITIZE_ADDRESS               // sanitize_address
	SANITIZE_HWADDRESS             // sanitize_hwaddress
	SANITIZE_MEMORY                // sanitize_memory
	SANITIZE_MEMTAG                // sanitize_memtag
	SANITIZE_THREAD                // sanitize_thread
	SDIV                           // sdiv
	SECTION                        // section
	SELECT                         // select
	SEQ_CST                        // seq_cst
	SEXT                           // sext
	SGE                            // sge
	SGT                            // sgt
	SHADOWCALLSTACK                // shadowcallstack
	SHL                            // shl
	SHUFFLEVECTOR                  // shufflevector
	SIDEEFFECT                     // sideeffect
	SIGNEXT                        // signext
	SINGLETHREAD                   // singlethread
	SITOFP                         // sitofp
	SLE                            // sle
	SLT                            // slt
	SOURCE_FILENAME                // source_filename
	SPECULATABLE                   // speculatable
	SPECULATIVE_LOAD_HARDENING     // speculative_load_hardening
	SPIR_FUNC                      // spir_func
	SPIR_KERNEL                    // spir_kernel
	SREM                           // srem
	SRET                           // sret
	SSP                            // ssp
	SSPREQ                         // sspreq
	SSPSTRONG                      // sspstrong
	STORE                          // store
	STRICTFP                       // strictfp
	SUB                            // sub
	SWIFTCC                        // swiftcc
	SWIFTERROR                     // swifterror
	SWIFTSELF                      // swiftself
	SWITCH                         // switch
	SYNCSCOPE                      // syncscope
	TAIL                           // tail
	TARGET                         // target
	THREAD_LOCAL                   // thread_local
	TO                             // to
	TOKEN                          // token
	TRIPLE                         // triple
	TRUE                           // true
	TRUNC                          // trunc
	TYPE                           // type
	UDIV                           // udiv
	UEQ                            // ueq
	UGE                            // uge
	UGT                            // ugt
	UITOFP                         // uitofp
	ULE                            // ule
	ULT                            // ult
	UMAX                           // umax
	UMIN                           // umin
	UNDEF                          // undef
	UNE                            // une
	UNNAMED_ADDR                   // unnamed_addr
	UNO                            // uno
	UNORDERED                      // unordered
	UNREACHABLE                    // unreachable
	UNWIND                         // unwind
	UREM                           // urem
	USELISTORDER                   // uselistorder
	USELISTORDER_BB                // uselistorder_bb
	UWTABLE                        // uwtable
	VA_ARG                         // va_arg
	VOID                           // void
	VOLATILE                       // volatile
	VSCALE                         // vscale
	WEAK                           // weak
	WEAK_ODR                       // weak_odr
	WEBKIT_JSCC                    // webkit_jscc
	WILLRETURN                     // willreturn
	WIN64CC                        // win64cc
	WITHIN                         // within
	WRITEONLY                      // writeonly
	CHAR_X                         // x
	X86_64_SYSVCC                  // x86_64_sysvcc
	X86_FASTCALLCC                 // x86_fastcallcc
	X86_FP80                       // x86_fp80
	X86_INTRCC                     // x86_intrcc
	X86_MMX                        // x86_mmx
	X86_REGCALLCC                  // x86_regcallcc
	X86_STDCALLCC                  // x86_stdcallcc
	X86_THISCALLCC                 // x86_thiscallcc
	X86_VECTORCALLCC               // x86_vectorcallcc
	XCHG                           // xchg
	XOR                            // xor
	ZEROEXT                        // zeroext
	ZEROINITIALIZER                // zeroinitializer
	ZEXT                           // zext
	EXCLDIBASICTYPE                // !DIBasicType
	EXCLDICOMMONBLOCK              // !DICommonBlock
	EXCLDICOMPILEUNIT              // !DICompileUnit
	EXCLDICOMPOSITETYPE            // !DICompositeType
	EXCLDIDERIVEDTYPE              // !DIDerivedType
	EXCLDIENUMERATOR               // !DIEnumerator
	EXCLDIEXPRESSION               // !DIExpression
	EXCLDIFILE                     // !DIFile
	EXCLDIGLOBALVARIABLE           // !DIGlobalVariable
	EXCLDIGLOBALVARIABLEEXPRESSION // !DIGlobalVariableExpression
	EXCLDIIMPORTEDENTITY           // !DIImportedEntity
	EXCLDILABEL                    // !DILabel
	EXCLDILEXICALBLOCK             // !DILexicalBlock
	EXCLDILEXICALBLOCKFILE         // !DILexicalBlockFile
	EXCLDILOCALVARIABLE            // !DILocalVariable
	EXCLDILOCATION                 // !DILocation
	EXCLDIMACRO                    // !DIMacro
	EXCLDIMACROFILE                // !DIMacroFile
	EXCLDIMODULE                   // !DIModule
	EXCLDINAMESPACE                // !DINamespace
	EXCLDIOBJCPROPERTY             // !DIObjCProperty
	EXCLDISUBPROGRAM               // !DISubprogram
	EXCLDISUBRANGE                 // !DISubrange
	EXCLDISUBROUTINETYPE           // !DISubroutineType
	EXCLDITEMPLATETYPEPARAMETER    // !DITemplateTypeParameter
	EXCLDITEMPLATEVALUEPARAMETER   // !DITemplateValueParameter
	EXCLGENERICDINODE              // !GenericDINode
	ALIGNCOLON                     // align:
	ARGCOLON                       // arg:
	ATTRIBUTESCOLON                // attributes:
	BASETYPECOLON                  // baseType:
	CCCOLON                        // cc:
	CHECKSUMCOLON                  // checksum:
	CHECKSUMKINDCOLON              // checksumkind:
	COLUMNCOLON                    // column:
	CONFIGMACROSCOLON              // configMacros:
	CONTAININGTYPECOLON            // containingType:
	COUNTCOLON                     // count:
	DEBUGBASEADDRESSCOLON          // debugBaseAddress:
	DEBUGINFOFORPROFILINGCOLON     // debugInfoForProfiling:
	DECLARATIONCOLON               // declaration:
	DIRECTORYCOLON                 // directory:
	DISCRIMINATORCOLON             // discriminator:
	DWARFADDRESSSPACECOLON         // dwarfAddressSpace:
	DWOIDCOLON                     // dwoId:
	ELEMENTSCOLON                  // elements:
	EMISSIONKINDCOLON              // emissionKind:
	ENCODINGCOLON                  // encoding:
	ENTITYCOLON                    // entity:
	ENUMSCOLON                     // enums:
	EXPORTSYMBOLSCOLON             // exportSymbols:
	EXPRCOLON                      // expr:
	EXTRADATACOLON                 // extraData:
	FILECOLON                      // file:
	FILENAMECOLON                  // filename:
	FLAGSCOLON                     // flags:
	GETTERCOLON                    // getter:
	GLOBALSCOLON                   // globals:
	HEADERCOLON                    // header:
	IDENTIFIERCOLON                // identifier:
	IMPORTSCOLON                   // imports:
	INCLUDEPATHCOLON               // includePath:
	INLINEDATCOLON                 // inlinedAt:
	ISDEFINITIONCOLON              // isDefinition:
	ISIMPLICITCODECOLON            // isImplicitCode:
	ISLOCALCOLON                   // isLocal:
	ISOPTIMIZEDCOLON               // isOptimized:
	ISUNSIGNEDCOLON                // isUnsigned:
	ISYSROOTCOLON                  // isysroot:
	LANGUAGECOLON                  // language:
	LINECOLON                      // line:
	LINKAGENAMECOLON               // linkageName:
	LOWERBOUNDCOLON                // lowerBound:
	MACROSCOLON                    // macros:
	NAMECOLON                      // name:
	NAMETABLEKINDCOLON             // nameTableKind:
	NODESCOLON                     // nodes:
	OFFSETCOLON                    // offset:
	OPERANDSCOLON                  // operands:
	PRODUCERCOLON                  // producer:
	RETAINEDNODESCOLON             // retainedNodes:
	RETAINEDTYPESCOLON             // retainedTypes:
	RUNTIMELANGCOLON               // runtimeLang:
	RUNTIMEVERSIONCOLON            // runtimeVersion:
	SCOPECOLON                     // scope:
	SCOPELINECOLON                 // scopeLine:
	SETTERCOLON                    // setter:
	SIZECOLON                      // size:
	SOURCECOLON                    // source:
	SPFLAGSCOLON                   // spFlags:
	SPLITDEBUGFILENAMECOLON        // splitDebugFilename:
	SPLITDEBUGINLININGCOLON        // splitDebugInlining:
	TAGCOLON                       // tag:
	TEMPLATEPARAMSCOLON            // templateParams:
	THISADJUSTMENTCOLON            // thisAdjustment:
	THROWNTYPESCOLON               // thrownTypes:
	TYPECOLON                      // type:
	TYPESCOLON                     // types:
	UNITCOLON                      // unit:
	VALUECOLON                     // value:
	VARCOLON                       // var:
	VIRTUALINDEXCOLON              // virtualIndex:
	VIRTUALITYCOLON                // virtuality:
	VTABLEHOLDERCOLON              // vtableHolder:
	COMMA
	EXCL
	DOTDOTDOT // ...
	LPAREN
	RPAREN
	LBRACK
	RBRACK
	LBRACE
	RBRACE
	MULT
	LT
	ASSIGN
	GT
	OR1

	NumTokens
)

var tokenStr = [...]string{
	"EOI",
	"INVALID_TOKEN",
	"COMMENT",
	"WHITESPACE",
	"GLOBAL_IDENT_TOK",
	"LOCAL_IDENT_TOK",
	"LABEL_IDENT_TOK",
	"ATTR_GROUP_ID_TOK",
	"COMDAT_NAME_TOK",
	"METADATA_NAME_TOK",
	"METADATA_ID_TOK",
	"DWARF_TAG_TOK",
	"DWARF_ATT_ENCODING_TOK",
	"DI_FLAG_TOK",
	"DISP_FLAG_TOK",
	"DWARF_LANG_TOK",
	"DWARF_CC_TOK",
	"CHECKSUM_KIND_TOK",
	"DWARF_VIRTUALITY_TOK",
	"DWARF_MACINFO_TOK",
	"DWARF_OP_TOK",
	"EMISSION_KIND_TOK",
	"NAME_TABLE_KIND_TOK",
	"INT_LIT_TOK",
	"FLOAT_LIT_TOK",
	"STRING_LIT_TOK",
	"INT_TYPE_TOK",
	"aarch64_vector_pcs",
	"acq_rel",
	"acquire",
	"add",
	"addrspace",
	"addrspacecast",
	"afn",
	"alias",
	"align",
	"alignstack",
	"alloca",
	"allocsize",
	"alwaysinline",
	"amdgpu_cs",
	"amdgpu_es",
	"amdgpu_gs",
	"amdgpu_hs",
	"amdgpu_kernel",
	"amdgpu_ls",
	"amdgpu_ps",
	"amdgpu_vs",
	"and",
	"any",
	"anyregcc",
	"appending",
	"arcp",
	"argmemonly",
	"arm_aapcs_vfpcc",
	"arm_aapcscc",
	"arm_apcscc",
	"ashr",
	"asm",
	"atomic",
	"atomicrmw",
	"attributes",
	"available_externally",
	"avr_intrcc",
	"avr_signalcc",
	"bitcast",
	"blockaddress",
	"br",
	"builtin",
	"byval",
	"c",
	"call",
	"callbr",
	"caller",
	"catch",
	"catchpad",
	"catchret",
	"catchswitch",
	"cc",
	"ccc",
	"cleanup",
	"cleanuppad",
	"cleanupret",
	"cmpxchg",
	"cold",
	"coldcc",
	"comdat",
	"common",
	"constant",
	"contract",
	"convergent",
	"cxx_fast_tlscc",
	"datalayout",
	"declare",
	"default",
	"define",
	"dereferenceable",
	"dereferenceable_or_null",
	"distinct",
	"dllexport",
	"dllimport",
	"double",
	"dso_local",
	"dso_preemptable",
	"eq",
	"exact",
	"exactmatch",
	"extern_weak",
	"external",
	"externally_initialized",
	"extractelement",
	"extractvalue",
	"fadd",
	"false",
	"fast",
	"fastcc",
	"fcmp",
	"fdiv",
	"fence",
	"filter",
	"float",
	"fmul",
	"fneg",
	"fp128",
	"fpext",
	"fptosi",
	"fptoui",
	"fptrunc",
	"frem",
	"from",
	"fsub",
	"gc",
	"getelementptr",
	"ghccc",
	"global",
	"half",
	"hhvm_ccc",
	"hhvmcc",
	"hidden",
	"icmp",
	"ifunc",
	"immarg",
	"inaccessiblemem_or_argmemonly",
	"inaccessiblememonly",
	"inalloca",
	"inbounds",
	"indirectbr",
	"initialexec",
	"inlinehint",
	"inrange",
	"inreg",
	"insertelement",
	"insertvalue",
	"intel_ocl_bicc",
	"inteldialect",
	"internal",
	"inttoptr",
	"invoke",
	"jumptable",
	"label",
	"landingpad",
	"largest",
	"linkonce",
	"linkonce_odr",
	"load",
	"local_unnamed_addr",
	"localdynamic",
	"localexec",
	"lshr",
	"max",
	"metadata",
	"min",
	"minsize",
	"module",
	"monotonic",
	"msp430_intrcc",
	"mul",
	"musttail",
	"naked",
	"nand",
	"ne",
	"nest",
	"ninf",
	"nnan",
	"noalias",
	"nobuiltin",
	"nocapture",
	"nocf_check",
	"noduplicate",
	"noduplicates",
	"nofree",
	"noimplicitfloat",
	"noinline",
	"none",
	"nonlazybind",
	"nonnull",
	"norecurse",
	"noredzone",
	"noreturn",
	"nosync",
	"notail",
	"nounwind",
	"nsw",
	"nsz",
	"null",
	"nuw",
	"oeq",
	"oge",
	"ogt",
	"ole",
	"olt",
	"one",
	"opaque",
	"optforfuzzing",
	"optnone",
	"optsize",
	"or",
	"ord",
	"partition",
	"personality",
	"phi",
	"ppc_fp128",
	"prefix",
	"preserve_allcc",
	"preserve_mostcc",
	"private",
	"prologue",
	"protected",
	"ptrtoint",
	"ptx_device",
	"ptx_kernel",
	"readnone",
	"readonly",
	"reassoc",
	"release",
	"resume",
	"ret",
	"returned",
	"returns_twice",
	"safestack",
	"samesize",
	"sanitize_address",
	"sanitize_hwaddress",
	"sanitize_memory",
	"sanitize_memtag",
	"sanitize_thread",
	"sdiv",
	"section",
	"select",
	"seq_cst",
	"sext",
	"sge",
	"sgt",
	"shadowcallstack",
	"shl",
	"shufflevector",
	"sideeffect",
	"signext",
	"singlethread",
	"sitofp",
	"sle",
	"slt",
	"source_filename",
	"speculatable",
	"speculative_load_hardening",
	"spir_func",
	"spir_kernel",
	"srem",
	"sret",
	"ssp",
	"sspreq",
	"sspstrong",
	"store",
	"strictfp",
	"sub",
	"swiftcc",
	"swifterror",
	"swiftself",
	"switch",
	"syncscope",
	"tail",
	"target",
	"thread_local",
	"to",
	"token",
	"triple",
	"true",
	"trunc",
	"type",
	"udiv",
	"ueq",
	"uge",
	"ugt",
	"uitofp",
	"ule",
	"ult",
	"umax",
	"umin",
	"undef",
	"une",
	"unnamed_addr",
	"uno",
	"unordered",
	"unreachable",
	"unwind",
	"urem",
	"uselistorder",
	"uselistorder_bb",
	"uwtable",
	"va_arg",
	"void",
	"volatile",
	"vscale",
	"weak",
	"weak_odr",
	"webkit_jscc",
	"willreturn",
	"win64cc",
	"within",
	"writeonly",
	"x",
	"x86_64_sysvcc",
	"x86_fastcallcc",
	"x86_fp80",
	"x86_intrcc",
	"x86_mmx",
	"x86_regcallcc",
	"x86_stdcallcc",
	"x86_thiscallcc",
	"x86_vectorcallcc",
	"xchg",
	"xor",
	"zeroext",
	"zeroinitializer",
	"zext",
	"!DIBasicType",
	"!DICommonBlock",
	"!DICompileUnit",
	"!DICompositeType",
	"!DIDerivedType",
	"!DIEnumerator",
	"!DIExpression",
	"!DIFile",
	"!DIGlobalVariable",
	"!DIGlobalVariableExpression",
	"!DIImportedEntity",
	"!DILabel",
	"!DILexicalBlock",
	"!DILexicalBlockFile",
	"!DILocalVariable",
	"!DILocation",
	"!DIMacro",
	"!DIMacroFile",
	"!DIModule",
	"!DINamespace",
	"!DIObjCProperty",
	"!DISubprogram",
	"!DISubrange",
	"!DISubroutineType",
	"!DITemplateTypeParameter",
	"!DITemplateValueParameter",
	"!GenericDINode",
	"align:",
	"arg:",
	"attributes:",
	"baseType:",
	"cc:",
	"checksum:",
	"checksumkind:",
	"column:",
	"configMacros:",
	"containingType:",
	"count:",
	"debugBaseAddress:",
	"debugInfoForProfiling:",
	"declaration:",
	"directory:",
	"discriminator:",
	"dwarfAddressSpace:",
	"dwoId:",
	"elements:",
	"emissionKind:",
	"encoding:",
	"entity:",
	"enums:",
	"exportSymbols:",
	"expr:",
	"extraData:",
	"file:",
	"filename:",
	"flags:",
	"getter:",
	"globals:",
	"header:",
	"identifier:",
	"imports:",
	"includePath:",
	"inlinedAt:",
	"isDefinition:",
	"isImplicitCode:",
	"isLocal:",
	"isOptimized:",
	"isUnsigned:",
	"isysroot:",
	"language:",
	"line:",
	"linkageName:",
	"lowerBound:",
	"macros:",
	"name:",
	"nameTableKind:",
	"nodes:",
	"offset:",
	"operands:",
	"producer:",
	"retainedNodes:",
	"retainedTypes:",
	"runtimeLang:",
	"runtimeVersion:",
	"scope:",
	"scopeLine:",
	"setter:",
	"size:",
	"source:",
	"spFlags:",
	"splitDebugFilename:",
	"splitDebugInlining:",
	"tag:",
	"templateParams:",
	"thisAdjustment:",
	"thrownTypes:",
	"type:",
	"types:",
	"unit:",
	"value:",
	"var:",
	"virtualIndex:",
	"virtuality:",
	"vtableHolder:",
	"COMMA",
	"EXCL",
	"...",
	"LPAREN",
	"RPAREN",
	"LBRACK",
	"RBRACK",
	"LBRACE",
	"RBRACE",
	"MULT",
	"LT",
	"ASSIGN",
	"GT",
	"OR1",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}

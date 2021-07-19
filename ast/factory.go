// generated by Textmapper; DO NOT EDIT

package ast

import (
	"fmt"

	"github.com/llir/ll"
)

func ToLlvmNode(n *Node) LlvmNode {
	switch n.Type() {
	case ll.APINotesField:
		return &APINotesField{n}
	case ll.AShrExpr:
		return &AShrExpr{n}
	case ll.AShrInst:
		return &AShrInst{n}
	case ll.AddExpr:
		return &AddExpr{n}
	case ll.AddInst:
		return &AddInst{n}
	case ll.AddrSpace:
		return &AddrSpace{n}
	case ll.AddrSpaceCastExpr:
		return &AddrSpaceCastExpr{n}
	case ll.AddrSpaceCastInst:
		return &AddrSpaceCastInst{n}
	case ll.Align:
		return &Align{n}
	case ll.AlignField:
		return &AlignField{n}
	case ll.AlignPair:
		return &AlignPair{n}
	case ll.AlignStack:
		return &AlignStack{n}
	case ll.AlignStackPair:
		return &AlignStackPair{n}
	case ll.AlignStackTok:
		return &AlignStackTok{n}
	case ll.AllocSize:
		return &AllocSize{n}
	case ll.AllocaInst:
		return &AllocaInst{n}
	case ll.AllocatedField:
		return &AllocatedField{n}
	case ll.AndExpr:
		return &AndExpr{n}
	case ll.AndInst:
		return &AndInst{n}
	case ll.Arg:
		return &Arg{n}
	case ll.ArgField:
		return &ArgField{n}
	case ll.Args:
		return &Args{n}
	case ll.ArrayConst:
		return &ArrayConst{n}
	case ll.ArrayType:
		return &ArrayType{n}
	case ll.AssociatedField:
		return &AssociatedField{n}
	case ll.Atomic:
		return &Atomic{n}
	case ll.AtomicOp:
		return &AtomicOp{n}
	case ll.AtomicOrdering:
		return &AtomicOrdering{n}
	case ll.AtomicRMWInst:
		return &AtomicRMWInst{n}
	case ll.AttrGroupDef:
		return &AttrGroupDef{n}
	case ll.AttrGroupID:
		return &AttrGroupID{n}
	case ll.AttrPair:
		return &AttrPair{n}
	case ll.AttrString:
		return &AttrString{n}
	case ll.AttributesField:
		return &AttributesField{n}
	case ll.BaseTypeField:
		return &BaseTypeField{n}
	case ll.BasicBlock:
		return &BasicBlock{n}
	case ll.BitCastExpr:
		return &BitCastExpr{n}
	case ll.BitCastInst:
		return &BitCastInst{n}
	case ll.BlockAddressConst:
		return &BlockAddressConst{n}
	case ll.BoolConst:
		return &BoolConst{n}
	case ll.BoolLit:
		return &BoolLit{n}
	case ll.BrTerm:
		return &BrTerm{n}
	case ll.ByRefAttr:
		return &ByRefAttr{n}
	case ll.Byval:
		return &Byval{n}
	case ll.CCField:
		return &CCField{n}
	case ll.CallBrTerm:
		return &CallBrTerm{n}
	case ll.CallInst:
		return &CallInst{n}
	case ll.CallingConvEnum:
		return &CallingConvEnum{n}
	case ll.CallingConvInt:
		return &CallingConvInt{n}
	case ll.Case:
		return &Case{n}
	case ll.CatchPadInst:
		return &CatchPadInst{n}
	case ll.CatchRetTerm:
		return &CatchRetTerm{n}
	case ll.CatchSwitchTerm:
		return &CatchSwitchTerm{n}
	case ll.CharArrayConst:
		return &CharArrayConst{n}
	case ll.ChecksumField:
		return &ChecksumField{n}
	case ll.ChecksumKind:
		return &ChecksumKind{n}
	case ll.ChecksumkindField:
		return &ChecksumkindField{n}
	case ll.Clause:
		return &Clause{n}
	case ll.ClauseType:
		return &ClauseType{n}
	case ll.Cleanup:
		return &Cleanup{n}
	case ll.CleanupPadInst:
		return &CleanupPadInst{n}
	case ll.CleanupRetTerm:
		return &CleanupRetTerm{n}
	case ll.CmpXchgInst:
		return &CmpXchgInst{n}
	case ll.ColumnField:
		return &ColumnField{n}
	case ll.Comdat:
		return &Comdat{n}
	case ll.ComdatDef:
		return &ComdatDef{n}
	case ll.ComdatName:
		return &ComdatName{n}
	case ll.CondBrTerm:
		return &CondBrTerm{n}
	case ll.ConfigMacrosField:
		return &ConfigMacrosField{n}
	case ll.ContainingTypeField:
		return &ContainingTypeField{n}
	case ll.CountField:
		return &CountField{n}
	case ll.DIBasicType:
		return &DIBasicType{n}
	case ll.DICommonBlock:
		return &DICommonBlock{n}
	case ll.DICompileUnit:
		return &DICompileUnit{n}
	case ll.DICompositeType:
		return &DICompositeType{n}
	case ll.DIDerivedType:
		return &DIDerivedType{n}
	case ll.DIEnumerator:
		return &DIEnumerator{n}
	case ll.DIExpression:
		return &DIExpression{n}
	case ll.DIFile:
		return &DIFile{n}
	case ll.DIFlagEnum:
		return &DIFlagEnum{n}
	case ll.DIFlagInt:
		return &DIFlagInt{n}
	case ll.DIFlags:
		return &DIFlags{n}
	case ll.DIGlobalVariable:
		return &DIGlobalVariable{n}
	case ll.DIGlobalVariableExpression:
		return &DIGlobalVariableExpression{n}
	case ll.DIImportedEntity:
		return &DIImportedEntity{n}
	case ll.DILabel:
		return &DILabel{n}
	case ll.DILexicalBlock:
		return &DILexicalBlock{n}
	case ll.DILexicalBlockFile:
		return &DILexicalBlockFile{n}
	case ll.DILocalVariable:
		return &DILocalVariable{n}
	case ll.DILocation:
		return &DILocation{n}
	case ll.DIMacro:
		return &DIMacro{n}
	case ll.DIMacroFile:
		return &DIMacroFile{n}
	case ll.DIModule:
		return &DIModule{n}
	case ll.DINamespace:
		return &DINamespace{n}
	case ll.DIObjCProperty:
		return &DIObjCProperty{n}
	case ll.DISPFlagEnum:
		return &DISPFlagEnum{n}
	case ll.DISPFlagInt:
		return &DISPFlagInt{n}
	case ll.DISPFlags:
		return &DISPFlags{n}
	case ll.DISubprogram:
		return &DISubprogram{n}
	case ll.DISubrange:
		return &DISubrange{n}
	case ll.DISubroutineType:
		return &DISubroutineType{n}
	case ll.DITemplateTypeParameter:
		return &DITemplateTypeParameter{n}
	case ll.DITemplateValueParameter:
		return &DITemplateValueParameter{n}
	case ll.DLLStorageClass:
		return &DLLStorageClass{n}
	case ll.DataLocationField:
		return &DataLocationField{n}
	case ll.DebugInfoForProfilingField:
		return &DebugInfoForProfilingField{n}
	case ll.DeclarationField:
		return &DeclarationField{n}
	case ll.DefaultedField:
		return &DefaultedField{n}
	case ll.Dereferenceable:
		return &Dereferenceable{n}
	case ll.DereferenceableOrNull:
		return &DereferenceableOrNull{n}
	case ll.DirectoryField:
		return &DirectoryField{n}
	case ll.DiscriminatorField:
		return &DiscriminatorField{n}
	case ll.DiscriminatorIntField:
		return &DiscriminatorIntField{n}
	case ll.Distinct:
		return &Distinct{n}
	case ll.DwarfAddressSpaceField:
		return &DwarfAddressSpaceField{n}
	case ll.DwarfAttEncodingEnum:
		return &DwarfAttEncodingEnum{n}
	case ll.DwarfAttEncodingInt:
		return &DwarfAttEncodingInt{n}
	case ll.DwarfCCEnum:
		return &DwarfCCEnum{n}
	case ll.DwarfCCInt:
		return &DwarfCCInt{n}
	case ll.DwarfLangEnum:
		return &DwarfLangEnum{n}
	case ll.DwarfLangInt:
		return &DwarfLangInt{n}
	case ll.DwarfMacinfoEnum:
		return &DwarfMacinfoEnum{n}
	case ll.DwarfMacinfoInt:
		return &DwarfMacinfoInt{n}
	case ll.DwarfOp:
		return &DwarfOp{n}
	case ll.DwarfTagEnum:
		return &DwarfTagEnum{n}
	case ll.DwarfTagInt:
		return &DwarfTagInt{n}
	case ll.DwarfVirtualityEnum:
		return &DwarfVirtualityEnum{n}
	case ll.DwarfVirtualityInt:
		return &DwarfVirtualityInt{n}
	case ll.DwoIdField:
		return &DwoIdField{n}
	case ll.ElementsField:
		return &ElementsField{n}
	case ll.Ellipsis:
		return &Ellipsis{n}
	case ll.EmissionKindEnum:
		return &EmissionKindEnum{n}
	case ll.EmissionKindField:
		return &EmissionKindField{n}
	case ll.EmissionKindInt:
		return &EmissionKindInt{n}
	case ll.EncodingField:
		return &EncodingField{n}
	case ll.EntityField:
		return &EntityField{n}
	case ll.EnumsField:
		return &EnumsField{n}
	case ll.Exact:
		return &Exact{n}
	case ll.ExceptionArg:
		return &ExceptionArg{n}
	case ll.ExportSymbolsField:
		return &ExportSymbolsField{n}
	case ll.ExprField:
		return &ExprField{n}
	case ll.ExternLinkage:
		return &ExternLinkage{n}
	case ll.ExternallyInitialized:
		return &ExternallyInitialized{n}
	case ll.ExtraDataField:
		return &ExtraDataField{n}
	case ll.ExtractElementExpr:
		return &ExtractElementExpr{n}
	case ll.ExtractElementInst:
		return &ExtractElementInst{n}
	case ll.ExtractValueExpr:
		return &ExtractValueExpr{n}
	case ll.ExtractValueInst:
		return &ExtractValueInst{n}
	case ll.FAddExpr:
		return &FAddExpr{n}
	case ll.FAddInst:
		return &FAddInst{n}
	case ll.FCmpExpr:
		return &FCmpExpr{n}
	case ll.FCmpInst:
		return &FCmpInst{n}
	case ll.FDivExpr:
		return &FDivExpr{n}
	case ll.FDivInst:
		return &FDivInst{n}
	case ll.FMulExpr:
		return &FMulExpr{n}
	case ll.FMulInst:
		return &FMulInst{n}
	case ll.FNegExpr:
		return &FNegExpr{n}
	case ll.FNegInst:
		return &FNegInst{n}
	case ll.FPExtExpr:
		return &FPExtExpr{n}
	case ll.FPExtInst:
		return &FPExtInst{n}
	case ll.FPToSIExpr:
		return &FPToSIExpr{n}
	case ll.FPToSIInst:
		return &FPToSIInst{n}
	case ll.FPToUIExpr:
		return &FPToUIExpr{n}
	case ll.FPToUIInst:
		return &FPToUIInst{n}
	case ll.FPTruncExpr:
		return &FPTruncExpr{n}
	case ll.FPTruncInst:
		return &FPTruncInst{n}
	case ll.FPred:
		return &FPred{n}
	case ll.FRemExpr:
		return &FRemExpr{n}
	case ll.FRemInst:
		return &FRemInst{n}
	case ll.FSubExpr:
		return &FSubExpr{n}
	case ll.FSubInst:
		return &FSubInst{n}
	case ll.FastMathFlag:
		return &FastMathFlag{n}
	case ll.FenceInst:
		return &FenceInst{n}
	case ll.FileField:
		return &FileField{n}
	case ll.FilenameField:
		return &FilenameField{n}
	case ll.FlagsField:
		return &FlagsField{n}
	case ll.FlagsStringField:
		return &FlagsStringField{n}
	case ll.FloatConst:
		return &FloatConst{n}
	case ll.FloatKind:
		return &FloatKind{n}
	case ll.FloatLit:
		return &FloatLit{n}
	case ll.FloatType:
		return &FloatType{n}
	case ll.FreezeInst:
		return &FreezeInst{n}
	case ll.FuncAttr:
		return &FuncAttr{n}
	case ll.FuncBody:
		return &FuncBody{n}
	case ll.FuncDecl:
		return &FuncDecl{n}
	case ll.FuncDef:
		return &FuncDef{n}
	case ll.FuncHeader:
		return &FuncHeader{n}
	case ll.FuncType:
		return &FuncType{n}
	case ll.GCNode:
		return &GCNode{n}
	case ll.GEPIndex:
		return &GEPIndex{n}
	case ll.GenericDINode:
		return &GenericDINode{n}
	case ll.GetElementPtrExpr:
		return &GetElementPtrExpr{n}
	case ll.GetElementPtrInst:
		return &GetElementPtrInst{n}
	case ll.GetterField:
		return &GetterField{n}
	case ll.GlobalDecl:
		return &GlobalDecl{n}
	case ll.GlobalIdent:
		return &GlobalIdent{n}
	case ll.GlobalsField:
		return &GlobalsField{n}
	case ll.Handlers:
		return &Handlers{n}
	case ll.HeaderField:
		return &HeaderField{n}
	case ll.ICmpExpr:
		return &ICmpExpr{n}
	case ll.ICmpInst:
		return &ICmpInst{n}
	case ll.IPred:
		return &IPred{n}
	case ll.IdentifierField:
		return &IdentifierField{n}
	case ll.Immutable:
		return &Immutable{n}
	case ll.ImportsField:
		return &ImportsField{n}
	case ll.InAlloca:
		return &InAlloca{n}
	case ll.InBounds:
		return &InBounds{n}
	case ll.InRange:
		return &InRange{n}
	case ll.Inc:
		return &Inc{n}
	case ll.IncludePathField:
		return &IncludePathField{n}
	case ll.IndirectBrTerm:
		return &IndirectBrTerm{n}
	case ll.IndirectSymbolDef:
		return &IndirectSymbolDef{n}
	case ll.IndirectSymbolKind:
		return &IndirectSymbolKind{n}
	case ll.InlineAsm:
		return &InlineAsm{n}
	case ll.InlinedAtField:
		return &InlinedAtField{n}
	case ll.InsertElementExpr:
		return &InsertElementExpr{n}
	case ll.InsertElementInst:
		return &InsertElementInst{n}
	case ll.InsertValueExpr:
		return &InsertValueExpr{n}
	case ll.InsertValueInst:
		return &InsertValueInst{n}
	case ll.IntConst:
		return &IntConst{n}
	case ll.IntLit:
		return &IntLit{n}
	case ll.IntToPtrExpr:
		return &IntToPtrExpr{n}
	case ll.IntToPtrInst:
		return &IntToPtrInst{n}
	case ll.IntType:
		return &IntType{n}
	case ll.IntelDialect:
		return &IntelDialect{n}
	case ll.InvokeTerm:
		return &InvokeTerm{n}
	case ll.IsDeclField:
		return &IsDeclField{n}
	case ll.IsDefinitionField:
		return &IsDefinitionField{n}
	case ll.IsImplicitCodeField:
		return &IsImplicitCodeField{n}
	case ll.IsLocalField:
		return &IsLocalField{n}
	case ll.IsOptimizedField:
		return &IsOptimizedField{n}
	case ll.IsUnsignedField:
		return &IsUnsignedField{n}
	case ll.LShrExpr:
		return &LShrExpr{n}
	case ll.LShrInst:
		return &LShrInst{n}
	case ll.Label:
		return &Label{n}
	case ll.LabelIdent:
		return &LabelIdent{n}
	case ll.LabelType:
		return &LabelType{n}
	case ll.LandingPadInst:
		return &LandingPadInst{n}
	case ll.LanguageField:
		return &LanguageField{n}
	case ll.LineField:
		return &LineField{n}
	case ll.Linkage:
		return &Linkage{n}
	case ll.LinkageNameField:
		return &LinkageNameField{n}
	case ll.LoadInst:
		return &LoadInst{n}
	case ll.LocalDefInst:
		return &LocalDefInst{n}
	case ll.LocalDefTerm:
		return &LocalDefTerm{n}
	case ll.LocalIdent:
		return &LocalIdent{n}
	case ll.LowerBoundField:
		return &LowerBoundField{n}
	case ll.MDString:
		return &MDString{n}
	case ll.MDTuple:
		return &MDTuple{n}
	case ll.MMXType:
		return &MMXType{n}
	case ll.MacrosField:
		return &MacrosField{n}
	case ll.MetadataAttachment:
		return &MetadataAttachment{n}
	case ll.MetadataDef:
		return &MetadataDef{n}
	case ll.MetadataID:
		return &MetadataID{n}
	case ll.MetadataName:
		return &MetadataName{n}
	case ll.MetadataType:
		return &MetadataType{n}
	case ll.Module:
		return &Module{n}
	case ll.ModuleAsm:
		return &ModuleAsm{n}
	case ll.MulExpr:
		return &MulExpr{n}
	case ll.MulInst:
		return &MulInst{n}
	case ll.NameField:
		return &NameField{n}
	case ll.NameTableKindEnum:
		return &NameTableKindEnum{n}
	case ll.NameTableKindField:
		return &NameTableKindField{n}
	case ll.NameTableKindInt:
		return &NameTableKindInt{n}
	case ll.NamedMetadataDef:
		return &NamedMetadataDef{n}
	case ll.NamedType:
		return &NamedType{n}
	case ll.NodesField:
		return &NodesField{n}
	case ll.NoneConst:
		return &NoneConst{n}
	case ll.NullConst:
		return &NullConst{n}
	case ll.NullLit:
		return &NullLit{n}
	case ll.OffsetField:
		return &OffsetField{n}
	case ll.OpaqueType:
		return &OpaqueType{n}
	case ll.OperandBundle:
		return &OperandBundle{n}
	case ll.OperandsField:
		return &OperandsField{n}
	case ll.OrExpr:
		return &OrExpr{n}
	case ll.OrInst:
		return &OrInst{n}
	case ll.OverflowFlag:
		return &OverflowFlag{n}
	case ll.PackedStructType:
		return &PackedStructType{n}
	case ll.Param:
		return &Param{n}
	case ll.ParamAttr:
		return &ParamAttr{n}
	case ll.Params:
		return &Params{n}
	case ll.Partition:
		return &Partition{n}
	case ll.Personality:
		return &Personality{n}
	case ll.PhiInst:
		return &PhiInst{n}
	case ll.PointerType:
		return &PointerType{n}
	case ll.PoisonConst:
		return &PoisonConst{n}
	case ll.Preallocated:
		return &Preallocated{n}
	case ll.Preemption:
		return &Preemption{n}
	case ll.Prefix:
		return &Prefix{n}
	case ll.ProducerField:
		return &ProducerField{n}
	case ll.Prologue:
		return &Prologue{n}
	case ll.PtrToIntExpr:
		return &PtrToIntExpr{n}
	case ll.PtrToIntInst:
		return &PtrToIntInst{n}
	case ll.RangesBaseAddressField:
		return &RangesBaseAddressField{n}
	case ll.RankField:
		return &RankField{n}
	case ll.ResumeTerm:
		return &ResumeTerm{n}
	case ll.RetTerm:
		return &RetTerm{n}
	case ll.RetainedNodesField:
		return &RetainedNodesField{n}
	case ll.RetainedTypesField:
		return &RetainedTypesField{n}
	case ll.ReturnAttr:
		return &ReturnAttr{n}
	case ll.RuntimeLangField:
		return &RuntimeLangField{n}
	case ll.RuntimeVersionField:
		return &RuntimeVersionField{n}
	case ll.SDKField:
		return &SDKField{n}
	case ll.SDivExpr:
		return &SDivExpr{n}
	case ll.SDivInst:
		return &SDivInst{n}
	case ll.SExtExpr:
		return &SExtExpr{n}
	case ll.SExtInst:
		return &SExtInst{n}
	case ll.SIToFPExpr:
		return &SIToFPExpr{n}
	case ll.SIToFPInst:
		return &SIToFPInst{n}
	case ll.SPFlagsField:
		return &SPFlagsField{n}
	case ll.SRemExpr:
		return &SRemExpr{n}
	case ll.SRemInst:
		return &SRemInst{n}
	case ll.ScalableVectorType:
		return &ScalableVectorType{n}
	case ll.ScopeField:
		return &ScopeField{n}
	case ll.ScopeLineField:
		return &ScopeLineField{n}
	case ll.Section:
		return &Section{n}
	case ll.SelectExpr:
		return &SelectExpr{n}
	case ll.SelectInst:
		return &SelectInst{n}
	case ll.SelectionKind:
		return &SelectionKind{n}
	case ll.SetterField:
		return &SetterField{n}
	case ll.ShlExpr:
		return &ShlExpr{n}
	case ll.ShlInst:
		return &ShlInst{n}
	case ll.ShuffleVectorExpr:
		return &ShuffleVectorExpr{n}
	case ll.ShuffleVectorInst:
		return &ShuffleVectorInst{n}
	case ll.SideEffect:
		return &SideEffect{n}
	case ll.SizeField:
		return &SizeField{n}
	case ll.SourceField:
		return &SourceField{n}
	case ll.SourceFilename:
		return &SourceFilename{n}
	case ll.SplitDebugFilenameField:
		return &SplitDebugFilenameField{n}
	case ll.SplitDebugInliningField:
		return &SplitDebugInliningField{n}
	case ll.StoreInst:
		return &StoreInst{n}
	case ll.StrideField:
		return &StrideField{n}
	case ll.StringLit:
		return &StringLit{n}
	case ll.StructConst:
		return &StructConst{n}
	case ll.StructRetAttr:
		return &StructRetAttr{n}
	case ll.StructType:
		return &StructType{n}
	case ll.SubExpr:
		return &SubExpr{n}
	case ll.SubInst:
		return &SubInst{n}
	case ll.SwiftError:
		return &SwiftError{n}
	case ll.SwitchTerm:
		return &SwitchTerm{n}
	case ll.SyncScope:
		return &SyncScope{n}
	case ll.SysrootField:
		return &SysrootField{n}
	case ll.TLSModel:
		return &TLSModel{n}
	case ll.TagField:
		return &TagField{n}
	case ll.Tail:
		return &Tail{n}
	case ll.TargetDataLayout:
		return &TargetDataLayout{n}
	case ll.TargetTriple:
		return &TargetTriple{n}
	case ll.TemplateParamsField:
		return &TemplateParamsField{n}
	case ll.ThisAdjustmentField:
		return &ThisAdjustmentField{n}
	case ll.ThreadLocal:
		return &ThreadLocal{n}
	case ll.ThrownTypesField:
		return &ThrownTypesField{n}
	case ll.TokenType:
		return &TokenType{n}
	case ll.TruncExpr:
		return &TruncExpr{n}
	case ll.TruncInst:
		return &TruncInst{n}
	case ll.TypeConst:
		return &TypeConst{n}
	case ll.TypeDef:
		return &TypeDef{n}
	case ll.TypeField:
		return &TypeField{n}
	case ll.TypeMacinfoField:
		return &TypeMacinfoField{n}
	case ll.TypeValue:
		return &TypeValue{n}
	case ll.TypesField:
		return &TypesField{n}
	case ll.UDivExpr:
		return &UDivExpr{n}
	case ll.UDivInst:
		return &UDivInst{n}
	case ll.UIToFPExpr:
		return &UIToFPExpr{n}
	case ll.UIToFPInst:
		return &UIToFPInst{n}
	case ll.URemExpr:
		return &URemExpr{n}
	case ll.URemInst:
		return &URemInst{n}
	case ll.UintLit:
		return &UintLit{n}
	case ll.UndefConst:
		return &UndefConst{n}
	case ll.UnitField:
		return &UnitField{n}
	case ll.UnnamedAddr:
		return &UnnamedAddr{n}
	case ll.UnreachableTerm:
		return &UnreachableTerm{n}
	case ll.UnwindToCaller:
		return &UnwindToCaller{n}
	case ll.UpperBoundField:
		return &UpperBoundField{n}
	case ll.UseListOrder:
		return &UseListOrder{n}
	case ll.UseListOrderBB:
		return &UseListOrderBB{n}
	case ll.VAArgInst:
		return &VAArgInst{n}
	case ll.ValueField:
		return &ValueField{n}
	case ll.ValueIntField:
		return &ValueIntField{n}
	case ll.ValueStringField:
		return &ValueStringField{n}
	case ll.VarField:
		return &VarField{n}
	case ll.VectorConst:
		return &VectorConst{n}
	case ll.VectorType:
		return &VectorType{n}
	case ll.VirtualIndexField:
		return &VirtualIndexField{n}
	case ll.VirtualityField:
		return &VirtualityField{n}
	case ll.Visibility:
		return &Visibility{n}
	case ll.VoidType:
		return &VoidType{n}
	case ll.Volatile:
		return &Volatile{n}
	case ll.VtableHolderField:
		return &VtableHolderField{n}
	case ll.Weak:
		return &Weak{n}
	case ll.XorExpr:
		return &XorExpr{n}
	case ll.XorInst:
		return &XorInst{n}
	case ll.ZExtExpr:
		return &ZExtExpr{n}
	case ll.ZExtInst:
		return &ZExtInst{n}
	case ll.ZeroInitializerConst:
		return &ZeroInitializerConst{n}
	case ll.NoType:
		return nilInstance
	}
	panic(fmt.Errorf("ast: unknown node type %v", n.Type()))
	return nil
}

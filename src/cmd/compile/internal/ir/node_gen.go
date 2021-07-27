// Code generated by mknode.go. DO NOT EDIT.

package ir

import "fmt"

func (n *AddStringExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AddStringExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	return &c
}
func (n *AddStringExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *AddStringExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *AddrExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AddrExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *AddrExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *AddrExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *ArrayType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ArrayType) copy() Node {
	c := *n
	return &c
}
func (n *ArrayType) doChildren(do func(Node) bool) bool {
	if n.Len != nil && do(n.Len) {
		return true
	}
	if n.Elem != nil && do(n.Elem) {
		return true
	}
	return false
}
func (n *ArrayType) editChildren(edit func(Node) Node) {
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Elem != nil {
		n.Elem = edit(n.Elem).(Ntype)
	}
}

func (n *AssignListStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AssignListStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Lhs = copyNodes(c.Lhs)
	c.Rhs = copyNodes(c.Rhs)
	return &c
}
func (n *AssignListStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.Lhs, do) {
		return true
	}
	if doNodes(n.Rhs, do) {
		return true
	}
	return false
}
func (n *AssignListStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Lhs, edit)
	editNodes(n.Rhs, edit)
}

func (n *AssignOpStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AssignOpStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *AssignOpStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *AssignOpStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *AssignStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AssignStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *AssignStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *AssignStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *BasicLit) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BasicLit) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *BasicLit) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *BasicLit) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *BinaryExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BinaryExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *BinaryExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *BinaryExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *BlockStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BlockStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	return &c
}
func (n *BlockStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	return false
}
func (n *BlockStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
}

func (n *BranchStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BranchStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *BranchStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *BranchStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *CallExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CallExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Args = copyNodes(c.Args)
	c.KeepAlive = copyNames(c.KeepAlive)
	return &c
}
func (n *CallExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if doNodes(n.Args, do) {
		return true
	}
	if doNames(n.KeepAlive, do) {
		return true
	}
	return false
}
func (n *CallExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	editNodes(n.Args, edit)
	editNames(n.KeepAlive, edit)
}

func (n *CaseClause) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CaseClause) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *CaseClause) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Var != nil && do(n.Var) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	return false
}
func (n *CaseClause) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Var != nil {
		n.Var = edit(n.Var).(*Name)
	}
	editNodes(n.List, edit)
	editNodes(n.Body, edit)
}

func (n *ChanType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ChanType) copy() Node {
	c := *n
	return &c
}
func (n *ChanType) doChildren(do func(Node) bool) bool {
	if n.Elem != nil && do(n.Elem) {
		return true
	}
	return false
}
func (n *ChanType) editChildren(edit func(Node) Node) {
	if n.Elem != nil {
		n.Elem = edit(n.Elem).(Ntype)
	}
}

func (n *ClosureExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ClosureExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ClosureExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *ClosureExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *CommClause) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CommClause) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *CommClause) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Comm != nil && do(n.Comm) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	return false
}
func (n *CommClause) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Comm != nil {
		n.Comm = edit(n.Comm).(Node)
	}
	editNodes(n.Body, edit)
}

func (n *CompLitExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CompLitExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	return &c
}
func (n *CompLitExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Ntype != nil && do(n.Ntype) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *CompLitExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Ntype != nil {
		n.Ntype = edit(n.Ntype).(Ntype)
	}
	editNodes(n.List, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *ConstExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ConstExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ConstExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *ConstExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *ConvExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ConvExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ConvExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *ConvExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *Decl) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *Decl) copy() Node {
	c := *n
	return &c
}
func (n *Decl) doChildren(do func(Node) bool) bool {
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *Decl) editChildren(edit func(Node) Node) {
	if n.X != nil {
		n.X = edit(n.X).(*Name)
	}
}

func (n *DynamicTypeAssertExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *DynamicTypeAssertExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *DynamicTypeAssertExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.T != nil && do(n.T) {
		return true
	}
	return false
}
func (n *DynamicTypeAssertExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.T != nil {
		n.T = edit(n.T).(Node)
	}
}

func (n *ForStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ForStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Late = copyNodes(c.Late)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *ForStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Cond != nil && do(n.Cond) {
		return true
	}
	if doNodes(n.Late, do) {
		return true
	}
	if n.Post != nil && do(n.Post) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	return false
}
func (n *ForStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Cond != nil {
		n.Cond = edit(n.Cond).(Node)
	}
	editNodes(n.Late, edit)
	if n.Post != nil {
		n.Post = edit(n.Post).(Node)
	}
	editNodes(n.Body, edit)
}

func (n *Func) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }

func (n *FuncType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *FuncType) copy() Node {
	c := *n
	c.Recv = copyField(c.Recv)
	c.Params = copyFields(c.Params)
	c.Results = copyFields(c.Results)
	return &c
}
func (n *FuncType) doChildren(do func(Node) bool) bool {
	if doField(n.Recv, do) {
		return true
	}
	if doFields(n.Params, do) {
		return true
	}
	if doFields(n.Results, do) {
		return true
	}
	return false
}
func (n *FuncType) editChildren(edit func(Node) Node) {
	editField(n.Recv, edit)
	editFields(n.Params, edit)
	editFields(n.Results, edit)
}

func (n *GoDeferStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *GoDeferStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *GoDeferStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Call != nil && do(n.Call) {
		return true
	}
	return false
}
func (n *GoDeferStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Call != nil {
		n.Call = edit(n.Call).(Node)
	}
}

func (n *Ident) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *Ident) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *Ident) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *Ident) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *IfStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *IfStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	c.Else = copyNodes(c.Else)
	return &c
}
func (n *IfStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Cond != nil && do(n.Cond) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	if doNodes(n.Else, do) {
		return true
	}
	return false
}
func (n *IfStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Cond != nil {
		n.Cond = edit(n.Cond).(Node)
	}
	editNodes(n.Body, edit)
	editNodes(n.Else, edit)
}

func (n *IndexExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *IndexExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *IndexExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Index != nil && do(n.Index) {
		return true
	}
	return false
}
func (n *IndexExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Index != nil {
		n.Index = edit(n.Index).(Node)
	}
}

func (n *InlineMarkStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InlineMarkStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *InlineMarkStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *InlineMarkStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *InlinedCallExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InlinedCallExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	c.ReturnVars = copyNodes(c.ReturnVars)
	return &c
}
func (n *InlinedCallExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	if doNodes(n.ReturnVars, do) {
		return true
	}
	return false
}
func (n *InlinedCallExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Body, edit)
	editNodes(n.ReturnVars, edit)
}

func (n *InstExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InstExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Targs = copyNodes(c.Targs)
	return &c
}
func (n *InstExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if doNodes(n.Targs, do) {
		return true
	}
	return false
}
func (n *InstExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	editNodes(n.Targs, edit)
}

func (n *InterfaceType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InterfaceType) copy() Node {
	c := *n
	c.Methods = copyFields(c.Methods)
	return &c
}
func (n *InterfaceType) doChildren(do func(Node) bool) bool {
	if doFields(n.Methods, do) {
		return true
	}
	return false
}
func (n *InterfaceType) editChildren(edit func(Node) Node) {
	editFields(n.Methods, edit)
}

func (n *KeyExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *KeyExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *KeyExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Key != nil && do(n.Key) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	return false
}
func (n *KeyExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Key != nil {
		n.Key = edit(n.Key).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}

func (n *LabelStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *LabelStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *LabelStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *LabelStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *LinksymOffsetExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *LinksymOffsetExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *LinksymOffsetExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *LinksymOffsetExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *LogicalExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *LogicalExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *LogicalExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *LogicalExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *MakeExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *MakeExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *MakeExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Len != nil && do(n.Len) {
		return true
	}
	if n.Cap != nil && do(n.Cap) {
		return true
	}
	return false
}
func (n *MakeExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Cap != nil {
		n.Cap = edit(n.Cap).(Node)
	}
}

func (n *MapType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *MapType) copy() Node {
	c := *n
	return &c
}
func (n *MapType) doChildren(do func(Node) bool) bool {
	if n.Key != nil && do(n.Key) {
		return true
	}
	if n.Elem != nil && do(n.Elem) {
		return true
	}
	return false
}
func (n *MapType) editChildren(edit func(Node) Node) {
	if n.Key != nil {
		n.Key = edit(n.Key).(Ntype)
	}
	if n.Elem != nil {
		n.Elem = edit(n.Elem).(Ntype)
	}
}

func (n *Name) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }

func (n *NilExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *NilExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *NilExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *NilExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *ParenExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ParenExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ParenExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *ParenExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *PkgName) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *PkgName) copy() Node {
	c := *n
	return &c
}
func (n *PkgName) doChildren(do func(Node) bool) bool {
	return false
}
func (n *PkgName) editChildren(edit func(Node) Node) {
}

func (n *RangeStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *RangeStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *RangeStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Key != nil && do(n.Key) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *RangeStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Key != nil {
		n.Key = edit(n.Key).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
	editNodes(n.Body, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *RawOrigExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *RawOrigExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *RawOrigExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *RawOrigExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *ResultExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ResultExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ResultExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *ResultExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *ReturnStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ReturnStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Results = copyNodes(c.Results)
	return &c
}
func (n *ReturnStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.Results, do) {
		return true
	}
	return false
}
func (n *ReturnStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Results, edit)
}

func (n *SelectStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SelectStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Cases = copyCommClauses(c.Cases)
	c.Compiled = copyNodes(c.Compiled)
	return &c
}
func (n *SelectStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doCommClauses(n.Cases, do) {
		return true
	}
	if doNodes(n.Compiled, do) {
		return true
	}
	return false
}
func (n *SelectStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editCommClauses(n.Cases, edit)
	editNodes(n.Compiled, edit)
}

func (n *SelectorExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SelectorExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SelectorExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *SelectorExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *SendStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SendStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SendStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Chan != nil && do(n.Chan) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	return false
}
func (n *SendStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Chan != nil {
		n.Chan = edit(n.Chan).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}

func (n *SliceExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SliceExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SliceExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Low != nil && do(n.Low) {
		return true
	}
	if n.High != nil && do(n.High) {
		return true
	}
	if n.Max != nil && do(n.Max) {
		return true
	}
	return false
}
func (n *SliceExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Low != nil {
		n.Low = edit(n.Low).(Node)
	}
	if n.High != nil {
		n.High = edit(n.High).(Node)
	}
	if n.Max != nil {
		n.Max = edit(n.Max).(Node)
	}
}

func (n *SliceHeaderExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SliceHeaderExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SliceHeaderExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Ptr != nil && do(n.Ptr) {
		return true
	}
	if n.Len != nil && do(n.Len) {
		return true
	}
	if n.Cap != nil && do(n.Cap) {
		return true
	}
	return false
}
func (n *SliceHeaderExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Ptr != nil {
		n.Ptr = edit(n.Ptr).(Node)
	}
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Cap != nil {
		n.Cap = edit(n.Cap).(Node)
	}
}

func (n *SliceType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SliceType) copy() Node {
	c := *n
	return &c
}
func (n *SliceType) doChildren(do func(Node) bool) bool {
	if n.Elem != nil && do(n.Elem) {
		return true
	}
	return false
}
func (n *SliceType) editChildren(edit func(Node) Node) {
	if n.Elem != nil {
		n.Elem = edit(n.Elem).(Ntype)
	}
}

func (n *StarExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *StarExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *StarExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *StarExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *StructKeyExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *StructKeyExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *StructKeyExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	return false
}
func (n *StructKeyExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}

func (n *StructType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *StructType) copy() Node {
	c := *n
	c.Fields = copyFields(c.Fields)
	return &c
}
func (n *StructType) doChildren(do func(Node) bool) bool {
	if doFields(n.Fields, do) {
		return true
	}
	return false
}
func (n *StructType) editChildren(edit func(Node) Node) {
	editFields(n.Fields, edit)
}

func (n *SwitchStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SwitchStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Cases = copyCaseClauses(c.Cases)
	c.Compiled = copyNodes(c.Compiled)
	return &c
}
func (n *SwitchStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Tag != nil && do(n.Tag) {
		return true
	}
	if doCaseClauses(n.Cases, do) {
		return true
	}
	if doNodes(n.Compiled, do) {
		return true
	}
	return false
}
func (n *SwitchStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Tag != nil {
		n.Tag = edit(n.Tag).(Node)
	}
	editCaseClauses(n.Cases, edit)
	editNodes(n.Compiled, edit)
}

func (n *TailCallStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *TailCallStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *TailCallStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Target != nil && do(n.Target) {
		return true
	}
	return false
}
func (n *TailCallStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Target != nil {
		n.Target = edit(n.Target).(*Name)
	}
}

func (n *TypeAssertExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *TypeAssertExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *TypeAssertExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Ntype != nil && do(n.Ntype) {
		return true
	}
	return false
}
func (n *TypeAssertExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Ntype != nil {
		n.Ntype = edit(n.Ntype).(Ntype)
	}
}

func (n *TypeSwitchGuard) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *TypeSwitchGuard) copy() Node {
	c := *n
	return &c
}
func (n *TypeSwitchGuard) doChildren(do func(Node) bool) bool {
	if n.Tag != nil && do(n.Tag) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *TypeSwitchGuard) editChildren(edit func(Node) Node) {
	if n.Tag != nil {
		n.Tag = edit(n.Tag).(*Ident)
	}
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *UnaryExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *UnaryExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *UnaryExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *UnaryExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *typeNode) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *typeNode) copy() Node {
	c := *n
	return &c
}
func (n *typeNode) doChildren(do func(Node) bool) bool {
	return false
}
func (n *typeNode) editChildren(edit func(Node) Node) {
}

func copyCaseClauses(list []*CaseClause) []*CaseClause {
	if list == nil {
		return nil
	}
	c := make([]*CaseClause, len(list))
	copy(c, list)
	return c
}
func doCaseClauses(list []*CaseClause, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editCaseClauses(list []*CaseClause, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(*CaseClause)
		}
	}
}

func copyCommClauses(list []*CommClause) []*CommClause {
	if list == nil {
		return nil
	}
	c := make([]*CommClause, len(list))
	copy(c, list)
	return c
}
func doCommClauses(list []*CommClause, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editCommClauses(list []*CommClause, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(*CommClause)
		}
	}
}

func copyNames(list []*Name) []*Name {
	if list == nil {
		return nil
	}
	c := make([]*Name, len(list))
	copy(c, list)
	return c
}
func doNames(list []*Name, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editNames(list []*Name, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(*Name)
		}
	}
}

func copyNodes(list []Node) []Node {
	if list == nil {
		return nil
	}
	c := make([]Node, len(list))
	copy(c, list)
	return c
}
func doNodes(list []Node, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editNodes(list []Node, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(Node)
		}
	}
}

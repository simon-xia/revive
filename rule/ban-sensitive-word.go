package rule

import (

"fmt"
"github.com/simon-xia/revive/lint"
"go/ast"
"go/token"
"regexp"
)

// BanSensitiveWordRule lints sensitive words in function, var etc
type BanSensitiveWordRule struct{}

// Apply applies the rule to given file.
func (r *BanSensitiveWordRule) Apply(file *lint.File, arguments lint.Arguments) []lint.Failure {

	var failures []lint.Failure

	banList := make(map[string]*regexp.Regexp)
	for _, arg := range []interface{}(arguments) {
		if 	banStr, ok := arg.(string); ok {
			if reg, er := regexp.Compile(banStr); er == nil {
				banList[banStr] = reg
			}
		}
	}

	fileAst := file.AST
	walker := lintBanSensitiveWord{
		file:    file,
		fileAst: fileAst,
		banList: banList,
		onFailure: func(failure lint.Failure) {
			failures = append(failures, failure)
		},
	}

	ast.Walk(&walker, fileAst)

	return failures
}
// Name returns the rule name.
func (r *BanSensitiveWordRule) Name() string {
	return "ban-sensitive-word"
}


type lintBanSensitiveWord struct {
	file                   *lint.File
	fileAst                *ast.File
	banList 			   map[string]*regexp.Regexp
	onFailure              func(lint.Failure)
}

func (w *lintBanSensitiveWord) Visit(n ast.Node) ast.Visitor {
	switch v := n.(type) {
	case *ast.AssignStmt:
		if v.Tok == token.ASSIGN {
			return w
		}
		for _, exp := range v.Lhs {
			if id, ok := exp.(*ast.Ident); ok {
				banCheck(id, w)
			}
		}
	case *ast.FuncDecl:
		banCheck(v.Name,w)
		banCheckList(v.Type.Params,w)
		banCheckList(v.Type.Results,w)
	case *ast.GenDecl:
		if v.Tok == token.IMPORT {
			return w
		}

		for _, spec := range v.Specs {
			switch s := spec.(type) {
			case *ast.TypeSpec:
				banCheck(s.Name,w)
			case *ast.ValueSpec:
				for _, id := range s.Names {
					banCheck(id,w)
				}
			}
		}
	case *ast.InterfaceType:
		for _, x := range v.Methods.List {
			ft, ok := x.Type.(*ast.FuncType)
			if !ok {
				continue
			}
			banCheckList(ft.Params,w)
			banCheckList(ft.Results,w)
		}
	case *ast.RangeStmt:
		if v.Tok == token.ASSIGN {
			return w
		}
		if id, ok := v.Key.(*ast.Ident); ok {
			banCheck(id,w)
		}
		if id, ok := v.Value.(*ast.Ident); ok {
			banCheck(id,w)
		}
	case *ast.StructType:
		for _, f := range v.Fields.List {
			for _, id := range f.Names {
				banCheck(id,  w)
			}
		}
	}
	return w
}

func banCheck(id *ast.Ident, w *lintBanSensitiveWord) {
	if id.Name == "_" {
		return
	}

	for s, reg := range w.banList {
		if reg.MatchString(id.Name) {
			w.onFailure(lint.Failure{
				Failure:    fmt.Sprintf("sensitive word %s found in name %s", s, id.Name),
				Confidence: 1,
				Node:       id,
				Category:   "naming",
			})
		}
	}
}

func banCheckList(fl *ast.FieldList, w *lintBanSensitiveWord) {
	if fl == nil {
		return
	}

	for _, f := range fl.List {
		for _, id := range f.Names {
			banCheck(id, w)
		}
	}
}

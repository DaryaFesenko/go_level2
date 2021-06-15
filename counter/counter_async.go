package counter

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type counter struct {
	count int
}

func CountAsyncFunc(fileName string, funcName string) (int, error) {
	fset := token.NewFileSet()

	result, err := parser.ParseFile(fset, fileName, nil, 0)

	if err != nil {
		return 0, err
	}

	count := new(counter)

	ast.Inspect(result, func(x ast.Node) bool {
		ast_func, ok := x.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if ast_func.Name.Name != funcName {
			return true
		}

		for _, node := range ast_func.Body.List {
			checkGo(node, count)
		}
		return false
	})

	return count.count, nil
}

func checkGo(node ast.Stmt, count *counter) {
	switch res := node.(type) {
	case *ast.GoStmt:
		count.count++
	case *ast.BlockStmt:
		for _, node := range res.List {
			checkGo(node, count)
		}
	case *ast.CaseClause:
		for _, node := range res.Body {
			checkGo(node, count)
		}
	case *ast.CommClause:
		for _, node := range res.Body {
			checkGo(node, count)
		}
	case *ast.ForStmt:
		for _, node := range res.Body.List {
			checkGo(node, count)
		}
	case *ast.IfStmt:
		for _, node := range res.Body.List {
			checkGo(node, count)
		}
	case *ast.RangeStmt:
		for _, node := range res.Body.List {
			checkGo(node, count)
		}
	case *ast.SelectStmt:
		for _, node := range res.Body.List {
			checkGo(node, count)
		}
	case *ast.SwitchStmt:
		for _, node := range res.Body.List {
			checkGo(node, count)
		}
	case *ast.TypeSwitchStmt:
		for _, node := range res.Body.List {
			checkGo(node, count)
		}
	default:
		return
	}
}

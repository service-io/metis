package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestPrintAst(t *testing.T) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "first/module/user/repository/account/autogen.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	err = ast.Print(fileSet, node)
	// err = printer.Fprint(os.Stdout, fileSet, node)
	if err != nil {
		fmt.Println(err)
	}
}

package test

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"os"
	"path/filepath"
	"testing"
)

func TestJennifer0(t *testing.T) {
	f := NewFile("account")
	f.PackageComment("ll")
	f.PackageComment("ll")
	f.PackageComment("ll")
	f.HeaderComment("This file is generated - do not edit.")
	f.ImportName("metis/test/first/entity", "entity")
	f.Line()
	//init := Func().Id("init").Params().Block(
	//	Qual("a.b/c", "Foo").Call().Comment("Local package - name is omitted."),
	//	Qual("d.e/f", "Bar").Call().Comment("Import is automatically added."),
	//	Qual("g.h/f", "Baz").Call().Comment("Colliding package name is renamed."),
	//)
	f.Comment("dfsa")
	f.Type().Id("Name").Struct(Id("A").Id("dto").Dot("Account").Tag(map[string]string{"json": "a"}))

	f.Type().Id("Name0").Interface(Id("SelectByID").Params(Id("id").Id("int64")).Params(Id("dto").Dot("Account")))
	f.Type().Id("Name1").Interface(
		Id("SelectByID").Params(Id("id").Id("int64")).Params(Qual("metis/test/first/model/dto", "Account")),
		f.Line(),
		Id("SelectByIDS").Params(Id("id").Id("int64")).Params(Qual("metis/test/first/entity", "Account")),
	)
	//Id("SelectByID").Params(Id("id").Id("int64")).Params(Id("dto").Dot("Account"))
	//f.Add(init)
	fmt.Printf("%#v\n", f)

	autogenFilePath := "second/module/user/repository/account/autogen1.go"
	if err := os.MkdirAll(filepath.Dir(autogenFilePath), 0766); err != nil {
		panic(err)
	}
	wr, err := os.OpenFile(autogenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	err = f.Render(wr)
}

func TestJennifer1(t *testing.T) {
	f := genRepositoryFile("account", nil)
	fmt.Printf("%#v\n", f)
	autogenFilePath := "second/module/user/repository/account/autogen.go"
	if err := os.MkdirAll(filepath.Dir(autogenFilePath), 0766); err != nil {
		panic(err)
	}
	wr, err := os.OpenFile(autogenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	err = f.Render(wr)
}

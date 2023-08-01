package test

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func genDeclDtoStruct(table string, columns []Column) jen.Code {
	camel := strcase.ToCamel(table)
	return jen.Type().Id(camel).Add(useEntity(camel))
}

func genDtoFile(table string, columns []Column) *jen.File {
	f := jen.NewFile("dto")
	f.ImportName("context", "context")
	f.ImportName("database/sql", "sql")
	f.ImportName("errors", "errors")
	f.ImportName("fmt", "fmt")
	f.ImportName("time", "time")
	f.ImportName("github.com/jinzhu/copier", "copier")
	f.ImportName("github.com/gin-gonic/gin", "gin")
	f.ImportName("go.uber.org/zap", "zap")
	f.ImportName("metis/config/constant", "constant")
	f.ImportName("metis/database", "database")
	f.ImportName("metis/test/second/model/dto", "dto")
	f.ImportName("metis/test/second/model/entity", "entity")
	f.ImportName("metis/util", "util")
	f.ImportName("metis/util/logger", "logger")
	f.ImportName("strings", "strings")

	f.HeaderComment("Code generated by tabuyos. DO NOT EDIT!")
	f.PackageComment("Package dto")
	f.PackageComment("@author tabuyos")
	f.PackageComment("@since " + time.Now().Format("2006/01/02"))
	f.PackageComment("@description " + table)

	f.Add(genDeclDtoStruct(table, columns))
	return f
}

func TestDto(t *testing.T) {
	strcase.ConfigureAcronym("ID", "id")
	strcase.ConfigureAcronym("id", "ID")

	table := "account"
	columns := getColumns(table)

	f := genDtoFile(table, columns)
	fmt.Printf("%#v\n", f)
	autogenFilePath := "second/model/dto/" + table + ".go"
	if err := os.MkdirAll(filepath.Dir(autogenFilePath), 0766); err != nil {
		panic(err)
	}
	wr, err := os.OpenFile(autogenFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	err = f.Render(wr)
}

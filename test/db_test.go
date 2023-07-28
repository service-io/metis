// Package test
// @author tabuyos
// @since 2023/7/28
// @description test
package test

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"metis/util"
	"os"
	"strings"
	"testing"
)

var TypeMappingMysqlToGo = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"bigint":             "int64",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int8",
	"smallint unsigned":  "int16",
	"mediumint unsigned": "int32",
	"bigint unsigned":    "int64",
	"bit":                "int8",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // string
	"datetime":           "time.Time", // string
	"timestamp":          "time.Time", // string
	"time":               "time.Time", // string
	"float":              "float32",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "[]byte",
	"varbinary":          "[]byte",
}

func TestMyAst(t *testing.T) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "entity/my_test.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	err = ast.Print(fileSet, node)
	// err = printer.Fprint(os.Stdout, fileSet, node)
	if err != nil {
		fmt.Println(err)
	}
}

func TestDbGen(t *testing.T) {
	columns := make([]Column, 0)

	columns = append(columns, Column{
		ColumnName:      "id",
		Type:            "bigint",
		Nullable:        "NO",
		TableName:       "survey",
		ColumnComment:   "主键",
		Tag:             "id",
		MaxLength:       0,
		NumberPrecision: 19,
		ColumnType:      "bigint",
		ColumnKey:       "PRI",
		Default:         "",
	})
	columns = append(columns, Column{
		ColumnName:      "title",
		Type:            "varchar",
		Nullable:        "NO",
		TableName:       "survey",
		ColumnComment:   "主题",
		Tag:             "title",
		MaxLength:       255,
		NumberPrecision: 0,
		ColumnType:      "varchar(1000)",
		ColumnKey:       "",
		Default:         "",
	})
	columns = append(columns, Column{
		ColumnName:      "start_at",
		Type:            "timestamp",
		Nullable:        "NO",
		TableName:       "survey",
		ColumnComment:   "开始时间",
		Tag:             "startAt",
		MaxLength:       0,
		NumberPrecision: 0,
		ColumnType:      "timestamp",
		ColumnKey:       "",
		Default:         "CURRENT_TIMESTAMP",
	})

	isTimeFn := func(col Column) bool {
		switch col.Type {
		case "date":
			return true
		case "datetime":
			return true
		case "timestamp":
			return true
		case "time":
			return true
		default:
			return false
		}
	}

	existTimeFn := func(cols []Column) bool {
		for _, col := range cols {
			if isTimeFn(col) {
				return true
			}
		}
		return false
	}

	var wr io.Writer

	f, err := os.OpenFile("entity/entity_test.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("openFile err: ", err.Error())
	}
	defer util.DeferClose(f)

	wr = f

	cm := "// Code generated by log-gen. DO NOT EDIT.\n"
	ff := &ast.File{
		Package: token.Pos(len(cm) + 1),
		Name: &ast.Ident{
			Name:    "entity",
			NamePos: token.Pos(len("package") + 2),
		}}

	ff.Comments = append(ff.Comments, &ast.CommentGroup{
		List: []*ast.Comment{
			{
				Text: "// " + "hellllllo",
			},
		},
	})

	if existTimeFn(columns) {
		ff.Decls = append(ff.Decls, &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf("\"%s\"", "time"),
					},
				},
			},
		})
	}

	ff.Decls = append(ff.Decls, Create(columns), makeFunc("Survey", "Doi", "id"))

	if _, err := wr.Write([]byte(cm)); err != nil {
	}
	fileSet := token.NewFileSet()
	// _ = ast.Print(fileSet, ff)
	if err := format.Node(wr, fileSet, ff); err != nil {
	} else {
	}

}

func Create(columns []Column) *ast.GenDecl {
	strcase.ConfigureAcronym("ID", "id")
	strcase.ConfigureAcronym("id", "ID")

	var structName string

	fields := make([]*ast.Field, 0, len(columns))
	for _, column := range columns {
		structName = column.TableName
		fields = append(fields, &ast.Field{
			Names: []*ast.Ident{
				{
					Name: strcase.ToCamel(column.ColumnName),
				},
			},
			Type: &ast.Ident{
				Name: TypeMappingMysqlToGo[column.Type],
			},
			Tag: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("`json:\"%s\"`", strcase.ToLowerCamel(strings.ToUpper(column.ColumnName))),
			},

			Comment: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "// " + column.ColumnComment,
					},
				},
			},
		})
	}

	ret := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: &ast.Ident{
					Name: strcase.ToCamel(structName),
				},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: fields,
					},
				},
			},
		},
	}

	return ret
}

func makeFunc(structName, funcName, ret string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: strings.ToLower("rec"),
						},
					},
					Type: &ast.StarExpr{
						X: &ast.Ident{
							Name: structName,
						},
					},
				},
			},
		},
		Name: &ast.Ident{
			Name: funcName,
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "a",
							},
						},
						Type: &ast.StarExpr{
							X: &ast.Ident{
								Name: "string",
							},
						},
					},
					{
						Names: []*ast.Ident{
							{
								Name: "b",
							},
						},
						Type: &ast.StarExpr{
							X: &ast.Ident{
								Name: "int",
							},
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							{
								Name: "s",
							},
							{
								Name: "x",
							},
						},
						Type: &ast.Ident{
							Name: "string",
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.Ident{
							Name: fmt.Sprintf("\"%s\"", ret),
						},
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: fmt.Sprintf("\"%s\"", ret),
						},
					},
				},
			},
		},
	}
}

type Column struct {
	ColumnName      string
	Type            string
	Nullable        string
	TableName       string
	ColumnComment   string
	Tag             string
	MaxLength       int64
	NumberPrecision int64
	ColumnType      string
	ColumnKey       string
	Default         interface{}
}

package main

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"golang.org/x/tools/go/packages"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("GOFILE"))
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  cwd = %s\n", cwd)
	fmt.Printf("  os.Args = %#v\n", os.Args)

	for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
		fmt.Println("  ", ev, "=", os.Getenv(ev))
	}

	fset := token.NewFileSet()
	fileName := os.Getenv("GOFILE")

	packages, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedFiles,
		Fset: fset,
	}, ".")
	if err != nil {
		panic(err)
	}
	for _, p := range packages {
		// fmt.Println(p)
		for _, file := range p.Syntax {
			// fmt.Println(file)
			for _, d := range file.Decls {
				funcDecl, ok := d.(*ast.FuncDecl)
				if !ok {
					continue
				}
				fmt.Println(funcDecl)
				fmt.Println(funcDecl.Body.Pos())
				fmt.Println(funcDecl.Body.Lbrace)
				funcDecl.Body.List = append(
					[]ast.Stmt{&ast.ExprStmt{
						X: &ast.CallExpr{
							Fun:  ast.NewIdent("trace.Trace"),
							Args: []ast.Expr{},
						},
					}},
					funcDecl.Body.List...,
				)
			}
			// Print the modified AST back to a Go source file
			f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
			if err != nil {
				fmt.Println("Error creating output file:", err)
				return
			}
			defer f.Close()
			printer.Fprint(f, fset, file)
		}
	}
}

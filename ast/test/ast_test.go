package test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestAst(t *testing.T) {
	// src is the input for which we want to print the AST.
	src := `package main
func main() {
	println("{Hello%s}, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	log.Fatal()
	// Print the AST.
	ast.Print(fset, f)

}

package main

import (
	"log"

	"bramp.net/antlr4/plsql"                   // The parser
	"github.com/antlr/antlr4/runtime/Go/antlr" // The antlr library
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

type plsqlListener struct {
	*plsql.BasePlSqlParserListener
}

func (p *plsqlListener) EnterAnonymous_block(ctx *plsql.Anonymous_blockContext) {
	log.Println("Enter Anonymous block", ctx)
	p.BasePlSqlParserListener.EnterAnonymous_block(ctx)
}
func (p *plsqlListener) ExitBlock(ctx *plsql.BlockContext) {
	log.Println("Exit Block", ctx)
	log.Println(ctx.ToStringTree(nil, nil))
	p.BasePlSqlParserListener.ExitBlock(ctx)
}

func (p *plsqlListener) ExitSql_statement(ctx *plsql.Sql_statementContext) {
	log.Println("Exit ExitSql_statement ", ctx)
	p.BasePlSqlParserListener.ExitSql_statement(ctx)
}

// Main shows how to use the Pl/SQL lexer and parser.
func Main() error {
	// Setup the input
	is := antlr.NewInputStream(`
DECLARE
  NUX NUMBER := 1;
BEGIN
  DBMS_OUTPUT.PUT_LINE('Hello');
EXCEPTION
  WHEN OTHERS THEN
    DBMS_OUTPUT.PUT_LINE('Something went wrong!');
END;
`)

	// Create the Lexer
	lexer := plsql.NewPlSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := plsql.NewPlSqlParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Finally walk the tree
	tree := p.Block()
	antlr.ParseTreeWalkerDefault.Walk(&plsqlListener{}, tree)
	log.Printf("Tree: %+v", tree)

	// Output:
	// Object: {"example":"json","with":["an","array"]}
	// Pair: "example":"json"
	// Pair: "with":["an","array"]
	// Array: ["an","array"]
	return nil
}

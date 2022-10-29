package parser

import "github.com/gotid/ddl-parser/gen"

// VisitRoot visits a parse tree produced by MySqlParser#root.
func (v *visitor) VisitRoot(ctx *gen.RootContext) interface{} {
	v.trace("VisitRoot")
	if ctx.SqlStatements() != nil {
		return ctx.SqlStatements().Accept(v)
	}

	return nil
}

// VisitSqlStatements visits a parse tree produced by MySqlParser#sqlStatements.
func (v *visitor) VisitSqlStatements(ctx *gen.SqlStatementsContext) interface{} {
	v.trace("VisitSqlStatements")
	var createTables []*CreateTable
	for _, e := range ctx.AllSqlStatement() {
		ret := e.Accept(v)
		if ret == nil {
			continue
		}

		if data, ok := ret.(*CreateTable); ok {
			createTables = append(createTables, data)
		}
	}

	return createTables
}

// VisitSqlStatement visits a parse tree produced by MySqlParser#sqlStatement.
func (v *visitor) VisitSqlStatement(ctx *gen.SqlStatementContext) interface{} {
	v.trace("VisitSqlStatement")
	if ctx.DdlStatement() != nil {
		return ctx.DdlStatement().Accept(v)
	}

	return nil
}

// VisitDdlStatement visits a parse tree produced by MySqlParser#ddlStatement.
func (v *visitor) VisitDdlStatement(ctx *gen.DdlStatementContext) interface{} {
	v.trace("VisitDdlStatement")
	if ctx.CreateTable() != nil {
		return v.visitCreateTable(ctx.CreateTable())
	}

	return nil
}

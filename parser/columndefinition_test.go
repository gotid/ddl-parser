package parser

import (
	"testing"

	"github.com/gotid/ddl-parser/gen"
	"github.com/stretchr/testify/assert"
)

func TestVisitor_VisitColumnDefinition(t *testing.T) {
	p := NewParser(WithDebugMode(true))
	accept := func(p *gen.MySqlParser, visitor *visitor) interface{} {
		definition := p.ColumnDefinition()
		ctx := definition.(*gen.ColumnDefinitionContext)
		return visitor.VisitColumnDefinition(ctx)
	}

	v, err := p.testMysqlSyntax("test.sql", accept, `bigint(20) NOT NULL DEFAULT 'test default' PRIMARY KEY COMMENT 'test comment'`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	columnDefinition := v.(*ColumnDefinition)

	assert.Equal(t, ColumnConstraint{
		NotNull:         true,
		HasDefaultValue: true,
		Primary:         true,
		Comment:         "test comment",
	}, *columnDefinition.ColumnConstraint)

	v, err = p.testMysqlSyntax("test.sql", accept, `bigint(20) NULL KEY`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	columnDefinition = v.(*ColumnDefinition)

	assert.Equal(t, ColumnConstraint{
		Key: true,
	}, *columnDefinition.ColumnConstraint)

	v, err = p.testMysqlSyntax("test.sql", accept, `bigint(20) NULL AUTO_INCREMENT UNIQUE KEY`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	columnDefinition = v.(*ColumnDefinition)

	assert.Equal(t, ColumnConstraint{
		AutoIncrement: true,
		Unique:        true,
	}, *columnDefinition.ColumnConstraint)

	v, err = p.testMysqlSyntax("test.sql", accept, `bigint(20) NULL DEFAULT NULL AUTO_INCREMENT UNIQUE KEY`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	columnDefinition = v.(*ColumnDefinition)

	assert.Equal(t, ColumnConstraint{
		AutoIncrement: true,
		Unique:        true,
	}, *columnDefinition.ColumnConstraint)

	v, err = p.testMysqlSyntax("test.sql", accept, `varchar(20) DEFAULT '' AUTO_INCREMENT UNIQUE KEY`)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	columnDefinition = v.(*ColumnDefinition)

	assert.Equal(t, ColumnConstraint{
		HasDefaultValue: true,
		AutoIncrement:   true,
		Unique:          true,
	}, *columnDefinition.ColumnConstraint)
}

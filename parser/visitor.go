package parser

import (
	"fmt"

	"github.com/gotid/ddl-parser/console"
	"github.com/gotid/ddl-parser/gen"
)

type visitor struct {
	gen.BaseMySqlParserVisitor
	prefix string
	debug  bool
	logger console.Console
}

func (v *visitor) trace(msg ...interface{}) {
	if v.debug {
		v.logger.Debug("Visit Trace: " + fmt.Sprint(msg...))
	}
}

func (v *visitor) panicWithExpr(expr Token, msg string) {
	if len(v.prefix) == 0 {
		err := fmt.Errorf("%v:%v %s", expr.GetLine(), expr.GetColumn(), msg)
		if v.debug {
			v.logger.Error(err)
		}

		panic(err)
		return
	}

	err := fmt.Errorf("%v line %v:%v %s", v.prefix, expr.GetLine(), expr.GetColumn(), msg)
	if v.debug {
		v.logger.Error(err)
	}

	panic(err)
}

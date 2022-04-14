package math

import (
	"fmt"
	"strings"
)

type ExpressionPrinter struct {
	sb *strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{
		sb: &strings.Builder{},
	}
}

func (p *ExpressionPrinter) VisitDoubleExpression(ex *DoubleExpression) {
	p.sb.WriteString(fmt.Sprintf("%g", ex.Value))
}

func (p *ExpressionPrinter) VisitAdditionExpression(ex *AdditionExpression) {
	p.sb.WriteRune('(')
	ex.Left.Accept(p)
	p.sb.WriteRune('+')
	ex.Right.Accept(p)
	p.sb.WriteRune(')')
}

func (p *ExpressionPrinter) Print(ex Expression) {
	ex.Accept(p)

	fmt.Println(p.sb.String())
}

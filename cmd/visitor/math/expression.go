package math

type Expression interface {
	Accept(ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitDoubleExpression(*DoubleExpression)
	VisitAdditionExpression(*AdditionExpression)
}

type DoubleExpression struct {
	Value float64
}

func (e *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(e)
}

type AdditionExpression struct {
	Left, Right Expression
}

func (e *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(e)
}

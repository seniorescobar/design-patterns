package main

import (
	"github.com/seniorescobar/design-patterns/cmd/visitor/math"
)

func main() {
	ex := &math.AdditionExpression{
		&math.DoubleExpression{1},
		&math.AdditionExpression{
			&math.DoubleExpression{2},
			&math.DoubleExpression{3},
		},
	}

	ep := math.NewExpressionPrinter()
	ep.Print(ex)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations=[]Calculation{}

func calculateExpression(expression string) (string, error) {
	expr,err:= govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "",err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "",err
	}

	return fmt.Sprintf("%v", result), err
}

func getCalculations(c echo.Context) error{
	return c.JSON(http.StatusOK, calculations)
}

func postCalculations(c echo.Context) error{
	return c.JSON(http.StatusOK, calculations)
}

func main() {
	e:=echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)

	e.Start("localhost:8080")
	fmt.Println("Hello World")
}

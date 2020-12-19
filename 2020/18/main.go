package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type expression struct {
	expr1    *expression
	val1     *int
	operator *byte
	expr2    *expression
	val2     *int
}

func parseExpressionPart1(line string, start int) (expression, int) {
	expr := expression{}
	i := start
	for ; i < len(line) && line[i] != ')'; i++ {
		if line[i] == '(' {
			var subExpr expression
			subExpr, i = parseExpressionPart1(line, i+1)
			if expr.val1 == nil && expr.expr1 == nil {
				expr.expr1 = &subExpr
			} else {
				expr.expr2 = &subExpr
			}
		} else if line[i] == '+' || line[i] == '*' {
			op := line[i]
			if expr.operator == nil {
				expr.operator = &op
			} else {
				// complete expression => nest
				tmp := expr
				expr = expression{expr1: &tmp, operator: &op}
			}
		} else {
			val, err := strconv.Atoi(line[i : i+1])
			if err == nil {
				if expr.val1 == nil && expr.expr1 == nil {
					expr.val1 = &val
				} else {
					expr.val2 = &val
				}
			}
		}
	}
	return expr, i
}

func parseExpressionPart2(line string, start int) (expression, int) {
	expr := expression{}
	i := start
	for ; i < len(line) && line[i] != ')'; i++ {
		if line[i] == '(' {
			var subExpr expression
			subExpr, i = parseExpressionPart2(line, i+1)
			if expr.val1 == nil && expr.expr1 == nil {
				expr.expr1 = &subExpr
			} else {
				expr.expr2 = &subExpr
			}
		} else if line[i] == '*' {
			op := line[i]
			var subExpr expression
			subExpr, i = parseExpressionPart2(line, i+1)
			// complete expression => nest
			tmp := expr
			expr = expression{expr1: &tmp, operator: &op}
			if subExpr.operator == nil {
				if subExpr.val1 != nil {
					expr.val2 = subExpr.val1
				} else {
					expr.expr2 = subExpr.expr1
				}
			} else {
				expr.expr2 = &subExpr
			}
			break
		} else if line[i] == '+' {
			op := line[i]
			if expr.operator == nil {
				expr.operator = &op
			} else {
				// complete expression => nest
				tmp := expr
				expr = expression{expr1: &tmp, operator: &op}
			}
		} else {
			val, err := strconv.Atoi(line[i : i+1])
			if err == nil {
				if expr.val1 == nil && expr.expr1 == nil {
					expr.val1 = &val
				} else {
					expr.val2 = &val
				}
			}
		}
	}
	return expr, i
}

func (e *expression) compute() int {
	val1 := 0
	if e.expr1 != nil {
		val1 = e.expr1.compute()
	} else {
		val1 = *e.val1
	}
	if e.operator == nil {
		return val1
	} else {
		if *e.operator == '+' {
			if e.expr2 != nil {
				return val1 + e.expr2.compute()
			} else {
				return val1 + *e.val2
			}
		} else {
			if e.expr2 != nil {
				return val1 * e.expr2.compute()
			} else {
				return val1 * *e.val2
			}
		}
	}
}

func (e *expression) string() string {
	v1 := ""
	v2 := ""
	if e.val1 != nil {
		v1 = fmt.Sprintf("%d", *e.val1)
	} else if e.expr1 != nil {
		v1 = fmt.Sprintf("(%s)", e.expr1.string())
	}
	if e.val2 != nil {
		v2 = fmt.Sprintf("%d", *e.val2)
	} else if e.expr2 != nil {
		v2 = fmt.Sprintf("(%s)", e.expr2.string())
	}
	if e.operator != nil {
		return fmt.Sprintf("%s %c %s", v1, *e.operator, v2)
	}
	return v1
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		expr, _ := parseExpressionPart1(line, 0)
		res += expr.compute()
	}
	fmt.Println(res)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()
		expr, _ := parseExpressionPart2(line, 0)
		res += expr.compute()
	}
	fmt.Println(res)
}

func test1() {
	type testCase struct {
		val string
		res int
	}
	cases := []testCase{
		{"1+2*3+4*5+6", 71},
		{"2*3", 6},
		{"(4*5)", 20},
		{"4*5", 20},
		{"(2*3)+(4*5)", 26},
		{"2*3+4*5", 50},
		{"2*3+(4*5)", 26},
		{"5+(8*3+9+3*4*3)", 437},
		{"5*9*(7*3*3+9*3+(8+6*4))", 12240},
		{"((2+4*9)*(6+9*8+6)+6)+2+4*2", 13632},
		{"1+(2*3)+(4*(5+6))", 51},
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"2 * 3", 6},
		{"(4 * 5)", 20},
		{"4 * 5", 20},
		{"(2 * 3) + (4 * 5)", 26},
		{"2 * 3 + 4 * 5", 50},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
	}
	for _, c := range cases {
		fmt.Printf("original: %s\n", c.val)
		expr, _ := parseExpressionPart1(c.val, 0)
		fmt.Printf("generated: %s\n", expr.string())
		res := expr.compute()
		if res == c.res {
			fmt.Printf("\033[32m%s == %d\033[0m\n", c.val, c.res)
		} else {
			fmt.Printf("\033[31m%s == %d, expected %d\033[0m\n", c.val, res, c.res)
		}
	}
}

func test2() {
	type testCase struct {
		val string
		res int
	}
	cases := []testCase{
		{"1+2*3+4*5+6", 231},
		{"1+(2*3)+(4*(5+6))", 51},
		{"2*3+(4*5)", 46},
		{"5+(8*3+9+3*4*3)", 1445},
		{"5*9*(7*3*3+9*3+(8+6*4))", 669060},
		{"((2+4*9)*(6+9*8+6)+6)+2+4*2", 23340},
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
		{"(9 + 4 * 9 * 4) + 3 + 7 * 8", 3824},
	}
	for _, c := range cases {
		fmt.Printf("original: %s\n", c.val)
		expr, _ := parseExpressionPart2(c.val, 0)
		fmt.Printf("generated: %s\n", expr.string())
		res := expr.compute()
		if res == c.res {
			fmt.Printf("\033[32m%s == %d\033[0m\n", c.val, c.res)
		} else {
			fmt.Printf("\033[31m%s == %d, expected %d\033[0m\n", c.val, res, c.res)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go [1|2]")
		os.Exit(1)
	}
	if os.Args[1] == "test1" {
		test1()
	} else if os.Args[1] == "test2" {
		test2()
	} else if os.Args[1] == "1" {
		part1()
	} else if os.Args[1] == "2" {
		part2()
	}
}

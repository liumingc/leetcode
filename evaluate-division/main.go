package main

import "fmt"

func main() {
	eqns := [][]string{
		//{"x1", "x2"}, {"x2", "x3"}, {"x1", "x4"}, {"x2", "x5"},
		{"a", "b"}, {"b", "c"},
	}
	values := []float64{
		//3.0, 0.5, 3.4, 5.6,
		2.0, 3.0,
	}
	queries := [][]string{
		//{"x2", "x4"}, {"x1", "x5"}, {"x1", "x3"}, {"x5", "x5"}, {"x5", "x1"}, {"x3", "x4"}, {"x4", "x3"}, {"x6", "x6"}, {"x0", "x0"},
		{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"},
	}
	res := calcEquation(eqns, values, queries)
	fmt.Println("eqns=", eqns)
	fmt.Println("vals=", values)
	fmt.Println("queries=", queries)
	fmt.Println(res)
}

func printVarMap(varMap map[string]*vari) {
	for _, v := range varMap {
		printVar(v)
	}
}

func printVar(v *vari) {
	fmt.Printf("var: %s, q=%d\n", v.sym, v.q)
	for x := range v.neib {
		fmt.Println(x)
	}
	fmt.Println("====")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andy/gopl.io/ch7/eval"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter expression")
	expr, _ := reader.ReadString('\n')
	fmt.Println(expr)
	fmt.Println("Enter vars each separated by space, e.g x=7 y=48")
	vars, _ := reader.ReadString('\n')
	fmt.Println(vars)
	p, err := eval.Parse(expr)
	if err != nil {
		fmt.Errorf("could not parse expression %q", err)
		os.Exit(1)
	}
	//get the Envs
	varss := strings.Fields(vars)
	env := eval.Env{}
	for i, v := range varss {
		fmt.Printf("var %d, %s\n", i, v)
		f := strings.Split(v, "=")
		if len(f) != 2 {
			fmt.Errorf("must enter vars as x=n format")
			os.Exit(1)
		}
		fv, err := strconv.ParseFloat(f[1], 64)
		if err != nil {
			fmt.Errorf("could not parse convert string to float %q", err)
			os.Exit(1)

		}
		env[eval.Var(f[0])] = fv
	}

	//calc

	fmt.Println(p.String())
	fmt.Println(p.Eval(env))
}

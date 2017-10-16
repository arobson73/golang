package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/andy/gopl.io/ch7/eval"
)

func main() {
	http.HandleFunc("/calc", calc)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

type Res struct {
	Result float64
	Expr   string
	Env    string
}

//Note this blocked the ability of writing back to the
//html.
/*
var temp = template.Must(template.New("temp").Parse(`
<!DOCTYPE html>
<html>
<head>
<title></title>
</head>
<body>
<style>
  .hide { position:absolute; top:-1px; left:-1px; width:1px; height:1px; }
</style>
<iframe name="hiddenFrame" class="hide"></iframe>
<form action="/calc" method="post" target="hiddenFrame">
    Expr:<input type="text" name="expr">
    Env:<input type="text" name="env">
    <input type="submit" value="Calc">
</form>
<h5>{{.Result}}</h5>
</body>
</html>
</html>
`))
*/
var temp = template.Must(template.New("temp").Parse(`
<!DOCTYPE html>
<html>
<head>
<title></title>
</head>
<body>
<form action="/calc" method="post">
    Expr:<input type="text" name="expr">
    Env:<input type="text" name="env">
    <input type="submit" value="Calc">
</form>
<h4>Result:</h4>
<h5>{{.Result}}</h5>	
</body>
</html>
</html>
`))

func calc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := temp
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		var rs Res
		//should do size checks here
		if len(r.Form["expr"]) != 1 || len(r.Form["env"]) != 1 {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return

		}
		//fmt.Println(len(r.Form["expr"]))
		rs.Expr = r.Form["expr"][0]
		pr, err := eval.Parse(rs.Expr)
		if err != nil {
			http.Error(w, "can't parse expression", http.StatusBadRequest)
			return
		}
		rs.Env = r.Form["env"][0]
		vals := getVals(rs.Env)
		vars := getVars(rs.Env)
		if len(vals) == 0 || len(vars) == 0 || len(vals) != len(vars) {
			http.Error(w, "no valid numbers / vars entered", http.StatusBadRequest)
			return
		}
		valsf := converttoFloat(vals)
		env := eval.Env{}
		for i := 0; i < len(vars); i++ {
			env[eval.Var(vars[i])] = valsf[i]
		}

		rs.Result = pr.Eval(env)

		display(w, rs)

	}
}

func getVars(in string) []string {
	//	fmt.Println("in getVars ", in)
	var reg = regexp.MustCompile(`[A-z]+`)
	r := reg.FindAllString(in, -1)
	//	fmt.Println(r)
	return r

}

func getVals(in string) []string {
	var valid = regexp.MustCompile(`[+-]?([0-9]*[.])?[0-9]+`)
	r := valid.FindAllString(in, -1)
	//	fmt.Println(r)
	return r

}

func converttoFloat(in []string) []float64 {
	var res []float64
	for _, v := range in {
		fv, _ := strconv.ParseFloat(v, 64)
		res = append(res, fv)
	}
	return res
}

func display(w io.Writer, d Res) {
	if err := temp.Execute(w, d); err != nil {
		fmt.Println("Error writing to template")

	}

}

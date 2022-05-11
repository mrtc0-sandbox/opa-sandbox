package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"context"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	input := map[string]interface{}{
		"method": "GET",
		"path": []interface{}{"slary", "bob"},
		"subject": map[string]interface{}{
			"user": "bob",
			"groups": []interface{}{"sales", "marketing"},
		},
	}

	module, err := loadRego("example.rego")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.TODO()
	query, err := rego.New(
		rego.Query("x = data.example.authz.allow"),
		rego.Module("example.rego", module),
	).PrepareForEval(ctx)

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} 

	if !results.Allowed() {
		fmt.Printf("%#v\n", results)
	}
}

func loadRego(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	} 

	return string(bytes), err
}

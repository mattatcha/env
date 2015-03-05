package main

import (
	"log"
	"os"
	"text/template"
)

var genTemplate = `// generated code -- DO NOT EDIT
package envconfig

{{range .}}
func {{.Upper}}(name string, defaultVal {{.Lower}}, usage string) {{.Lower}} {
	v := Var(new{{.Upper}}Value(defaultVal), name, usage)
	return v.Value.Get().({{.Lower}})
}
{{end}}`

func main() {

	// Prepare some data to insert into the template.
	type Type struct {
		Upper, Lower string
	}
	var types = []Type{
		{"String", "string"},
		{"Bool", "bool"},
		{"Float64", "float64"},
		{"Int", "int"},
		{"Int64", "int64"},
	}

	t := template.Must(template.New("envconfig").Parse(genTemplate))

	err := t.Execute(os.Stdout, types)
	if err != nil {
		log.Println("executing template:", err)
	}

}

package main

import (
	"log"
	"os"
	"text/template"
)

var genTemplate = `// generated code -- DO NOT EDIT
package envconfig

import "net"

{{range .}}
func {{.Upper}}(name string, defaultVal {{.Lower}}, usage string) {{.Lower}} {
	v := Var(new{{.Upper}}Value(defaultVal), name, usage)
	return v.Value.Get().({{.Lower}})
}
{{if .Options}}
func {{.Upper}}Option(name string, defaultVal {{.Lower}}, options []{{.Lower}}, usage string) {{.Lower}} {
	v := Var(new{{.Upper}}Value(defaultVal), name, usage)
	for _,option := range options {
		if option == v.Value.Get().({{.Lower}}) {
			return v.Value.Get().({{.Lower}})
		}
	}
	return nil
}
{{end}}
{{end}}`

func main() {

	// Prepare some data to insert into the template.
	type Type struct {
		Upper, Lower string
		Options      bool
	}
	var types = []Type{
		{"String", "string", false},
		{"Bool", "bool", false},
		{"Float64", "float64", false},
		{"Int", "int", false},
		{"Int64", "int64", false},
		{"Uint", "uint", false},
		{"Uint64", "uint64", false},
		{"IP", "net.IP", false},
	}

	t := template.Must(template.New("envconfig").Parse(genTemplate))

	err := t.Execute(os.Stdout, types)
	if err != nil {
		log.Println("executing template:", err)
	}

}

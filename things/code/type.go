package code

import "github.com/dpb587/go-schemaorg"

// // Computer programming source code. Example: Full (compile ready) solutions,
// code snippet samples, scripts, templates.
var Type = schemaorg.NewDataType("http://schema.org", "Code")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

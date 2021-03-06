package action

import "github.com/dpb587/go-schemaorg"

// // An action performed by a direct agent and indirect participants upon a direct
// object. Optionally happens at a location with the help of an inanimate
// instrument. The execution of the action may produce a result. Specific action
// sub-type documentation specifies the exact expectation of each
// argument/role.</p>
// 
// <p>See also <a
// href="http://blog.schema.org/2014/04/announcing-schemaorg-actions.html">blog
// post</a> and <a href="http://schema.org/docs/actions.html">Actions overview
// document</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Action")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

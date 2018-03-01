package userpagevisits

import "github.com/dpb587/go-schemaorg"

// // UserInteraction and its subtypes is an old way of talking about users
// interacting with pages. It is generally better to use <a class="localLink"
// href="http://schema.org/Action">Action</a>-based vocabulary, alongside types
// such as <a class="localLink" href="http://schema.org/Comment">Comment</a>.
var Type = schemaorg.NewDataType("http://schema.org", "UserPageVisits")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

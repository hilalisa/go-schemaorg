package hotel

import "github.com/dpb587/go-schemaorg"

// // A hotel is an establishment that provides lodging paid on a short-term basis
// (Source: Wikipedia, the free encyclopedia, see
// http://en.wikipedia.org/wiki/Hotel).
// <br /><br />
// See also the <a href="/docs/hotels.html">dedicated document on the use of
// schema.org for marking up hotels and other forms of accommodations</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Hotel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}

package blog

import "github.com/dpb587/go-schemaorg"

var (
	// The postings that are part of this blog.
	BlogPosts = schemaorg.NewProperty("blogPosts")

	// A posting that is part of this blog.
	BlogPost = schemaorg.NewProperty("blogPost")

	// The International Standard Serial Number (ISSN) that identifies this serial
	// publication. You can repeat this property to identify different formats of,
	// or the linking ISSN (ISSN-L) for, this serial publication.
	Issn = schemaorg.NewProperty("issn")
)

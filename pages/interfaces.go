package pages

import (
	"io"

	"github.com/osteele/gojekyll/pipelines"
	"github.com/osteele/gojekyll/templates"
)

// Page is a Jekyll page.
type Page interface {
	// Paths
	SiteRelPath() string // relative to the site source directory
	Permalink() string   // relative URL path
	OutputExt() string

	// Output
	Published() bool
	Static() bool
	Write(RenderingContext, io.Writer) error

	// Variables
	PageVariables() templates.VariableMap

	// internal
	initPermalink() error
}

// RenderingContext provides context information to a Page.
type RenderingContext interface {
	RenderingPipeline() pipelines.PipelineInterface
	SiteVariables() templates.VariableMap // value of the "site" template variable
}

// Container is the Page container; either the Site or Collection.
type Container interface {
	OutputExt(pathname string) string
	PathPrefix() string // PathPrefix is the relative prefix, "" for the site and "_coll/" for a collection
}
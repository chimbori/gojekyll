package liquid

import (
	"fmt"
	"io"
	"strings"

	"github.com/acstech/liquid"
	"github.com/acstech/liquid/core"
)

func init() {
	liquid.Tags["link"] = linkFactory
}

// A LinkTagHandler given an include tag file name returns a URL.
type LinkTagHandler func(string) (string, bool)

var currentLinkHandler LinkTagHandler

// SetLinkHandler sets the function that resolves an include tag file name to a URL.
func SetLinkHandler(h LinkTagHandler) {
	currentLinkHandler = h
}

// Link tag data, for passing information from the factory to Execute
type linkData struct {
	standaloneTag
	filename string
}

// linkFactory creates a link tag
func linkFactory(p *core.Parser, config *core.Configuration) (core.Tag, error) {
	start := p.Position
	p.SkipPastTag()
	end := p.Position - 2
	filename := strings.TrimSpace(string(p.Data[start:end]))
	return &linkData{standaloneTag{"link"}, filename}, nil
}

// Execute is required by the Liquid tag interface
func (l *linkData) Execute(writer io.Writer, data map[string]interface{}) core.ExecuteState {
	url, ok := currentLinkHandler(l.filename)
	if !ok {
		panic(fmt.Errorf("link tag: %s not found", l.filename))
	}
	if _, err := writer.Write([]byte(url)); err != nil {
		panic(err)
	}
	return core.Normal
}

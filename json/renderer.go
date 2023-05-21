package json

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Renderer struct {
	heading bool
}

func NewRenderer() renderer.NodeRenderer {
	r := &Renderer{}

	return r
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindHeading, r.renderHeading)

	reg.Register(ast.KindText, r.renderText)
}

func (r *Renderer) renderHeading(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("\n\"heading\": \"")
		r.heading = true
	} else {
		_, _ = w.WriteString("\"\n")
		r.heading = false
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	segment := n.Segment

	v := segment.Value(source)

	w.Write(v)

	if r.heading {
		w.WriteString(".")
	}

	return ast.WalkContinue, nil
}

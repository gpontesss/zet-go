package md

import mdast "github.com/yuin/goldmark/ast"

// FirstHeading walks a markdown node tree and returns the first matching
// heading.
func FirstHeading(tree mdast.Node) mdast.Node {
	var heading mdast.Node
	mdast.Walk(tree, func(node mdast.Node, _ bool) (mdast.WalkStatus, error) {
		if node.Kind().String() == "Heading" {
			heading = node
			return mdast.WalkStop, nil
		}
		return mdast.WalkContinue, nil
	})
	return heading
}

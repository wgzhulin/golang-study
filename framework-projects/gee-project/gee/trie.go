package gee

import "strings"

type node struct {
	pattern  string
	part     string
	isWild   bool
	children []*node
}

func parsePattern(pattern string) []string {
	parts := strings.Split(pattern, "/")

	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		result = append(result, part)
		if part[0] == '*' {
			break
		}
	}

	return result
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for i, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, n.children[i])
		}
	}

	return nodes
}

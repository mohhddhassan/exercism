
package pov

type Tree struct {
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{
		value:    value,
		children: children,
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}

	result := tr.Value()

	if len(tr.Children()) == 0 {
		return result
	}

	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}

	return "(" + result + ")"
}

// findNode searches for a node with the given value.
func findNode(node *Tree, value string) *Tree {
	if node == nil {
		return nil
	}

	if node.value == value {
		return node
	}

	for _, child := range node.children {
		if found := findNode(child, value); found != nil {
			return found
		}
	}

	return nil
}

// buildGraph converts the tree into an undirected graph.
func buildGraph(node *Tree, graph map[*Tree][]*Tree) {
	if node == nil {
		return
	}

	for _, child := range node.children {
		graph[node] = append(graph[node], child)
		graph[child] = append(graph[child], node)

		buildGraph(child, graph)
	}
}

// FromPov returns the tree from the specified node's perspective.
func (tr *Tree) FromPov(from string) *Tree {
	if tr == nil {
		return nil
	}

	// Find the requested root.
	root := findNode(tr, from)
	if root == nil {
		return nil
	}

	// Convert the original tree into an undirected graph.
	graph := make(map[*Tree][]*Tree)
	buildGraph(tr, graph)

	// Rebuild the tree from the new root.
	var rebuild func(node, previous *Tree) *Tree

	rebuild = func(node, previous *Tree) *Tree {
		newNode := &Tree{
			value: node.value,
		}

		for _, neighbor := range graph[node] {
			// Don't travel back to the previous node.
			if neighbor == previous {
				continue
			}

			newNode.children = append(
				newNode.children,
				rebuild(neighbor, node),
			)
		}

		return newNode
	}

	return rebuild(root, nil)
}

// PathTo returns the shortest path between two nodes.
func (tr *Tree) PathTo(from, to string) []string {
	if tr == nil {
		return nil
	}

	start := findNode(tr, from)
	end := findNode(tr, to)

	if start == nil || end == nil {
		return nil
	}

	// Build an undirected graph.
	graph := make(map[*Tree][]*Tree)
	buildGraph(tr, graph)

	// BFS queue.
	queue := []*Tree{start}

	visited := map[*Tree]bool{
		start: true,
	}

	// Store the previous node for each visited node.
	parent := make(map[*Tree]*Tree)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range graph[current] {
			if visited[neighbor] {
				continue
			}

			visited[neighbor] = true
			parent[neighbor] = current
			queue = append(queue, neighbor)
		}
	}

	if !visited[end] {
		return nil
	}

	// Reconstruct the shortest path.
	var path []string

	for current := end; current != nil; current = parent[current] {
		path = append(path, current.value)

		if current == start {
			break
		}
	}

	// Reverse the path.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

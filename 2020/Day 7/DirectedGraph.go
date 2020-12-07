package main // directed weighted graph with optional color field for each vertex

type DirectedGraph struct {
	vertices map[string]*Vertex
}

type Vertex struct {
	name, color        string
	incoming, outgoing []*Edge
}

type Edge struct {
	head, tail *Vertex
	weight     int
}

func (graph DirectedGraph) AddVertex(vertex *Vertex) {
	graph.vertices[vertex.name] = vertex
}

func (graph DirectedGraph) AddEdge(from, to *Vertex, weight int) {
	newEdge := Edge{
		head:   to,
		tail:   from,
		weight: weight,
	}
	from.outgoing = append(from.outgoing, &newEdge)
	to.incoming = append(to.incoming, &newEdge)
}

func (graph DirectedGraph) GetVertex(name string) *Vertex {
	return graph.vertices[name]
}

func (graph DirectedGraph) GetOrCreateVertex(name string) *Vertex {
	if !graph.HasVertex(name) {
		graph.AddVertex(NewVertex(name))
	}
	return graph.vertices[name]
}

func (graph DirectedGraph) HasVertex(name string) bool {
	return graph.vertices[name] != nil
}

func NewGraph() DirectedGraph {
	return DirectedGraph{vertices: make(map[string]*Vertex)}
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		name:     name,
		color:    "",
		incoming: []*Edge{},
		outgoing: []*Edge{},
	}
}

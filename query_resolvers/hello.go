package query_resolvers

import "context"

type Query struct{}

type PageInfo struct {
	startCursor string
	endCursor string
	hasNextPage bool
	hasPreviousPage bool
}

type Hello struct {
	text string
}

type HelloEdge struct {
	cursor string
	node *Hello
}

type HelloConnection struct {
	totalCount int32
	edges *[]*HelloEdge
	pageInfo *PageInfo
}

func (_ *Query) HelloCollection() *HelloConnection {
	// TODO: inline this
	t := HelloEdge{cursor: "<some cursor>", node: &Hello{text: "Some hello"}}
	res := []*HelloEdge{}
	res = append(res, &t)

	return &HelloConnection{
		pageInfo: &PageInfo{
			startCursor: "<startcursor>",
			endCursor: "<endcursor>",
			hasNextPage: true,
			hasPreviousPage: false,
		},
		totalCount: 100500,
		edges: &res,
	}
}

// HelloConnection

func (p *HelloConnection) PageInfo(ctx context.Context) *PageInfo {
	return p.pageInfo
}

func (p *HelloConnection) TotalCount(ctx context.Context) int32 {
	return p.totalCount
}

func (p *HelloConnection) Edges(ctx context.Context) *[]*HelloEdge {
	return p.edges
}

// PageInfo

func (p *PageInfo) StartCursor (ctx context.Context) *string {
	return &p.startCursor
}

func (p *PageInfo) EndCursor (ctx context.Context) *string {
	return &p.endCursor
}

func (p *PageInfo) HasNextPage (ctx context.Context) bool {
	return p.hasNextPage
}

func (p *PageInfo) HasPreviousPage (ctx context.Context) bool {
	return p.hasPreviousPage
}


// HelloEdge

func (p *HelloEdge) Cursor(ctx context.Context) *string {
	return &p.cursor
}

func (p *HelloEdge) Node(ctx context.Context) *Hello {
	return p.node
}

// Hello

func (p *Hello) Text(ctx context.Context) *string {
	return &p.text
}
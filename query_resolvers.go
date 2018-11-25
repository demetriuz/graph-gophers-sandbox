package main

import "context"

type query struct{}

type Resp struct {
	Name string
}

type Subresp struct {
	Name string
}

func (p *Resp) NAME(ctx context.Context) *string {
	return &p.Name
}

func (p *Resp) SUBRESP(ctx context.Context) *Subresp {
	return &Subresp{Name: "Subresp name"}
}

func (p *Subresp) NAME(ctx context.Context) *string {
	return &p.Name
}


func (_ *query) Hello() string { return "Hello, world!" }
func (_ *query) Bye(ctx context.Context, args struct{ Name string }) (*Resp, error) {
	return &Resp{Name: args.Name}, nil
}

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type SingleGood struct {
	itemID string
	name   string
	amount int
}

func ChangeGood(ctx context.Context, c *app.RequestContext) {
	//id := c.Query("id")

}

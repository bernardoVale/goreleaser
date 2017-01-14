package repos

import (
	"strings"

	"github.com/goreleaser/releaser/context"
)

// Pipe for brew deployment
type Pipe struct{}

// Description of the pipe
func (Pipe) Description() string {
	return "Filling repositories data..."
}

// Run the pipe
func (Pipe) Run(ctx *context.Context) (err error) {
	owner, name := split(ctx.Config.Repo)
	ctx.Repo = &context.Repo{
		Owner: owner,
		Name:  name,
	}
	owner, name = split(ctx.Config.Brew.Repo)
	ctx.BrewRepo = &context.Repo{
		Owner: owner,
		Name:  name,
	}
	return
}

func split(pair string) (string, string) {
	parts := strings.Split(pair, "/")
	return parts[0], parts[1]
}
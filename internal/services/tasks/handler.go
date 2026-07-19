package tasks

import service "github.com/elum2b/services/tasks"

func handler(ctx service.Context) error {

	return ctx.Successful()

}

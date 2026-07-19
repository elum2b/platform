package payment

import service "github.com/elum2b/services/payment"

func handler(ctx service.Context) error {

	return ctx.Successful()

}

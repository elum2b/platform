package services

import (
	"context"

	"github.com/elum2b/services/calendar"
	"github.com/elum2b/services/control"
	controlapi "github.com/elum2b/services/control/service/internalapi"
	"github.com/elum2b/services/cpa"
	"github.com/elum2b/services/delivery"
	"github.com/elum2b/services/payment"
	"github.com/elum2b/services/promo"
	"github.com/elum2b/services/reference"
	"github.com/elum2b/services/tasks"
)

var (
	Calendar  = calendar.New()
	Control   = control.New()
	CPA       = cpa.New()
	Payment   = payment.New()
	Promo     = promo.New()
	Reference = reference.New()
	Tasks     = tasks.New()

	Delivery = func() *delivery.Delivery {
		value, err := delivery.New(delivery.ResolverFunc(
			func(
				ctx context.Context,
				destination delivery.Destination,
			) (delivery.Endpoint, error) {
				endpoint, err := Control.Internal.GetApplicationDeliveryEndpoint(
					ctx,
					controlapi.ApplicationDeliveryRequest{
						WorkspaceID: destination.WorkspaceID,
						AppID:       destination.AppID,
						PlatformID:  destination.PlatformID,
					},
				)
				if err != nil {
					return delivery.Endpoint{}, err
				}

				return delivery.Endpoint{
					URL:       endpoint.URL,
					Secret:    endpoint.Secret,
					IsEnabled: endpoint.IsEnabled,
				}, nil
			},
		))
		if err != nil {
			panic(err)
		}

		return value
	}()
)

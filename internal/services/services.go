package services

import (
	"github.com/elum2b/services/calendar"
	"github.com/elum2b/services/control"
	"github.com/elum2b/services/cpa"
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
)

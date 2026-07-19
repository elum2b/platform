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
	Calendar  *calendar.Calendar
	Control   *control.Control
	CPA       *cpa.CPA
	Payment   *payment.Payment
	Promo     *promo.Promo
	Reference *reference.Reference
	Tasks     *tasks.Tasks
)

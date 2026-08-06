package rule

import (
	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveRequest struct {
	WorkspaceID  string  `json:"workspace_id"   validate:"required,uuid"`
	Provider     string  `json:"provider"       validate:"required"`
	GroupKey     string  `json:"group_key"      validate:"required,max=255"`
	ExternalType string  `json:"external_type"  validate:"required"`
	Key          string  `json:"key"            validate:"required,max=255"`
	Type         string  `json:"type"           validate:"required,max=255"`
	Quantity     int64   `json:"quantity"`
	Scale        uint16  `json:"scale"`
	Unit         *string `json:"unit,omitempty"`
	Position     int32   `json:"position"       validate:"required,min=1"`
	IsEnabled    bool    `json:"is_enabled"`
}

var (
	rewardRuleSaveKey         = "tasks.partner.reward_rule.save"
	rewardRuleSaveDescription = `
Creates or updates a partner reward rule. Requires the
'tasks.partner.reward_rule.save' permission in the target workspace.`
)

// Save exposes the partner reward rule upsert method.
var Save = adapter.Method[SaveRequest, struct{}]{
	Key:         rewardRuleSaveKey,
	Description: rewardRuleSaveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(rewardRuleSaveKey),
	},
	Handler: func(ctx *adapter.Context, data SaveRequest) (struct{}, error) {
		err := services.Tasks.Admin.SavePartnerRewardRule(
			ctx.Context,
			tadmin.SavePartnerRewardRuleParams{
				WorkspaceID:  data.WorkspaceID,
				Provider:     data.Provider,
				GroupKey:     data.GroupKey,
				ExternalType: data.ExternalType,
				Reward: tadmin.RewardModel{
					Key:      data.Key,
					Type:     data.Type,
					Quantity: data.Quantity,
					Scale:    data.Scale,
					Unit:     data.Unit,
				},
				Position:  data.Position,
				IsEnabled: data.IsEnabled,
			},
		)

		return struct{}{}, err
	},
}

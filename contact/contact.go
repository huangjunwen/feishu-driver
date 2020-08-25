package contact

import (
	"context"

	"github.com/huangjunwen/feishu-driver/conf"
	"github.com/huangjunwen/feishu-driver/utils"
)

// ContactScopeResult 包含通讯录授权范围
type ContactScopeResult struct {
	utils.APIResultBase

	Data struct {
		AuthedDepartments     []string `json:"authed_departments"`
		AuthedOpenDepartments []string `json:"authed_open_departments"`
		AuthedEmployeeIds     []string `json:"authed_employee_ids"`
		AuthedOpenIds         []string `json:"authed_open_ids"`
	} `json:"data"`
}

// GetContactScope 获得通讯录授权范围
func GetContactScope(ctx context.Context, provider conf.TenantAccessTokenProvider) (*ContactScopeResult, error) {
	result := &ContactScopeResult{}
	err := utils.GetJSONWithTenantAccessToken(ctx, "/contact/v1/scope/get", provider, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

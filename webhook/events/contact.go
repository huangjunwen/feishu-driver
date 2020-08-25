package events

// UserAdd 员工加入企业
type UserAdd struct {
	AppId      string `json:"app_id"`
	TenantKey  string `json:"tenant_key"`
	OpenId     string `json:"open_id"`
	EmployeeId string `json:"employee_id"`
	UnionId    string `json:"union_id"`
}

// UserUpdate 个人信息发生变化
type UserUpdate struct {
	AppId      string `json:"app_id"`
	TenantKey  string `json:"tenant_key"`
	OpenId     string `json:"open_id"`
	EmployeeId string `json:"employee_id"`
	UnionId    string `json:"union_id"`
}

// UserLeave 离职
type UserLeave struct {
	AppId      string `json:"app_id"`
	TenantKey  string `json:"tenant_key"`
	OpenId     string `json:"open_id"`
	EmployeeId string `json:"employee_id"`
	UnionId    string `json:"union_id"`
}

// DeptAdd 新建部门
type DeptAdd struct {
	AppId            string `json:"app_id"`
	TenantKey        string `json:"tenant_key"`
	OpenDepartmentId string `json:"open_department_id"`
}

// DeptUpdate 修改部门
type DeptUpdate struct {
	AppId            string `json:"app_id"`
	TenantKey        string `json:"tenant_key"`
	OpenDepartmentId string `json:"open_department_id"`
}

// DeptDelete 删除部门
type DeptDelete struct {
	AppId            string `json:"app_id"`
	TenantKey        string `json:"tenant_key"`
	OpenDepartmentId string `json:"open_department_id"`
}

// UserStatusChange 用户状态变更
type UserStatusChange struct {
	AppId        string `json:"app_id"`
	TenantKey    string `json:"tenant_key"`
	OpenId       string `json:"open_id"`
	EmployeeId   string `json:"employee_id"`
	UnionId      string `json:"union_id"`
	BeforeStatus struct {
		IsActive   bool `json:"is_active"`
		IsFrozen   bool `json:"is_frozen"`
		IsResigned bool `json:"is_resigned"`
	} `json:"before_status"`
	CurrentStatus struct {
		IsActive   bool `json:"is_active"`
		IsFrozen   bool `json:"is_frozen"`
		IsResigned bool `json:"is_resigned"`
	} `json:"current_status"`
	ChangeTime string `json:"change_time"`
}

// ContactScopeChange 授权范围变更
type ContactScopeChange struct {
	AppId     string `json:"app_id"`
	TenantKey string `json:"tenant_key"`
}

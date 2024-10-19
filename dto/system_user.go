package dto

// 用户筛选条件请求
type SystemUserSearchRequest struct {
	Username      string `json:"username"`
	CNName        string `json:"cnName"`
	ENName        string `json:"enName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Gender        *uint  `json:"gender"`
	NativePlace   string `json:"nativePlace"`
	Department    string `json:"department"`
	JobPosition   string `json:"jobPosition"`
	JobId         string `json:"jobId"`
	OfficeAddress string `json:"officeAddress"`
	OfficeStation string `json:"officeStation"`
	Status        *uint  `json:"status"`
	CreatorId     *uint  `json:"creatorId"`
	SystemRoleId  *uint  `json:"systemRoleId"`
}

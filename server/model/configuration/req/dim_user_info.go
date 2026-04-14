package req

type DimUserInfoSaveRequest struct {
	ID         int      `json:"id"`
	UserName   string   `json:"userName"`
	UserType   int      `json:"userType"`
	UserRemark string   `json:"userRemark"`
	IpAddress  []string `json:"ipAddress"`
}

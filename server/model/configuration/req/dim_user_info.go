package req

type DimUserInfoSaveRequest struct {
	ID        int      `json:"id"`
	UserName  string   `json:"userName"`
	UserType  int      `json:"userType"`
	Remark    string   `json:"remark"`
	IpAddress []string `json:"ipAddress"`
}

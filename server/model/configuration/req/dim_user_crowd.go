package req

type DimUserCrowdSaveRequest struct {
	ID        int    `json:"id"`
	CrowdName string `json:"crowdName"`
	Remark    string `json:"remark"`
	UserIds   []int  `json:"userIds"`
}

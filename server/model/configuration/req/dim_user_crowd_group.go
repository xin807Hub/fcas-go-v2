package req

type DimUserCrowdGroupSaveRequest struct {
	ID        int    `json:"id"`
	GroupName string `json:"groupName"`
	Remark    string `json:"remark"`
	CrowdIds  []int  `json:"crowdIds"`
}

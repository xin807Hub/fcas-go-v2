package traffic

type UserCommonRankItemName interface {
	GetID() int64
	SetName(string)
}

type UserCommonRankParams struct {
	// 选择使用什么数据库和粒度表所需参数
	Particle      int    `json:"-"`
	TimeRangeType string `json:"-"`

	// 查询条件参数
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`

	UserIdList  []int64 `json:"userIdList"`  // for user
	CrowdIdList []int64 `json:"crowdIdList"` // for user_crowd
	GroupIdList []int64 `json:"groupIdList"` // for user_crowd_group

	LinkIdList  []int64  `json:"linkIdList"`
	AppIdList   []int64  `json:"appIdList"`
	IspNameList []string `json:"ispNameList"` // 运营商

	// param related to level2
	TopN int `json:"topN"`

	// params related to level3
	SrcIp string `json:"srcIp"`
}

type RankLevel1Base struct {
	ID           int64   `json:"id" gorm:"column:user_id"`
	StartTime    string  `json:"startTime" gorm:"column:start_time"`
	TrafficUp    uint64  `json:"trafficUp" gorm:"column:traffic_up"`
	TrafficDn    uint64  `json:"trafficDn" gorm:"column:traffic_dn"`
	TrafficTotal uint64  `json:"trafficTotal" gorm:"column:traffic_total"`
	SpeedUp      float64 `json:"speedUp" gorm:"column:speed_up"`
	SpeedDn      float64 `json:"speedDn" gorm:"column:speed_dn"`

	Name string `json:"name" gorm:"column:-"` // 用于展示
}

func (r *RankLevel1Base) GetID() int64 {
	return r.ID
}

func (r *RankLevel1Base) SetName(name string) {
	r.Name = name
}

// RankLevel1AggTraffic 根据 user_id 聚合后的用户的流量数据
type RankLevel1AggTraffic struct {
	ID           int64  `json:"id" gorm:"column:user_id"`
	TrafficUp    uint64 `json:"trafficUp" gorm:"column:traffic_up"`
	TrafficDn    uint64 `json:"trafficDn" gorm:"column:traffic_dn"`
	TrafficTotal uint64 `json:"trafficTotal" gorm:"column:traffic_total"`

	MaxSpeedUp float64 `json:"maxSpeedUp"`
	MaxSpeedDn float64 `json:"maxSpeedDn"`
	AvgSpeedUp float64 `json:"avgSpeedUp"`
	AvgSpeedDn float64 `json:"avgSpeedDn"`

	Count uint64 `json:"-" gorm:"-"`           // 用于平均值计算
	Name  string `json:"name" gorm:"column:-"` // 用户名，用于展示
}

func (r *RankLevel1AggTraffic) GetID() int64 {
	return r.ID
}

func (r *RankLevel1AggTraffic) SetName(name string) {
	r.Name = name
}

type RankLevel2Table struct {
	SrcIp           string  `json:"srcIp" gorm:"column:src_ip"`
	TrafficUp       uint64  `json:"trafficUp" gorm:"column:traffic_up"`
	TrafficDn       uint64  `json:"trafficDn" gorm:"column:traffic_dn"`
	TrafficTotal    uint64  `json:"trafficTotal" gorm:"column:traffic_total"`
	AvgSpeedUp      float64 `json:"avgSpeedUp" gorm:"column:avg_speed_up"`
	AvgSpeedDn      float64 `json:"avgSpeedDn" gorm:"column:avg_speed_dn"`
	TotalProportion float64 `json:"totalProportion" gorm:"-"`
}

type RankLevel3Table struct {
	StartTime  string  `json:"startTime" gorm:"column:start_time"`
	SrcIp      string  `json:"srcIp" gorm:"column:src_ip"`
	SpeedUp    float64 `json:"speedUp" gorm:"column:speed_up"`
	SpeedDn    float64 `json:"speedDn" gorm:"column:speed_dn"`
	TotalSpeed float64 `json:"totalSpeed" gorm:"-"`
}

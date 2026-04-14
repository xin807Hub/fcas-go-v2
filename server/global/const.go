package global

const (
	DatabaseV1 = "bigdata_fcas"
	DatabaseV2 = "bigdata_fcas_v2"

	Interval10mParticle = 600   // 10min 60s * 10
	Interval1hParticle  = 3600  // 1h 60s * 1 * 60
	Interval1dParticle  = 86400 // 1day 60s * 1 * 60 *  24

	OneDays   = 1  // 查询开始时间、结束时间的跨度大于1天，要从小时表里取数据
	TwoDays   = 2  // 查询开始时间、结束时间的跨度大于2天，要从天表里取数据
	SevenDays = 7  // 查询开始时间、结束时间的跨度大于2天，要从小时表里取数据
	OneMonth  = 31 // 查询开始时间距离现在大于30天要查询小时表

	DateTimeLayout = "2006-01-02 15:04:05" // 定义时间格式
	TimeLayout     = "15:04:05"            // 定义时间格式

	LevelOne   = "level1"
	LevelTwo   = "level2"
	LevelThree = "level3"

	QueryOld          = "old"
	QueryCrossOld2New = "cross"
	QueryNew          = "new"

	OtherPart = "其余"

	DictTypeAllIsp        = "isp"
	DictTypeSelectIsp     = "ispSelect"
	DictTypeAppType       = "appType"
	DictTypeAppId         = "appId"
	DictTypeAppTypeIdTree = "appTypeIdTree"

	Bind   = "all"
	UnBind = "NA"
)

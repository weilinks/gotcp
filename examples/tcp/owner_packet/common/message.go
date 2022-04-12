package common

const (
	MsgLogin 				= 110 // 登录请求 client -> server
	MsgLoginResp			= 111 // 登录响应 server -> client
	MsgDeviceInfo			= 310 // 设备属性上报 client -> server
	MsgDeviceInfoResp 		= 311 // 设备属性上报响应 server -> client
	MsgDeviceAlarm 			= 410 // 设备告警上报 client -> server
	MsgDeviceAlarmResp 		= 411 // 设备告警响应 server -> client
	MsgRemoteCtrl 			= 500 // 远程控制 server -> client
	MsgRemoteCtrlResp 		= 501 // 远程控制响应 client -> server
	MsgConfigQuery 			= 210 // 远程配置查询 server -> client
	MsgConfigQueryResp 		= 211 // 远程配置查询响应 client -> server
)

// 换电柜状态
type CabinetState string
// 设备类型
type DeviceType int

// 换电柜状态
const (
	INIT 		CabinetState	= "0"		// 初始化
	FREE 		CabinetState 	= "1"		// 空闲
	EXCHANGE 	CabinetState	= "2" 		// 换电
	RETURN 		CabinetState	= "3"		// 还电
	TAKE 		CabinetState	= "4" 		// 取电
	EXCEPTION 	CabinetState 	= "5"		// 异常
	FORBIDDEN 	CabinetState 	= "6"		// 禁用
)

// 设备类型
const (
	BATTERY  			DeviceType		= 	1			// 电池
	EXCHANGE_CABINET	DeviceType		=	2			// 换电柜
	CHARGING_CABINET  	DeviceType		= 	3			// 充电柜
)

// 通用响应结果 对应响应的result字段
type FixtureResponse int

// 通用响应result字段
const (
	Failure		FixtureResponse = 0 // 失败
	Success		FixtureResponse = 1 // 成功
)

// 远程控制设备端响应result字段
var ResponseResult = map[FixtureResponse]string {
	0 	: "失败(未知错误)",
	1 	: "成功",
	2 	: "指令不支持",
	3 	: "前一个流程未结束",
	4 	: "当前没有空柜门",
	5 	: "当前没有满电电池",
	6 	: "当前没有对应电压等级电池",
	16 	: "换电柜正在升级",
	17 	: "换电柜暂停使用",
	40 	: "指令超时",
}

func (l FixtureResponse) String () string {
	return ResponseResult[l]
}

// 设备告警信号量类型
// 设备告警上报涉及的信号指令、远程控制指令的值类型
type CommandType string

func (c CommandType) Get() map[string]string  {
	return SignalText[c]
}

func (c CommandType) String(signalValue string) string {
	return SignalText[c][signalValue]
}

// 设备告警信号
const (
	BmsFault			CommandType  	= "bmsFault"		// BMS故障
	BmsAlarm			CommandType		= "bmsAlarm"		// BMS告警
	CabAlarm			CommandType		= "cabAlarm"		// 整柜告警
	CabFault			CommandType		= "cabFault"		// 整柜故障
	BoxAlarm			CommandType		= "boxAlarm"		// 单仓告警
	BoxFault			CommandType		= "boxFault"		// 单仓故障
	SwitchFinish		CommandType		= "switchFinish"	// 换电柜换电流程终止事件
	SwitchControl		CommandType		= "switchControl"	// 控制换电柜命令 远程下发指令
	// ====================================远程控制指令,其中switchControl告警、下发都有========================
	Handle				CommandType		= "handle"			// 握手
	SwCabVolControl		CommandType		= "swCabVolControl" // 换电柜语音音量  值0-100
	SwCabTempControl	CommandType		= "swCabTempControl"// 换电柜温度阈值控制
	SwCabSocControl		CommandType		= "swCabSocControl" // 换电柜满电阈值
	SwCabReset			CommandType		= "swCabReset"		// 换电柜重启
	SwCabTcpPort		CommandType		= "swCabTcpPort"	// 配置上报的服务器地址和端口  "value":"test.abcd.com,4800"
	StartHeat			CommandType		= "startHeat"		// 开始加热温度
	StopHeat			CommandType		= "stopHeat"		// 停止加热温度
	MaxChgCurrent		CommandType		= "maxChgCurrent"	// 最大充电电流
	OverTemp			CommandType		= "overTemp"		// 整机过温保护值
	RecOverTemp			CommandType		= "recOverTemp"		// 整机过温保护恢复值
	AlarmTemp			CommandType		= "alarmTemp"		// 整机过温告警值
	RecalarmTemp		CommandType		= "recalarmTemp"	// 整机过温告警恢复值
	// ====================================远程查询指令,其中多数指令与远程控制指令相同========================
	SwCabIPPort			CommandType		= "swCabIPPort"		// 查询当前设备的IP和端口
)

var SignalText = map[CommandType](map[string]string) {
	BmsAlarm: BmsAlarmSignalText,
	BmsFault: BmsFaultSignalText,
	CabAlarm: CabAlarmSignalText,
	CabFault: CabFaultSignalText,
	BoxAlarm: BmsAlarmSignalText,
	BoxFault: BoxFaultSignalText,
	SwitchControl: SwitchControlSignalText,
	SwitchFinish: SwitchFinishSignalText,
	Handle: HandleSignalText,
}

// BMS故障
var BmsFaultSignalText = map[string]string {
	"00":	"短路保护",
	"01":	"单芯欠压保护",
	"02":	"单芯过压保护",
	"03":	"放电过流保护",
	"04":	"充电过流保护",
	"05":	"低温保护",
	"06":	"过温保护",
	"07":	"状态异常保护",
	"08":	"MOS异常",
	"09":	"总电压过呀保护",
	"10":	"总电压欠压保护",
	"11":	"单芯间压差过大",
}

// BMS告警
var BmsAlarmSignalText = map[string]string {
	"00": "单芯电压低告警",
	"01": "单芯电压高告警",
	"02": "电芯低温告警",
	"03": "电芯高温告警",
	"04": "总电压高告警",
	"05": "总电压低告警",
	"06": "单芯压差过大告警",
	"07": "MOS高温告警",
	"08": "环境低温告警",
	"09": "环境高温告警",
}

// 整柜告警
var CabAlarmSignalText = map[string]string {
	"01": "过温告警",
	"02": "功率告警",
	"03": "水浸告警",
	"04": "烟雾告警",
	"05": "市电低压告警",
	"06": "市电高压告警",
}

// 整柜故障
var CabFaultSignalText = map[string]string {
	"01": "水浸故障",
	"02": "消防故障",
	"03": "过温故障",
	"04": "市电故障",
	"05": "整柜散热风扇故障",
	"06": "12V开关电源故障",
	"07": "控制板通讯故障",
	"08": "机柜空开故障",
}

// 单仓告警
var BoxAlarmSignalText = map[string]string {
	"01": "消防故障",
	"02": "过温故障",
	"03": "充电器故障",
	"04": "仓锁故障",
	"05": "电池通信故障",
	"06": "充电器短路保护",
	"07": "充电过流",
	"08": "充电过压",
	"09": "指示灯故障",
	"10": "电池异常(无法充电)",
	"11": "仓内照明灯故障",
}

// 单仓故障
var BoxFaultSignalText = map[string]string {
	"01": "消防故障",
	"02": "过温故障",
	"03": "充电器故障",
	"04": "仓锁故障",
	"05": "电池通信故障",
	"06": "充电器短路保护",
	"07": "充电过流",
	"08": "充电过压",
	"09": "指示灯故障",
	"10": "电池异常(无法充电)",
	"11": "仓内照明灯故障",
}

// 控制换电柜命令
var SwitchControlSignalText = map[string]string {
	"00": "设置换电柜不可用",
	"01": "换电",
	"02": "放电(平台控制换电逻辑)",
	"03": "取电(平台控制换电逻辑)",
	"04": "开启柜门",
	"06": "设为柜门不可用",
	"07": "设为柜门可用",
	"08": "柜门绑定电池序列号",
	"09": "柜门解绑电池序列号",
	"10": "设置换电柜可用",
	"11": "租用电池(首放)",
	"12": "退还电池",
}

var SwitchFinishSignalText = map[string]string {
	"21": "用户没有放入电池，终止流程",
	"22": "用户与放入的电池不匹配，终止流程",
	"23": "无法识别电池，终止流程",
	"24": "满电仓门开启失败，终止流程",
	"25": "用户取出电池,未关仓门，流程正常结束",
	"26": "用户取出电池,已关仓门，流程正常结束",
}

var HandleSignalText = map[string]string {
	"01": "换电",
	"11": "租用电池(首放)",
	"12": "退还电池",
}


// tcp 数据包中的指令类型
type MsgType struct {
	MsgType int `json:msgType`				// 报文类型
}


// 盾创设备登录 msgType: 110
// client -> server
type Login struct {
	*MsgType
	Imsi				string			`json:"imsi"` 			// SIM卡的IMSI信息
	Ccid				string 			`json:"ccid"`			// SIM卡的CCID信息
	Imei				string			`json:"imei"`			// IEMI信息
	CabSta				CabinetState 	`json:"cabSta"`			// 换电柜状态
	HardVersion			string			`json:"hardVersion"`	// 硬件版本
	SoftVersion			string			`json:"softVersion"`	// 软件版本
	DevId				string			`json:"devId"`			// 设备ID
	ProtocolVersion		string			`json:"protocolVersion"`	// 协议版本
	DevType				DeviceType		`json:"devType"`		// 设备类型
	TxnNo				string			`json:"txnNo"`			// 流水号
}

// 登录请求响应 msgType: 111
// server -> client
type LoginResp struct {
	*MsgType
	Result				int				`json:"result"`		// 响应结果  0:失败 1:成功
	Value				int				`json:"value"`		// 结果描述	1：成功 2： 登录超时 3：认证失败
	TxnNo				string			`json:"txnNo"`		// 流水号
	DevId				string			`json:devId`		// 设备ID
}

// 属性上报请求 msgType: 310
// client -> server
type DeviceInfo struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	TxnNo				string			`json:"txnNo"`		// 流水号
	IsFull				int				`json:"isFull"`		// 是否全量上报  0：增量  1：全量
	CabList				[]CabinetInfo	`json:"cabList"`	// 电柜属性列表
	BoxList				[]BoxInfo		`json:"boxList"`	// 仓属性列表
	BatList				[]BatteryInfo	`json:"batList"`	// 电池属性列表
}

// 电柜整机信息
type CabinetInfo struct {
	LocationSta			string 			`json:"locationSta"`		// GSM小区信息
	DBM					string			`json:"dBM"`				// GSM信号强度
	CabSta				string			`json:"cabSta"`				// 换电柜状态	换电、放电、取电动作
	CabVol				string			`json:"cabVol"`				// 换电柜总电压
	CabCur				string			`json:"cabCur"`				// 换电柜总电流
	EmKwh				string			`json:"emKwh"`				// 柜子总用电量 kwh
	CabT				string 			`json:"cabT"`				// 换电柜温度
	CabEnable			string 			`json:"cabEnable"`			// 柜体是否禁用 启用
	BatNum				string			`json:"batNum"`				// 在柜电池个数
	BatFullA			string			`json:"batFullA"`			// 48V可用电池个数
	BatFullB			string			`json:"batFullB"`			// 60V可用电池个数
	BatFullC			string			`json:"batFullC"`			// 72V可用电池个数
	CabFault			[]string		`json:"cabFault"`			// 换电柜故障(整机),000:无，多个:["01", "02"]
	CabAlarm			[]string		`json:"cabAlarm"`			// 换电柜告警(整机),000:无，多个:["01", "02"]
}

// 电柜电池仓信息
type BoxInfo struct {
	DoorId				string				`json:"doorId"`			// 柜门编号
	DoorSta				string				`json:"doorSta"`		// 柜门状态 0：关  1：开
	BoxSta				string				`json:"boxSta"`			// 柜体状态 1:电池正在充电 2：电池充满 5：异常 6 有电池未识别编号（避免把有电池的仓门当空仓处理） 7  待充电 0：无电池
	BoxEnable			string				`json:"boxEnable"`		// 仓是否禁用
	BoxT				string				`json:"boxT"`			// 仓温度
	BoxChgSta			string				`json:"boxChgSta"`		// 仓充电器状态
	BoxHeatSta			string				`json:"boxHeatSta"`		// 加热器状态
	BoxLockSta			string				`json:"boxLockSta"`		// 锁状态
	BoxFireSta			string				`json:"boxFireSta"`		// 灭火器状态
	LowpowerSta			string				`json:"lowpowerSta"`	// 12V状态
	BoxAlarm			string				`json:"boxAlarm"`		// 单仓告警 000：无
	BoxFault			[]string			`json:"boxFault"`		// 单仓故障 000：无
	BatteryId			[]string			`json:"batteryId"`		// 仓内电池编码
}

// 电池仓中电池信息
type BatteryInfo struct {
	DoorId				string				`json:"doorId"`			// 电池在柜内编号
	BatteryId			string				`json:"batteryId"`		// 电池编号
	Soc					string				`json:"soc"`			// 电池SOC
	Soh					string				`json:"soh"`			// 电池SOH
	BmsT				string				`json:"bmsT"`			// 电池BMS板温度
	BatT				string				`json:"batT"`			// 电池电芯温度
	EnvT				string				`json:"envT"`			// 电池环境温度
	BatchgTime			string				`json:"batchgTime"`		// 电池充电时长
	BatFulTime			string				`json:"batFulTime"`		// 电池距离满电时长
	TotalAH				string				`json:"totalAH"`		// 电池容量Ah
	BatSta				string				`json:"batSta"`			// 电池状态  0： 移动 1：静止 2：存储 3：休眠
	BatCtrl				string				`json:"batCtrl"`		// 电池控制 0：放电状态 1：充电状态 2：负载在位 3：充电在位状态 4：空载状态
	CellNum				string				`json:"cellNum"`		// 电芯数量
	BatVol				string				`json:"batVol"`			// 电池电压
	BatCycle			string				`json:"batCycle"`		// 电池循环次数
	ChgCur				string				`json:"chgCur"`			// 电池充电电流
	CellVol				[]string			`json:"cellCol"`		// 电芯电压（数组）
	BmsFault			[]string			`json:"bmsFault"`		// bms故障  000：无
	BmsAlarm			[]string			`json:"bmsAlarm"`		// bms告警  000：无
}

// 属性上报响应 msgType: 311
// server -> client
type DeviceInfoResp struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	Result				string			`json:"result"`		// 结果 1：成功 0：失败
	TxnNo				string			`json:"txnNo"`		// 流水号
}

// 告警上报请求 msgType: 410
// client -> server
type DeviceAlarm struct {
	*MsgType
	DevId				string				`json:"devId"`		// 设备Id
	TxnNo				string				`json:"txnNo"`		// 流水号
	AlarmList			[]AlarmEvent		`json:"alarmList"`	// 告警列表
}

// 告警具体信息
type AlarmEvent struct {
	Id					string				`json:"id"`					// 信号量ID
	AlarmTime			int					`json:"alarmTime"`			// 告警时间  格式：13位Unix时间戳
	AlarmDesc			CommandType			`json:"alarmDesc"`			// 告警事件描述
	AlarmFlag			int					`json:"alarmFlag"`			// 告警标识 1：开始 0：结束
	DoorId				int					`json:"doorId"`				// 仓门ID  对应信号量类别属于电池，柜门，字段必填，否则省略
	BatteryId			string				`json:"batteryId"`			// 电池设备ID， 对应信号量类别属于电池，必填，否则省略
	UserId				string				`json:"userId"`				// 用户ID，非必填
	EmptBatteryId		string				`json:"emptBatteryId"`		// 亏电电池编码  只在换电结束上报
	EmptDoorId			string				`json:"emptDoorId"`			// 亏电电池所在仓编号  只在换电结束上报
	EmptBatsoc			string				`json:"emptBatsoc"`			// 亏电电池soc  只在换电结束上报
	FullBatteryId		string				`json:"fullBatteryId"`		// 满电电池编码  只在换电结束上报
	FullDoorID			string				`json:"fullDoorID"`			// 满电电池所在仓编号  只在换电结束上报
	FullBatsoc			string				`json:"fullBatsoc"`			// 满电电池soc  只在换电结束上报
}

// 告警上报响应 msgType: 411
// server -> client
type DeviceAlarmResp struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	Result				string			`json:"result"`		// 结果 1：成功 0：失败
	TxnNo				string			`json:"txnNo"`		// 流水号
}

// 远程控制请求 msgType: 500
// server -> client
type RemoteCtrl struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	TxnNo				string			`json:"txnNo"`		// 流水号
	ParamList			[]CtrlParams	`json:"paramList"`	// 控制参数列表
}

// 远程控制参数
type CtrlParams struct {
	Id					string			`json:"id"`				// 信号量ID
	Value				string			`json:"value"`			// 参数值
	DoorId				int				`json:"doorId"`			// 柜门ID 如果控制的是柜门，字段必填，否则可以省略
	BatteryId			string			`json:"batteryId"`		// 电池设备ID 如果控制的是电池，必填，否则可以省略
	Voltage				string			`json:"voltage"`		// 电池电压 首放流程中的这个字段必填
	ScanBattery			int				`json:"scanBattery"`	// 电池是否支持通讯  扫换电柜自个字段必填 0、支持通讯 1、不支持通讯
	UserId				string			`json:"userId"`			// 用户ID
}

// 远程控制响应 msgType: 501
// client -> server
// Result: 0： 未知错误 1： 成功 2： 指令不支持 3： 前一个流程没结束 4： 当前没有空柜门 5： 当前没有满电电池
// 6： 当前没有对应电压等级电池 16： 换电柜正在升级 17： 换电柜暂停使用 40：指令超时
type RemoteCtrlResp struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	Result				string			`json:"result"`		// 结果 1：成功 0：失败
	TxnNo				string			`json:"txnNo"`		// 流水号
}

// 配置查询请求 msgType: 210
// server -> client
type ConfigQuery struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	TxnNo				string			`json:"txnNo"`		// 流水号
	ParamList			[]ConfigParams	`json:"paramList"`	// 查询参数列表
}

// 配置查询参数
type ConfigParams struct {
	Id					string			`json:"id"`				// 信号量ID
}

// 配置查询响应 msgType: 211
// client -> server
type ConfigQueryResp struct {
	*MsgType
	DevId				string			`json:"devId"`		// 设备ID
	TxnNo				string			`json:"txnNo"`		// 流水号
	Result				int				`json:"result"`		// 结果 0-失败 1-成功
	ResultList			[]ConfigRespResult	`json:"resultList"`	// 配置列表
}

// 配置查询响应结果
type ConfigRespResult struct {
	Id					string			`json:"id"`				// 信号量ID
	Value				string			`json:"value"`			// 配置值
}
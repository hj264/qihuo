package cons

const (
	EURUSD     = "eurusd"
	GBPUSD     = "gbpusd"
	USDJPY     = "usdjpy"
	AUDUSD     = "audusd"
	USDCHF     = "usdchf"
	USDCAD     = "usdcad"
	USDHKD     = "usdhkd"
	USDCNY     = "usdcny"
	USDCNH     = "usdcnh"
	USDKRW     = "usdkrw"
	NZDCAD     = "nzdcad"
	EURNZD     = "eurnzd"
	AUDJPY     = "audjpy"
	AUDNZD     = "audnzd"
	AUDCAD     = "audcad"
	CHFJPY     = "chfjpy"
	AUDCHF     = "audchf"
	GBPCAD     = "gbpcad"
	EURGBP     = "eurgbp"
	CADCHF     = "cadchf"
	GBPJPY     = "gbpjpy"
	EURJPY     = "eurjpy"
	EURCHF     = "eurchf"
	NZDJPY     = "nzdjpy"
	GBPCHF     = "gbpchf"
	NZDCHF     = "nzdchf"
	CADJPY     = "cadjpy"
	GBPAUD     = "gbpaud"
	EURAUD     = "euraud"
	GBPNZD     = "gbpnzd"
	EURCAD     = "eurcad"
	USDMXN     = "usdmxn"
	USDTRY     = "usdtry"
	USDZAR     = "usdzar"
	USDDKK     = "usddkk"
	USDNOK     = "usdnok"
	USDRUB     = "usdrub"
	USDSEK     = "usdsek"
	USDSGD     = "usdsgd"
	NZDUSD     = "nzdusd"
	AUDCNY     = "audcny"
	CNYHKD     = "cnyhkd"
	GBPCNY     = "gbpcny"
	EURCNY     = "eurcny"
	EURTRY     = "eurtry"
	CADCNY     = "cadcny"
	CNYJPY     = "cnyjpy"
	USDTWD     = "usdtwd"
	XAU        = "xau"
	XAG        = "xag"
	XPD        = "xpd"
	XAP        = "xap"
	COPPER     = "copper"
	WTICRUDE   = "wticrude"
	OIL        = "oil"
	NATURALGAS = "naturalgas"
	USWHEAT    = "uswheat"
	// 股票
	INTC   = "intc"
	MSFT   = "msft"
	ADEL   = "adel"
	APLH   = "aplh"
	AJ     = "aj"
	AH     = "ah"
	HK0001 = "hk0001"
	HK0002 = "hk0002"
	JP1332 = "jp1332"
	DH6434 = "dh6434"
	DH5276 = "dh5276"
	AMZN   = "amzn"
	AAPL   = "aapl"
	TSLA   = "tsla"
	GOOGL  = "googl"
	META   = "meta"
	LLY    = "lly"
	AVGO   = "avgo"
	JPM    = "jpm"
	NVDA   = "nvda"
	GBI    = "gbi"
)

// GetVariety 获取品种id
func GetVariety(name string) int64 {
	var varietyId int64
	switch name {
	// 外汇
	case EURUSD:
		varietyId = 1
	case GBPUSD:
		varietyId = 2
	case USDJPY:
		varietyId = 3
	case AUDUSD:
		varietyId = 5
	case USDCHF:
		varietyId = 4
	case USDCAD:
		varietyId = 7
	case USDHKD:
		varietyId = 155
	case USDCNY:
		varietyId = 2111
	case USDCNH:
		varietyId = 961728
	case USDKRW:
		varietyId = 650
	case NZDCAD:
		varietyId = 56
	case EURNZD:
		varietyId = 52
	case AUDJPY:
		varietyId = 49
	case AUDNZD:
		varietyId = 50
	case AUDCAD:
		varietyId = 47
	case CHFJPY:
		varietyId = 13
	case AUDCHF:
		varietyId = 48
	case GBPCAD:
		varietyId = 54
	case EURGBP:
		varietyId = 6
	case CADCHF:
		varietyId = 14
	case GBPJPY:
		varietyId = 11
	case EURJPY:
		varietyId = 9
	case EURCHF:
		varietyId = 10
	case NZDJPY:
		varietyId = 58
	case GBPCHF:
		varietyId = 12
	case NZDCHF:
		varietyId = 57
	case CADJPY:
		varietyId = 51
	case GBPAUD:
		varietyId = 53
	case EURAUD:
		varietyId = 15
	case GBPNZD:
		varietyId = 55
	case EURCAD:
		varietyId = 16
	case USDMXN:
		varietyId = 39
	case USDTRY:
		varietyId = 18
	case USDZAR:
		varietyId = 17
	case USDDKK:
		varietyId = 43
	case USDNOK:
		varietyId = 59
	case USDRUB:
		varietyId = 2186
	case USDSEK:
		varietyId = 41
	case USDSGD:
		varietyId = 42
	case NZDUSD:
		varietyId = 8
	case AUDCNY:
		varietyId = 1486
	case CNYHKD:
		varietyId = 1564
	case GBPCNY:
		varietyId = 1740
	case EURCNY:
		varietyId = 1623
	case EURTRY:
		varietyId = 66
	case CADCNY:
		varietyId = 1524
	case CNYJPY:
		varietyId = 1565
	case USDTWD:
		varietyId = 2206
	// 外汇END

	// 期货
	case XAU: //黄金
		varietyId = 68
	case XAG: //白银
		varietyId = 69
	case XPD: //钯
		varietyId = 1043108
	case XAP: //铂金
		varietyId = 1199886
	case COPPER: //铜
		varietyId = 8831
	case WTICRUDE: //WTI原油
		varietyId = 1043109
	case OIL: //伦敦布伦特原油
		varietyId = 8833
	case NATURALGAS: //天然气
		varietyId = 8862
	case USWHEAT: //美国小麦
		varietyId = 8917

	// 股票
	case INTC: //英特尔
		varietyId = 251
	case MSFT: // 微软
		varietyId = 252
	case ADEL: //阿达尼企业有限公司
		varietyId = 17984
	case APLH: //阿波罗医院企业有限公司
		varietyId = 18004
	case AJ: //AJ 塑料 PCL
		varietyId = 102393
	case AH: //美国亚太高科技股份有限公司
		varietyId = 102394
	case HK0001:
		varietyId = 8564
	case HK0002:
		varietyId = 8563
	case JP1332:
		varietyId = 952851
	case DH6434:
		varietyId = 1012787
	case DH5276:
		varietyId = 1012560
	case AMZN: //亚马逊-纳斯达克
		varietyId = 6435
	case AAPL: //苹果-纳斯达克
		varietyId = 6408
	case TSLA: //特斯拉-纳斯达克
		varietyId = 13994
	case GOOGL: //亚马逊-纳斯达克
		varietyId = 6369
	case META: //Meta-纳斯达克
		varietyId = 26490
	case LLY: ////礼来公司-纳斯达克
		varietyId = 273
	case AVGO: //博通公司-纳斯达克
		varietyId = 6623
	case JPM: //摩根大通-纳斯达克
		varietyId = 267
	case NVDA: //英伟达公司-纳斯达克
		varietyId = 6497
	case GBI: //高盛
		varietyId = 266
	default:
		varietyId = 0
	}

	return varietyId
}

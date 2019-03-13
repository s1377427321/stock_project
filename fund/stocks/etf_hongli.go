package stocks

import . "fund/common/structs"

//var Hongli_etf []*StockItem = []*StockItem{
//	&StockItem{
//		Code:    "601088",
//		Name:    "中国神华",
//		BuyNums: 3981075,
//	},
//	&StockItem{
//		Code:    "600104",
//		Name:    "上汽集团",
//		BuyNums: 2029019,
//	},
//	&StockItem{
//		Code:    "600398",
//		Name:    "海澜之家",
//		BuyNums: 4212900,
//	},
//	&StockItem{
//		Code:    "601939",
//		Name:    "建设银行",
//		BuyNums: 7996853,
//	},
//	&StockItem{
//		Code:    "600011",
//		Name:    "华能国际",
//		BuyNums: 7790895,
//	},
//	&StockItem{
//		Code:    "600741",
//		Name:    "华域汽车",
//		BuyNums: 2077229,
//	},
//	&StockItem{
//		Code:    "600066",
//		Name:    "宇通客车",
//		BuyNums: 2534422,
//	},
//	&StockItem{
//		Code:    "601288",
//		Name:    "农业银行",
//		BuyNums: 13667222,
//	},
//	&StockItem{
//		Code:    "601398",
//		Name:    "工商银行",
//		BuyNums: 8403404,
//	},
//	&StockItem{
//		Code:    "601006",
//		Name:    "大秦铁路",
//		BuyNums: 5340962,
//	},
//	&StockItem{
//		Code:    "601988",
//		Name:    "中国银行",
//		BuyNums: 12083161,
//	},
//	&StockItem{
//		Code:    "600900",
//		Name:    "长江电力",
//		BuyNums: 2593892,
//	},
//	&StockItem{
//		Code:    "601328",
//		Name:    "交通银行",
//		BuyNums: 7194681,
//	},
//	&StockItem{
//		Code:    "600028",
//		Name:    "中国石化",
//		BuyNums: 6231407,
//	},
//	&StockItem{
//		Code:    "600660",
//		Name:    "福耀玻璃",
//		BuyNums: 1558998,
//	},
//	&StockItem{
//		Code:    "600027",
//		Name:    "华电国际",
//		BuyNums: 10091659,
//	},
//	&StockItem{
//		Code:    "600177",
//		Name:    "雅戈尔",
//		BuyNums: 5024038,
//	},
//	&StockItem{
//		Code:    "600383",
//		Name:    "金地集团",
//		BuyNums: 3741835,
//	},
//	&StockItem{
//		Code:    "600023",
//		Name:    "浙能电力",
//		BuyNums: 8100818,
//	},
//	&StockItem{
//		Code:    "600325",
//		Name:    "华发股份",
//		BuyNums: 4653802,
//	},
//	&StockItem{
//		Code:    "600036",
//		Name:    "招商银行",
//		BuyNums: 1291357,
//	},
//	&StockItem{
//		Code:    "600886",
//		Name:    "国投电力",
//		BuyNums: 4490600,
//	},
//	&StockItem{
//		Code:    "601998",
//		Name:    "中信银行",
//		BuyNums: 5255900,
//	},
//	&StockItem{
//		Code:    "601818",
//		Name:    "光大银行",
//		BuyNums: 8758400,
//	},
//	&StockItem{
//		Code:    "600048",
//		Name:    "保利地产",
//		BuyNums: 2623997,
//	},
//	&StockItem{
//		Code:    "601166",
//		Name:    "兴业银行",
//		BuyNums: 2177768,
//	},
//	&StockItem{
//		Code:    "601158",
//		Name:    "重庆水务",
//		BuyNums: 5590484,
//	},
//	&StockItem{
//		Code:    "600642",
//		Name:    "申能股份",
//		BuyNums: 5956941,
//	},
//	&StockItem{
//		Code:    "600016",
//		Name:    "民生银行",
//		BuyNums: 4052200,
//	},
//	&StockItem{
//		Code:    "600033",
//		Name:    "福建高速",
//		BuyNums: 9249800,
//	},
//	&StockItem{
//		Code:    "600887",
//		Name:    "伊利股份",
//		BuyNums: 979300,
//	},
//	&StockItem{
//		Code:    "600674",
//		Name:    "川投能源",
//		BuyNums: 3117400,
//	},
//	&StockItem{
//		Code:    "600585",
//		Name:    "海螺水泥",
//		BuyNums: 802000,
//	},
//	&StockItem{
//		Code:    "600183",
//		Name:    "生益科技",
//		BuyNums: 2901194,
//	},
//	&StockItem{
//		Code:    "600795",
//		Name:    "国电电力",
//		BuyNums: 10013793,
//	},
//	&StockItem{
//		Code:    "603328",
//		Name:    "依顿电子",
//		BuyNums: 2386500,
//	},
//	&StockItem{
//		Code:    "601009",
//		Name:    "南京银行",
//		BuyNums: 3321255,
//	},
//	&StockItem{
//		Code:    "603001",
//		Name:    "奥康国际",
//		BuyNums: 2076321,
//	},
//	&StockItem{
//		Code:    "600312",
//		Name:    "平高电气",
//		BuyNums: 4313200,
//	},
//	&StockItem{
//		Code:    "600894",
//		Name:    "广日股份",
//		BuyNums: 4001483,
//	},
//	&StockItem{
//		Code:    "600269",
//		Name:    "赣粤高速",
//		BuyNums: 5715900,
//	},
//	&StockItem{
//		Code:    "600704",
//		Name:    "物产中大",
//		BuyNums: 4414119,
//	},
//	&StockItem{
//		Code:    "601668",
//		Name:    "中国建筑",
//		BuyNums: 4137199,
//	},
//	&StockItem{
//		Code:    "601601",
//		Name:    "中国太保",
//		BuyNums: 696624,
//	},
//	&StockItem{
//		Code:    "600018",
//		Name:    "上港集团",
//		BuyNums: 3595711,
//	},
//	&StockItem{
//		Code:    "600004",
//		Name:    "白云机场",
//		BuyNums: 1591328,
//	},
//	&StockItem{
//		Code:    "600015",
//		Name:    "华夏银行",
//		BuyNums: 2743588,
//	},
//	&StockItem{
//		Code:    "600583",
//		Name:    "海油工程",
//		BuyNums: 3881300,
//	},
//	&StockItem{
//		Code:    "600350",
//		Name:    "山东高速",
//		BuyNums: 5006800,
//	},
//	&StockItem{
//		Code:    "600376",
//		Name:    "首开股份",
//		BuyNums: 2645280,
//	},
//	//&StockItem{
//	//	Code:    "601066",
//	//	Name:    "中信建投",
//	//	BuyNums: 18690,
//	//},
//	//&StockItem{
//	//	Code:    "603587",
//	//	Name:    "地素时尚",
//	//	BuyNums: 2442,
//	//},
//	//&StockItem{
//	//	Code:    "601330",
//	//	Name:    "绿色动力",
//	//	BuyNums: 4971,
//	//},
//	//&StockItem{
//	//	Code:    "603650",
//	//	Name:    "彤程新材",
//	//	BuyNums: 2376,
//	//},
//	//&StockItem{
//	//	Code:    "603706",
//	//	Name:    "东方环宇",
//	//	BuyNums: 1765,
//	//},
//	//&StockItem{
//	//	Code:    "603105",
//	//	Name:    "芯能科技",
//	//	BuyNums: 3410,
//	//},
//}

var Hongli_etf []*StockItem = []*StockItem{
	&StockItem{
		Code:    "600664",
		Name:    "哈药股份",
		BuyNums: 26696700,
	},
	&StockItem{
		Code:    "601088",
		Name:    "中国神华",
		BuyNums: 5650075,
	},
	&StockItem{
		Code:    "600507",
		Name:    "方大特钢",
		BuyNums: 7337417,
	},
	&StockItem{
		Code:    "601566",
		Name:    "九牧王",
		BuyNums: 5318419,
	},
	&StockItem{
		Code:    "600104",
		Name:    "上汽集团",
		BuyNums: 2552719,
	},
	&StockItem{
		Code:    "603328",
		Name:    "依顿电子",
		BuyNums: 6736300,
	},
	&StockItem{
		Code:    "600873",
		Name:    "梅花生物",
		BuyNums: 14803600,
	},
	&StockItem{
		Code:    "600028",
		Name:    "中国石化",
		BuyNums: 11774307,
	},
	&StockItem{
		Code:    "600376",
		Name:    "首开股份",
		BuyNums: 8014980,
	},
	&StockItem{
		Code:    "600383",
		Name:    "金地集团",
		BuyNums: 5901635,
	},
}

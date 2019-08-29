package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"strconv"
	"net/http"
	. "email_notice/common"
	"github.com/astaxie/beego/logs"
	"time"
	"strings"
)

func AddNoticeStock(code string, heightPrice, lowPrice, money float64) {
	mx.Lock()
	var s *Stock
	var isNew = false
	if oldStock, ok := NoticeStockS[code]; ok {
		s = oldStock
	} else {
		s = &Stock{
			BuyMoney:       money,
			Code:           code,
			NoticeCallBack: NoticeEmail,
			NoticeLimit:    NoticeLimit,
		}
		NoticeStockS[code] = s
		isNew = true
	}

	logs.Info("---AddNoticeStock  ", code, heightPrice, lowPrice)
	if strings.Index(code, "0") == 0 {
		s.Url = fmt.Sprintf(mainUrl, "sz"+code)
	} else if strings.Index(code, "6") == 0 || strings.Index(code, "5") == 0 {
		s.Url = fmt.Sprintf(mainUrl, "sh"+code)
	}

	//s.Url = fmt.Sprintf(mainUrl, code)
	s.LowPrice = lowPrice
	s.HightPrice = heightPrice
	s.Count = 0
	mx.Unlock()
	if isNew {
		s.Start()
	}
}

func AddBuyStock(code string, price, money float64, copies int) {
	bmx.Lock()
	var s *BuyStock
	var isNew = false
	if oldStock, ok := BuyStocks[code]; ok {
		s = oldStock
	} else {
		s = &BuyStock{
			StockName:                  code,
			BuyPrice:                   price,
			AllMoney:                   money,
			NumberOfCopies:             copies,
			NumberOfCopiesPice:         make(map[float64]*BuyLimit, 0),
			OrderNumberOfCopiesPiceKey: make([]float64, 0),
			AddNoticeFunc:              AddNoticeStock,
			DeleteNoticeFunc:           DeleteNoticeStock,
		}
		BuyStocks[code] = s
		isNew = true
	}

	s.StockUrl = fmt.Sprintf(mainUrl, code)

	bmx.Unlock()
	if isNew {
		s.Start()
	}
}

func DeleteBuyStock(code string) bool {
	if old, ok := BuyStocks[code]; ok {
		mx.Lock()
		old.Close()
		delete(BuyStocks, code)
		mx.Unlock()
		return true
	} else {
		return false
	}
}

func DeleteNoticeStock(code string) bool {
	if old, ok := NoticeStockS[code]; ok {
		mx.Lock()
		old.Close()
		delete(NoticeStockS, code)
		mx.Unlock()
		return true
	} else {
		return false
	}
}

//localhost:4661/delete?code=sh511990
func deleteNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")

	isDelete := DeleteNoticeStock(code)

	if isDelete {
		return c.String(http.StatusOK, "Delete OK")
	} else {
		return c.String(http.StatusOK, "Delete Object Not Exist")
	}
}

//120.79.154.53:4661/add?code=sh511990&height=100.05&low=99.905&money=130000
//netstat -aon|findstr "40051"
//localhost:4661/add?code=sh511990&height=100.05&low=99&money=130000
//localhost:4661/add?code=sh511990&height=100.05&low=99.905&money=130000
//http://swjswj.vip:4661/add?code=sz000895&height=26.00&low=23.70&money=160000
func addNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")
	heightPrice, _ := strconv.ParseFloat(c.QueryParam("height"), 64)
	lowPrice, _ := strconv.ParseFloat(c.QueryParam("low"), 64)
	money, _ := strconv.ParseFloat(c.QueryParam("money"), 64)

	AddNoticeStock(code, heightPrice, lowPrice, money)

	return c.String(http.StatusOK, "Add OK")
}

//localhost:4661/addstock?code=sh600004&money=130000&copies=20&price=14.844
func addStock(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	money, _ := strconv.ParseFloat(c.QueryParam("money"), 64)
	copies, _ := strconv.Atoi(c.QueryParam("copies"))

	AddBuyStock(code, price, money, copies)

	return c.String(http.StatusOK, "Add Stock OK")
}

//localhost:4661/deletestock?code=sh600004
func deleteStock(c echo.Context) error {
	code := c.QueryParam("code")

	isSuccess := DeleteBuyStock(code)

	if isSuccess {
		return c.String(http.StatusOK, "Delete OK")
	} else {
		return c.String(http.StatusOK, "Delete Object Not Exist")
	}

}

func addStockStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	name := c.QueryParam("name")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)

	magnification, _ := strconv.Atoi(c.QueryParam("magnification"))

	InstanceStopWinLoseManage().Add(code, name, price, magnification)

	result := InstanceStopWinLoseManage().ShowItems(code)

	return c.String(http.StatusOK, result+"\n  add OK"+time.Now().Format("2006-01-02 15:04:05"))
	//return c.String(http.StatusOK, "add OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func deleteStockStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")

	result := InstanceStopWinLoseManage().ShowItems(code)

	InstanceStopWinLoseManage().DeleteStock(code)
	return c.String(http.StatusOK, result+"\n  deleteStockStopWinLose OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func addLowSuctionStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	buy_num, _ := strconv.Atoi(c.QueryParam("buy_num"))
	InstanceStopWinLoseManage().AddLowSuction(code, buy_num, price)

	result := InstanceStopWinLoseManage().ShowItems(code)
	return c.String(http.StatusOK, result+"\n addLowSuctionStopWinLose OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func reduceHighThrowStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	buy_num, _ := strconv.Atoi(c.QueryParam("buy_num"))
	InstanceStopWinLoseManage().ReduceHighThrow(code, buy_num, price)

	result := InstanceStopWinLoseManage().ShowItems(code)

	return c.String(http.StatusOK, result+"\n reduceHighThrowStopWinLose OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func showStocksStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	result := InstanceStopWinLoseManage().ShowItems(code)
	if result == "" {
		return c.String(http.StatusOK, "Nothing "+time.Now().Format("2006-01-02 15:04:05"))
	}

	return c.HTML(http.StatusOK, result)
}

func RunHttpServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))

	e.Static("/", "./assets")

	e.GET("/add", addNoticeStock)
	e.GET("/delete", deleteNoticeStock)
	e.GET("/addstock", addStock)
	e.GET("/deletestock", deleteStock)

	e.GET("/addswl", addStockStopWinLose)
	e.GET("/deleteswl", deleteStockStopWinLose)
	e.GET("/addlowsuctionswl", addLowSuctionStopWinLose)
	e.GET("/reducehighthrowswl", reduceHighThrowStopWinLose)
	e.GET("/showstocksswl", showStocksStopWinLose)
	//e.GET("/getstockbuyNum",getStockBuyNum)
	fmt.Println("RunHttpServer ----------------- ")
	err := e.Start(httpPort)
	if err != nil {
		panic(err.Error())
	}
}

var ttt = `
<!DOCTYPE HTML>
<html lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=gbk">
    <meta charset="gbk" />
    <title>JQuery JSONView</title>
    <link rel="stylesheet" href="jquery.jsonview.css" />
    <style type="text/css">
        #songResJson{
            border: 1px solid #8d8c8c;
            padding: 5px;
            white-space: pre-wrap;
            word-wrap: break-word;
        }
    </style>
</head>
<body>
    <div class="apiResult">
    <h2> json 数据</h2>
    <pre id="songResJson"></pre> 
    </div> 
 <script>
    var data = %s;
    $("#songResJson").text(JSON.stringify(data,null,2)); 
</script>

</body>
`

var tttt = `<!DOCTYPE HTML>
<html lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=gbk">
    <meta charset="gbk" />
    <title>JQuery JSONView</title>
    <link rel="stylesheet" href="jquery.jsonview.css" />
    <script type="text/javascript" src="jquery-1.6.2.min.js"></script>
    <script type="text/javascript" src="jquery.jsonview.js"></script>
    <script type="text/javascript">
		/* 模拟数据 */
        var json2 = {
            "id": 61891,
            "subLemmaId": 14022133,
            "newLemmaId": 1122445,
            "key": "中国",
            "desc": "世界四大文明古国之一",
            "title": "中国",
            "card": [
    {
        "key": "m7_nameC",
        "name": "中文名称",
        "value": [
        "中国"
      ],
        "format": [
        "中国"
      ]
    },
    {
        "key": "m7_nameE",
        "name": "英文名称",
        "value": [
        "China"
      ],
        "format": [
        "China"
      ]
    },
    {
        "key": "m7_nameOther",
        "name": "简称",
        "value": [
        "华、夏"
      ],
        "format": [
        "华、夏"
      ]
    },
    {
        "key": "m7_continent",
        "name": "所属洲",
        "value": [
        "亚洲"
      ],
        "format": [
        "亚洲"
      ]
    },
    {
        "key": "m7_country1",
        "name": "首都",
        "value": [
        "<a target=_blank href=\"/item/%E9%95%BF%E5%AE%89/31540\" data-lemmaid=\"31540\">长安</a>、<a target=_blank href=\"/item/%E6%B4%9B%E9%98%B3/125712\" data-lemmaid=\"125712\">洛阳</a>、<a target=_blank href=\"/item/%E5%8D%97%E4%BA%AC/23952\" data-lemmaid=\"23952\">南京</a>、<a target=_blank href=\"/item/%E5%8C%97%E4%BA%AC/128981\" data-lemmaid=\"128981\">北京</a>等"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E9%95%BF%E5%AE%89/31540\" data-lemmaid=\"31540\">长安</a>、<a target=_blank href=\"/item/%E6%B4%9B%E9%98%B3/125712\" data-lemmaid=\"125712\">洛阳</a>、<a target=_blank href=\"/item/%E5%8D%97%E4%BA%AC/23952\" data-lemmaid=\"23952\">南京</a>、<a target=_blank href=\"/item/%E5%8C%97%E4%BA%AC/128981\" data-lemmaid=\"128981\">北京</a>等"
      ]
    },
    {
        "key": "m7_country6",
        "name": "官方语言",
        "value": [
        "<a target=_blank href=\"/item/%E6%B1%89%E8%AF%AD\">汉语</a>"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E6%B1%89%E8%AF%AD\">汉语</a>"
      ]
    },
    {
        "key": "m7_country12",
        "name": "主要民族",
        "value": [
        "<a target=_blank href=\"/item/%E6%B1%89%E6%97%8F\">汉族</a>"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E6%B1%89%E6%97%8F\">汉族</a>"
      ]
    },
    {
        "key": "m7_country13",
        "name": "主要宗教",
        "value": [
        "<a target=_blank href=\"/item/%E5%84%92%E6%95%99\">儒教</a>、<a target=_blank href=\"/item/%E4%BD%9B%E6%95%99\">佛教</a>、<a target=_blank href=\"/item/%E9%81%93%E6%95%99\">道教</a>、<a target=_blank href=\"/item/%E4%BC%8A%E6%96%AF%E5%85%B0%E6%95%99\">伊斯兰教</a>"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E5%84%92%E6%95%99\">儒教</a>、<a target=_blank href=\"/item/%E4%BD%9B%E6%95%99\">佛教</a>、<a target=_blank href=\"/item/%E9%81%93%E6%95%99\">道教</a>、<a target=_blank href=\"/item/%E4%BC%8A%E6%96%AF%E5%85%B0%E6%95%99\">伊斯兰教</a>"
      ]
    },
    {
        "key": "m7_ext_0",
        "name": "通用文字",
        "value": [
        "<a target=_blank href=\"/item/%E6%B1%89%E5%AD%97\">汉字</a>"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E6%B1%89%E5%AD%97\">汉字</a>"
      ]
    },
    {
        "key": "m7_ext_1",
        "name": "代表人物",
        "value": [
        "<a target=_blank href=\"/item/%E7%82%8E%E5%B8%9D/17732\" data-lemmaid=\"17732\">炎帝</a>、<a target=_blank href=\"/item/%E9%BB%84%E5%B8%9D\">黄帝</a>、<a target=_blank href=\"/item/%E5%AD%94%E5%AD%90/1584\" data-lemmaid=\"1584\">孔子</a>、<a target=_blank href=\"/item/%E5%AD%99%E4%B8%AD%E5%B1%B1/128084\" data-lemmaid=\"128084\">孙中山</a>等"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E7%82%8E%E5%B8%9D/17732\" data-lemmaid=\"17732\">炎帝</a>、<a target=_blank href=\"/item/%E9%BB%84%E5%B8%9D\">黄帝</a>、<a target=_blank href=\"/item/%E5%AD%94%E5%AD%90/1584\" data-lemmaid=\"1584\">孔子</a>、<a target=_blank href=\"/item/%E5%AD%99%E4%B8%AD%E5%B1%B1/128084\" data-lemmaid=\"128084\">孙中山</a>等"
      ]
    },
    {
        "key": "m7_ext_2",
        "name": "代表事物",
        "value": [
        "<a target=_blank href=\"/item/%E9%95%BF%E5%9F%8E/14251\" data-lemmaid=\"14251\">长城</a>、<a target=_blank href=\"/item/%E6%95%85%E5%AE%AB/9326\" data-lemmaid=\"9326\">故宫</a>、<a target=_blank href=\"/item/%E9%BE%99/13027234\" data-lemmaid=\"13027234\">龙</a>、<a target=_blank href=\"/item/%E5%87%A4/2910854\" data-lemmaid=\"2910854\">凤</a>、<a target=_blank href=\"/item/%E4%B9%A6%E6%B3%95/177069\" data-lemmaid=\"177069\">书法</a>等"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E9%95%BF%E5%9F%8E/14251\" data-lemmaid=\"14251\">长城</a>、<a target=_blank href=\"/item/%E6%95%85%E5%AE%AB/9326\" data-lemmaid=\"9326\">故宫</a>、<a target=_blank href=\"/item/%E9%BE%99/13027234\" data-lemmaid=\"13027234\">龙</a>、<a target=_blank href=\"/item/%E5%87%A4/2910854\" data-lemmaid=\"2910854\">凤</a>、<a target=_blank href=\"/item/%E4%B9%A6%E6%B3%95/177069\" data-lemmaid=\"177069\">书法</a>等"
      ]
    },
    {
        "key": "m7_ext_3",
        "name": "文化思想",
        "value": [
        "<a target=_blank href=\"/item/%E8%AF%B8%E5%AD%90%E7%99%BE%E5%AE%B6/16808\" data-lemmaid=\"16808\">诸子百家</a>、<a target=_blank href=\"/item/%E6%B1%89%E6%9C%8D%E6%96%87%E5%8C%96\">汉服文化</a>、<a target=_blank href=\"/item/%E5%94%90%E8%AF%97%E5%AE%8B%E8%AF%8D/8584359\" data-lemmaid=\"8584359\">唐诗宋词</a>等"
      ],
        "format": [
        "<a target=_blank href=\"/item/%E8%AF%B8%E5%AD%90%E7%99%BE%E5%AE%B6/16808\" data-lemmaid=\"16808\">诸子百家</a>、<a target=_blank href=\"/item/%E6%B1%89%E6%9C%8D%E6%96%87%E5%8C%96\">汉服文化</a>、<a target=_blank href=\"/item/%E5%94%90%E8%AF%97%E5%AE%8B%E8%AF%8D/8584359\" data-lemmaid=\"8584359\">唐诗宋词</a>等"
      ]
    }
  ],
            "image": "http://f.hiphotos.baidu.com/baike/pic/item/37d12f2eb9389b50ba8e5dec8235e5dde7116e1f.jpg",
            "src": "37d12f2eb9389b50ba8e5dec8235e5dde7116e1f",
            "imageHeight": 606,
            "imageWidth": 770,
            "isSummaryPic": "y",
            "abstract": "中国，是以华夏文明为源泉、中华文化为基础，并以汉族为主体民族的多民族国家，通用汉语、汉字，汉族与少数民族被统称为“中华民族”，又自称为炎黄子孙、龙的传人。中国是世界四大文明古国之一，有着悠久的历史，距今约5000年前，以中原地区为中心开始出现聚落组织进而形成国家，后历经多次民族交融和朝代更迭，直至形成多民族国家的大一统局面。20世纪初辛亥革命后，君主政体退出历史舞台，共和政体建立。1949年中华人民共和国成立后，在中国大陆建立了人民代表大会制度的政体。中国疆域辽阔、民族众多，先秦时期的华夏族在中原地区繁衍生息，到了汉代通过文化交融使汉族正式成型，奠定了中国主体民族的基础。后又通过与周边民族的交融，逐步形成统一多民族国家的局面，而人口也不断攀升，宋代中国人口突破一亿，清朝时期人口突破四亿，到2005年中国人口已突破十三亿。中国文化渊远流长、博大精深、绚烂多彩，是东亚文化圈的文化宗主国，在世界...",
            "moduleIds": [
					137685135,
					137685136,
					137685137,
					137685138,
					137685139,
					137685140,
					137685141,
					137685142,
					137685143,
					137685144,
					137685145,
					137685146,
					137685147,
					137685148,
					137685149,
					137685150,
					137685151,
					137685152,
					137685153,
					137685154,
					137685155,
					137685156,
					137685157,
					137685158,
					137685159,
					137685160,
					137685161,
					137685162,
					137685163,
					137685164,
					137685165,
					137685166,
					137685167,
					137685168,
					137685169,
					137685170,
					137685171,
					137685172,
					137685173,
					137685174
				  ],
            "url": "http://baike.baidu.com/subview/61891/14022133.htm",
            "wapUrl": "http://wapbaike.baidu.com/item/%E4%B8%AD%E5%9B%BD/1122445",
            "hasOther": 1,
            "totalUrl": "http://baike.baidu.com/view/61891.htm",
            "catalog": [
    "<a href='http://baike.baidu.com/subview/61891/14022133.htm#1'>词义</a>",
    "<a href='http://baike.baidu.com/subview/61891/14022133.htm#2'>历史</a>",
    "<a href='http://baike.baidu.com/subview/61891/14022133.htm#3'>地理</a>",
    "<a href='http://baike.baidu.com/subview/61891/14022133.htm#4'>政治</a>"
  ],
            "wapCatalog": [
    "<a href='http://wapbaike.baidu.com/item/%E4%B8%AD%E5%9B%BD/1122445#1'>词义</a>",
    "<a href='http://wapbaike.baidu.com/item/%E4%B8%AD%E5%9B%BD/1122445#2'>历史</a>",
    "<a href='http://wapbaike.baidu.com/item/%E4%B8%AD%E5%9B%BD/1122445#3'>地理</a>",
    "<a href='http://wapbaike.baidu.com/item/%E4%B8%AD%E5%9B%BD/1122445#4'>政治</a>"
  ],
            "logo": "http://img.baidu.com/img/baike/logo-baike.gif",
            "copyrights": "以上内容来自百度百科平台，由百度百科网友创作。",
            "customImg": "",
            "redirect": []
        };
 
		/* 模拟数据 */
		// var json3 = {"id": {"first":1, "second":2, "third":3, "fourth": [4, 5, 6, 7], "fifth": 5}, "name": {"ZS": "张三", "LS": "李四"}};
 
		var flag = false;
 
		$(function() {
			/* JSONView第一个参数就是需要转换的json数据 */
			$("#json").JSONView(json2, {collapsed: false, nl2br: true, recursive_collapser: true });
		});
 
 
		function expand() {
			$("#json").JSONView(json2, {collapsed: false, nl2br: true, recursive_collapser: true });
			flag = false;
		};
 
		function collapse() {
			$("#json").JSONView(json2, {collapsed: true, nl2br: true, recursive_collapser: true });
			flag = true;
		};
 
		function tog() {
			if(flag == false) {
				collapse();
				flag = true;
			} else {
				expand();
				flag = false;
			}
		};
    </script>
</head>
<body>
    <h2>数据</h2>
    <button id="collapse-btn" οnclick="collapse()">
        折叠
    </button>
 
    <button id="expand-btn"  οnclick="expand()">
        展开
    </button>
    <button id="toggle-btn" οnclick="tog()">
        切换
    </button>
 
    <div id="json"></div>
    <hr />
 
</body>
`

var bbb = `<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <style type="text/css">
        #songResJson{
            border: 1px solid #8d8c8c;
            padding: 5px;
            white-space: pre-wrap;
            word-wrap: break-word;
        }
    </style>
</head>
<body>
    <div class="apiResult">
    <h2> json 数据</h2>
    <pre id="songResJson"></pre>  
</div>
<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.2.6/jquery.min.js" type="text/javascript"></script>
<script>
    var data = %s;
    $("#songResJson").text(JSON.stringify(data,null,2)); 
</script>
</body>
</html>`

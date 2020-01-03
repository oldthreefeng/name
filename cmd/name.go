package cmd

import (
	"crypto/tls"
	_ "fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	MaxIdleConnections int = 20
	RequestTimeout   int = 30
)

var (
	httpClient *http.Client
)

func init() {
	httpClient = createHttpClient()
}

func createHttpClient() *http.Client {
	//proxy
	proxy, _ := url.Parse("http://192.168.0.88:38080")
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	return client
}

func Name() {

	datas := make([]NameInfo, 10)
	Names := getWaterWords(Wuxing)
	c  := make(chan NameInfo, 1000)
	for _, name := range Names {
		go Meimingteng(httpClient, getName(string(name)), c)
	}

	for i := 1; i <= 25; i++ {
		a, ok := GetMess(c)
		if ok {
			datas = append(datas, a)
		}
	}

	sort.Sort(NameInfoSlice(datas))
	for _, data := range datas {
		appendFile(data.name + " " + strconv.Itoa(data.scope) + "\n")
	}
}

//获取名字，可以自己组合
func getName(water string) string {
	return "铭" + water
}

//获取水属性的汉字列表
func getWaterWords(wuxing string) (words string) {
	//水属性的字
	switch wuxing {
	case "J":
		words = strings.Replace(string(J), " ", "", -1)
	case "M":
		words = strings.Replace(string(M), " ", "", -1)
	case "S":
		words = strings.Replace(string(S), " ", "", -1)
	case "H":
		words = strings.Replace(string(H), " ", "", -1)
	case "T":
		words = strings.Replace(string(T), " ", "", -1)
	default:
		words = strings.Replace(string(G), " ", "", -1)
	}

	words = strings.Replace(words, "\n", "", -1)
	return words
}

func Meimingteng(client *http.Client, name string, c chan NameInfo) () {
	Time = "2019-12-21-12-30"
	var (
		TimeSplit = strings.Split(Time, "-")
		year      = TimeSplit[0]
		month     = TimeSplit[1]
		day       = TimeSplit[2]
		hour      = TimeSplit[3]
		minute    = TimeSplit[4]
	)
	data := url.Values{}
	//必须要有，反爬虫
	data.Set("__EVENTTARGET", "ctl00$ContentPlaceHolder1$InputBasicInfo1$btNext")
	data.Set("__EVENTARGUMENT", "")
	data.Set("__VIEWSTATE", "/wEPDwULLTEyNjU5OTUwOTBkGAEFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYeBTtjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYlNwZWNpZnlCaXJ0aGRheQU9Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJTcGVjaWZ5TGluQ2hhblFpbgU9Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJTcGVjaWZ5TGluQ2hhblFpbgU+Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJOb3RTcGVjaWZ5QmlydGhkYXkFPmN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJiTm90U3BlY2lmeUJpcnRoZGF5BTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYlNvbGFyBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYkx1bmFyBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYkx1bmFyBTdjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRjYklzTGVhcE1vbnRoBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRsYnROb25lBTFjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRsYnROb25lBTJjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYnRMdW5ZdQUyY3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkcmJ0THVuWXUFNGN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFNoaUppbmcFNGN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFNoaUppbmcFMWN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFBvZW0FMWN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidFBvZW0FMmN0bDAwJENvbnRlbnRQbGFjZUhvbGRlcjEkSW5wdXRCYXNpY0luZm8xJHJidElkaW9tBTJjdGwwMCRDb250ZW50UGxhY2VIb2xkZXIxJElucHV0QmFzaWNJbmZvMSRyYnRJZGlvbQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMgU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkMwU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNgU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkNwU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOAU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOQU6Y3RsMDAkQ29udGVudFBsYWNlSG9sZGVyMSRJbnB1dEJhc2ljSW5mbzEkY2JsUGVyc29uYWxpdHkkOc+RBppac3/CXaY8AJwjzwaxBvxk")
	data.Set("__VIEWSTATEGENERATOR", "9F5AD4C7")
	data.Set("__EVENTVALIDATION", "/wEWuwICkof8EwK19uvYAgKlkY7lBAKBmOYKAoCY5goCqcPP0AcCqMPP0AcCq8PP0AcCnd3BDAKVuaPyDgKV99WYCAL9kPD0BgKmsv/0AQKL2c+pBQKMlpPbCAKMloemAQKMluuBCgKMlt/sAgKMlsO3CwKMlreTBAKMlpv+DAKMls+WCgKMlrPyAgLhr8HqBQLhr7W2DgLhr5mRBwLhr438DwLhr/HHCALhr+WiAQLhr8mNCgLhr73pAgLhr+EBAuGv1ewIAvq448ULAvq416AEAvq4u4wNAvq4r9cFAvq4k7IOAvq4h50HAvq46/gPAvq438MIAvq4g/wFAvq498cOAt/RhLABAt/R6JsKAt/R3OYCAt/RwMELAt/RtK0EAt/RmIgNAt/RjNMFAt/R8L4OAt/RpNcLAt/RiLIEArDrpqsHArDrivYPArDr/tEIArDr4rwBArDr1ocKArDruuMCArDrrs4LArDrkqkEArDrxsEBArDrqq0KApWEuIYNApWErOEFApWEkMwOApWEhJcHApWE6PIPApWE3N0IApWEwLgBApWEtIQKApWE2LwHApWEzAcC7p3a8AIC7p3O2wsC7p2ypwQC7p2mgg0C7p2K7QUC7p3+yA4C7p3ikwcC7p3W/g8C7p36lw0C7p3u8gUCw7b86wgCw7bgtgECw7bUkQoCw7a4/QICw7as2AsCw7aQowQCw7aEjg0Cw7bo6QUCw7acggMCw7aA7QsC9Ny8qgEC9Nyg9QkC9NyU0AIC9Nz4uwsC9NzshgQC9NzQ4QwC9NzEzAUC9NyoqA4C9NzcwAsC9NzAqwQCyfXehAcCyfXC7w8CyfW2ywgCyfWalgECyfWO8QkCyfXy3AICyfXmpwsCyfXKggQCyfX+uwECyfXihgoCjZaL8A8CjZb/2wgCjZbjpgECjZbXgQoCjZa77QICjZavyAsCjZaTkwQCjZaH/gwCjZarlwoCjZaf8gIC5q+t6wUC5q+Rtg4C5q+FkQcC5q/p/A8C5q/dxwgC5q/BogEC5q+1jgoC5q+Z6QIC5q/NAQLmr7HtCAL7uM/FCwL7uLOhBAL7uKeMDQLyr4XtDALzr4XtDALwr4XtDALxr4XtDAL2r4XtDAL3r4XtDAL0r4XtDALlr4XtDALqr4XtDALyr8XuDALyr8nuDALyr83uDAKVwsetDgKUwsetDgKXwsetDgKWwsetDgKRwsetDgKQwsetDgKTwsetDgKCwsetDgKNwsetDgKVwoeuDgKVwouuDgKVwo+uDgKVwrOuDgKVwreuDgKVwruuDgKVwr+uDgKVwqOuDgKVwuetDgKVwuutDgKUwoeuDgKUwouuDgKUwo+uDgKUwrOuDgKUwreuDgKUwruuDgKUwr+uDgKUwqOuDgKUwuetDgKUwuutDgKXwoeuDgKXwouuDgKArtmFCAKTrpWGCAKMrpWGCAKNrpWGCAKOrpWGCAKPrpWGCAKIrpWGCAKJrpWGCAKKrpWGCAKbrpWGCAKUrpWGCAKMrtWFCAKMrtmFCAKMrt2FCAKMruGFCAKMruWFCAKMrumFCAKMru2FCAKMrvGFCAKMrrWGCAKMrrmGCAKNrtWFCAKNrtmFCAKNrt2FCAKNruGFCALCsOLeDALdsOLeDALcsOLeDALfsOLeDALesOLeDALZsOLeDALYsOLeDALbsOLeDALKsOLeDALFsOLeDALdsKLdDALdsK7dDALdsKrdDALdsJbdDALdsJLdDALdsJ7dDALdsJrdDALdsIbdDALdsMLeDALdsM7eDALcsKLdDALcsK7dDALcsKrdDALcsJbdDALcsJLdDALcsJ7dDALcsJrdDALcsIbdDALcsMLeDALcsM7eDALfsKLdDALfsK7dDALfsKrdDALfsJbdDALfsJLdDALfsJ7dDALfsJrdDALfsIbdDALfsMLeDALfsM7eDALesKLdDALesK7dDALesKrdDALesJbdDALesJLdDALesJ7dDALesJrdDALesIbdDALesMLeDALesM7eDALZsKLdDALZsK7dDALZsKrdDALZsJbdDALZsJLdDALZsJ7dDALZsJrdDALZsIbdDALZsMLeDALZsM7eDALq3rerBALCuMjbDwKi2vn0BQK1gMHLAgKy5bvtBQK236G8AQLV9/naAgKgqtmeCgLVnOXeBwL62evxBgL22aPyBgL32aPyBgL02aPyBgL12aPyBgLy2aPyBgLz2aPyBgLw2aPyBgLh2aPyBgLu2aPyBgL22ePxBgL22e/xBgL22evxBgL22dfxBgL22dPxBgL22dvxBgL22cfxBgL62e/xBgL/l8bNAwL/l8rNAwL/l77NAwL/l8LNAwL/l9bNAwL/l9rNAwL/l87NAwL/l9LNAwL/l6bNAwL/l6rNAwL9lKKdCgKWk7qbCgL1g6aRDwLXw/fkCALvyL/mDALev/OZDgLL5ZPJCAKi/N2EAQKD5eu3CwKgl5eyCgKeoZaXAwKd8qmzBALwr6/EDwLS6rKCBxK3OW282UF90M07q6v0IUMlKeiv")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbXing", LastName)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbMingWords", name)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlGenders", Gender)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$SPECIFY_BIRHDAY", "rbSpecifyBirthday")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$CalendarType", "rbSolar")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlYear", year)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlMonth", month)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlDay", day)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlHour", hour)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlMinute", minute)
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbCountry", "中国")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbProvince", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbCity", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbOtherHopes", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$ddlCareer", "-2")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbFather", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbMother", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbAvoidWords", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$tbAvoidSimpParts", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbUserName", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbPwd", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$tbVCode", "")
	data.Set("ctl00$ContentPlaceHolder1$InputBasicInfo1$LoginAnywhere1$loginParam", "2")

	req, _ := http.NewRequest("POST", "https://www.meimingteng.com/Naming/Default.aspx?Tag=4", strings.NewReader(data.Encode()))

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,fr;q=0.7,tr;q=0.6,zh-TW;q=0.5")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "Params=%26Xing%3d%e6%9d%8e%26Gender%3d1%26Year%3d2018%26Month%3d3%26Day%3d13%26Hour%3d13%26Minute%3d40%26IsSolarCalendar%3d1%26IsLeapMonth%3d0%26NameType%3d2%26ReiterativeLocution%3d0%26Location%3d%e4%b8%ad%e5%9b%bd++%26Career%3d-2%26Personality%3d%26Father%3d%26Mother%3d%26SpecifiedName%3d%26SpecifiedNameIndex%3d0%26OtherHopes%3d%26AvoidWords%3d%26AvoidSimpParts%3d%26SpecifiedMing1SimpParts%3d%26SpecifiedMing2SimpParts%3d%26SpecifiedMing1Stroke%3d%26SpecifiedMing2Stroke%3d%26Tag%3d4%7c2%26LinChanQi%3dFalse%26NamingByCategoryCategoryID%3d-1%26SM1S%3d%26SM2S%3d%26SM1T%3d%26SM2T%3d%26SM1M%3d%26SM2M%3d%26RN%3d%26SpecifiedMing1Spell%3d%26SpecifiedMing2Spell%3d%26SM1Y%3d%26SM2Y%3d%26FA%3d%e4%b8%8b%e5%8d%88++++%e6%98%a5%e5%ad%a3++%e6%ad%a3%e6%9c%88%26LOCATION_COUNTY%3d%e4%b8%ad%e5%9b%bd%26LOCATION_PROVINCE%3d%26LOCATION_CITY%3d%26MING_WORDS%3d%e9%93%ad%e6%ba%a5; ASP.NET_SessionId=tgy22x55l3jdvk2ci4gjpxru; mmtsuser=1; ckcookie=chcookie; HELLO_USER=1; Hm_lvt_637e96da78d1c6c8f8a218c811dea5fb=1521614202; Hm_lpvt_637e96da78d1c6c8f8a218c811dea5fb=1521614588; qrcode=1")
	req.Header.Set("Host", "www.meimingteng.com")
	req.Header.Set("Origin", "https://www.meimingteng.com")
	req.Header.Set("Referer", "https://www.meimingteng.com/Naming/Default.aspx?Tag=4")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 8_10_2) AppleWebKit/456.20 (KHTML, like Gecko) Chrome/45.0.14.25 Safari/567.24")

	//反安全策略，随机IP
	ip := RandomIp()
	req.Header.Set("CLIENT-IP", ip)
	req.Header.Set("X-FORWARDED-FOR", ip)

	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return
	}
	//doc,_ := goquery.NewDocumentFromReader(resp.Body)
	doc, _ := goquery.NewDocumentFromResponse(resp)

	//获取文化、五行、生肖、五格分数
	wenhua := doc.Find("#ctl00_ContentPlaceHolder1_ShowNameDetails1_lbNameScore > font:nth-child(5) > b").Text()
	wuxing := doc.Find("#bdAppSummDiv > table:nth-child(7) > tbody > tr > td > font:nth-child(5) > b").Text()
	shengxiao := doc.Find("#bdAppSummDiv > table:nth-child(7) > tbody > tr > td > font:nth-child(9) > b").Text()
	wuge := doc.Find("#bdAppSummDiv > table:nth-child(7) > tbody > tr > td > font:nth-child(13) > b").Text()

	//fmt.Println(wenhua, wuxing, shengxiao, wuge)
	wenhuaI, _ := strconv.Atoi(wenhua)
	wuxingI, _ := strconv.Atoi(wuxing)
	shengxiaoI, _ := strconv.Atoi(shengxiao)
	wugeI, _ := strconv.Atoi(wuge)
	c <- NameInfo{name: LastName + name, scope: (wenhuaI + wuxingI + shengxiaoI + wugeI) / 4}
}

type NameInfo struct {
	name  string
	scope int
}

//排序使用
type NameInfoSlice []NameInfo

func (p NameInfoSlice) Len() int           { return len(p) }
func (p NameInfoSlice) Less(i, j int) bool { return p[i].scope > p[j].scope }
func (p NameInfoSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func appendFile(data string) {
	file, _ := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.WriteString(data)
	defer file.Close()
}

func GetMess(mess chan NameInfo) (v NameInfo, b bool) {
	for {
		select {
		case v = <-mess:
			return v, true
		case <-time.After(time.Second * 1):
			return
		}
	}
}

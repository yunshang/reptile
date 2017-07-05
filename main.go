package main

import (
	"github.com/PuerkitoBio/goquery"
  "github.com/gin-gonic/gin"
  "fmt"
  "strings"
  "strconv"
	"log"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Qb struct {
	Id      int    `json:"id"`
	Brand_id      string    `json:"brand_id"`
	Content string `json:"content"`
}

type Seriesb struct {
	Series_id      string    `json:"series_id"`
	Series_name string `json:"series_name"`
	Series_group_name string `json:"series_group_name"`
}

ifewfwef
type Modelb struct {
  Discharge_standard    string `json:"discharge_standard"`
  Gear_type    string `json:"Gear_type"`
  Liter    string `json:"Liter"`
  Liter_type    string `json:"Liter_type"`
	Model_id      string    `json:"model_id"`
	Model_name string `json:"model_name"`
	Model_price string `json:"model_price"`
	Model_year string `json:"model_year"`
}

type City struct {
  City_code    string `json:"city_code"`
  City_id    string `json:"city_id"`
  City_name    string `json:"city_name"`
	Display_order      string    `json:"display_order"`
	Enabled string `json:"enabled"`
	Hot_level string `json:"hot_level"`
	Initial string `json:"initial"`
	Lat string `json:"lat"`
	Lot string `json:"lot"`
	Prov_id string `json:"prov_id"`
	Sell_enabled string `json:"sell_enabled"`
	Sld string `json:"sld"`
	Zone_id string `json:"zone_id"`
}

type Pricehistory struct {
  Date  string `json:"date"`
  Eval_price  string `json:"eval_price"`
}

type Pricefuture struct {
  register_year  string `json:"register_year"`
  eval_price  string `json:"eval_price"`
}

type Provprices struct {
  Bid  string `json:"bid"`
  Body_type  string `json:"body_type"`
  Brand_name  string `json:"brand_name"`
  Discharge_standard  string `json:"discharge_standard"`
  Door_number  string `json:"Door_number"`
  Drive_name  string `json:"drive_name"`
  Engine_power  string `json:"engine_power"`
  Gear_type  string `json:"gear_type"`
  // Highlight_config  []string `json:"highlight_config"`
  Id  string `json:"id"`
  Level  string `json:"level"`
  Level_id  string `json:"level_id"`
  Liter  string `json:"Liter"`
  Maker_name  string `json:"maker_name"`
  Maker_type  string `json:"maker_type"`
  Market_date  string `json:"market_date"`
  Max_reg_year  string `json:"max_reg_year"`
  Min_reg_year  string `json:"min_reg_year"`
  Model_status  string `json:"model_status"`
  Name  string `json:"name"`
  Price  string `json:"price"`
  Short_name  string `json:"short_name"`
  Sid  string `json:"sid"`
  Sname  string `json:"sname"`
  Ssname  string `json:"ssname"`
  Star  string `json:"star"`
  Stop_make_year  string `json:"stop_make_year"`
  Year  string `json:"year"`
  Liter_turbo  string `json:"liter_turbo"`
}

type Pinggu struct {
  Title    string `json:"title"`
  // Address    []string `json:"address"`
  Address    string `json:"address"`
  Card_time    string `json:"card_time"`
	Mileage      string    `json:"mileage"`
	Transmission string `json:"transmission"`
	Displacement string `json:"displacement"`
	Emission_standards string `json:"emission_standards"`
	Price string `json:"price"`
}

type Prices struct {
	Provid      string    `json:"provid"`
	Price string `json:"price"`
	Provname string `json:"provname"`
}

func main() {
  // getFutureHistory()
  // getPriceHistory()
  fmt.Println(allProvPrices())
  r := gin.Default()
  r.LoadHTMLGlob("public/*")
	r.GET("/", Index)
  r.GET("/series/:seriesID", Findseries)
  r.GET("/models/:modelsID", Findmodels)
	r.Run()
}


func Index(c *gin.Context) {
	city_result := getCity()
	brand_result := getBrand()
  series_result := getseriesBrand(1)
  model_result := getmodelBrand(13)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"cities":    city_result,
		"brands":    brand_result,
		"series":    series_result,
		"models":    model_result,
	})
	return
}

func Findseries(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
  id := c.Param("seriesID")
  seriesID, _ := strconv.Atoi(id)
  series_result := getseriesBrand(seriesID)
  c.JSON(200, gin.H{
    "data":  series_result,
  })
  return 
}

func Findmodels(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
  id := c.Param("modelsID")
  modelsID, _ := strconv.Atoi(id)
  model_result := getmodelBrand(modelsID)
  c.JSON(200, gin.H{
    "data":  model_result,
  })
  return 
}

func getBrand() []Qb {
	qburl := "https://www.che300.com/pinggu"

	doc, err := goquery.NewDocument(qburl)
	if err != nil {
		log.Fatal(err)
	}

	var qb []Qb
	doc.Find(".ucarselecttype_pinpaibottom_ul .list_1").Each(func(i int, s *goquery.Selection) {
    stringid, _ := s.Attr("id")
    id, _ := strconv.Atoi(stringid)
    content := s.Text()
    brand_id, _ := s.Attr("rel")
    qb = append(qb, Qb{Id: id, Brand_id: brand_id, Content: content})
	})
	return qb
}

func getseriesBrand(p int) []Seriesb {
  var props []Seriesb
	qburl := "https://ssl-meta.che300.com/meta/series/series_brand" + strconv.Itoa(p) + ".json"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := []byte(body)

    er := json.Unmarshal(bodystr, &props)
    if er != nil {
      panic(er)
    } else {
      // fmt.Println(props)
    }
  }
	return props
}

func getmodelBrand(p int) []Modelb {
  var props []Modelb
	qburl := "https://ssl-meta.che300.com/meta/model/model_series" + strconv.Itoa(p) + ".json"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := []byte(body)

    er := json.Unmarshal(bodystr, &props)
    if er != nil {
      panic(er)
    } else {
      // fmt.Println(props)
    }
  }

	return props
}

func getCity() []City {
  var props []City
	qburl := "https://ssl-meta.che300.com/location/all_city.json"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body))
    bodystr := []byte(body)

    er := json.Unmarshal(bodystr, &props)
    if er != nil {
      panic(er)
    } else {
      // fmt.Println(props)
    }
  }

	return props
}


func getFutureHistory() []Pricefuture {
  var props []Pricefuture
	qburl := "https://dingjia.che300.com/app/EvalResult/getFuturePriceTrend?callback=jQuery183048247267398977844_1499239194771&provId=4&cityId=4&seriesId=77&modelId=30304&year=2017&month=6&mile=10&_=1499239195245"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    split_string := strings.Split(strings.Split(string(body),"([{")[1],"}])")[0]
    sss  := strings.Split(split_string,"},{")
    for _, element := range sss {
      props = append(props, Pricefuture{register_year: strings.Split(strings.Split(element, ":")[1],",")[0], eval_price: strings.Split(element, ":")[2]})
    }
    fmt.Println(props)
  }

	return props
}

func getPriceHistory() []Pricehistory {
  var props []Pricehistory
	qburl := "https://dingjia.che300.com/app/EvalResult/getPriceHistory?callback=jQuery183048247267398977844_1499239194770&provId=4&cityId=4&modelId=30304&regYear=2017&regMonth=6&mileAge=10&price=20.246191981464&_=1499239195244"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    split_string := strings.Split(strings.Split(string(body),":[{")[1],"}]})")[0]
    sss  := strings.Split(split_string,"},{")
    for _, element := range sss {
      props = append(props, Pricehistory{Eval_price: strings.Split(lstrings.Split(element, ":")[1],",")[0], Date: strings.Split(element, ":")[2]})
    }
    fmt.Println(props)
  }

	return props
}

func allProvPrices() ([]Provprices, []Prices){
  var props []Provprices
  var props2 []Prices
	qburl := "https://dingjia.che300.com/app/EvalResult/allProvPrices?callback=jQuery183048247267398977844_1499239194773&brand=6&series=77&model=30304&regDate=2017-6&mile=10&_=1499239195252"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    str := strings.Split((strings.Split(string(body),":{")[2]),"},\"prices\":[{")
    str2 := strings.Split(str[0],",")
    props = append(props, Provprices{Id: strings.Split(str2[0], ":")[1],
                                     Drive_name: strings.Split(str2[1], ":")[1],
                                     Model_status: strings.Split(str2[2], ":")[1],
                                     Name: strings.Split(str2[21], ":")[1],
                                     Short_name: strings.Split(str2[22], ":")[1],
                                     Market_date: strings.Split(str2[23], ":")[1],
                                     Stop_make_year: strings.Split(str2[24], ":")[1],
                                     Level_id: strings.Split(str2[25], ":")[1],
                                     Level: strings.Split(str2[26], ":")[1],
                                     Sname: strings.Split(str2[27], ":")[1],
                                     Ssname: strings.Split(str2[28], ":")[1],
                                     Maker_name: strings.Split(str2[29], ":")[1],
                                     Maker_type: strings.Split(str2[30], ":")[1],
                                     Star: strings.Split(str2[31], ":")[1],
                                     Sid: strings.Split(str2[32], ":")[1],
                                     Bid: strings.Split(str2[33], ":")[1],
                                     Brand_name: strings.Split(str2[34], ":")[1],
                                     Year: strings.Split(str2[35], ":")[1],
                                     Price: strings.Split(str2[36], ":")[1],
                                     Discharge_standard: strings.Split(str2[37], ":")[1],
                                     Gear_type: strings.Split(str2[38], ":")[1],
                                     Liter: strings.Split(str2[39], ":")[1],
                                     Liter_turbo: strings.Split(str2[40], ":")[1],
                                     Door_number: strings.Split(str2[41], ":")[1],
                                     Body_type: strings.Split(str2[42], ":")[1],
                                     Min_reg_year: strings.Split(str2[43], ":")[1],
                                     Max_reg_year: strings.Split(str2[44], ":")[1],
                                     Engine_power: strings.Split(str2[45], ":")[1],
                                   })
    str_price := strings.Split(str[1],"}]}})")
    strr := strings.Split(str_price[0],"},{")
    for _, element := range strr{
      props2 = append(props2, Prices{Price: strings.Split(strings.Split(element,",")[2],":")[1],
                            Provid: strings.Split(strings.Split(element,",")[0],":")[1],
                            Provname: strings.Split(strings.Split(element,",")[1],":")[1],
                          })
    }
  }

	return props,props2

}

func getPingguinfo() []Pinggu {
	qburl := "https://www.che300.com/pinggu/v10c40m30616r2017-6g20"

	doc, err := goquery.NewDocument(qburl)
	if err != nil {
		log.Fatal(err)
	}

	var pinggu []Pinggu
	doc.Find(".result").Each(func(i int, s *goquery.Selection) {
     // title := s.Find(".rh-wrap h1").Text()
     price, _ := s.Find("#price").Attr("value")
     fmt.Println(price)
    // pinggu = append(pinggu, Pinggu{})
	})
	return pinggu
}

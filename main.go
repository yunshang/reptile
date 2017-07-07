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

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

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
  Register_year  string `json:"register_year"`
  Eval_price  string `json:"eval_price"`
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
  r := gin.Default()
  r.LoadHTMLGlob("public/*")
	r.GET("/", Index)
	r.POST("/info", Info)
  r.GET("/series/:seriesID", Findseries)
  r.GET("/models/:modelsID", Findmodels)
	r.Run()
}

func Info(c *gin.Context) {
  cityid := c.PostForm("city_id")
  provid := c.PostForm("prov_id")
  brandid := c.PostForm("brand_id")
  seriesid := c.PostForm("series_id")
  modalid := c.PostForm("modal_id")
  date := c.PostForm("date")
  mile := c.PostForm("mile")
  year := strings.Split(date,"-")[0]
  month := strings.Split(date,"-")[1]
  ci,_ :=  strconv.Atoi(cityid)
  pr,_ :=  strconv.Atoi(provid)
  br,_ :=  strconv.Atoi(brandid)
  se,_ :=  strconv.Atoi(seriesid)
  mo,_ :=  strconv.Atoi(modalid)
  mi,_ :=  strconv.Atoi(mile)
  ye,_ :=  strconv.Atoi(year)
  mot,_ :=  strconv.Atoi(month)
  future_result := getFutureHistory(ci,pr,se,mo,ye,mot,mi)
  history_prices := getPriceHistory(pr,ci,mo,ye,mot,mi,future_result[0].Eval_price)
  modelInfo, prices_all := allProvPrices(br,se,mo,date,mi)
	c.HTML(http.StatusOK, "info.html", gin.H{
    "future_result":  future_result,
		"history_prices":    history_prices,
		"modalinfo":    modelInfo,
		"prices_all":    prices_all,
	})
}

func Index(c *gin.Context) {
	city_result := getCity()
	brand_result := getBrand()
	fmt.Println(brand_result)
	fmt.Println(city_result)
 //  series_result := getseriesBrand(1)
 //  model_result := getmodelBrand(13)
	// c.HTML(http.StatusOK, "index.html", gin.H{
 //    "proid":     city_result[0].Prov_id,
	// 	"cities":    city_result,
	// 	"brands":    brand_result,
	// 	"series":    series_result,
	// 	"models":    model_result,
	// })
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

 // db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
 // defer db.Close()
 // if err != nil{
 //  log.Fatalln(err)
 // }
	//
 // db.SetMaxIdleConns(20)
 // db.SetMaxOpenConns(20)
	//

  var qb []Qb
	doc.Find(".ucarselecttype_pinpaibottom_ul .list_1").Each(func(i int, s *goquery.Selection) {
    stringid, _ := s.Attr("id")
    id, _ := strconv.Atoi(stringid)
    content := s.Text()
    brand_id, _ := s.Attr("rel")
    qb = append(qb, Qb{Id: id, Brand_id: brand_id, Content: content})
		// rs, err := db.Exec("INSERT INTO che300_brand(bid, name) VALUES (?, ?)", id, content)
  //   fmt.Println(rs)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
	})
	// for _, b := range qb {
 //     getseriesBrand(b.Id)
	// }
	return qb
}

func getseriesBrand(p int) []Seriesb {
  var props []Seriesb
	qburl := "https://ssl-meta.che300.com/meta/series/series_brand" + strconv.Itoa(p) + ".json"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	// defer db.Close()
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	// db.SetMaxIdleConns(20)
	// db.SetMaxOpenConns(20)
	//

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
	// for _, b := range props {
	// 	rs, err := db.Exec("INSERT INTO che300_series(uid, name, bid) VALUES (?, ?, ?)", b.Series_id, b.Series_name, p)
	// 	modelsID, _ := strconv.Atoi(b.Series_id)
	// 	getmodelBrand(modelsID)
	// 	fmt.Println(rs)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	return props
}

func getmodelBrand(p int) []Modelb {
  var props []Modelb
	qburl := "https://ssl-meta.che300.com/meta/model/model_series" + strconv.Itoa(p) + ".json"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	// defer db.Close()
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	// db.SetMaxIdleConns(20)
	// db.SetMaxOpenConns(20)
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
	// for _, b := range props {
	// 	rs, err := db.Exec("INSERT INTO che300_models(uid, name, sid) VALUES (?, ?, ?)", b.Model_id, b.Model_name, p)
	// 	fmt.Println(rs)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

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
	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	// defer db.Close()
	// if err != nil{
	// 	log.Fatalln(err)
	// }
 // db.SetMaxIdleConns(20)
 // db.SetMaxOpenConns(20)
	// for _, b := range props {
	// 	rs, err := db.Exec("INSERT INTO che300_city(uid, name, pid) VALUES (?, ?, ?)", b.City_id, b.City_name, b.Prov_id)
	// 	fmt.Println(rs)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	return props
}


func getFutureHistory(provid int,cityid int,seriesid int,modelid int,year int,month int,mile int) []Pricefuture {
  var props []Pricefuture
	qburl := "https://dingjia.che300.com/app/EvalResult/getFuturePriceTrend?callback=jQuery183048247267398977844_1499239194771&provId=" + strconv.Itoa(provid) + "&cityId=" + strconv.Itoa(cityid) + "&seriesId=" + strconv.Itoa(seriesid) + "&modelId=" + strconv.Itoa(modelid) + "&year=" + strconv.Itoa(year) + "&month=" + strconv.Itoa(month) + "&mile= " + strconv.Itoa(mile) + "&_=1499239195245"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    split_string := strings.Split(strings.Split(string(body),"([{")[1],"}])")[0]
    sss  := strings.Split(split_string,"},{")
    for _, element := range sss {
      props = append(props, Pricefuture{Register_year: strings.Split(strings.Split(element, ":")[1],",")[0], Eval_price: strings.Split(element, ":")[2]})
    }
  }

	return props
}

func getPriceHistory(provid int,cityid int,modelid int,regyear int,regmonth int,mileage int,pri string) []Pricehistory {
  var props []Pricehistory
	qburl := "https://dingjia.che300.com/app/EvalResult/getPriceHistory?callback=jQuery183048247267398977844_1499239194770&provId=" + strconv.Itoa(provid) + "&cityId=" + strconv.Itoa(cityid) + "&modelId=" + strconv.Itoa(modelid) + "&regYear=" + strconv.Itoa(regyear) + "&regMonth=" + strconv.Itoa(regmonth) + "&mileAge=" + strconv.Itoa(mileage) + "&price=" + pri + "&_=1499239195244"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    split_string := strings.Split(strings.Split(string(body),":[{")[1],"}]})")[0]
    fmt.Println(split_string)
    sss  := strings.Split(split_string,"},{")
    for _, element := range sss {
      props = append(props, Pricehistory{Eval_price: strings.Split(strings.Split(element, ":")[1],",")[0], Date: strings.Split(element, ":")[2]})
    }
    fmt.Println(props)
  }

	return props
}

func allProvPrices(brand,series,model int,regDate string,mile int) ([]Provprices, []Prices){
  var props []Provprices
  var props2 []Prices
	qburl := "https://dingjia.che300.com/app/EvalResult/allProvPrices?callback=jQuery183048247267398977844_1499239194773&brand=" + strconv.Itoa(brand) + "&series=" + strconv.Itoa(series) + "&model=" + strconv.Itoa(model) + "&regDate=" +regDate+ "&mile=" + strconv.Itoa(mile) +"&_=1499239195252"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    str := strings.Split((strings.Split(string(body),":{")[2]),"},\"prices\":[{")
    str2 := strings.Split(str[0],",")
    fmt.Println(str2)
    str3 := strings.Split(str[0],",\"highlight_config")
    str5 := strings.Split(str3[0],",")
    str4 := strings.Split(str3[1],"}\",")
    str6 := strings.Split(str4[1],",")

    props = append(props, Provprices{Id: strings.Split(str5[0], ":")[1],
                                     Drive_name: strings.Split(str5[1], ":")[1],
                                     Model_status: strings.Split(str5[2], ":")[1],
                                     Name: strings.Split(str6[0], ":")[1],
                                     Short_name: strings.Split(str6[1], ":")[1],
                                     Market_date: strings.Split(str6[2], ":")[1],
                                     Stop_make_year: strings.Split(str6[3], ":")[1],
                                     Level_id: strings.Split(str6[4], ":")[1],
                                     Level: strings.Split(str6[5], ":")[1],
                                     Sname: strings.Split(str6[6], ":")[1],
                                     Ssname: strings.Split(str6[7], ":")[1],
                                     Maker_name: strings.Split(str6[8], ":")[1],
                                     Maker_type: strings.Split(str6[9], ":")[1],
                                     Star: strings.Split(str6[10], ":")[1],
                                     Sid: strings.Split(str6[11], ":")[1],
                                     Bid: strings.Split(str6[12], ":")[1],
                                     Brand_name: strings.Split(str6[13], ":")[1],
                                     Year: strings.Split(str6[14], ":")[1],
                                     Price: strings.Split(str6[15], ":")[1],
                                     Discharge_standard: strings.Split(str6[16], ":")[1],
                                     Gear_type: strings.Split(str6[17], ":")[1],
                                     Liter: strings.Split(str6[18], ":")[1],
                                     Liter_turbo: strings.Split(str6[19], ":")[1],
                                     Door_number: strings.Split(str6[20], ":")[1],
                                     Body_type: strings.Split(str6[21], ":")[1],
                                     Min_reg_year: strings.Split(str6[22], ":")[1],
                                     Max_reg_year: strings.Split(str6[23], ":")[1],
                                     Engine_power: strings.Split(str6[24], ":")[1],
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
  fmt.Println(props)
  fmt.Println(props2)

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

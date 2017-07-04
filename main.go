package main

import (
	"github.com/PuerkitoBio/goquery"
  "github.com/gin-gonic/gin"
  "fmt"
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

func main() {
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
      fmt.Println(props)
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
      fmt.Println(props)
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
    bodystr := []byte(body)

    er := json.Unmarshal(bodystr, &props)
    if er != nil {
      panic(er)
    } else {
      fmt.Println(props)
    }
  }

	return props
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

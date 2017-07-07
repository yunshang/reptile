package main
import (
	"github.com/PuerkitoBio/goquery"
  "fmt"
  "strings"
  "strconv"
	"log"
	iconv "github.com/djimenez/iconv-go"
  "net/http"
  "io/ioutil"
  // "encoding/json"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Jingbrand struct {
	Id      int    `json:"id"`
	Text      string    `json:"text"`
}

type Jingseries struct {
	Id      int    `json:"id"`
	Text      string    `json:"text"`
	Bid      int    `json:"bid"`
}

type Jingmodel struct {
	Id      int    `json:"id"`
	Text      string    `json:"text"`
	Sid      int    `json:"bid"`
}

type Prov2 struct {
  Name    string `json:"name"`
	Id int `json:"id"`
}

type Prov3 struct {
  Name    string `json:"name"`
	Id int `json:"id"`
}

type City2 struct {
  Id    string `json:"id"`
  Name    string `json:"name"`
	Pid string `json:"pid"`
}

type City3 struct {
  Id    string `json:"id"`
  Name    string `json:"name"`
	Pid string `json:"pid"`
}

func main() {
  // getCity()
	// getPro3()
  // getCity3()
    getbrand()
}

func getbrand() []Jingbrand {
  var props []Jingbrand
  qburl := "http://www.jingzhengu.com/Resources/ajax/PingGuHandlerV5.ashx?op=getAppointYearBeforeMake&year=2017"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	//

  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := []byte(body)
    str := strings.Split(string(bodystr), "[{")
    str1 := strings.Split(str[1],"}]")
		fmt.Println(str1[0])
    str2 := strings.Split(str1[0],"},{")
    for _, element := range str2 {
      id2 := strings.Split(element, ":")[5]
      id3,_ := strconv.Atoi(strings.Split(id2,"\"")[1])
      props = append(props, Jingbrand{Id: id3, Text: strings.Split(strings.Split(strings.Split(element, ":")[3],",")[0],"\"")[1]})
    }

		fmt.Println(props)
    // er := json.Unmarshal([]byte(str1[0]), &props)
    // if er != nil {
    //   panic(er)
    // } else {
    //   fmt.Println(props)
    // }
  }
	for _, b := range props {
		// rs, err := db.Exec("INSERT INTO chejing_brand(name, bid) VALUES (?, ?)", b.Text, b.Id)
		modelsID :=  b.Id
		getseries(modelsID)
		// fmt.Println(rs)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
	}
	return props
}

func getseries(brand int) []Jingseries {
  var props []Jingseries
  qburl := "http://www.jingzhengu.com/Resources/ajax/PingGuHandlerV5.ashx?op=getAppointYearBeforeModel&makeid=" + strconv.Itoa(brand) + "&year=2017"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	//

  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := []byte(body)
    str := strings.Split(string(bodystr), "[{")
    str1 := strings.Split(str[1],"}]")
		fmt.Println(str1[0])
    str2 := strings.Split(str1[0],"},{")
    for _, element := range str2 {
      id2 := strings.Split(element, ":")[5]
      id3,_ := strconv.Atoi(strings.Split(id2,"\"")[1])
      props = append(props, Jingseries{Id: id3, Text: strings.Split(strings.Split(strings.Split(element, ":")[3],",")[0],"\"")[1], Bid: brand})
    }

		fmt.Println(props)
    // er := json.Unmarshal([]byte(str1[0]), &props)
    // if er != nil {
    //   panic(er)
    // } else {
    //   fmt.Println(props)
    // }
  }
	for _, b := range props {
		// rs, err := db.Exec("INSERT INTO chejing_series(name, bid, uid) VALUES (?, ?, ?)", b.Text, brand, b.Id)
		// modelsID, _ := strconv.Atoi(b.Series_id)
		// getmodelBrand(modelsID)
		// fmt.Println(rs)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
    getmodel(b.Id)
	}
	return props
}

func getmodel(series int) []Jingmodel {
  var props []Jingmodel
  qburl := "http://www.jingzhengu.com/Resources/Ajax/PingGuHandlerV5.ashx?op=getAppointYearBeforeStyle&modelid=" + strconv.Itoa(series) + "&year=2017"
  client := &http.Client{}
  reqest, _ := http.NewRequest("GET", qburl, nil)
  response,_ := client.Do(reqest)
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
  defer db.Close()
  if err != nil{
    log.Fatalln(err)
  }
  db.SetMaxIdleConns(20)
  db.SetMaxOpenConns(20)
  //

  if response.StatusCode == 200 {
    body, _ := ioutil.ReadAll(response.Body)
    bodystr := []byte(body)
    if string(bodystr) != "" {
      str := strings.Split(string(bodystr), "[{")
      str1 := strings.Split(str[1],"}]")
      fmt.Println(str1[0])
      str2 := strings.Split(str1[0],"},{")
      for _, element := range str2 {
        id2 := strings.Split(element, ":")[5]
        id3,_ := strconv.Atoi(strings.Split(id2,"\"")[1])
        props = append(props, Jingmodel{Id: id3, Text: strings.Split(strings.Split(strings.Split(element, ":")[3],",")[0],"\"")[1], Sid: series})
      }

      fmt.Println(123)
      fmt.Println(props)
      for _, b := range props {
        // rs, err := db.Exec("INSERT INTO chejing_model(name, sid, uid) VALUES (?, ?, ?)", b.Text, series, b.Id)
        // fmt.Println(rs)
        if err != nil {
          log.Fatalln(err)
        }
      }
    }
      fmt.Println(321)
  }
  return props
}

func getCity3() []City3 {
  qburl := "http://www.jingzhengu.com/"

	doc, err := goquery.NewDocument(qburl)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	//

	var qb []City3
	doc.Find("#syytopcityid dd a").Each(func(i int, s *goquery.Selection) {
    id, _ := s.Attr("data")
    uid, _ := strconv.Atoi(id)
    content := s.Text()
    pid, _ := s.Attr("onclick")
    fmt.Println(uid)
    fmt.Println(content)
    fmt.Println(strings.Split(pid,"'")[1])
    // qb = append(qb, Prov3{Id: id, Pid: pid, Name: output})
    // rs, err := db.Exec("INSERT INTO chejing_city(uid, name, pid) VALUES (?, ?, ?)", uid, content,strings.Split(pid,"'")[1])
    // fmt.Println(rs)
    if err != nil {
      log.Fatalln(err)
    }
	})
	fmt.Println(qb)
	return qb
}


func getPro3() []Prov3 {
  qburl := "http://www.jingzhengu.com/"

	doc, err := goquery.NewDocument(qburl)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	//

	var qb []Prov3
	doc.Find("#syytopcityid dt").Each(func(i int, s *goquery.Selection) {
		if(i != 1) {
			id, _ := s.Attr("data")
			uid, _ := strconv.Atoi(id)
			content := s.Text()
			// qb = append(qb, Prov3{Id: id, Pid: pid, Name: output})
			// rs, err := db.Exec("INSERT INTO chejing_province(uid, name) VALUES (?, ?)", uid, content)
			// fmt.Println(rs)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})
	fmt.Println(qb)
	return qb
}

func getCity() []City2 {
  qburl := "http://www.che168.com/pinggu/#pvareaid=102140"

	doc, err := goquery.NewDocument(qburl)
	if err != nil {
		log.Fatal(err)
	}

 db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/reptile?parseTime=true")
 defer db.Close()
 if err != nil{
  log.Fatalln(err)
 }
 db.SetMaxIdleConns(20)
 db.SetMaxOpenConns(20)
	//

  var qb []City2
	doc.Find("#div_Area .city a").Each(func(i int, s *goquery.Selection) {
    pid, _ := s.Attr("pid")
    stringid, _ := s.Attr("id")
    content := s.Text()
		output,_ := iconv.ConvertString(content, "gb2312", "utf-8")
    qb = append(qb, City2{Id: stringid, Pid: pid, Name: output})
		// rs, err := db.Exec("INSERT INTO che168_city(uid, name, pid) VALUES (?, ?, ?)", stringid, output, pid)
  //   fmt.Println(rs)
		if err != nil {
			log.Fatalln(err)
		}
	})
	fmt.Println(qb)
	return qb
}

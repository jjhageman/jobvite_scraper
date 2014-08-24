package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

type Results struct {
  Results string
}


type Job struct {
  jobtitle string
}

func getJson(url string) ([]byte, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  return body, err
}

func printResponse(resp interface{}) {
  m := resp.(map[string]interface{})
  for k, v := range m {
    switch vv := v.(type) {
      case string:
        fmt.Println(k, "is string", vv)
      case int:
        fmt.Println(k, "is int", vv)
      case []interface{}:
        fmt.Println(k, "is an array:\n")
        printResponse(vv)
      default:
        fmt.Println(k, "is of a type I don't know how to handle")
    }
  }  
}

func decode(resp []byte, jobs *interface{}) (error) {
  err := json.Unmarshal(resp, &jobs)
  return err
}


func main() {
//GET /api/jobsearch?q=design+manager&radius=30&limit=10&start=0&latitude=37.7749295&longitude=-122.41941550000001&_=1408687065572 HTTP/1.1
//Host: search.jobvite.com
//Connection: keep-alive
//Accept: application/json, text/javascript, */*; q=0.01
//X-Requested-With: XMLHttpRequest
//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36
//Content-Type: application/json
//Referer: http://search.jobvite.com/web/modules/layout/home.htm
//Accept-Encoding: gzip,deflate,sdch
//Accept-Language: en-US,en;q=0.8
//Cookie: s_cc=true; s_campaign=www.indeed.com; s_sq=%5B%5BB%5D%5D; server=prd-fjweb04; __utma=246400283.1066620650.1408674419.1408674419.1408674419.1; __utmc=246400283; __utmz=246400283.1408674419.1.1.utmcsr=recruiting.jobvite.com|utmccn=(referral)|utmcmd=referral|utmcct=/; uidc=sa5Xafw6; wwwtidc=BD8C90962E8ECC30334444DC166EA925B7C09B9013B76703DB0653EF43306F4B; _mkto_trk=id:703-ISJ-362&token:_mch-jobvite.com-1408072765492-75000; __unam=de28ef2-147d7ad54fa-6342c0b0-9; linkedin_oauth_nr1knao7jpcn_crc=null; __utmd=1; __utma=251856523.1863009445.1408072770.1408682383.1408682510.5; __utmb=251856523.1.9.1408687065562; __utmc=251856523; __utmz=251856523.1408682383.4.4.utmcsr=recruiting.jobvite.com|utmccn=(referral)|utmcmd=referral|utmcct=/
  
  //url := "http://search.jobvite.com/api/jobsearch?q=design+manager&radius=30&limit=10&start=0&latitude=37.7749295&longitude=-122.41941550000001&_=1408684755349"
  url := "http://search.jobvite.com/api/jobsearch?q=design+manager&radius=100&limit=10&start=0&_=1408684755349"

  fmt.Printf("Getting jobs at: %s\n", url)

  resp, getErr := getJson(url)
  if getErr != nil {
    fmt.Printf("Sorry, something went wrong getting the url: %q\n", getErr)
  } else {



    fmt.Printf("Raw response: %q\n\n", resp)

    //var dat map[string]interface{}
    //if parseErr := json.Unmarshal(resp, &dat); parseErr != nil {
    //    panic(parseErr)
    //}
    //fmt.Println(dat)
    //results := dat["results"].([]interface{})
    //fmt.Println(co)
    //parseErr := decode(resp, &dat)
  }
}

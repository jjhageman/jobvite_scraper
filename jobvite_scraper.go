package main

import (
  "net/http"
  "fmt"
  "io"
  "io/ioutil"
  "bufio"
  "os"
  "encoding/json"
  "strings"
)

type Results struct {
  Results []Job
  Query string
  ResultsRaw interface{} `json:"totalResults"`

  TotalString string
  TotalInt    uint64
}

type Job struct {
  Jobtitle string
  FormattedLocation string
  Url string
  JobId string
  Modified string
  Company string
  Date string
  City string
  State string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func returnErr(e error) {
    if e != nil {
      return
    }
}

func (j Job) printJob() {
  fmt.Println("----------")
  fmt.Printf("%v - %s - %s\n", j.Modified, j.Company, j.Jobtitle)
  fmt.Println(j.FormattedLocation)
  fmt.Println(j.JobId)
}

func Decode(r io.Reader) (x *Results, err error) {
    x = new(Results)
    if err = json.NewDecoder(r).Decode(x); err != nil {
        return
    }
    switch t := x.ResultsRaw.(type) {
    case string:
        x.TotalString = t
    case float64:
        x.TotalInt = uint64(t)
    }
    return
}

func getJobs(url string) (r *Results, err error) {
  resp, err := http.Get(url)
  if err != nil {
    return
  }

  //fmt.Printf("Raw: %v\n\n", resp)

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return
  }

  json.Unmarshal(body, &r)

  return
}

func writeJobs(r *Results) (err error) {
  f, err := os.Create("/tmp/dat2")
  returnErr(err)

  defer f.Close()

  w := bufio.NewWriter(f)

  for _, job := range r.Results {
    if strings.Contains(job.City, "Francisco") || strings.Contains(job.FormattedLocation, "Francisco") {
      _, err := w.WriteString("buffered\n")
      returnErr(err)
    }
  }

  w.Flush()
  return
}

func main() {
  //url := "http://search.jobvite.com/api/jobsearch?q=project+manager&radius=30&limit=25&start=0&end=100&latitude=37.7749295&longitude=-122.41941550000001&_=1408684755349"
  url := "http://search.jobvite.com/api/jobsearch?q=project+manager&radius=100&limit=100&start=0&_=1408684755349"

  fmt.Printf("Getting jobs at: %s\n", url)

  var r Results
  _, err := getJobs(url)
  check(err)

  fmt.Printf("Searching %d/%v results on query '%s':\n\n", len(r.Results), r.ResultsRaw, r.Query)
  err2 := writeJobs(&r)
  check(err2)
}


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

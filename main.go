package main
import (
    "fmt"
    "log"
    "strings"
    "net/http"
	"github.com/PuerkitoBio/goquery"
	"os"
)
var links[] string 
//get all link
func scraper(){
    url_links := [...]string{
        "http://thegioisongngu.com/category/%E4%B8%AD%E5%9B%BD/",
        "http://thegioisongngu.com/category/%E4%B8%AD%E5%9B%BD/page/2/",
        "http://thegioisongngu.com/category/%E4%B8%AD%E5%9B%BD/page/3/", 
        "http://thegioisongngu.com/category/%E4%B8%AD%E5%9B%BD/page/4/",
    }
	for _, urls := range url_links {   
    resp, err := http.Get(urls)
    if err != nil{
        log.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200{
        log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
    }
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil{
        log.Fatal(err)
    }
    // fmt.Println(doc.Find("title").Text())

    doc.Find("div.entry").Each(func(i int, s *goquery.Selection){
        href, _ := s.Find("p span a").First().Attr("href")
        // if has_attr{
        //     fmt.Println(href)
        // }
        links=append(links,href)

    })
}
}
//get data
func scrapers(){
// fmt.Println(links)
file, _ := os.Create("result.txt")
defer file.Close()
for _, v := range links {   

    resp, err := http.Get(v)
    if err != nil{
        log.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200{
        log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
    }
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(doc.Find("title").Text())

    doc.Find("tr td").Each(func(i int, s *goquery.Selection){
        lines := "\n"+ strings.TrimSpace(s.Text())
        // fmt.Fprintf(file, strings.TrimSpace(s.Find("td").Text()) + "\n")
        fmt.Fprintf(file, lines + "\n")

        })
    }
}
func main(){
    scraper()
    scrapers()	
}

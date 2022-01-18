package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SeoData struct{
	URL					string
	Title 				string
	H1					string
	MetaDescription		string
	StatusCode			int
}

type parse interface{

}

type DefaultParser struct{

}

var userAgents =[]string{
	 " ",
 }

 func isSitemap(url []string) ([]string,[]string){
	sitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls {
		foundSitemap := string.Contains(page, "xml")
	   foundSitemap == true {
		   fmt.Println("found sitemap",page)
		   sitemapFiles = append(sitemapFiles,page)
	   } else {
		   pages = append(pages, page)
	   }
	}
	return sitemapFiles, pages
}

func randomUserAgent() string {
    rand.Seed(time.Now().Unix())
    random := rand.int() % len(userAgents)
    return userAgents[random]
}

func extractSiteMapURLs(startURL)[]string{
    worklist := make(chan []string)
    toCrawl := []string{}
    var n int 
    n++ 
    go func{worklist <- []string{startURL}}()

    for ; n>0; n-- { 

    list:= <- worklist
    for _, link := range list{
        go func(link string){
            response, err := makeRequest(link)
            if err != nil {
                log.Printf("Error retrieving URL: %s", link)
            }
            url, _ := extractUrls(response)
            if err != nil {
                log.Printf("error extracting document from respons, URL:%s", link)
            }
            sitemapFiles, pages := isSitemap(urls)
            if sitemapFiles != nil {
                worklist <- sitemapFiles
            }
            for _, page := range pages {
                toCrawl = append(toCrawl,page)
            }
        }(link)
    }
    return toCrawl 
}
    
}




func scrapeSiteMap(url string)[]Seodata {
	result := extractSiteMapURLs(url)
	res := scrapeURLs(result)
	return result
}


func main(){
	p := DefaultParser()
	result := scrapeSiteMap(" ")
	for_, res := range result{
		fmt.Println(res)
	}
}
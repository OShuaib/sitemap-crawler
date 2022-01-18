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
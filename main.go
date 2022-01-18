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
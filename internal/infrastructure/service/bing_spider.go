package service

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/internal/domain"
)

const BingBaseURL = "https://www.bing.com/search?q="

type BingSpider struct {
	baseUrl string
	ofType  domain.SearchEngineType
}

func NewBingSpider() *BingSpider {
	return &BingSpider{
		baseUrl: BingBaseURL,
		ofType:  domain.BING,
	}
}

func (b *BingSpider) GetSearchEngineType() domain.SearchEngineType {
	return b.ofType
}

func (b *BingSpider) Query(keyword domain.Keyword) (*domain.SearchEngine, error) {

	doc := b.fetchFromInternet(keyword.String())
	resultsData := b.parseDocumentData(doc)
	return domain.NewSearchEngine(b.ofType, resultsData), nil
}

func (b *BingSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(b.baseUrl + keyword)
	if err != nil {
		fmt.Println("Error fetching data from bing.com!")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return doc
}

func (b *BingSpider) parseDocumentData(doc *goquery.Document) *domain.QueryResults {
	resultsData := domain.EmptyQueryResult()
	doc.Find(".b_algo").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("p").Text()
		time := "unknown"
		resultsData.Add(b.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (b *BingSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.UrlBaseResultItem {
	return domain.NewResultItem(time, title, description, url)
}

package scrape

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"reflect"

	"github.com/gocolly/colly"
)

type Page struct {
	url   *url.URL
	colly *colly.Collector
	html  *colly.HTMLElement
}

func New(urlString string) (p *Page, err error) {
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		msg := fmt.Sprintf("URL: %s could not parsed, because %s", urlString, err)
		return nil, errors.New(msg)
	}

	p = &Page{
		colly: colly.NewCollector(),
		url:   u,
	}

	p.colly.OnHTML("html", func(b *colly.HTMLElement) {
		p.html = b
	})

	p.colly.Visit(p.url.String())

	return p, nil
}

func (p *Page) Scrape(collectors ...interface{}) {
	for _, c := range collectors {
		rv := reflect.ValueOf(&c).Elem()
		if rv.Kind() != reflect.Interface {
			log.Fatalf("collectors must be pointer, but %v", rv.Kind())
			return
		}

		re := rv.Elem()
		rt := reflect.Indirect(re).Type()
		tmp := reflect.New(rv.Elem().Type()).Elem()
		tmp.Set(re)
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			name := f.Name
			sel := f.Tag.Get("selector")
			if sel != "" {
				txt := p.html.DOM.Find(sel).Text()
				tmp.Elem().FieldByName(name).SetString(txt)
			}
		}

		rv.Set(tmp)
	}
}

func (p *Page) ScrapeTable(trSelector string, collectors interface{}) {
	st := reflect.TypeOf(collectors).Elem()
	sv := reflect.ValueOf(collectors).Elem()
	if st.Kind() != reflect.Slice {
		log.Fatalf("collectors must be slice, but %v", st.Kind())
		return
	}

	rt := st.Elem()
	var isPtr bool
	if rt.Kind() == reflect.Ptr {
		isPtr = true
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		log.Fatalf("a member of collectors must be struct, but %v", rt.Kind())
		return
	}

	p.html.ForEach(trSelector, func(i int, e *colly.HTMLElement) {
		if e.Name != "tr" {
			log.Printf("selector '%s' could not found <tr />, found <%s />", trSelector, e.Name)
			return
		}

		rp := reflect.New(rt)
		rv := rp.Elem()
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			name := f.Name
			sel := f.Tag.Get("selector")
			if sel != "" {
				txt := e.DOM.Find(sel).Text()
				rv.FieldByName(name).SetString(txt)
			}
		}

		if isPtr {
			sv.Set(
				reflect.Append(sv, rp),
			)
			return
		}

		sv.Set(
			reflect.Append(sv, rv),
		)
	})
}

func (p *Page) NextLink(selector string) (nextPageURL string) {
	return p.html.DOM.Find(selector).AttrOr("href", "")
}

func (p *Page) NextLinks(selector string) (nextPageURLs []string) {
	p.html.ForEach(selector, func(_ int, e *colly.HTMLElement) {
		nextPageURLs = append(nextPageURLs, e.Attr("href"))
	})

	return nextPageURLs
}

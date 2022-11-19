package main

import (
	"encoding/json"
	"log"
	"sync"
	"syscall/js"
	"wasm_sample/calc"
	"wasm_sample/scrape"
)

func main() {
	ch := make(chan struct{})

	exeFn := js.FuncOf(exec)
	defer exeFn.Release()

	fibFn := js.FuncOf(fib)
	defer fibFn.Release()

	fibMemFn := js.FuncOf(fibMem)
	defer fibMemFn.Release()

	js.Global().Set("scrapeExec", exeFn)
	js.Global().Set("fibonacci", fibFn)
	js.Global().Set("fibonacciMemorized", fibMemFn)
	<-ch
}

func exec(js.Value, []js.Value) interface{} {
	var wg sync.WaitGroup

	pages := getTargets()
	for i := range pages {
		wg.Add(1)
		go scraping(pages, i, &wg)
	}

	wg.Wait()
	return stringify(pages)
}

func fib(_ js.Value, args []js.Value) interface{} {
	return calc.Fibonacci(args[0].Int())
}

func fibMem(_ js.Value, args []js.Value) interface{} {
	return calc.FibonacciMemorized(args[0].Int())
}

func scraping(ps []page, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	var p = ps[i]
	s, err := scrape.New(p.URL)
	if err != nil {
		log.Fatal(err)
	}

	s.Scrape(&p.Result)
	ps[i] = p
}

type pageList []page

type collector struct {
	Title string `selector:"title" json:"title"`
	H1    string `selector:"h1" json:"h1"`
}

type page struct {
	Name   string    `json:"name"`
	URL    string    `json:"url"`
	Result collector `json:"result"`
}

func stringify(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func getTargets() []page {
	return []page{
		{
			Name: "北海道",
			URL:  "https://www.pref.hokkaido.lg.jp/",
		},
		{
			Name: "青森",
			URL:  "https://www.pref.aomori.lg.jp/",
		},
		{
			Name: "岩手",
			URL:  "https://www.pref.iwate.jp/",
		},
		{
			Name: "宮城",
			URL:  "https://www.pref.miyagi.jp/",
		},
		{
			Name: "秋田",
			URL:  "https://www.pref.akita.lg.jp/",
		},
		{
			Name: "山形",
			URL:  "https://www.pref.yamagata.jp/",
		},
		{
			Name: "福島",
			URL:  "https://www.pref.fukushima.lg.jp/",
		},
		{
			Name: "茨城",
			URL:  "https://www.pref.ibaraki.jp/",
		},
		{
			Name: "栃木",
			URL:  "https://www.pref.tochigi.lg.jp/",
		},
		{
			Name: "群馬",
			URL:  "https://www.pref.gunma.jp/",
		},
		{
			Name: "埼玉",
			URL:  "https://www.pref.saitama.lg.jp/",
		},
		{
			Name: "千葉",
			URL:  "https://www.pref.chiba.lg.jp/",
		},
		{
			Name: "東京",
			URL:  "https://www.metro.tokyo.lg.jp/",
		},
		{
			Name: "神奈川",
			URL:  "https://www.pref.kanagawa.jp/",
		},
		{
			Name: "新潟",
			URL:  "https://www.pref.niigata.lg.jp/",
		},
		{
			Name: "富山",
			URL:  "https://www.pref.toyama.jp/",
		},
		{
			Name: "石川",
			URL:  "https://www.pref.ishikawa.lg.jp/",
		},
		{
			Name: "福井",
			URL:  "https://www.pref.fukui.lg.jp/",
		},
		{
			Name: "山梨",
			URL:  "https://www.pref.yamanashi.jp/",
		},
		{
			Name: "長野",
			URL:  "https://www.pref.nagano.lg.jp/",
		},
		{
			Name: "岐阜",
			URL:  "https://www.pref.gifu.lg.jp/",
		},
		{
			Name: "静岡",
			URL:  "https://www.pref.shizuoka.jp/",
		},
		{
			Name: "愛知",
			URL:  "https://www.pref.aichi.jp/",
		},
		{
			Name: "三重",
			URL:  "https://www.pref.mie.lg.jp/",
		},
		{
			Name: "滋賀",
			URL:  "https://www.pref.shiga.lg.jp/",
		},
		{
			Name: "京都",
			URL:  "https://www.pref.kyoto.jp/",
		},
		{
			Name: "大阪",
			URL:  "https://www.pref.osaka.lg.jp/",
		},
		{
			Name: "奈良",
			URL:  "https://www.pref.nara.jp/",
		},
		{
			Name: "兵庫",
			URL:  "https://web.pref.hyogo.lg.jp/",
		},
		{
			Name: "和歌山",
			URL:  "https://www.pref.wakayama.lg.jp/",
		},
	}
}

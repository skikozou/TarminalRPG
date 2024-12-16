package syosetu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	TarminalRPG "TarminalRPG/src"

	"github.com/PuerkitoBio/goquery"
)

func GetSyosetu(ncode string, page int) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ncode.syosetu.com/%s/%s", ncode, strconv.Itoa(page)), nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Set("User-Agent", "none")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}

	doc, err := extractSyosetu(resp)
	if err != nil {
		return []string{}, err
	}

	return doc, nil
}

func GetSyosetuList(ncode string) (SyosetuList, error) {
	var lists SyosetuList
	var pages int
	res, err := GetNovelAPI(ncode)
	pages = res[1].GeneralAllNo
	if err != nil {
		return SyosetuList{}, err
	}
	for i := 0; i < (pages-pages%100)/100+1; i++ {
		req, err := http.NewRequest("GET", fmt.Sprintf("https://ncode.syosetu.com/%s/?p=%d", ncode, i+1), nil)
		if err != nil {
			return SyosetuList{}, err
		}
		req.Header.Set("User-Agent", "none")
		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			return SyosetuList{}, err
		}

		list, err := extractSyosetuList(resp)
		if err != nil {
			return SyosetuList{}, err
		}

		lists = append(lists, list...)
	}

	return lists, nil
}

func GetNovelAPI(ncode string) (NovelAPIResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.syosetu.com/novelapi/api/?out=json&ncode=%s", ncode), nil)
	if err != nil {
		return NovelAPIResponse{}, err
	}
	req.Header.Set("User-Agent", "none")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return NovelAPIResponse{}, err
	}

	var resData NovelAPIResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NovelAPIResponse{}, err
	}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return NovelAPIResponse{}, err
	}

	return resData, nil
}

func extractSyosetu(res *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return []string{}, err
	}

	var novel []string
	doc.Find("div.js-novel-text.p-novel__text").Each(func(i int, s *goquery.Selection) {
		children := s.Children()
		children.Each(func(j int, child *goquery.Selection) {
			novel = append(novel, child.Text())
		})
	})

	return novel, nil
}

func extractSyosetuList(res *http.Response) (SyosetuList, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return SyosetuList{}, err
	}

	var list SyosetuList
	var Err error
	var EachFunc = func(i int, s *goquery.Selection) {
		val, _ := s.Attr("class")
		if val == "p-eplist__sublist" {
			page, err := strconv.Atoi(strings.Split(s.Children().Nodes[0].Attr[0].Val, "/")[2])
			title := strings.Split(s.Children().Nodes[0].FirstChild.Data, "\n")[1]
			if err != nil {
				return
			}

			list = append(list, struct {
				page  int
				title string
			}{page, title})
		}
	}
	doc.Find("div.p-eplist > div").Each(EachFunc)

	if Err != nil {
		return SyosetuList{}, Err
	}

	return list, nil
}

func NoveltoObject(list SyosetuList) TarminalRPG.Project {
	return TarminalRPG.Project{}
}

func NoveltoJson(list SyosetuList) string {
	return ""
}

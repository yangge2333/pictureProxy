package api

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type loliconRsp struct {
	Error string `json:"error"`
	Data  []Pic  `json:"data"`
}
type Pic struct {
	Urls Url `json:"urls"`
}
type Url struct {
	Original string `json:"original"`
}

var Hello = helloApi{}

type helloApi struct{}

type PictureMap struct {
	picMap *sync.Map
	getPic *chan int
}

var GlobalR18PictureMap *PictureMap
var GlobalPictureMap *PictureMap
var GlobalSnbPictureMap *PictureMap

func init() {
	r18PicChan := make(chan int, 10)
	picChan := make(chan int, 10)
	snbChan := make(chan int, 10)
	GlobalR18PictureMap = &PictureMap{
		picMap: &sync.Map{},
		getPic: &r18PicChan,
	}
	GlobalPictureMap = &PictureMap{
		picMap: &sync.Map{},
		getPic: &picChan,
	}
	GlobalSnbPictureMap = &PictureMap{
		picMap: &sync.Map{},
		getPic: &snbChan,
	}
	go GlobalR18PictureMap.AddR18Picture()
	go GlobalPictureMap.AddPicture()
	go GlobalSnbPictureMap.AddSnbPicture()
}

func (m *PictureMap) AddR18Picture() {
	for {
		*m.getPic <- 1
		rsp := loliconApi(`{"r18" :1}`)
		if rsp == nil && len(rsp.Data) != 0 {
			return
		}
		m.picMap.Store(rsp.Data[0].Urls.Original, getPicFromUrl(rsp.Data[0].Urls.Original))
	}
}

func (m *PictureMap) AddPicture() {
	for {
		*m.getPic <- 1
		rsp := loliconApi(`{"r18" :0}`)
		if rsp == nil && len(rsp.Data) != 0 {
			return
		}
		m.picMap.Store(rsp.Data[0].Urls.Original, getPicFromUrl(rsp.Data[0].Urls.Original))
	}
}

func (m *PictureMap) AddSnbPicture() {
	for {
		*m.getPic <- 1
		rsp := loliconApi(`{"tag": ["久岐忍"]}`)
		if rsp == nil && len(rsp.Data) != 0 {
			return
		}
		m.picMap.Store(rsp.Data[0].Urls.Original, getPicFromUrl(rsp.Data[0].Urls.Original))
	}
}

func (m *PictureMap) GetPicture() []byte {
	result := make([]byte, 0)
	select {
	case <-*m.getPic:
		m.picMap.Range(func(key, value any) bool {
			result = value.([]byte)
			m.picMap.Delete(key)
			return false
		})
	}
	return result
}

func (*helloApi) R18(r *ghttp.Request) {
	r.Response.WriteExit(GlobalR18PictureMap.GetPicture())
}

func (*helloApi) Normal(r *ghttp.Request) {
	r.Response.WriteExit(GlobalPictureMap.GetPicture())
}

func (*helloApi) KukiShinobu(r *ghttp.Request) {
	r.Response.WriteExit(GlobalSnbPictureMap.GetPicture())
}

func loliconApi(params string) *loliconRsp {
	url := "https://api.lolicon.app/setu/v2"
	method := "POST"

	payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))

	rsp := &loliconRsp{}
	json.Unmarshal(body, rsp)
	return rsp
}

func getPicFromUrl(url string) []byte {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

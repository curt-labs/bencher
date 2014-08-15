package bencher

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type HttpRoute struct {
	Url     string       `json:"url"`
	Method  string       `json:"method"`
	Params  []RouteParam `json:"params"`
	Results HttpResults  `json:"result"`
}

type HttpResults struct {
	BenchResult testing.BenchmarkResult `json:"bench"`
	Results     []HttpResult            `json:"results"`
}

type HttpResult struct {
	StatusCode int    `json:"code"`
	Body       string `json:"body"`
	Url        string `json:"url"`
	Type       string `json:"type"`
}

type RouteParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (r *HttpRoute) Benchmark() HttpResult {

	var result HttpResult
	params := url.Values{}
	for _, param := range r.Params {
		params.Add(param.Key, param.Value)
	}

	client := &http.Client{}
	var req *http.Request
	var err error
	if strings.ToUpper(r.Method) == "GET" {
		req, err = http.NewRequest("GET", r.Url, bytes.NewBufferString(params.Encode()))
	} else {
		req, err = http.NewRequest(strings.ToUpper(r.Method), r.Url, bytes.NewBufferString(params.Encode()))
	}

	if err != nil {
		return result
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return result
	}

	result.StatusCode = resp.StatusCode

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result
	}
	result.Body = string(buf)

	if resp.Header.Get("Conent-Type") == "application/json" {
		result.Type = "json"
	}

	return result
}

func Run(c *gin.Context) {

	if err := c.Req.ParseMultipartForm(10000); err != nil {
		c.Fail(500, err)
		return
	}

	m := c.Req.MultipartForm
	if m == nil {
		c.Fail(500, errors.New("failed to parse data"))
		return
	}

	dt := m.File["file"][0]
	file, err := dt.Open()
	if err != nil {
		c.Fail(500, err)
		return
	}
	defer file.Close()

	f, err := ioutil.ReadAll(file)
	if err != nil {
		c.Fail(500, err)
		return
	}

	var routes []HttpRoute
	err = json.Unmarshal(f, &routes)
	if err != nil {
		c.Fail(500, err)
		return
	}

	for i, route := range routes {
		results := make([]HttpResult, 0)
		benchResult := testing.Benchmark(func(bk *testing.B) {
			bk.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					results = append(results, route.Benchmark())
					i++
					if i > 50 {
						break
					}
				}

			})
		})

		routes[i].Results = HttpResults{
			Results:     results,
			BenchResult: benchResult,
		}
	}

	c.JSON(200, routes)
	return
}

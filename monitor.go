package ndrf

import (
	"gopkg.in/resty.v1"
	"sync/atomic"
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

//resty monitor

type status struct {
	Code1xx int64
	Code2xx int64
	Code3xx int64
	Code4xx int64
	Code5xx int64
	Codesum int64
}

type restyKin struct {
	status status
	useTime time.Duration
}

var restyKins map[string]*restyKin

func init() {
	restyKins = make(map[string]*restyKin, 0)
	handelrStatuss = make(map[string]*handlerStatus, 0)
}

func RestyMonitor(client *resty.Client, response *resty.Response) error {
	code := response.RawResponse.StatusCode
	switch  {
	case 100 <= code && code < 200:
		atomic.AddInt64(&(restyKins[client.HostURL].status.Code1xx), 1)
	case 200 <= code && code < 300:
		atomic.AddInt64(&(restyKins[client.HostURL].status.Code2xx), 1)
	case 300 <= code && code < 400:
		atomic.AddInt64(&(restyKins[client.HostURL].status.Code3xx), 1)
	case 400 <= code && code < 500:
		atomic.AddInt64(&(restyKins[client.HostURL].status.Code4xx), 1)
	case 500 <= code && code < 600:
		atomic.AddInt64(&(restyKins[client.HostURL].status.Code5xx), 1)
	}
	atomic.AddInt64(&(restyKins[client.HostURL].status.Codesum), 1)
	restyKins[client.HostURL].useTime = restyKins[client.HostURL].useTime + response.Time()
	return nil
}

func OpenMonitorResty(r *gin.Engine, serviceName string) {
	r.GET("/monitor/resty", func(c *gin.Context) {
		html := ""
		for k,v := range restyKins {
			html1xx := fmt.Sprintf(`http_request_status{target="%s", code="%s", service="%s"} %d` + "\n",k, "1xx", serviceName, atomic.LoadInt64(&v.status.Code1xx))
			html2xx := fmt.Sprintf(`http_request_status{target="%s", code="%s", service="%s"} %d` + "\n",k, "2xx", serviceName, atomic.LoadInt64(&v.status.Code2xx))
			html3xx := fmt.Sprintf(`http_request_status{target="%s", code="%s", service="%s"} %d` + "\n",k, "3xx", serviceName, atomic.LoadInt64(&v.status.Code3xx))
			html4xx := fmt.Sprintf(`http_request_status{target="%s", code="%s", service="%s"} %d` + "\n",k, "4xx", serviceName, atomic.LoadInt64(&v.status.Code4xx))
			html5xx := fmt.Sprintf(`http_request_status{target="%s", code="%s", service="%s"} %d` + "\n",k, "5xx", serviceName, atomic.LoadInt64(&v.status.Code5xx))
			htmlTime := fmt.Sprintf(`http_request_time{target="%s", service="%s"} %d` + "\n",k, serviceName, v.useTime)
			html = html + html1xx + html2xx + html3xx + html4xx + html5xx + htmlTime
		}
		c.String(http.StatusOK,html)
	})

}

func CreateResty(host string) *resty.Client {
	restyKins[host] = new(restyKin)
	r := resty.New()
	r.SetHostURL(host)
	r.OnAfterResponse(RestyMonitor)
	return r
}


//service monitor
type handlerStatus struct {
	method string
	path string
	status status
	size int64
	time time.Duration
}
var handelrStatuss map[string]*handlerStatus

func ServiceCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		if _,ok := handelrStatuss[c.HandlerName()]; ok {
			fmt.Println("进入循环")
			switch {
			case 100 <= c.Writer.Status() && c.Writer.Status() < 200:
				atomic.AddInt64(&(handelrStatuss[c.HandlerName()].status.Code1xx), 1)
			case 200 <= c.Writer.Status() && c.Writer.Status() < 300:
				atomic.AddInt64(&(handelrStatuss[c.HandlerName()].status.Code2xx), 1)
			case 300 <= c.Writer.Status() && c.Writer.Status() < 400:
				atomic.AddInt64(&(handelrStatuss[c.HandlerName()].status.Code3xx), 1)
			case 400 <= c.Writer.Status() && c.Writer.Status() < 500:
				atomic.AddInt64(&(handelrStatuss[c.HandlerName()].status.Code4xx), 1)
			case 500 <= c.Writer.Status() && c.Writer.Status() < 600:
				atomic.AddInt64(&(handelrStatuss[c.HandlerName()].status.Code5xx), 1)
			}
			atomic.AddInt64(&(handelrStatuss[c.HandlerName()].status.Codesum), 1)
			userTime := time.Now().Sub(startTime)
			handelrStatuss[c.HandlerName()].time = handelrStatuss[c.HandlerName()].time + (userTime / 1000)
			atomic.AddInt64(&(handelrStatuss[c.HandlerName()].size), int64(c.Writer.Size()))
			fmt.Println("fasd", handelrStatuss)
		}
	}
}

func OpenMonitorService(r *gin.Engine, serviceName string) {
	for _, v := range r.Routes() {
		h := new(handlerStatus)
		h.path = v.Path
		h.method = v.Method
		h.status.Codesum = 0
		h.status.Code1xx = 0
		h.status.Code2xx = 0
		h.status.Code3xx = 0
		h.status.Code4xx = 0
		h.status.Code5xx = 0
		h.time = 0
		h.size = 0
		handelrStatuss[v.Handler] = h
	}
	r.GET("/monitor/service", func(c *gin.Context) {
		html := ""
		for _,v := range handelrStatuss{
			html1xx := fmt.Sprintf(`http_response_status{method="%s",path="%s", code="%s", service="%s"} %d` + "\n",v.method, v.path, "1xx",serviceName, atomic.LoadInt64(&v.status.Code1xx))
			html2xx := fmt.Sprintf(`http_response_status{method="%s",path="%s", code="%s", service="%s"} %d` + "\n",v.method, v.path, "2xx",serviceName, atomic.LoadInt64(&v.status.Code2xx))
			html3xx := fmt.Sprintf(`http_response_status{method="%s",path="%s", code="%s", service="%s"} %d` + "\n",v.method, v.path, "3xx",serviceName, atomic.LoadInt64(&v.status.Code3xx))
			html4xx := fmt.Sprintf(`http_response_status{method="%s",path="%s", code="%s", service="%s"} %d` + "\n",v.method, v.path, "4xx",serviceName, atomic.LoadInt64(&v.status.Code4xx))
			html5xx := fmt.Sprintf(`http_response_status{method="%s",path="%s", code="%s", service="%s"} %d` + "\n",v.method, v.path, "5xx",serviceName, atomic.LoadInt64(&v.status.Code5xx))
			htmlTime := fmt.Sprintf(`http_response_time{method="%s",path="%s", service="%s"} %d` + "\n",v.method, v.path,serviceName, v.time)
			htmlSize := fmt.Sprintf(`http_response_size{method="%s",path="%s", service="%s"} %d` + "\n",v.method, v.path,serviceName, atomic.LoadInt64(&v.size))

			html = html + html1xx + html2xx + html3xx + html4xx + html5xx + htmlTime + htmlSize
		}
		c.String(http.StatusOK,html)
	})
}
package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/transport/server/interfaces"
	"time"

	"github.com/gorilla/mux"
)

// Channel -
type Channel struct {
	handlers   map[string]interfaces.IHandler
	conf       *configuration.ChannelConfig
	httpServer *http.Server
}

// SetConfig -
func (it *Channel) SetConfig(config *configuration.ChannelConfig) {
	it.conf = config
}

// RegisteHandler - channel和handler是一对一关系
func (it *Channel) RegisteHandler(rule string, handler interfaces.IHandler) {
	if it.handlers == nil {
		it.handlers = make(map[string]interfaces.IHandler)
	}
	it.handlers[rule] = handler
}

// MakeRouter - 统一一个接口，但是请求由不同的携程处理
func (it *Channel) MakeRouter() {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 10
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/api/{rule:.*}", it.handleRequest).Methods("POST")
	it.httpServer = &http.Server{
		Addr:           it.conf.LoaclAddress,
		Handler:        muxRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

// Open -
func (it *Channel) Open() error {
	if err := it.httpServer.ListenAndServe(); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Close -
func (it *Channel) Close() error {
	return it.httpServer.Close()
}

func (it *Channel) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	rule := params["rule"]
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	// w.Header().Set("content-type","text/plain")
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	go func(data []byte) {
		// 没有数据处理器，发送错误
		if it.handlers == nil {
			io.WriteString(w, "404")
			return
		}
		handler := it.handlers[rule]
		if handler == nil {
			io.WriteString(w, "404")
			return
		}

		// 将数据发送给用户定义的处理器，如果处理出错，返回错误
		if request, err := handler.OnReceive(data); err == nil {
			io.WriteString(w, string(request[:]))
		} else {
			io.WriteString(w, err.Error())
		}
	}(body)
}

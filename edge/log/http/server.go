package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"star-edge-cloud/edge/log/implemention"
	"star-edge-cloud/edge/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// LogServer - 日志服务
type LogServer struct {
	Logger     *implemention.Logger
	ServerAddr string
}

// Start - 启动服务
func (ls *LogServer) Start() error {
	mux := ls.makeMuxRouter()
	log.Println("服务启动，开始监听端口", ls.ServerAddr)
	s := &http.Server{
		Addr:           ls.ServerAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// create handlers
func (ls *LogServer) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/api/log/write", ls.handleWriteLog).Methods("POST")
	muxRouter.HandleFunc("/api/log/remove", ls.handleRemoveLog).Methods("POST")
	muxRouter.HandleFunc("/api/log/load", ls.handleLoadLog).Methods("POST")
	return muxRouter
}

func (ls *LogServer) handleWriteLog(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var info models.LogInfo
	err := json.Unmarshal(body, &info)
	if err != nil {
		// log.Println(err.Error())
		io.WriteString(w, err.Error())
		return
	}
	ls.Logger.Write(&info)
	io.WriteString(w, "success")
}

func (ls *LogServer) handleRemoveLog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	id := r.Form.Get("id")
	if err := ls.Logger.Delete(id); err != nil {
		io.WriteString(w, "success")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (ls *LogServer) handleLoadLog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	// start := r.Form.Get("start")
	// end := r.Form.Get("end")
	// level := r.Form.Get("level")
	// start1, _ := time.Parse("2006-01-02 15:04:05", start)
	// end1, _ := time.Parse("2006-01-02 15:04:05", end)
	// level1, _ := strconv.Atoi(level)

	top := r.Form.Get("top")
	top1, _ := strconv.Atoi(top)
	infos, err := ls.Logger.QueryTop(top1)
	data, err := json.Marshal(infos)
	io.WriteString(w, string(data[:]))
}

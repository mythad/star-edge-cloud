package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/interfaces"
	"time"

	"github.com/gorilla/mux"
)

// RestServer -
type RestServer struct {
	requestHandler      interfaces.IRequestHandler
	alarmHandler        interfaces.IAlarmHandler
	commandHandler      interfaces.ICommandHandler
	faultHandler        interfaces.IFaultHandler
	realtimeDataHandler interfaces.IRealtimeDataHandler
	resultHandler       interfaces.IResultHandler
	ServerAddr          string
}

// Start - 启动服务
func (cs *RestServer) Start() error {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 10
	// cs.logger.LogDBPath = "../../../data/sqlite3/log.db"
	// info := &model.LogInfo{ID: common.GetGUID(), Message: "服务启动:" + cs.ServerAddr, Level: 3, Time: time.Now()}
	// cs.logger.Write(info)
	mux := cs.makeMuxRouter()
	// godotenv.Load("conf.env")
	// httpPort := os.Getenv("PORT")
	// log.Println("服务启动:", cs.ServerAddr)
	s := &http.Server{
		Addr:           cs.ServerAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		// common.AppendToFile("/root/a.txt", err.Error())
		os.Exit(10)
		return err
	}

	return nil
}

// SetConfig -
func (cs *RestServer) SetConfig(conf string) {

}

// Stop - 停止服务
func (cs *RestServer) Stop() error {
	return nil
}

// create handlers
func (cs *RestServer) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/api/transport/request", cs.handleRequest).Methods("POST")
	muxRouter.HandleFunc("/api/transport/command", cs.handleCommand).Methods("POST")
	muxRouter.HandleFunc("/api/transport/realtimedata", cs.handleRealtimeData).Methods("POST")
	muxRouter.HandleFunc("/api/transport/alarm", cs.handleAlarm).Methods("POST")
	muxRouter.HandleFunc("/api/transport/fault", cs.handleFault).Methods("POST")
	muxRouter.HandleFunc("/api/transport/result", cs.handleResult).Methods("POST")
	return muxRouter
}

// SetRequestHandler -
func (cs *RestServer) SetRequestHandler(handler interfaces.IRequestHandler) {
	cs.requestHandler = handler
}

// SetCommandHandler -
func (cs *RestServer) SetCommandHandler(handler interfaces.ICommandHandler) {
	cs.commandHandler = handler
}

// SetRealtimeDataHandler -
func (cs *RestServer) SetRealtimeDataHandler(handler interfaces.IRealtimeDataHandler) {
	cs.realtimeDataHandler = handler
}

// SetAlarmHandler -
func (cs *RestServer) SetAlarmHandler(handler interfaces.IAlarmHandler) {
	cs.alarmHandler = handler
}

// SetFaultHandler -
func (cs *RestServer) SetFaultHandler(handler interfaces.IFaultHandler) {
	cs.faultHandler = handler
}

// SetResultHandler -
func (cs *RestServer) SetResultHandler(handler interfaces.IResultHandler) {
	cs.resultHandler = handler
}

func (cs *RestServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if cs.requestHandler == nil {
		io.WriteString(w, "no request data handler.")
		return
	}

	request := &models.Request{}
	response := &models.Response{}
	if err := json.Unmarshal(body, &request); err == nil {
		go cs.requestHandler.Handle(request, response)
		if data, err := json.Marshal(response); err == nil {
			io.WriteString(w, string(data[:]))
			return
		}
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleCommand(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if cs.commandHandler == nil {
		io.WriteString(w, "no command data handler.")
		return
	}

	request := &models.Command{}
	response := &models.Response{}
	if err := json.Unmarshal(body, &request); err == nil {
		go cs.commandHandler.Handle(request, response)
		if data, err := json.Marshal(response); err == nil {
			io.WriteString(w, string(data[:]))
			return
		}
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleRealtimeData(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if cs.realtimeDataHandler == nil {
		io.WriteString(w, "no realtime data handler.")
		r.Body.Close()
		return
	}

	request := &models.RealtimeData{}
	response := &models.Response{}
	if err := json.Unmarshal(body, &request); err == nil {
		go cs.realtimeDataHandler.Handle(request, response)
		if data, err := json.Marshal(response); err == nil {
			io.WriteString(w, string(data[:]))
			return
		}
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleAlarm(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if cs.alarmHandler == nil {
		io.WriteString(w, "no alarm data handler.")
		return
	}

	request := &models.Alarm{}
	response := &models.Response{}
	if err := json.Unmarshal(body, &request); err == nil {
		go cs.alarmHandler.Handle(request, response)
		if data, err := json.Marshal(response); err == nil {
			io.WriteString(w, string(data[:]))
			return
		}
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleFault(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if cs.faultHandler == nil {
		io.WriteString(w, "no fault data handler.")
		return
	}

	request := &models.Fault{}
	response := &models.Response{}
	if err := json.Unmarshal(body, &request); err == nil {
		go cs.faultHandler.Handle(request, response)
		if data, err := json.Marshal(response); err == nil {
			io.WriteString(w, string(data[:]))
			return
		}
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleResult(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if cs.resultHandler == nil {
		io.WriteString(w, "no result data handler.")
		return
	}

	request := &models.Result{}
	response := &models.Response{}
	if err := json.Unmarshal(body, &request); err == nil {
		go cs.resultHandler.Handle(request, response)
		if data, err := json.Marshal(response); err == nil {
			io.WriteString(w, string(data[:]))
			return
		}
		io.WriteString(w, err.Error())
	} else {
		io.WriteString(w, err.Error())
	}
}

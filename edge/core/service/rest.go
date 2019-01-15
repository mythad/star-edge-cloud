package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"star_cloud/edge/core/config"
	"star_cloud/edge/core/device"
	"star_cloud/edge/core/extension"
	"star_cloud/edge/models"
	"star_cloud/edge/utils/common"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CoreServer - 元数据服务
type CoreServer struct {
	DevManager   device.DeviceManager
	ExtManager   extension.ExtentionManager
	StoreManager extension.StoreManager
	LogManager   extension.LogManager
	Conf         *config.Config
}

// Start - 启动服务
func (cs *CoreServer) Start() error {
	mux := cs.makeMuxRouter()
	// cs.Loghelper.Write(&models.LogInfo{ID: common.GetGUID(), Message: "Core服务启动，开始监听端口 :" + cs.Conf.Port, Level: 1, Time: time.Now()})
	log.Println("服务启动，开始监听端口 :", cs.Conf.Port)
	s := &http.Server{
		Addr:           ":" + cs.Conf.Port,
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
func (cs *CoreServer) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/{category:html|js|css|images}/{name:.*}", cs.handleStaticResource)
	muxRouter.HandleFunc("/api/device/add", cs.handleAddDevice).Methods("POST")
	muxRouter.HandleFunc("/api/device/remove", cs.handleRemoveDevice).Methods("POST")
	muxRouter.HandleFunc("/api/device/run", cs.handleRunDevice).Methods("POST")
	muxRouter.HandleFunc("/api/device/all", cs.handleAllDevice).Methods("POST")
	muxRouter.HandleFunc("/api/device/count", cs.handleCountDevice).Methods("POST")
	muxRouter.HandleFunc("/api/device/stop", cs.handleStopDevice).Methods("POST")
	muxRouter.HandleFunc("/api/device/status", cs.handleGetDeviceStatus).Methods("POST")
	muxRouter.HandleFunc("/api/extension/add", cs.handleAddExtension).Methods("POST")
	muxRouter.HandleFunc("/api/extension/remove", cs.handleRemoveExtension).Methods("POST")
	muxRouter.HandleFunc("/api/extension/run", cs.handleRunExtension).Methods("POST")
	muxRouter.HandleFunc("/api/extension/stop", cs.handleStopExtension).Methods("POST")
	muxRouter.HandleFunc("/api/extension/all", cs.handleAllExtension).Methods("POST")
	muxRouter.HandleFunc("/api/extension/count", cs.handleCountExtension).Methods("POST")
	muxRouter.HandleFunc("/api/store/run", cs.handleRunStore).Methods("POST")
	muxRouter.HandleFunc("/api/store/stop", cs.handleStopStore).Methods("POST")
	muxRouter.HandleFunc("/api/store/status", cs.handleStoreStatus).Methods("POST")
	muxRouter.HandleFunc("/api/log/run", cs.handleRunLog).Methods("POST")
	muxRouter.HandleFunc("/api/log/stop", cs.handleStopLog).Methods("POST")
	muxRouter.HandleFunc("/api/log/status", cs.handleLogStatus).Methods("POST")
	muxRouter.HandleFunc("/api/help/content", cs.handleHelp).Methods("Get")
	// muxRouter.HandleFunc("/api/{name:.*}", cs.handle)
	return muxRouter
}

func (cs *CoreServer) handleStaticResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	path := fmt.Sprintf("./website/%[1]s", name)
	content := common.ReadString(path)
	io.WriteString(w, content)
}

func (cs *CoreServer) handleAddDevice(w http.ResponseWriter, r *http.Request) {
	// 创建设备目录和设备运行文件
	file, head, err := r.FormFile("file")
	if err == nil {
		io.WriteString(w, "未上传文件")
	}
	defer file.Close()

	// 读取设备信息
	name := r.Form.Get("dev_name")
	conf := r.Form.Get("conf")
	listeners := r.Form.Get("listeners")
	cmdaddr := r.Form.Get("cmd_addr")
	logurl := r.Form.Get("logurl")
	now := time.Now().Format("2006-01-02 15:04:05")
	device := &models.Device{
		ID:            common.GetGUID(),
		Name:          name,
		Conf:          conf,
		FileName:      head.Filename,
		ServerAddress: cmdaddr,
		Listeners:     listeners,
		RegistryTime:  now,
		LogBaseURL:    logurl}
	if err := cs.DevManager.AddDevice(file, device); err == nil {
		io.WriteString(w, "success.")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleRemoveDevice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	if err := cs.DevManager.RemoveDevice(id); err == nil {
		io.WriteString(w, "success.")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleRunDevice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	err = cs.DevManager.Run(id)
	if err == nil {
		io.WriteString(w, "running")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleAllDevice(w http.ResponseWriter, r *http.Request) {
	devices, _ := cs.DevManager.QueryAllDevice()
	// 更新状态
	for index, item := range devices {
		if status := cs.DevManager.GetStatus(&item); status != item.Status {
			devices[index].Status = status
			cs.DevManager.UpdateDevice(&devices[index])
		}
	}
	data, _ := json.Marshal(devices)
	io.WriteString(w, string(data[:]))
}

func (cs *CoreServer) handleStopDevice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	err = cs.DevManager.Stop(id)
	if err == nil {
		io.WriteString(w, "stopped")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleCountDevice(w http.ResponseWriter, r *http.Request) {
	// id := r.Form.Get("id")
	// w.Header().Set("Content-Type", "text/plain")
	devices, _ := cs.DevManager.QueryAllDevice()
	var runNum, totalNum int
	for _, d := range devices {
		if d.Status == 2 {
			runNum++
		}
		totalNum++
	}
	json := fmt.Sprintf(`{"Running":%[1]d,"Total":%[2]d}`, runNum, totalNum)
	// json := fmt.Sprintf(`{"Running":%[1]d,"Total":%[2]d}`, 6, 5)
	io.WriteString(w, json)
}

func (cs *CoreServer) handleGetDeviceStatus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	dev, _ := cs.DevManager.GetDevice(id)
	status := cs.DevManager.GetStatus(dev)
	io.WriteString(w, strconv.Itoa(status))
}

func (cs *CoreServer) handleAddExtension(w http.ResponseWriter, r *http.Request) {
	// 创建设备目录和设备运行文件
	file, head, err := r.FormFile("file")
	if err == nil {
		io.WriteString(w, "未上传文件")
	}
	defer file.Close()

	// 读取服务信息
	name := r.Form.Get("ext_name")
	conf := r.Form.Get("conf")
	listeners := r.Form.Get("listeners")
	cmdaddr := r.Form.Get("cmd_addr")
	logurl := r.Form.Get("logurl")
	now := time.Now().Format("2006-01-02 15:04:05")

	ext := &models.Extension{
		ID:            common.GetGUID(),
		Name:          name,
		Conf:          conf,
		FileName:      head.Filename,
		ServerAddress: cmdaddr,
		Listeners:     listeners,
		RegistryTime:  now,
		Type:          "algorithm",
		LogBaseURL:    logurl}

	if err := cs.ExtManager.AddExtension(file, ext); err == nil {
		io.WriteString(w, "success.")
	} else {
		io.WriteString(w, err.Error())
	}

}

func (cs *CoreServer) handleRemoveExtension(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	if err := cs.ExtManager.RemoveExtension(id); err == nil {
		io.WriteString(w, "success.")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleRunExtension(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	err = cs.ExtManager.Run(id)
	if err == nil {
		io.WriteString(w, "running")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleCountExtension(w http.ResponseWriter, r *http.Request) {
	exts, _ := cs.ExtManager.QueryAllExtension("algorithm")
	var runNum, totalNum int
	for _, e := range exts {
		if e.Status == 2 {
			runNum++
		}
		totalNum++
	}
	json := fmt.Sprintf(`{"Running":%[1]d,"Total":%[2]d}`, runNum, totalNum)
	// json := fmt.Sprintf(`{"Running":%[1]d,"Total":%[2]d}`, 6, 5)
	io.WriteString(w, json)
}

func (cs *CoreServer) handleAllExtension(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// err := r.ParseForm()
	// if err != nil {
	// 	log.Println("解析表单数据失败!")
	// }
	// etype := r.Form.Get("type")
	exts, _ := cs.ExtManager.QueryAllExtension("algorithm")
	// 更新状态
	for index, item := range exts {
		if status := cs.ExtManager.GetStatus(&exts[index]); status != item.Status {
			exts[index].Status = status
			cs.ExtManager.UpdateExtension(&exts[index])
		}
	}
	data, _ := json.Marshal(exts)
	io.WriteString(w, string(data[:]))

}

func (cs *CoreServer) handleStopExtension(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("解析表单数据失败!")
	}
	id := r.Form.Get("id")
	err = cs.ExtManager.Stop(id)
	if err == nil {
		io.WriteString(w, "stopped")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleRunStore(w http.ResponseWriter, r *http.Request) {
	err := cs.StoreManager.Run()
	if err == nil {
		io.WriteString(w, "running")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleStopStore(w http.ResponseWriter, r *http.Request) {
	err := cs.StoreManager.Stop()
	if err == nil {
		io.WriteString(w, "stopped")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleStoreStatus(w http.ResponseWriter, r *http.Request) {
	status := cs.StoreManager.GetStatus()
	switch status {
	case 1:
		io.WriteString(w, "stopped")
	case 2:
		io.WriteString(w, "running")
	default:
		io.WriteString(w, "error")
	}
}

func (cs *CoreServer) handleRunLog(w http.ResponseWriter, r *http.Request) {
	err := cs.LogManager.Run()
	if err == nil {
		io.WriteString(w, "running")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleStopLog(w http.ResponseWriter, r *http.Request) {
	err := cs.LogManager.Stop()
	if err == nil {
		io.WriteString(w, "stopped")
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *CoreServer) handleLogStatus(w http.ResponseWriter, r *http.Request) {
	status := cs.LogManager.GetStatus()
	switch status {
	case 1:
		io.WriteString(w, "stopped")
	case 2:
		io.WriteString(w, "running")
	default:
		io.WriteString(w, "error")
	}
}

func (cs *CoreServer) handleHelp(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./doc/help.md")
	if err != nil {
		io.WriteString(w, "fail")
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	io.WriteString(w, string(data[:]))
}

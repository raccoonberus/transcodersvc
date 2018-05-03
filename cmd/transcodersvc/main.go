package main

import (
	"flag"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	"github.com/racoonberus/transcodersvc"
	"github.com/racoonberus/transcodersvc/command"
	"github.com/racoonberus/transcodersvc/convert"
	"github.com/racoonberus/transcodersvc/entity"
	"github.com/racoonberus/transcodersvc/handler"
	"github.com/racoonberus/transcodersvc/repository"
	"io/ioutil"
)

const ServiceName = "transcodersvc"

var (
	logFile         = flag.String("log-file", "/var/log/"+ServiceName+".log", "-log-file=/some/path/to/file.log - define log file path")
	uploadDir       = flag.String("upload-dir", "/bucket/", "-upload-dir=/some/directory/ - define uploads location")
	needClear       = flag.Bool("clear", false, "-clear=true if you want remove file after conversion")
	tasksBufferSize = flag.Int("tasks-buf-size", 100, "-tasks-buf-size=1000 - set task buffer to 1000")
)

func main() {

	convertionCtx := transcodersvc.VideoConvertContext{}
	convertionCtx.Add(convert.ThreeGPPRule{})
	convertionCtx.Add(convert.AviRule{})
	convertionCtx.Add(convert.VOBRule{})

	taskRepo := repository.TaskRepository{}

	tasks := make(chan entity.Task, 2*(*tasksBufferSize))
	go findTasks(taskRepo, *tasksBufferSize, tasks)
	go processTasks(taskRepo, tasks, runtime.NumCPU(), convertionCtx)

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("POST").Path("/file/").HandlerFunc(handler.SetUploadDir(*uploadDir)(handler.FileUpload))
	router.Methods("POST").Path("/task/").HandlerFunc(handler.TaskCreate)
	router.Methods("GET").Path("/task/{id}").HandlerFunc(handler.TaskInfo)

	http.ListenAndServe(":8080", router)
}

func findTasks(repo repository.TaskRepository, limit int, tasks chan<- entity.Task) {
	for {
		if cap(tasks) < limit {
			for _, t := range repo.Select(limit) {
				tasks <- t
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func processTasks(repo repository.TaskRepository, tasks <-chan entity.Task, concurrency int, ctx transcodersvc.VideoConvertContext) {
	sem := make(chan bool, concurrency)
	for {
		sem <- true
		go func() {
			task := <-tasks

			task.StartedAt = time.Now()
			repo.Update(task)

			cmd, err := ctx.GetCmd(task.Resource.InternalFilename, "mp4")
			if err != nil {
				// TODO: need logger
				// TODO: close task
				return
			}

			out, err := command.ShellExec(cmd)
			if err != nil {
				// TODO: need logger
				// TODO: close task
				return
			}

			task.FinishedAt = time.Now()
			task.Output = out

			res, _ := http.Get(task.CallbackUrl)
			defer res.Body.Close()

			resContent, _ := ioutil.ReadAll(res.Body)
			task.CallbackResponse = string(resContent)
			repo.Update(task)

			<-sem
		}()
	}
}

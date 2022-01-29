package main

import (
	"database/sql"
	"fmt"
	"github.com/aragorn-yang/go-camp-01/dao"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.HandleFunc("/", home)
	dao.Dao.Init()
	defer dao.Dao.Close()

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err.Error())
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-signals:
		fmt.Printf("get signal %s, application is gonna shut down", sig)
		os.Exit(0)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	var users = dao.UserDao
	user, err := users.GetById(1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writer.Write([]byte("<html><body>Not Found</body></html>"))
		}
		return
	}
	writer.Write([]byte("<html><body>" + user.Name + "</body></html>"))
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	Start "rwa/projectfile/internal/handlers/register"
)

// сюда код писать не надо

func main() {
	addr := ":8080"
	h := Start.GetApp()
	fmt.Println("start server at", addr)
	http.ListenAndServe(addr, h)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//if err := graceful.Shutdown(ctx); err != nil {
	//	fmt.Println("server graceful shutdown error:", err)
	//}
}

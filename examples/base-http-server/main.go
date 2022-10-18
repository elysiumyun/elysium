package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/elysiumyun/elysium"
)

func init() {
	elysium.Version.Print()
}
func main() {
	engine := elysium.New()

	engine.GET("/", whoamiHandler)
	engine.GET("/health", func(ctx *elysium.Context) {
		ctx.JSON(
			http.StatusOK,
			elysium.H{
				"service":      "base http server",
				"healthy":      true,
				"timestamp":    time.Now().Format(time.ANSIC),
				"dependencies": []elysium.H{{}},
			},
		)
	})

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: engine,
	}
	log.Printf("Server running at http://%v\n", server.Addr)

	panic(server.ListenAndServe())
}

func whoamiHandler(ctx *elysium.Context) {
	var buffer strings.Builder
	u, _ := url.Parse(ctx.R.URL.String())
	wait := u.Query().Get("wait")
	if len(wait) > 0 {
		duration, err := time.ParseDuration(wait)
		if err == nil {
			time.Sleep(duration)
		}
	}

	hostname, _ := os.Hostname()
	buffer.WriteString(fmt.Sprintf("Hostname: %s\n", hostname))

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			buffer.WriteString(fmt.Sprintf("IP: %s\n", ip))
		}
	}

	buffer.WriteString(fmt.Sprintf("RemoteAddr: %s\n", ctx.R.RemoteAddr))
	ctx.String(http.StatusOK, buffer.String())
	if err := ctx.R.Write(ctx.W); err != nil {
		http.Error(ctx.W, err.Error(), http.StatusInternalServerError)
		return
	}
}

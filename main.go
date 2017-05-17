/*
* @Author: Jim Weber
* @Date:   2017-03-21 13:30:41
* @Last Modified by:   Jim Weber
* @Last Modified time: 2017-03-21 13:37:29
 */

package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
        "time"
)

func serve(w http.ResponseWriter, r *http.Request) {

	log.Printf("%+v", r)
	var respBuf bytes.Buffer
	respBuf.WriteString("Method:")
	respBuf.WriteString(r.Method)
	respBuf.WriteString("\n")
	respBuf.WriteString("Proto:")
	respBuf.WriteString(r.Proto)
	respBuf.WriteString("\n")
	respBuf.WriteString("Headers:")
	for k, v := range r.Header {
		respBuf.WriteString(k)
		respBuf.WriteString(":")
		for _, value := range v {
			respBuf.WriteString(value)
			respBuf.WriteString("\n")
		}
		respBuf.WriteString("\n")
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, respBuf.String())
}

func main() {

	portPtr := flag.String("p", "8999", "Port to Listen on")
        startDelay := flag.Int("d", 0, "Seconds to Delay http server startup")
	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

        time.Sleep(time.Duration(*startDelay) * time.Second) 
	log.Println("Starting HTTP Echo Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", serve)
	http.ListenAndServe(":"+*portPtr, mux)

}

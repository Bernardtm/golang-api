package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "strconv"
        "syscall"
)

func health(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("UP")
        w.WriteHeader(http.StatusOK)
}

func main() {

        // Define command-line flags
        port := flag.Int("port", 8080, "Port to run the API on")
        path := flag.String("path", "", "Path to run the API on")
        pidFile := flag.String("pidfile", "api.pid", "Path to the PID file")
        flag.Parse()

        // Write the PID to the specified file
        err := ioutil.WriteFile(*pidFile, []byte(strconv.Itoa(os.Getpid())), 0644)
        if err != nil {
                fmt.Printf("Error writing PID file: %v\n", err)
                return
        }
        defer os.Remove(*pidFile) // Remove PID file when program exits

        http.HandleFunc("/"+*path+"/", health)
        http.HandleFunc("/"+*path+"/commons/up", health)

        // Start the API server
        addr := fmt.Sprintf(":%d", *port)
        fmt.Printf("Server listening on %s...\n", addr)
        if err := http.ListenAndServe(addr, nil); err != nil {
                fmt.Printf("Error starting server: %v\n", err)
                syscall.Exit(1)
        }
}

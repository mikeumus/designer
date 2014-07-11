package init

import (
  "log"
  "net/http"
)

func init() {
  // Simple static webserver:
  log.Fatal(http.ListenAndServe(":8088", http.FileServer(http.Dir("/home/mikeumus/Documents/MDM/Development/Google/Polymer/PolyGo/mikeDev"))))
}
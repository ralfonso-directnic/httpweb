package main

import (
  "fmt"
  "log"
  "net/http"
  "flag"
  "path/filepath"
  "os/user"
  "strings"
)

var port string
var dir string

func main() {

  
  flag.StringVar(&port,"port","80","Port")
  flag.StringVar(&dir,"dir","./","Directory to serve")

  flag.Parse()


  usr, _ := user.Current()
  udir := usr.HomeDir

  if strings.Contains(dir,"~") { 
  
    dir = strings.Replace(dir,"~/",fmt.Sprintf("%s/",udir),1)
 
  }


  ddir,err2 := filepath.Abs(dir)

  log.Printf("Serving %s\n",ddir)

  if(err2!=nil){
   log.Println(err2)
  }

  fs := http.FileServer(http.Dir(ddir))
  http.Handle("/", fs)

  log.Printf("Listening on port %s and serving %s",port,dir)
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    log.Fatal(err)
  }
}

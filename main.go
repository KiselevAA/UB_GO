package main

import (
  "sciter-sdk"
  "sciter-sdk/window"
  "fmt"
  "win"
  //"io/ioutil"
  //"log"
  //"os"
  //"path/filepath"
  //"strings"
)
type t_resolution struct {
  width int
  height int
}
const vVersion = "0.0.0"

func getResolution() t_resolution {
  var l_resolution t_resolution
  hDC := win.GetDC(0)
  defer win.ReleaseDC(0,hDC)
  l_resolution.width = int(win.GetDeviceCaps(hDC,win.HORZRES))
  l_resolution.height = int(win.GetDeviceCaps(hDC,win.VERTRES))
  return l_resolution
}
func main(){

  res := getResolution();

  rect := sciter.NewRect((res.height/2)-150,(res.width/2)-150,250,250)
  window,windowGenerateionError := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG,rect)

  if windowGenerateionError != nil {
    fmt.Println("failed to generate sciter window ", windowGenerateionError.Error())
  }
  window.LoadFile ("logon_main.html")
  window.SetTitle("UB_GO"+" "+vVersion)
  window.Show()
  window.Run()

}

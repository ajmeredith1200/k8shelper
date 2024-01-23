package linter

import (
  "fmt"
)

func init(){
  var path string
  kubeManifestFiles := make(map[string]string)
  fmt.Println("Manifests path: ")
  fmt.Scan(&path)
  entries, err := os.readDir(path)
  if err != nil{
    fmt.Errorf("Path Not Found")
  }
  for _, err in range kubeManifestFiles{
    counter := 0
    linter(kubeManifestFiles[counter])
    counter ++
  }
}
//TODO this wont work ^ but cant test readDir output on this laptop, 
// you need to get each path as its own index in the map then you can iterate over each 
//manifests at the start of the linter.
func linter(kubeManifestFiles){
  
}

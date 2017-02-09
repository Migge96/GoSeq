package main

import(
  "fmt"
  "io/ioutil"
  "Sequence-analysis/bio"
)

func check(e error){
  if e != nil {
    panic(e)
    }
}

func main() {
  //f, err := os.Open("test.txt")
  dat, err := ioutil.ReadFile("test.txt")
  check(err)
  bio.CountNuc(string(dat))
  fmt.Println("GC-Gehalt:", bio.GCContent(string(dat)))
  fmt.Println("A:", bio.A)
  fmt.Println("T:", bio.T)
  fmt.Println("G:", bio.G)
  fmt.Println("C:", bio.C)
  

}

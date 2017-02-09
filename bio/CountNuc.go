package bio

import (
    "strings"
)
var A,T,G,C int

func CountNuc(param string)  {
    A = strings.Count(param, "A")
    T = strings.Count(param, "T")
    G = strings.Count(param, "G")
    C = strings.Count(param, "C")
    
}

func GCContent(param string) float64 {
        CountNuc(param)
    content := float64(G+C) / float64(A+C+G+T)
    
    return content
}   
  


/*
func main(){
    countNuc("AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGCTTCTGAACTG")
    fmt.Println("A:",a, "T:",t, "G:",g, "C:",c)

}
*/


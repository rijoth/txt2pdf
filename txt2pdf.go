package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"github.com/jung-kurt/gofpdf"
)

// this function is used to check input from user
func getFile() (string, string){
    inputFile := flag.String("i", "", "Input file name (required)")
    outputFile := flag.String("o", "", "Output filename (optional)")

    //parsing flags
    flag.Parse()
    // check if input file name is empty
    if (len(strings.TrimSpace(*inputFile)) == 0) {
        log.Fatalf("Filename is required, please provide filename using -i <filename> !!")
    }
    //check if output filename is empty
    var outF string
    if (len(strings.TrimSpace(*outputFile)) == 0){
        outF = strings.TrimSuffix(*inputFile, filepath.Ext(*inputFile)) + ".pdf"
    } else {
        outF = *outputFile
    }

    return *inputFile, outF

}

// a function to generate pdf
func genPdf(input string, output string) string{
    content, err := ioutil.ReadFile(input)

    if(err != nil){
        log.Fatalf("%s file not found, please make sure that input file is in same directory !",input)
    }
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 14)
    pdf.MultiCell(190, 5, string(content), "0", "0", false)
    _ = pdf.OutputFileAndClose(output)
    return "pdf created !"
}

func main(){
    iFile, oFile := getFile()
    op := genPdf(iFile,oFile)
    fmt.Println(op)
}

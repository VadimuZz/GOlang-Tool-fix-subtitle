// The MIT License (MIT)
// 
// Copyright (c) 2015 Vadim Kudriavtcev [VadimuZ]
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "flag"
  "time"
  "bufio"
  "strings"
  "strconv"
)

var fileName string = "example.str";
var fileOutput string = "_";
var secondsToFix int = 5
var output bool = false

func main() {

  f := flag.String("f","","[required] Path to file with subtitles")
  s := flag.String("s","","[required] Seconds, shift all timeline. It's should be integer, or negative integer")
  o := flag.String("o","","Set \"true\" for output on display")
  n := flag.String("n","","New output file. Default name will be the same file name \"-f\" with prefix _")
    
  flag.Parse()

  fmt.Printf("Tool for fix subtitle v 1.0 by VadimuZ\n")

  flag.Usage = func() {
    fmt.Printf("It's a great tool for fix delay/hurry subtitle. Program can shift time to forward or back, for all expressions.\n")
    fmt.Printf("Where [flags] are:\n")
    flag.PrintDefaults()
  }

  if flag.NFlag() == 0 {
    flag.Usage()
    return
  }

  fileName = *f
  secondsToFix, _ = strconv.Atoi(*s)

  output, _ = strconv.ParseBool(*o)
  if(output == true) {
    output = true
  } else {
    output = false
  }

  if *n != "" && *n != *f  {
    fileOutput = *n
  } else {
    fileOutput = "_" + fileName
  }

  file, err := os.Open(fileName)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  newFile, err := os.Create(fileOutput)
  if err != nil {
    fmt.Println(err)
  }

  if output { 
    fmt.Println("--- --- Output start --- ---")
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {

      line := scanner.Text()

      if strings.Contains(line,"-->") {
        
        arr := strings.Split(line, "-->")

        from := addTime( getTimeArr( arr[0] ), secondsToFix)
        to := addTime( getTimeArr( arr[1] ), secondsToFix)
        if output {
          fmt.Println( from+" --> "+to )
        }
        n, err := io.WriteString(newFile, from+" --> "+to+"\n")
        if err != nil {
          fmt.Println(n, err)
        }
      } else {
        if output {
          fmt.Println(line)
        }
        n, err := io.WriteString(newFile, line+"\n")
        if err != nil {
          fmt.Println(n, err)
        } 
      }
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
  newFile.Close()

  if output { 
    fmt.Println("--- --- Output end --- ---")
  }

  fmt.Println("Done")
}

func getTimeArr(arr string) []string {
  from := strings.TrimSpace(arr)
  fromArr := strings.Split(from, ":")
  fromTempMs := strings.Split(fromArr[2], ",")
  fromArr[2] = fromTempMs[0]
  fromArr = append(fromArr, fromTempMs[1])

  return fromArr
}

func addTime(timeArr []string, additionTime int) string {
  hour, _ := strconv.Atoi(timeArr[0])
  minute, _ := strconv.Atoi(timeArr[1])
  second, _ := strconv.Atoi(timeArr[2])
  millisecond, _ := strconv.Atoi(timeArr[3])

  t := time.Date(2000, time.November, 1, hour, minute, second, millisecond, time.UTC)

  pastTime, _ := time.ParseDuration( strconv.Itoa(additionTime)+"s" );    
  newTimeFrom := t.Add( pastTime )

  return addZero( strconv.Itoa( newTimeFrom.Hour() )) +":"+ addZero(strconv.Itoa( newTimeFrom.Minute() ))+":"+addZero( strconv.Itoa( newTimeFrom.Second() ) )+","+strconv.Itoa( millisecond )
}

func addZero(str string) string {
  if(len(str) == 1) {
    return "0"+str
  } else {
    return str
  }
}
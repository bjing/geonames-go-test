package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Location struct {
  geonameid        int 
  name             string
  asciiname        string
  alternatenames   []string
  latitude         float64
  longitude        float64
  featureClass     string
  featureCode      string
  countryCode      string
  cc2              []string
  admin1Code       string
  admin2Code       string
  admin3Code       string
  admin4Code       string
  population       int
  elevation        int
  dem              int
  timezone         string
  modificationDate string
} 

func fromStringToInt(s string) int {
  if s == "" {
    return 0
  } else {
    res, err := strconv.Atoi(s)
    if err != nil {
      fmt.Println(err)
      os.Exit(2)
    }

    return res
  }
}

func fromStringToFloat(s string) float64 {
  if s == "" {
    return 0.0
  } else {
    res, err := strconv.ParseFloat(s, 64)
    if err != nil {
      fmt.Println(err)
      os.Exit(2)
    }
    return res
  }
}

func main() {
  files, err := ioutil.ReadDir("data/")
  if err != nil {
    log.Fatal(err)
  }

  
  allLocations := []Location{}
  for _, file := range files {
    fmt.Println("Processing file " + file.Name())

    if strings.HasSuffix(file.Name(), ".txt") {
      f, err := os.Open("data/" + file.Name())

      if err != nil {
        log.Fatal(err)
      }

      r := csv.NewReader(f)
      r.Comma = '\t'
      r.Comment = '#'

      records, err := r.ReadAll()

      if err != nil {
        log.Fatal(err)
      }

      for _, record := range records {
        location := Location {
          geonameid        : fromStringToInt(record[0]), 
          name             : record[1],  
          asciiname        : record[2],  
          alternatenames   : strings.Split(record[3], ","),  
          latitude         : fromStringToFloat(record[4]),  
          longitude        : fromStringToFloat(record[5]),  
          featureClass     : record[6],  
          featureCode      : record[7],  
          countryCode      : record[8],  
          cc2              : strings.Split(record[9], ","),  
          admin1Code       : record[10],  
          admin2Code       : record[11],  
          admin3Code       : record[12],  
          admin4Code       : record[13],  
          population       : fromStringToInt(record[14]),  
          elevation        : fromStringToInt(record[15]),  
          dem              : fromStringToInt(record[16]),  
          timezone         : record[17],  
          modificationDate : record[18], 
        }

        allLocations = append(allLocations, location)
      }
      fmt.Println()
    } else {
      msg := file.Name() + " is not a data file" 
      fmt.Println(msg)
    }
    fmt.Println("Total number of locations: ", len(allLocations))
    fmt.Println("First location: ", allLocations[0])
  }
}


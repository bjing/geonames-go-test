package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Geonames location data type
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
  defaultValue := 0
  if s == "" {
    return defaultValue
  } else {
    res, err := strconv.Atoi(s)
    if err != nil {
      fmt.Println(err)
      return defaultValue
    }

    return res
  }
}

func fromStringToFloat(s string) float64 {
  defaultValue := 0.0
  if s == "" {
    return defaultValue
  } else {
    res, err := strconv.ParseFloat(s, 64)
    if err != nil {
      fmt.Println(err)
      return defaultValue
    }
    return res
  }
}

/*
** Convert a csv record into a Geonames Location type
*/
func handleRecord(record []string) Location {
  return Location {
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
}

/*
** Get a list of data files
*/
func getDataFiles() []fs.FileInfo {
    // Get all the file names from data/
  files, err := ioutil.ReadDir("data/")
  if err != nil {
    log.Fatal(err)
  }
  return files
}

/*
** Load locations from a single data file
*/
func loadLocationsFromDataFile(fileName string) []Location {
  f, err := os.Open("data/" + fileName)
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

  locations := []Location{}
  for _, record := range records {
    locations = append(locations, handleRecord(record))
  }
  return locations
}


func main() {
  dataFiles := getDataFiles()
  
  allLocations := []Location{}
  for _, file := range dataFiles {
    fmt.Println("Processing file " + file.Name())

    if strings.HasSuffix(file.Name(), ".txt") {
      locations := loadLocationsFromDataFile(file.Name())
      allLocations = append(allLocations, locations...)
      fmt.Println()
    } else {
      msg := file.Name() + " is not a data file" 
      fmt.Println(msg)
    }
  }

  fmt.Println("Total number of locations: ", len(allLocations))
  fmt.Println("First location: ", allLocations[0])
}


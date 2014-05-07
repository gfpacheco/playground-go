package main

import (
    "fmt"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

type Beer struct {
    Name, Description, Country string
    Alcohol float64
}

func main() {
    session, err := mgo.Dial("localhost")

    if err != nil {
        panic(err)
    }

    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    beers := session.DB("workshop-floripa").C("beers")

    var result []Beer

    err = beers.Find(bson.M{}).All(&result)

    if err != nil {
        panic(err)
    }

    fmt.Println(result)
}

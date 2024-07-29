package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
)

func main() {
	//UUID v1
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("UUID v1", uuidV1)

	//UUID v4
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("UUID v4", uuidV4)

	//ULID
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	ulid, _ := ulid.New(ulid.Timestamp(t), entropy)
	fmt.Println("ULID   ", ulid)

	//XID
	xid := xid.New()
	fmt.Println("XID    ", xid)
}

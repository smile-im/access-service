package test

import (
	"log"
	"testing"

	"github.com/smile-im/microkit-client/client/access"
	"github.com/smile-im/microkit-client/proto/accesspb"
)

var (
	cl accesspb.AccessClient
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var err error
	cl, err = access.NewClient()
	if err != nil {
		log.Panicln(err)
	}
	m.Run()
}

package tests

import (
	"github.com/zhifeiji/leanote/app/db"
	"testing"
	//	. "github.com/zhifeiji/leanote/app/lea"
	//	"github.com/zhifeiji/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}

package itime

import (
	"go_trade/src/infa/ops"
	"log"
	"testing"
)

func Test_TimeFunc(t *testing.T) {
	s := "2026-06-24 11:11:11"

	t1, err := Parse(s)
	ops.MustNoErr(err)
	log.Printf("%+v", t1)
}
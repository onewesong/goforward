package models

import (
	"fmt"
	"strings"
)

type ForwardLink struct {
	ListenAddr Addr
	TargetAddr Addr
}

type ForwardLinks []ForwardLink

func NewForwardLink(f string) (fl ForwardLink, err error) {
	fs := strings.Split(f, "->")
	if len(fs) != 2 {
		err = fmt.Errorf("invalid forward link: %s", f)
		return
	}
	fl.ListenAddr, err = ParseAddr(fs[0])
	if err != nil {
		return
	}
	fl.TargetAddr, err = ParseAddr(fs[1])
	if err != nil {
		return
	}
	return
}

func NewForwardLinks(f string) (l ForwardLinks, err error) {
	var fl ForwardLink
	for _, i := range strings.Split(f, ",") {
		if fl, err = NewForwardLink(i); err != nil {
			return
		} else {
			l = append(l, fl)
		}
	}
	return
}

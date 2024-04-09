package engine

import "crypto/md5"

type Name struct {
	Hash string
}

func (a *Name) Hashed(md5name string) {
	h := md5.New()
	h.Write([]byte(md5name))
	a.Hash = string(h.Sum(nil))
}

package ast_test

import (
	"git.sr.ht/~nelsam/hel/pkg/pers"
	"github.com/poy/onpar/matchers"
)

var (
	not              = matchers.Not
	haveOccurred     = matchers.HaveOccurred
	haveLen          = matchers.HaveLen
	equal            = matchers.Equal
	beNil            = matchers.BeNil
	containSubstring = matchers.ContainSubstring
	endWith          = matchers.EndWith

	haveMethodExecuted = pers.HaveMethodExecuted
	returning          = pers.Returning
	storeArgs          = pers.StoreArgs
	within             = pers.Within
)

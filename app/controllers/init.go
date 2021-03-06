package controllers

import "github.com/revel/revel"

func init() {
        revel.OnAppStart(Init)
        revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
        revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
        revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

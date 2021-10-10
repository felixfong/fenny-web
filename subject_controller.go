package main

import "fenng-web/framework"

func SubjectListController(ctx *framework.Context) error {
	ctx.Json(200, "ok, subjectListController")
	return nil
}

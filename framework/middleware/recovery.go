package middleware

import "fenng-web/framework"

func Recovery() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		defer func() {
			if p := recover(); p != nil {
				ctx.Json(500, p)
			}
		}()
		ctx.Next()

		return nil
	}
}

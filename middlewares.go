package main

import "github.com/gin-gonic/gin"

func setAcceptLanguageToIndonesia(ctx *gin.Context) {
	header := map[string][]string{
		"Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"id-ID,id;q=0.5"},
	}

	ctx.Request.Header = header
	ctx.Next()
}

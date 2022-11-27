package handlers

import (
	"net/http"

	"github.com/aZ4zil/go_apotek/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetAllTransactions(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")
	if user == nil {
		flasher.Set("info", "Mohon login terlebih dahulu Anda sebelum mengakses halaman ini.")
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	if ctx.Request.Method == "GET" {
		defer flasher.Del()

		userID := GetSessionValue(user)["id"].(uint)
		transactions := models.GetAllTransactionByUserID(userID)
		ctx.HTML(http.StatusOK, "get_all_transactions", gin.H{
			"flasher":      flasher,
			"user":         user,
			"transactions": transactions,
		})
		return
	}
}

func CreateNewTransaction(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")
	if user == nil {
		flasher.Set("info", "Mohon login terlebih dahulu sebelum Anda mengakses halaman ini.")
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	if ctx.Request.Method == "GET" {
		defer flasher.Del()

		// userID := GetSessionValue(user)["id"].(uint)
		drugs := models.GetAllGrugs()

		ctx.HTML(http.StatusOK, "create_transactions", gin.H{
			"flasher": flasher,
			"user":    user,
			"drugs":   drugs,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		transactionRequest := &Transaction{}

		err := ctx.ShouldBindWith(transactionRequest, binding.FormPost)
		if err != nil {
			flasher.Set("danger", err.Error())
			ctx.Redirect(http.StatusFound, "/transactions/pay")
			return
		}

		transaction := models.Transaction{
			DrugID:      transactionRequest.DrugID,
			UserID:      transactionRequest.UserID,
			TotalGoods:  transactionRequest.TotalGoods,
			TotalPay:    transactionRequest.TotalPay,
			TotalPrice:  transactionRequest.TotalPrice,
			TotalRefund: transactionRequest.TotalRefund,
		}

		err = models.CreateNewTransaction(&transaction)
		if err != nil {
			flasher.Set("danger", "Mohon maaf proses yang Anda minta tidak dapat kami akses. Mohon tunggu sebentar atau laporkan masalah ini pada pihak layanan.")
			ctx.Redirect(http.StatusFound, "/transactions/pay")
			return
		}

		flasher.Set("success", "Permintaan berhasil kami proses, barang sedang di kemas.")
		ctx.Redirect(http.StatusFound, "/transactions/pay")
		return
	}
}

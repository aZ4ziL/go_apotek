package handlers

func GetSessionValue(s interface{}) map[string]interface{} {
	values, _ := s.(map[string]interface{})
	result := make(map[string]interface{})

	for k, v := range values {
		result[k] = v
	}
	return result
}

type Transaction struct {
	DrugCode    string `json:"drug_code" form:"drug_code"`
	UserID      uint   `json:"user_id" form:"user_id"`
	TotalGoods  uint   `json:"total_goods" form:"total_goods"`
	TotalPay    uint   `json:"total_pay" form:"total_pay"`
	TotalPrice  uint   `json:"total_price" form:"total_price"`
	TotalRefund uint   `json:"total_refund" form:"total_refund"`
}

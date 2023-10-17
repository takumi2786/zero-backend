package controller

/* Controllerの出力データの構造を定義します。*/

// ErrorResponseは、エラー時レスポンスの構造を定義します。
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccessResponseは、成功時レスポンスの構造を定義します。
type SuccessResponse struct {
	Code int `json:"code"`
	Body interface{}
}

package constants

type ReCode struct {
	Code    int            `json:"code"`
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

func (r *ReCode) Ok() {
	r.Code = Success
	r.Success = true
	r.Message = "成功"
}

func (r *ReCode) Err() {
	r.Code = Error
	r.Success = false
	r.Message = "失败"
}

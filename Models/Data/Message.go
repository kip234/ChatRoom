package Data

type Message struct {
	Typ int `json:"typ" binding:"required"`
	Owner int `json:"owner" binding:"required"`//所有者
	PoolName string `json:"poolname"`//来自某个群？
	Content []byte `json:"content" binding:"required"`
}

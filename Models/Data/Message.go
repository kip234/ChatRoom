package Data

type Message struct {
	Typ int `json:"typ" binding:"required"`
	PoolName string `json:"poolname"`//来自某个群？
	Content []byte `json:"content" binding:"required"`
}

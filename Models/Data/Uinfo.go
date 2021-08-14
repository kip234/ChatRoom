package Data

type Uinfo struct {
	Uid int			`json:"uid"`//因为Message含有Owner字段所以这里就不用了
	Name string		`json:"name"`
}

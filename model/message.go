package model

//详细信息
type Details struct {
	Name          string
	Age           string
	Sex           string
	Height        string
	MonthlyIncome string // 月收入
	Marriage      string //婚姻状况
	Education     string //教育程度
	WorkLocation  string //工作地
	Occupation    string //职业
	Constellation string //星座
	Birthplace    string
	Weight        string
}

//生活状况
type LivingCondition struct {
	House string
	Car   string
	Smoke string
	Drink string
}

//择偶条件
type SpouseCondition struct {
	Sex           string
	Age           string
	Height        string
	Education     string
	MonthlyIncome string
	Marriage      string
	WorkLocation  string
}

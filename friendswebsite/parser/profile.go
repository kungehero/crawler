package parser

import (
	"crawler-go/engine"
	"crawler-go/model"
	"regexp"
)

var (
	//Id            = regexp.MustCompile(`<p class="brief-info fs14 lh32 c9f">ID：([^<]+)<span class="brief-item">诚信值：<span class="red"></span></span>`)
	Name          = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
	Age           = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+)</td>`)
	Sex           = regexp.MustCompile(`<td><span class="label">性别：</span>([^<]+)</td>`)
	Height        = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
	MonthlyIncome = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	Marriage      = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	Education     = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	WorkLocation  = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
	Occupation    = regexp.MustCompile(`<td><span class="label">职业：</span>([^<]+)</td>`)
	Constellation = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
	Birthplace    = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	Weight        = regexp.MustCompile(`<td><span class="label">体重：</span>([^<]+)</td>`)

	/* House = regexp.MustCompile(`<td><span class="label">住房条件：</span>([^<]+)</td>`)
	Car   = regexp.MustCompile(`<td><span class="label">是否购车：</span>([^<]+)</td>`)
	Smoke = regexp.MustCompile(`<td><span class="label">是否吸烟：</span>([^<]+)</td>`)
	Drink = regexp.MustCompile(`<td><span class="label">是否喝酒：</span>([^<]+)</td>`)

	SpouseSex           = regexp.MustCompile(`<td><span class="label">性别：</span>([^<]+)</td>`)
	SpouseAge           = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+)</td>`)
	SpouseHeight        = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
	SpouseEducation     = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	SpouseMonthlyIncome = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	SpouseMarriage      = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	SpouseWorkLocation  = regexp.MustCompile(`<td><span class="label">工作地区：</span>([^<]+)</td>`) */
)

func ParserProfileMessage(urls []byte) engine.ParserResult {
	//sId := strings.TrimSpace(ExpanRegexp(urls, Id))
	sName := ExpanRegexp(urls, Name)
	sAge := ExpanRegexp(urls, Age)
	sSex := ExpanRegexp(urls, Sex)
	sHeight := ExpanRegexp(urls, Height)
	sMonthlyIncome := ExpanRegexp(urls, MonthlyIncome)
	sMarriage := ExpanRegexp(urls, Marriage)
	sEducation := ExpanRegexp(urls, Education)
	sWorkLocation := ExpanRegexp(urls, WorkLocation)
	sOccupation := ExpanRegexp(urls, Occupation)
	sConstellation := ExpanRegexp(urls, Constellation)
	sBirthplace := ExpanRegexp(urls, Birthplace)
	sWeight := ExpanRegexp(urls, Weight)

	/* sHouse := ExpanRegexp(urls, House)
	sCar := ExpanRegexp(urls, Car)
	sSmoke := ExpanRegexp(urls, Smoke)
	sDrink := ExpanRegexp(urls, Drink)

	sSpouseSex := ExpanRegexp(urls, SpouseSex)
	sSpouseAge := ExpanRegexp(urls, SpouseAge)
	sSpouseHeight := ExpanRegexp(urls, SpouseHeight)
	sSpouseEducation := ExpanRegexp(urls, SpouseEducation)
	sSpouseMonthlyIncome := ExpanRegexp(urls, SpouseMonthlyIncome)
	sSpouseMarriage := ExpanRegexp(urls, SpouseMarriage)
	sSpouseWorkLocation := ExpanRegexp(urls, SpouseWorkLocation) */

	profile := model.Details{}
	//profile.Id = sId
	profile.Name = sName
	profile.Age = sAge
	profile.Sex = sSex
	profile.Height = sHeight
	profile.MonthlyIncome = sMonthlyIncome
	profile.Marriage = sMarriage
	profile.Education = sEducation
	profile.WorkLocation = sWorkLocation
	profile.Occupation = sOccupation
	profile.Birthplace = sBirthplace
	profile.Weight = sWeight
	profile.Constellation = sConstellation
	result := engine.ParserResult{Items: []interface{}{profile}}
	return result
}

func ExpanRegexp(page []byte, re *regexp.Regexp) string {
	addr := re.FindSubmatch(page)
	if len(addr) >= 2 {
		return string(addr[1])
	} else {
		return ""
	}
}

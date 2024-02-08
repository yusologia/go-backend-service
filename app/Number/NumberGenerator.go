package Number

import (
	"gorm.io/gorm/schema"
	"service/App/Repository"
	"service/App/Service/Helper"
	"strconv"
	"strings"
	"time"
)

type NumberGeneratorInterface interface {
	Generate() string
	Prefix() string
}

type NumberGenerator struct {
	Model schema.Tabler
}

func (nm NumberGenerator) Generate(generator NumberGeneratorInterface) string {
	var number string

	increment := Repository.GetIncrementMonthly(nm.Model)

	number += time.Now().Format("010602")
	number += Helper.StrPadLeft(strconv.Itoa(int(increment)), 4, '0')
	number += strconv.Itoa(Helper.RandInt(111, 999))

	return strings.ToUpper(generator.Prefix() + number)
}

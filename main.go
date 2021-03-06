package main

import (
	"fmt"
	"github.com/arshiamidos/generator_polynomial_galois/gpg"
	"github.com/arshiamidos/qr/utils"
	"github.com/arshiamidos/qr_presenter/painter"
)

func main() {
	v := 1
	e := 'L'
	t := 'B'
	m := utils.Byte{"arash midos 666"}
	

	GPG1 := gpg.NewAntiLog(
		utils.ConvertToMessagePoly(
			utils.BreakUp8Bit(m.Data, m.Parse(), v, e, t)),
	)

	GPG1.SetGroupBlock(
		utils.BreakUp8Bit(m.Data, m.Parse(), v, e, t),
		utils.GetGroupBlock(v, e))

	GPG1.SetGroupBlockECC(
		utils.BreakUp8Bit(m.Data, m.Parse(), v, e, t),
		utils.GetGroupBlock(v, e),
		utils.GetECCWCount(v, e),
	)
	s := GPG1.Serialize(utils.GetGroupBlock(v, e)["GROUP1"][0]+utils.GetGroupBlock(v, e)["GROUP2"][0], v)
	fmt.Println(s)
	fmt.Println(GPG1.GroupBlockEC)
	painter.PaintV(v,e, 10, s)
}

package main

import (
	"emvco-parser/m/v2/emvco"
	"log"
)

func main() {
	emv := emvco.NewEMVCo()

	// log.Println(emv.Parse("00020101021126720020COM.BANKSINARMAS.WWW011893600153020185390502150020000000008540303UBE51440014ID.CO.QRIS.WWW0215ID20200253932040303UBE5204599953033605802ID5910RALALI.COM6013JAKARTA BARAT61051161062070703A0163041DA4"))
	// log.Println(emv.ToString())
	log.Println(emv.Parse("00020101021238560010A0000007270126000697041501121133666688880208QRIBFTTA53037045802VN630441CA"))
	log.Println(emv.ToString())
}

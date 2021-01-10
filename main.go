package main

import (
	"fmt"

	"github.com/yuanyu90221/golang_adapter_pattern_sample/printers"
)

func main() {
	msg := "Message Test"
	returnedMsg := printers.PrinterAdapter{OldPrinter: &printers.MyLegacyPrinter{}, Msg: msg}
	fmt.Println("with LegacyPrinter:", returnedMsg.PrintStored())
	returnedMsg = printers.PrinterAdapter{OldPrinter: nil, Msg: msg}
	fmt.Println("without LegacyPrinter:", returnedMsg.PrintStored())
}

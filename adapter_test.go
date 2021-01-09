package golang_adapter_pattern_sample

import (
	"testing"

	"github.com/yuanyu90221/golang_adapter_pattern_sample/printers"
)

func TestPrinterAdapter_PrintStored_test(t *testing.T) {
	msg := "Hello World!"
	adapter := printers.PrinterAdapter{OldPrinter: &printers.MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
	adapter = printers.PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}

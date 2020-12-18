package main

import (
	"fmt"
	"testing"
)

func TestModes(t *testing.T) {
	for m := AM_MODE; m <= THRBX_MODE; m++ {
		var i interface{} = m.String()
		r, ok := i.(string)
		if !ok {
			t.Errorf("ERROR: Expected type string got %v", r)
		}
	}
}

func TestImportOnly(t *testing.T) {

	for m := AM_MODE; m <= WSPR_MODE; m++ {
		if m.isImportOnly() {
			t.Errorf("ERROR: Did not expect %v to be Import Only", m)
		}
	}

	for m := AMTORFEC_MODE; m <= THRBX_MODE; m++ {
		if !m.isImportOnly() {
			t.Errorf("ERROR: Expected %v to be Import Only", m)
		}
	}

}

func TestSubmodes(t *testing.T) {

	for s := ASCI_SUBMODE; s <= USB_SUBMODE; s++ {
		var i interface{} = s.String()
		r, ok := i.(string)
		if !ok {
			t.Errorf("ERROR: Expected type string got %v", r)
		}
	}

	for m := AMTORFEC_MODE; m <= THRBX_MODE; m++ {
		submodes := m.Submodes()
		for _, submode := range submodes {
			fmt.Println(submode.String())
		}
	}

}

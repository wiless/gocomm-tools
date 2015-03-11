package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/qml.v1"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type InputPins struct {
	Names []string
	Types []string
}

type ChipInfo struct {
	Name        string
	InPinCount  float32
	OutPinCount float64
	ModuleCount float32
	Message     string
	ShowDialog  bool
	InPins      InputPins
}

func (i *InputPins) Update(count int, val string) {
	i.Names[count] = val
}

func (i *InputPins) Create(count int) {
	i.Names = make([]string, count)

	i.Types = make([]string, count)
	for cnt := 0; cnt < count; cnt++ {
		i.Names[cnt], i.Types[cnt] = fmt.Sprintf("Pin-%d", cnt), "BitChannel"
	}

}

func (i *InputPins) Name(index int) string {
	return i.Names[index]
}

func (i *InputPins) Type(index int) string {
	return i.Types[index]
}

func NewChipInfo() ChipInfo {
	var result ChipInfo
	result.InPinCount = 1
	result.InPins.Create(int(result.InPinCount))
	result.OutPinCount = 1
	result.ModuleCount = 1
	result.Name = "DummyChip"
	result.ShowDialog = false
	result.Message = "There is no message to co65mmunicate"
	return result
}

func (c *ChipInfo) CreateInputPins(cnt int) {
	c.InPinCount = float32(cnt)
	c.InPins.Create(cnt)
	qml.Changed(c, &c.InPinCount)
}

func (c *ChipInfo) Generate() {
	log.Printf("I will now generate data  :  %#v", c)

	if c.ModuleCount == 2 {
		c.ShowDialog = true
		qml.Changed(c, &c.ShowDialog)
	}

	output, err := json.Marshal(c)
	fmt.Printf("Err JSON : %v", err)
	fmt.Printf("Output JSON : %s", output)
	/// Generate the JSON file ..

}

func run() error {
	engine := qml.NewEngine()
	context := engine.Context()
	defaultchip := NewChipInfo()
	controls, err := engine.LoadFile("chipinfo.qml")
	context.SetVar("chipInfo", &defaultchip)
	context.SetVar("inputPins", &defaultchip.InPins)

	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

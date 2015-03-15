package main

import (
	"fmt"
	"os"

	"github.com/wiless/gocomm-tools"

	"gopkg.in/qml.v1"
)

var newqmlchip qmlChip
var inpins qmlInPins //(gocommui.DefaultInputPins(1, 0))

func main() {

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// type InputPins struct {
// 	Names []string
// 	Types []string
// }

// type ChipInfo struct {
// 	Name        string
// 	InPinCount  float32
// 	OutPinCount float64
// 	ModuleCount float32
// 	Message     string
// 	ShowDialog  bool
// 	InPins      InputPins
// }

// func (i *InputPins) Update(count int, val string) {
// 	i.Names[count] = val
// }

// func (i *InputPins) Create(count int) {
// 	i.Names = make([]string, count)

// 	i.Types = make([]string, count)
// 	for cnt := 0; cnt < count; cnt++ {
// 		i.Names[cnt], i.Types[cnt] = fmt.Sprintf("Pin-%d", cnt), "BitChannel"
// 	}

// }

// func (i *InputPins) Name(index int) string {
// 	return i.Names[index]
// }

// func (i *InputPins) Type(index int) string {
// 	return i.Types[index]
// }

// func NewChipInfo() ChipInfo {
// 	var result ChipInfo
// 	result.InPinCount = 1
// 	result.InPins.Create(int(result.InPinCount))
// 	result.OutPinCount = 1
// 	result.ModuleCount = 1
// 	result.Name = "DummyChip"
// 	result.ShowDialog = false
// 	result.Message = "There is no message to co65mmunicate"
// 	return result
// }

// func (c *ChipInfo) CreateInputPins(cnt int) {
// 	c.InPinCount = float32(cnt)
// 	c.InPins.Create(cnt)
// 	qml.Changed(c, &c.InPinCount)
// }

// func (c *ChipInfo) Generate() {
// 	log.Printf("I will now generate data  :  %#v", c)

// 	if c.ModuleCount == 2 {
// 		c.ShowDialog = true
// 		qml.Changed(c, &c.ShowDialog)
// 	}

// 	output, err := json.Marshal(c)
// 	fmt.Printf("Err JSON : %v", err)
// 	fmt.Printf("Output JSON : %s", output)
// 	/// Generate the JSON file ..

// }

type qmlInPins struct {
	data []gocommui.JsonPin
}

// func (q *qmlChip) SetInPinCount(val int) {
// 	// Overloading
// 	//fmt.Println("Overloading works")
// 	(*gocommui.JsonChip)(q).SetInPinCount(value)

// 	q.InPinCount = val
// }

type qmlChip gocommui.JsonChip

func CreateQmlChip() qmlChip {
	result := qmlChip(gocommui.NewJsonChip())
	result.Name = "Modem"

	return result
}

// func UpdateInputTab(value int) {
// 	newqmlchip.InPinCount = value
// 	inpins.Resize(value)

// 	// (*gocommui.JsonChip)(q).SetInPinCount(value)
// 	// q.Resize(value)
// 	qml.Changed(newqmlchip, &newqmlchip.InPinCount)
// }

func (q *qmlInPins) Create(cnt int) {
	(*q).data = gocommui.DefaultInputPins(cnt, 0)
	// (*q).data=gocommui.De
}

func (q *qmlInPins) Resize(cnt int) {
	fmt.Printf("\n before %v ", q.data)
	/// truncate
	if len(q.data) > cnt {
		q.data = q.data[0:cnt]
	}

	/// extend
	if len(q.data) < cnt {
		q.data = append(q.data, gocommui.DefaultInputPins(cnt-len(q.data), len(q.data))...)
	}
	newqmlchip.InPinCount = cnt
	fmt.Printf("\n after %v", q.data)
	qml.Changed(&newqmlchip, &newqmlchip.InPinCount)

}

func (q *qmlInPins) Name(index int) string {
	if len(q.data) <= index {
		return ""
	}
	return (*q).data[index].Name
}

func (q *qmlInPins) Update(index int, value string) {
	fmt.Println("Updated the name ")
	if len(q.data) < index {
		return
	}

	(*q).data[index].Name = value
}

// func (j *qmlChip) InPinName(cnt int) string {
// 	fmt.Printf("You asked for and  you got it")
// 	if cnt < len(j.Pins) {
// 		return j.Pins[cnt].Name
// 	} else {
// 		return ""
// 	}
// }

func run() error {

	engine := qml.NewEngine()
	context := engine.Context()
	// defaultchip := gocommui.NewJsonChip()

	newqmlchip = CreateQmlChip()

	// newqmlchip.JsonChip = qmlChip(gocommui.NewJsonChip())

	controls, err := engine.LoadFile("chipinfo.qml")
	context.SetVar("chipInfo", &newqmlchip)
	context.SetVar("inputPins", &inpins)

	// context.SetVar("inputPins", &defaultchip.InPins)

	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

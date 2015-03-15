package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wiless/gocomm-tools"

	"gopkg.in/qml.v1"
)

var newqmlchip qmlChip
var inpins qmlInPins
var outpins qmlOutPins
var modules qmlModules

func main() {

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type qmlInPins struct {
	data []gocommui.JsonPin
}

type qmlOutPins struct {
	data []gocommui.JsonPin
}

func (q *qmlOutPins) Create(cnt int) {
	(*q).data = gocommui.DefaultPins(cnt, 0, true)
	// (*q).data=gocommui.De
}

func (q *qmlOutPins) Resize(cnt int) {

	/// truncate
	if len(q.data) > cnt {
		q.data = q.data[0:cnt]
	}

	/// extend
	if len(q.data) < cnt {
		q.data = append(q.data, gocommui.DefaultPins(cnt-len(q.data), len(q.data), false)...)
	}
	newqmlchip.OutPinCount = cnt
	qml.Changed(&newqmlchip, &newqmlchip.OutPinCount)

}

func (q *qmlOutPins) Name(index int) string {
	if len(q.data) <= index {
		return ""
	}
	return (*q).data[index].Name
}

func (q *qmlOutPins) Update(index int, value string) {
	fmt.Println("Updated the name ")
	if len(q.data) < index {
		return
	}

	(*q).data[index].Name = value
}

type qmlChip gocommui.JsonChip

func (q *qmlChip) Generate() {

	q.Pins = append(inpins.data, outpins.data...)
	q.Modules = modules.data
	for i := 0; i < len(modules.data); i++ {
		q.Modules[i].InPins = strings.Split(modules.Inames[i], ",")
		q.Modules[i].OutPins = strings.Split(modules.Onames[i], ",")
	}

	// jsondata, _ := json.Marshal(*q)
	// fmt.Printf("\n %s", jsondata)
}

func (q *qmlChip) SaveAs(fname string) {

	q.Generate()

	jsondata, _ := json.Marshal(*q)
	// strings.HasPrefix(fname, "file://")
	fname = strings.TrimPrefix(fname, "file://")
	w, err := os.Create(fname)
	if err != nil {
		log.Printf("Unable to save to file %v, \n Error = %v", fname, err)
		return
	}
	fmt.Fprintf(w, "%s", jsondata)
}

func CreateQmlChip() qmlChip {
	result := qmlChip(gocommui.NewJsonChip())
	result.Name = "NewChip"
	return result
}

func (q *qmlInPins) Create(cnt int) {
	(*q).data = gocommui.DefaultPins(cnt, 0, true)

}

func (q *qmlInPins) Resize(cnt int) {
	/// truncate
	if len(q.data) > cnt {
		q.data = q.data[0:cnt]

	}

	/// extend
	if len(q.data) < cnt {
		q.data = append(q.data, gocommui.DefaultPins(cnt-len(q.data), len(q.data), true)...)
	}
	newqmlchip.InPinCount = cnt
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
	context.SetVar("outputPins", &outpins)
	context.SetVar("modules", &modules)

	// context.SetVar("inputPins", &defaultchip.InPins)

	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

type qmlModules struct {
	data   []gocommui.JsonModule
	Inames []string
	Onames []string
}

func (q *qmlModules) Create(cnt int) {
	(*q).data = gocommui.DefaultModules(cnt, 0)

}

func (q *qmlModules) Resize(cnt int) {
	currentcnt := len(q.data)
	/// truncate
	if currentcnt > cnt {
		q.data = q.data[0:cnt]
		q.Inames = q.Inames[0:cnt]
		q.Onames = q.Onames[0:cnt]

	} else if currentcnt < cnt {

		/// extend

		q.Inames = append(q.Inames, make([]string, cnt-currentcnt)...)
		q.Onames = append(q.Onames, make([]string, cnt-currentcnt)...)
		q.data = append(q.data, gocommui.DefaultModules(cnt-currentcnt, currentcnt)...)

	}
	newqmlchip.ModuleCount = cnt
	qml.Changed(&newqmlchip, &newqmlchip.ModuleCount)

}

func (q *qmlModules) IPins(index int) string {
	if len(q.Inames) <= index {
		return ""
	}
	return (*q).Inames[index]
}
func (q *qmlModules) Name(index int) string {
	if len(q.data) <= index {
		return ""
	}
	return (*q).data[index].Name
}
func (q *qmlModules) OPins(index int) string {
	if len(q.Onames) <= index {
		return ""
	}
	return (*q).Onames[index]
}

func (q *qmlModules) Update(index int, value string) {
	fmt.Println("Updated the name ")
	if len(q.data) < index {
		return
	}

	(*q).data[index].Name = value
}

func (q *qmlModules) UpdatePins(index int, value string, input bool) {
	fmt.Println("Updated the name ")
	if len(q.data) < index {
		return
	}
	if input {
		(*q).Inames[index] = value
	} else {
		(*q).Onames[index] = value
	}

}

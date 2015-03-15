import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
import QtQuick.Dialogs 1.1
import QtQuick.Controls.Styles 1.1


ApplicationWindow {
    visible: true
    title: "Basic layouts"
    property int margin: 11
    width: 517
    height: mainLayout.implicitHeight + 2 * margin
    minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin
    minimumHeight: mainLayout.Layout.minimumHeight + 2 * margin
    statusBar: StatusBar {
Timer {
        id:timer
                   interval: 1000; running:false; repeat: false
                    onTriggered:  mystatusbar.statustext="By @iamssk"
                }

        RowLayout {
            anchors.fill: parent

            Label {
                property string statustext:"By @iamssk"
                id:mystatusbar
                text: statustext
            }
            Text { id: time }
            Item {
                Timer {
                    interval: 500; running: true; repeat: true
                    onTriggered: time.text = Date().toString()
                }


            }
        }
    }
    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        anchors.margins: margin
        GroupBox {
            id: rowBox
            title: "Chip Name"
            Layout.fillWidth: true
            RowLayout {
                id: rowLayout
                anchors.fill: parent
                TextField {
                    id:chipname
                    font.bold: true
                    placeholderText: "This wants to grow horizontally"
                    Layout.fillWidth: true
                    text:chipInfo.name
                }
                Button {
                    id:okbutton
                    text: "Generate"
                    onClicked: {                        
                        inputPins.resize(spinInPinCount.value);
                        outputPins.resize(spinOutPinCount.value);
                        modules.resize(spinModuleCount.value);
                        fileDialog.open()
                        chipInfo.generate();   
                    }
                }
            }
        }

        GroupBox {
            id: gridBox
            title: "Grid layout12"
            Layout.fillWidth: true

            GridLayout {
                id: gridLayout
                rows: 3
                flow: GridLayout.TopToBottom          
                anchors.left: parent.left
                anchors.top: parent.top
                anchors.right: parent.right
                anchors.bottom: parent.bottom
                Label { width:200; id:numInputPins;text: "Input Pins" }
                Label { width:100;id:numOutputPins; text: "Output Pins" }
                Label { width:100;text: "Modules" }
                SpinBox {Layout.fillWidth:true;   id:spinInPinCount;value:chipInfo.inPinCount
                 onEditingFinished:
                 {
                     inputPins.resize(value)    


                 }
             }            
             SpinBox {
                 Layout.fillWidth:true;  id:spinOutPinCount;value:chipInfo.outPinCount
                 onEditingFinished:
                 {
                  outputPins.resize(value)
              }
          }

          SpinBox { Layout.fillWidth:true; id:spinModuleCount;value:chipInfo.moduleCount;minimumValue:1
              onEditingFinished:
              {
                 modules.resize(value)    
             }

         }

         Binding { target:chipInfo; property:"message"; value: messagebox.text }
         // Binding { target:chipInfo; property:"inPinCount"; value: spinInPinCount.value }
         // Binding { target:chipInfo; property:"outPinCount"; value: spinOutPinCount.value }
         // Binding { target:chipInfo; property:"moduleCount"; value: spinModuleCount.value }
         Binding { target:chipInfo; property:"name"; value: chipname.text }


     }
 }

 SplitView {
    Layout.fillWidth:true
    Layout.fillHeight:true
    // anchors.fill: parent
    orientation: Qt.Vertical

    TabView {
        antialiasing: true
        z: 0
        rotation: 0
        tabsVisible: true
        tabPosition: 1
        frameVisible: false
        opacity: 0.7
        Tab {
            title: "Input Pins"
            //     Rectangle { color: "red" }

            ListView {
                id:inputpinlist
                anchors.bottom: parent.bottom
                anchors.horizontalCenter: parent.horizontalCenter
                width:parent.width
                model: chipInfo.inPinCount
                delegate:
                Row {
                    id: rlayout1
                    // anchors.fill: parent
                    spacing: 6                            

                    TextField {

                        text: inputPins.name(index)
                        style: TextFieldStyle {
                            textColor: "black"
                            background: Rectangle {
                                radius: 5
                                color:"green"
                                implicitWidth: 100
                                implicitHeight: 24
                                border.color: "green"
                                border.width: 1
                            }
                        }

                        onEditingFinished:
                        {
                            console.debug("Calling update fucntion")
                            inputPins.update(index,text)
                        }
                    }
                    TextField { width:100;text: "ComplexBit"}
                }

            }

        }        
        Tab {
            title: "Output Pins"
            ListView {
                id:outputpinlist
                anchors.bottom: parent.bottom
                anchors.horizontalCenter: parent.horizontalCenter
                width:parent.width
                model: chipInfo.outPinCount
                delegate:
                Row {
                    id: rlayout2
                    // anchors.fill: parent
                    spacing: 6                            

                    TextField {

                        text: outputPins.name(index)
                        style: TextFieldStyle {
                            textColor: "black"
                            background: Rectangle {
                                radius: 5
                                color:"red"
                                implicitWidth: 100
                                implicitHeight: 24
                                border.color: "green"
                                border.width: 1
                            }
                        }

                        onEditingFinished:
                        {
                            console.debug("Calling update fucntion")
                            outputPins.update(index,text)
                        }
                    }
                    TextField {
                        width:100

                        text: "ComplexBit"                                

                        
                    }   
                }

            }
        }
        Tab {
            title: "Modules"                            
            ListView {
                id:modulelist
                anchors.bottom: parent.bottom
                anchors.horizontalCenter: parent.horizontalCenter
                width:parent.width
                model: chipInfo.moduleCount
                Row {

                   Label{text:"Name"}
                   Label{text:"InputPins"}
                   Label{text:"OutputPins"}
               }                       
               delegate:
               Row {
                id: rlayout3
                // anchors.fill: parent
                spacing: 6                            

                TextField {

                    text: modules.name(index)
                    style: TextFieldStyle {
                        textColor: "black"
                        background: Rectangle {
                            radius: 5
                            color:"yellow"
                            implicitWidth: 100
                            implicitHeight: 24
                            border.color: "blue"
                            border.width: 1
                        }
                    }
                    onEditingFinished:
                    {
                        console.debug("Calling Module fucntion")
                        modules.update(index,text)
                    }
                }
                TextField {
                    // width:100
                    text: modules.iPins(index)
                    onEditingFinished:
                    {
                        console.debug("Calling Module fucntion")
                        modules.updatePins(index,text,true)
                    }

                }   
                TextField {
                    // width:100
                    text: modules.oPins(index)                                
                    onEditingFinished:
                    {
                        console.debug("Calling Module fucntion")
                        modules.updatePins(index,text,false)
                    }
                }   
            }

        }

    }
}
TextArea {
    id:messagebox 
    text:"chipInfo.message "
    Layout.minimumHeight: 30
    Layout.fillHeight: true
    Layout.fillWidth: true
}

}


}

MessageDialog {
    id: messageDialog
    visible:false //chipInfo.showDialog

    title: "May I have your attention please"
    text: "It's so cool that you are using Qt Quick."
    onAccepted: {
        console.log("And of course you could only agree.")
        // Qt.quit()
    }
    // Component.onCompleted: visible = true
}

FileDialog {
    id: fileDialog
    title: "Save As"
    visible:false
    nameFilters : ["Json Files (*.json)"]
    selectFolder:false
    selectMultiple:false
    onAccepted: {
        console.log("You chose: " + fileDialog.fileUrl)
        timer.running=true
        mystatusbar.statustext="Saved to "+fileDialog.fileUrl
        


        chipInfo.saveAs(fileDialog.fileUrl)


        // Qt.quit()
    }
    
    onRejected: {
        console.log("Canceled")
        // Qt.quit()
    }
    // Component.onCompleted: visible = true
}
}




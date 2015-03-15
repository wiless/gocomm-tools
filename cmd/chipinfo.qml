import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
import QtQuick.Dialogs 1.1
import QtQuick.Controls.Styles 1.1


ApplicationWindow {
    visible: true
    title: "Basic layouts"
    property int margin: 11
    width: mainLayout.implicitWidth + 2 * margin
    height: mainLayout.implicitHeight + 2 * margin
    minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin
    minimumHeight: mainLayout.Layout.minimumHeight + 2 * margin

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
                        // chipInfo.generate();   
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
                SpinBox {Layout.fillWidth:true;  id:spinOutPinCount;value:chipInfo.outPinCount}
                SpinBox { Layout.fillWidth:true; id:spinModuleCount;value:chipInfo.moduleCount;minimumValue:1}

                Binding { target:chipInfo; property:"message"; value: messagebox.text }
                // Binding { target:chipInfo; property:"inPinCount"; value: spinInPinCount.value }
                Binding { target:chipInfo; property:"outPinCount"; value: spinOutPinCount.value }
                Binding { target:chipInfo; property:"moduleCount"; value: spinModuleCount.value }
                Binding { target:chipInfo; property:"name"; value: chipname.text }
                

                // TextArea {
                    //     text: "ABA This widget spans over three rows in the GridLayout.\n"
                    //         + "All items in the GridLayout are implicitly positioned from top to bottom."
                    //     Layout.rowSpan: 3
                    //     Layout.fillHeight: true
                    //     Layout.fillWidth: true
                    // }
                }
            }

            SplitView {
                Layout.fillWidth:true
                Layout.fillHeight:true
                // anchors.fill: parent
                orientation: Qt.Vertical

                TabView {
                    Tab {
                        title: "Input Pins"
                        Rectangle { color: "red" }

                        ListView {
                            id:inputpinlist
                            anchors.bottom: parent.bottom
                            anchors.horizontalCenter: parent.horizontalCenter
                            width:parent.width
                            model: chipInfo.inPinCount
                            delegate:
                            Row {
                                id: rlayout
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
                                TextField {
                                    width:100

                                    text: "ComplexBit"                                

                                    // onAccepted:
                                    // {
                                        //     console.debug("Calling update fucntion")
                                        //     inputPins.update(index,text)
                                        // }
                                    }   
                                }

                            }

                        }
                        Tab {
                            title: "Output Pins"
                            TextArea { textColor: "blue";text:"All Output Pin info goes here" }
                        }
                        Tab {
                            title: "Modules"
                            TextArea { textColor: "green";text:"All Module realted info goes here" }
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
        }

import React from 'react';
import { TreeTable, TreeState } from 'cp-react-tree-table';
import Popup from "reactjs-popup"

// backend
import {endpoint} from '../../../Protobuf/endpoint.json'
const { RegistrationRequest, ClassOperations, ReqOp, RegistrationResponse } = require('../../../Protobuf/CV/cvInterface_pb');
const { CourseValidationClient } = require('../../../Protobuf/CV/cvInterface_grpc_web_pb');
const { Student, Response } = require('../../../Protobuf/RV/rvInterface_pb.js');
const { RegistrationValidationClient } = require('../../../Protobuf/RV/rvInterface_grpc_web_pb.js');

class RegistrationTable extends React.Component {

    constructor(props) {
      super(props);
      
      this.state = {
        treeValue: TreeState.create(this.props.classes),
        enableRegistration: false,
        childrenFontSize: '15px'
      };

      this.onCourseDrop = this.onCourseDrop.bind(this);
      this.onCourseAdd = this.onCourseAdd.bind(this);
    }

    // deprecated, but still works haha might have to update l8r
    componentWillReceiveProps(prop) {
      this.setState({ treeValue: TreeState.create(prop.classes) });
    }

    render() {

      const { treeValue } = this.state;
      let totalHeight = 0;
      for(let i = 0; i < treeValue.data.length; i++){
        treeValue.data[i].metadata.height = 50;
        totalHeight += 50;
      }
      treeValue.height = totalHeight;

      return (
        <TreeTable
          value={treeValue}
          onChange={this.handleOnChange}
          headerHeight={50}
          height={treeValue.height}
          >
  
            <TreeTable.Column 
                renderCell={this.renderIndexCell}
                renderHeaderCell={() => 
                                    <div>
                                        <Popup overlayStyle={{backgroundColor:"#00000055"}} modal trigger={ <button class="courseTable-addBtn" style={{fontSize: this.state.childrenFontSize, backgroundColor:"#00000000", fontWeight:"600", color:"#0d6efd"}}>Add</button> }>
                                            <div class="registrationTable-popup">
                                                <div class="registrationTable-popupHeader" style={{marginBottom:"5%"}}>
                                                    <h5>Add Classes 📝</h5>
                                                    <hr style={{marginRight:"7.5%"}}></hr>
                                                </div>
                                                <div class="registrationTable-popupContent" style={{display:"flex"}}>
                                                    <div style={{flex: "50%"}}>
                                                        {this.indexComponent(1)}
                                                        {this.indexComponent(2)}
                                                        {this.indexComponent(3)}
                                                        {this.indexComponent(4)}
                                                    </div>
                                                    <div style={{flex: "50%", marginLeft:"auto"}}>
                                                        {this.indexComponent(5)}
                                                        {this.indexComponent(6)}
                                                        {this.indexComponent(7)}
                                                        {this.indexComponent(8)}
                                                    </div>
                                                </div>
                                                <div style={{display:"flex", marginTop:"2.5%", width: "100%", justifyContent:"flex-end", paddingRight:"7.5%"}}>
                                                    <button style={{fontSize:"14px"}} onClick={() => this.onCourseAdd()} type="button" class="btn btn-primary">Add</button>
                                                </div>
                                                <hr style={{marginRight:"7.5%", marginTop:"10%"}}></hr>
                                            </div>
                                        </Popup>
                                    </div>
                                }
                />
            
            <TreeTable.Column 
                renderCell={this.renderSecondCell}
                renderHeaderCell={() => <span>Course Code</span>}
                />

            <TreeTable.Column 
                renderCell={this.renderThirdCell}
                renderHeaderCell={() => <span>Course Number</span>}
                />

            <TreeTable.Column 
                renderCell={this.renderFourthCell}
                renderHeaderCell={() => <span>Course Name</span>}
                grow={2}
                />

            <TreeTable.Column 
                renderCell={this.renderFifthCell}
                renderHeaderCell={() => <span>Credits</span>}
                grow={0.5}
                />

            <TreeTable.Column 
                renderCell={this.renderSixthCell}
                renderHeaderCell={() => <span>Status</span>}
                />

        </TreeTable>
     );
    }
  
    onCourseAdd = () => {
        let courseList = []
        for(let i = 1; i <= 8; i++){
            let val = document.getElementById('index' + i).value

            if(val != '')
                courseList.push({Index : val, Op : "ADD"})
        }
        console.log(courseList)
    }

    onCourseDrop = (row) => {
        if(window.confirm('Are you sure you want to drop class \''+row.data.coursename+'\'?')){
            let index = row.metadata.index
            let data = this.state.treeValue.data
            data.splice(index, 1)


            this.handleOnChange(data)
        }
    }

    registrationValidation = () => {
        var client = new RegistrationValidationClient("http://" + endpoint)

        var request = new Student();
        request.setToken(sessionStorage.getItem("token"))
    
        client.checkRegVal(request, { "grpc_service" : "rv" }, (err, response) => {
            this.setState({enableRegistration : response.getEligible()})
        });
    }

    handleOnChange = (data) => {
        let structuredData = []
        for(let i = 0; i < data.length; i++){
            structuredData.push({data: data[i].data})
        }
        
        this.props.updateClasses(structuredData)
    }

    renderSixthCell = (row) => {
        return (
            <div>
                <span style={{fontWeight: "600"}}>{row.data.status}</span>
            </div>
        );
      }

    renderFifthCell = (row) => {
        return (
            <div>
                <span style={{fontWeight: "600"}}>{row.data.credits}</span>
            </div>
        );     
    }

    renderFourthCell = (row) => {
        return (
            <div>
                <span style={{fontWeight: "600"}}>{row.data.coursename}</span>
            </div>
        );
      }

    renderThirdCell = (row) => {
        return (
            <div>
                <span style={{fontWeight: "600"}}>{row.data.coursenumber}</span>
            </div>
        );
      }

    renderSecondCell = (row) => {
      return (
          <div>
                <span style={{fontWeight: "600"}}>{row.data.coursecode}</span>
          </div>
      );
    }

    renderIndexCell = (row) => {
        return (
          <div style={{ paddingLeft: (row.metadata.depth * 15) + 'px'}}>
            {(row.data.status == "Added!")
                ?
                    <button class="courseTable-deleteBtn" onClick={() => {this.onCourseDrop(row)}} style={{fontSize: this.state.childrenFontSize, backgroundColor:"#00000000", fontWeight:"600", color:"rgb(226, 28, 28)"}}>Drop</button>
                :
                    <div></div>
            }
          </div>
        );
    }
  
    indexComponent = (indexNumber) =>{
        return (
            <div style={{display:"flex"}}>
                <div class="input-group input-group-sm mb-3">
                    <div class="input-group-prepend">
                        <span class="input-group-text" id="inputGroup-sizing-sm" style={{fontSize:"14px"}}>Index {indexNumber}</span>
                    </div>
                    <input id={'index'+indexNumber} style={{fontSize:"14px", flex:"0.75"}} type="text" class="form-control" aria-label="Small" aria-describedby="inputGroup-sizing-sm"></input>
                </div>
            </div>
        );
    }

}

export default RegistrationTable;

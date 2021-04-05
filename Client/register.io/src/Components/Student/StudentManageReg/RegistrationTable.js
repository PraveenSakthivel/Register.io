import React from 'react';
import { TreeTable, TreeState } from 'cp-react-tree-table';
import Popup from "reactjs-popup"
import CountdownTimer from './CountdownTimer'

// backend
import { CVRequest } from '../../../Protobuf/RequestMaker'

class RegistrationTable extends React.Component {

    constructor(props) {
      super(props);
      
      this.state = {
        treeValue: TreeState.create(this.props.classes),
        enableRegister: false,
        registerTime: '',
        childrenFontSize: '15px',
        popupOpen: false
      };

      this.onCourseDrop = this.onCourseDrop.bind(this);
      this.onCourseAdd = this.onCourseAdd.bind(this);
    }

    componentDidMount = () => {
        this.setState({ treeValue: TreeState.create(this.props.classes) });
        this.setState({ enableRegister: this.props.enableRegister })
        this.setState({ registerTime: this.props.registerTime })
    }

    render() {

      const { treeValue } = this.state;

      if(treeValue.data.length == 0){
        this.setState({ treeValue : TreeState.create([ {data: { coursecode:'', coursenumber: '', coursename: '', credits: '', status: '' } }]) })
      }

      let totalHeight = 0;
      for(let i = 0; i < treeValue.data.length; i++){
        treeValue.data[i].metadata.height = 50;
        totalHeight += 50;
      }
      treeValue.height = totalHeight;

      let style = {}
      if(!this.state.enableRegister)
        style = {fontSize:"14px", pointerEvents:"none", height:"35px", backgroundColor:"grey"}
      else
        style = {fontSize:"14px"}

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
                                        <Popup open={this.state.popupOpen} overlayStyle={{backgroundColor:"#00000055"}} modal trigger={ <button class="courseTable-addBtn" style={{fontSize: this.state.childrenFontSize, backgroundColor:"#00000000", fontWeight:"600", color:"#0d6efd"}}>Add</button> }>
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
                                                <div style={{display:"flex", justifyContent:"flex-end", marginTop:"2.5%", width: "100%",  paddingRight:"7.5%"}}>
                                                    <div style={{paddingTop:"5px", paddingRight:"5%"}}>
                                                        <CountdownTimer date={new Date(this.state.registerTime)} />
                                                    </div>
                                                    <div>
                                                        <button style={style} onClick={() => this.onCourseAdd()} type="button" class="btn btn-primary">Add</button> 
                                                    </div>
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
        if(this.state.enableRegister){
            let courseList = []
            for(let i = 1; i <= 8; i++){
                let val = document.getElementById('index' + i).value

                if(val != ''){
                    courseList.push({val:val, reqop:'ADD'})
                }
            }
            console.log( CVRequest(courseList) )
        }
        else{
        }
    }

    onCourseDrop = (row) => {
        if(window.confirm('Are you sure you want to drop class \''+row.data.coursename+'\'?')){
            let index = row.metadata.index
            let data = this.state.treeValue.data
            data.splice(index, 1)


            this.handleOnChange(data)
        }
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

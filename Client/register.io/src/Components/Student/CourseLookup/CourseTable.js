import React from 'react';
import { TreeTable, TreeState } from 'cp-react-tree-table';
import Popup from "reactjs-popup"

// backend
import { CVRequest } from '../../../Protobuf/RequestMaker'

class CourseTable extends React.Component {

    constructor(props) {
      super(props);
      
      this.state = {
        treeValue: TreeState.create(this.props.data),
        childrenFontSize: "14px",
        enableRegister: this.props.enableRegister,
        heavyFontWeight: "500",
        fontWeight: 400,
        registrationsMap: this.registrationsToSet(),
        index:'',
        status:''
      };
    }

    registrationsToSet(){
      let registrations = this.props.studentRegistrations
      let set = new Set()

      for(let i = 0; i < registrations.length; i++){
        set.add(registrations[i].array[5])
      }

      return set
    }

    // deprecated, but still works haha might have to update l8r
    componentWillReceiveProps(prop) {
      this.setState({ treeValue: TreeState.create(prop.data) });
    }

    handleClose(){
      this.setState({index:'', status:''})
    }

    render() {
      const { treeValue } = this.state;
      
      if(treeValue.data.length == 0 && !this.props.dept){
        this.setState({ treeValue : TreeState.create([ {data: { coursecode:'', coursenumber: '', coursename: 'Select a Department to Begin', credits: '', status: '' } }]) })
      }

      let totalHeight = 0;
      for(let i = 0; i < treeValue.data.length; i++){
        if(treeValue.data[i].data.time != null){
            let height = this.returnDateHeight(treeValue.data[i].data);
            if(height != 1){
                treeValue.data[i].metadata.height = 27 * height;
                if(treeValue.data[i].$state.isVisible)
                  totalHeight += 27 * height;
            }
            else{
                treeValue.data[i].metadata.height = 37;
                if(treeValue.data[i].$state.isVisible)
                  totalHeight += 37;
            }
        }
        else{
            treeValue.data[i].metadata.height = 50;
            if(treeValue.data[i].$state.isVisible)
              totalHeight += 50;
        }
      }
      treeValue.height = totalHeight;
      var statusColors = (status) =>{
        if(status == "Added!")
            return { fontWeight:"500", color: "green"}
        else if(status == "Dropped!")
            return { fontWeight:"500", color: "green"}
        else
            return { fontWeight:"500", color: "red"}
      }

      return (
        <div>
          {(this.state.status != '')
                ?
                    <Popup open={true} modal onClose={() => this.handleClose()} overlayStyle={{backgroundColor:"#00000055"}} >
                        <div class="registrationTable-popup" style={{height:"fit-content"}}>
                            <div class="registrationTable-popupHeader" style={{marginBottom:"5%"}}>
                                <h5>Registration Results 🧾</h5>
                                <hr style={{marginRight:"7.5%"}}></hr>
                            </div>
                            
                                <div style={{display:"flex",  marginRight:"7.5%"}}>
                                    <p style={{overflow:"hidden",textOverflow: "ellipsis", width:"50%", whiteSpace:"nowrap"}}>
                                        <b>Index: </b>{this.state.index}</p>
                                    <p style={{textAlign:"right", flex:"1"}}>
                                        <b>Status: </b><b style={statusColors(this.state.status)}>{this.state.status}</b></p>
                                </div>
                            
                            <hr style={{marginRight:"7.5%", marginBottom:"7.5%"}}></hr>
                        </div>
                    </Popup>
                :
                    <div></div>
            }
          <TreeTable
            value={treeValue}
            onChange={this.handleOnChange}
            headerHeight={50}
            height={treeValue.height}
            >
    
            <TreeTable.Column grow={1.90}
              renderCell={this.renderIndexCell}
              renderHeaderCell={() => <span>Course</span>}/>

            <TreeTable.Column grow={0.5}
              renderCell={this.renderSecondCell}
              renderHeaderCell={() => <span>Code</span>}
              />

            <TreeTable.Column grow={0.35}
              renderCell={this.renderThirdCell}
              renderHeaderCell={() => <span>Credits</span>}
              />

            <TreeTable.Column 
              renderCell={this.renderFourthCell}
              renderHeaderCell={() => <span>Section</span>}
              />

            <TreeTable.Column 
              renderCell={this.renderFifthCell}
              renderHeaderCell={() => <span>Schedule</span>}
              />

            <TreeTable.Column 
              renderCell={this.renderSixthCell}
              renderHeaderCell={() => <span>Instructor</span>}
              />

            <TreeTable.Column grow={0.75}
              renderCell={this.renderAddCell}
              renderHeaderCell={() => <span></span>}
              />
              
          </TreeTable>
          </div>
     );
    }
  
    handleOnChange = (newValue) => {
      this.setState({ treeValue: newValue });
    }
  
    returnDateHeight = (dat) =>{
        let height = 0;
        let time = dat.time
        for(let i = 0; i < time.length; i++){
            let c = time.charAt(i);
            if(c == ')')
                height++;
        }
        let prof = dat.instructor
        let pHeight = 0
        for(let i = 0; i < prof.length; i++){
          let p = prof.charAt(i)
          if(p == ')')
            pHeight++
        }
        if(pHeight > height)
          height = pHeight
        return height;
    }

    parseDate = (time) =>{
        var timing = "";
        var content = [];
        let newSeg = false
        for(let i = 0; i < time.length; i++){
            let c = time.charAt(i);
            if(!newSeg && c == ','){
                content.push(<div><span style={{fontSize: this.state.childrenFontSize}}>{timing}</span><br></br></div>);
                timing = "";
            }
            else if(c != '(' && c != ')')
                timing+=c;
            else if(c == '('){
              newSeg = true
            }
            else if(c == ')'){
              newSeg = false
            }
        }
        content.push(<div key={timing}><span style={{fontSize: this.state.childrenFontSize}}>{timing}</span></div>);
        return <div>{content}</div>;
    }

    onCourseAdd = (index) => {
      let courseList = []
      for(let i = 1; i <= 8; i++){
          let val = index

          if(val != ''){
              courseList.push({val:val, reqop:'ADD'})
          }
      }

      CVRequest(courseList, this.courseChangeCallback) 
    }

    resultCodes(i){
      switch(i){
          case(1):
              return "Added!"
          case(2):
              return "Insufficient Prereqs"
          case(3):
              return "Timing Conflict"
          case(5):
              return "Invalid Index"
          case(6):
              return "Server Error (0)"
          case(7):
              return "Server Error (1)"
      }
    }
    courseChangeCallback = ( serverResponse, action ) =>{
      // eventually put in logic that will look at the serverResponse and decide whether to call validateLogin or just show error
      let responseMap = serverResponse
      let index
      let status
      for(const [key, value] of Object.entries(responseMap.map_)){
        index = key
        status = value  
      }
      let ind = status.key
      let stat = this.resultCodes(status.value)
      this.setState({index:ind, status:stat})
      if(status.value == 1){
        let map = this.state.registrationsMap
        map.add(ind)
        this.setState({registrationsMap:map})
      }
    }

    renderAddCell = (row) => {
      return(
        <div>
          {(row.data.coursename == "Select a Department to Begin")
          ?
            <div></div>
          :
          (row.data.name == null && this.state.enableRegister)
            ?
              (this.state.registrationsMap.has(row.data.index))
              ?
                <button class="courseTable-addBtn" style={{pointerEvents:"none", textDecoration:"none", border:"none", fontSize: this.state.childrenFontSize, backgroundColor:"#00000000", fontWeight:"600", color:"darkgreen"}}>Added!</button>
              :
                <button onClick={() => this.onCourseAdd(row.data.index)} class="courseTable-addBtn" style={{fontSize: this.state.childrenFontSize, backgroundColor:"#00000000", fontWeight:"600", color:"#0d6efd"}}>Add</button>
            :
              (row.data.name != null)
              ?
                <a style={{fontWeight:"600", textDecoration:"underline"}}>PreReqs</a>
              :
                <div></div>
          
          }
        </div>
      );
    }

    renderSixthCell = (row) => {
        return (
            <div>
              {(row.data.coursename == "Select a Department to Begin")
          ?
            <div></div>
          :
              (row.data.instructor != null)
                  ?
                  (
                      /*<span style={{fontSize: this.state.childrenFontSize}}>{row.data.instructor}</span>*/
                      <div>
                      {
                          this.parseDate(row.data.instructor)
                      }
                      </div>
                  )
                  :
                  (
                      <span style={{ paddingLeft: (10) + 'px'}}>👩‍🏫</span>
                  )
              }
            </div>
        );
      }

    renderFifthCell = (row) => {
        return (
            <div>
              {(row.data.coursename == "Select a Department to Begin")
          ?
            <div></div>
          :
              (row.data.time != null)
                  ?
                  (
                      (row.data.time != "By Arrangement")
                      ?
                        <div>
                        {
                            this.parseDate(row.data.time)
                        }
                        </div>
                      :
                      <span style={{fontSize: this.state.childrenFontSize}}>{row.data.time}</span>
                  )
                  :
                  (
                      <span style={{ paddingLeft: (10) + 'px'}}>📅</span>     
                  )
              }
            </div>
        );     
    }

    renderFourthCell = (row) => {
        return (
            <div>
              {(row.data.coursename == 'Select a Department to Begin')
          ?
            <div></div>
          :
              (row.data.status == null)
                  ?
                  (
                      <div>
                        <span style={{fontWeight: "600", color:"#009432"}}>Open: {row.data.openSections}</span>
                        <span> • </span>
                        <span style={{fontWeight: "600", color:"#d63031"}}>Closed: {row.data.closedSections}</span>
                      </div>
                  )
                  :
                  (
                    (row.data.status)
                        ?
                            <span style={{fontWeight: "500", fontSize: this.state.childrenFontSize, color:"#009432", paddingLeft: (row.metadata.depth * 10) + 'px'}}>Open</span>
                        :
                            <span style={{fontWeight: "500", fontSize: this.state.childrenFontSize, color:"#d63031", paddingLeft: (row.metadata.depth * 10) + 'px'}}>Closed</span>
                  )
              }
            </div>
        );
      }

    renderThirdCell = (row) => {
        return (
            <div>
              {(row.data.credits != null)
                  ?
                  (
                      <span style={{paddingLeft: "10px", fontWeight: this.state.heavyFontWeight}}>{row.data.credits}</span>
                  )
                  :
                  (
                      <span></span>     
                  )
              }
            </div>
        );
      }

    renderSecondCell = (row) => {
      return (
          <div>
            {(row.data.courseCode != null)
                ?
                (
                    <span style={{fontWeight: this.state.heavyFontWeight}}>{row.data.courseCode}</span>
                )
                :
                (
                    <span style={{fontSize: this.state.childrenFontSize}}>{row.data.index}</span>
                )
            }
          </div>
      );
    }

    renderIndexCell = (row) => {

        return (
          <div style={{ paddingLeft: (row.metadata.depth * 15) + 'px'}}
            className={row.metadata.hasChildren ? 'with-children' : 'without-children'}>
            
            {/*(row.metadata.hasChildren)
              ? (
                  <button className="toggle-button" onClick={row.toggleChildren}></button>
                )
              : ''
            
              */}

            {(row.data.coursename == 'Select a Department to Begin')
              ?
                <span style={{fontWeight:"500"}}>{row.data.coursename}</span>
              :
                <div></div>
            } 

            {(row.data.name != null)
              ? 
                <div onClick={row.toggleChildren} style={{cursor:"pointer"}}><button className="toggle-button" ></button><span class="courseTable-courseName" style={{fontWeight:"500"}}>{row.data.name}</span> </div>
              : 
                <span style={{fontSize: this.state.childrenFontSize}}>{row.data.section}</span>
              }
          </div>
        );
    }
  
}

export default CourseTable;
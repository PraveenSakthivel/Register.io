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
        popupOpen: false,
        registerResults:[],
        posResults: [],
        pends: []
      };

      this.onCourseDrop = this.onCourseDrop.bind(this);
      this.onCourseAdd = this.onCourseAdd.bind(this);
    }

    componentDidMount = () => {
        this.setState({ treeValue: TreeState.create(this.props.classes) });
        this.setState({ enableRegister: this.props.enableRegister })
        this.setState({ registerTime: this.props.registerTime })
        this.validate()
    }


    componentDidUpdate ( prevProps ) {
        if(this.props.classes.length != prevProps.classes.length){
            this.setState({treeValue: TreeState.create(this.props.classes)})
            //clearInterval(this.fetchRegistrationsCaller)
            //this.fetchRegistrationsCaller = undefined
        }
        if(this.props.classes.length == prevProps.classes.length + this.state.registerResults.length || this.props.classes.length == prevProps.classes.length -1){
            clearInterval(this.fetchRegistrationsCaller)
            this.fetchRegistrationsCaller = undefined
        }
    }

    componentWillUnmount () {
        if(this.fetchRegistrationsCaller != undefined)
            clearInterval(this.fetchRegistrationsCaller)
        this.fetchRegistrationsCaller = undefined
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

      // WHATS LEFT: JUST SHOW POPUP IF REGISTERRESULTS IS NOT EMPTY, THEN EMPTY REGISTERRESULTS AND THIS.RESULTMAP

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
            {(this.state.registerResults.length != 0)
                ?
                    <Popup open={true} modal onClose={() => this.handleClose()} overlayStyle={{backgroundColor:"#00000055"}} >
                        <div class="registrationTable-popup" style={{height:"fit-content"}}>
                            <div class="registrationTable-popupHeader" style={{marginBottom:"5%"}}>
                                <h5>Registration Results 🧾</h5>
                                <hr style={{marginRight:"7.5%"}}></hr>
                            </div>
                            
                            {this.state.registerResults.map(i => 
                                <div style={{display:"flex",  marginRight:"7.5%"}}>
                                    <p style={{overflow:"hidden",textOverflow: "ellipsis", width:"50%", whiteSpace:"nowrap"}}>
                                        <b>Index: </b>{i.index}</p>
                                    <p style={{textAlign:"right", flex:"1"}}>
                                        <b>Status: </b><b style={statusColors(i.status)}>{i.status}</b></p>
                                </div>
                            )}
                            
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


            <TreeTable.Column 
                renderCell={this.renderIndexCell}
                grow={0.75}
                renderHeaderCell={() => 
                                    <div>
                                        <Popup open={this.state.popupOpen} onOpen={() => this.handleOpen()} overlayStyle={{backgroundColor:"#00000055"}} modal trigger={ <button class="courseTable-addBtn" style={{fontSize: this.state.childrenFontSize, backgroundColor:"#ffffff00", fontWeight:"600", color:"#0d6efd"}}>Add</button> }>
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
                grow={0.75}
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
        
        </div>
     );
    }

    handleClose=()=>{
        this.setState({registerResults:[]})
        this.resultMap = []
    }

    onCourseAdd = (callback) => {
        if(this.state.enableRegister){
            let courseList = []
            for(let i = 1; i <= 8; i++){
                let val = document.getElementById('index' + i).value

                if(val != ''){
                    courseList.push({val:val, reqop:'ADD'})
                }
            }

            CVRequest(courseList, this.courseChangeCallback) 
        }
        this.setState({popupOpen:false})
    }

    handleOpen(){
        this.setState({popupOpen:true})
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

    resultMap = []
    posResultMap = []
    fetchReg = false

    logMapElements(value, key, map) {
        if(value != 1)
            this.resultMap.push({ index: key, status: this.resultCodes(value) })
        else{
            this.resultMap.push({ index: key, status: this.resultCodes(value) })
            this.fetchReg = true
            this.posResultMap.push({ index: key })
        }
    }


    courseChangeCallback = ( serverResponse, action ) =>{
        // eventually put in logic that will look at the serverResponse and decide whether to call validateLogin or just show error
        let responseMap = serverResponse
        this.posResultMap = []
        responseMap.forEach(this.logMapElements.bind(this))
        if(action == "DROP"){
            for(let i = 0; i < this.resultMap.length; i++){
             if(this.resultMap[i].status == "Added!"){
                    this.resultMap[i].status = "Dropped!"
                }
            }
        }
        this.setState({ registerResults : this.resultMap })

        if(this.fetchReg)
            this.fetchRegistrations( )
    }

    fetchRegistrationsCaller = undefined

    fetchRegistrations = () => {
        var startTime = new Date().getTime(); 
        this.fetchRegistrationsCaller = setInterval(() => {
            this.fetchReg = false
            console.log("called")
            this.validate()
            if(new Date().getTime() - startTime > 2001){
                clearInterval(this.fetchRegistrationsCaller);
                this.fetchRegistrationsCaller = undefined
                return;
            }
        
        }, 500);
        //setTimeout(function(){ clearInterval(this.fetchRegistrationsCaller); }, 2001);

    }

    validate(){
        this.props.validateLogin(window.sessionStorage.getItem('token'))
    }

    onCourseDrop = (row) => {
        if(window.confirm('Are you sure you want to drop class \''+row.data.coursename+'\'?')){
            let courseList = []
            courseList.push({val:row.data.coursecode, reqop:'DROP'})
            CVRequest(courseList, this.courseChangeCallback)
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

        let color
        switch(row.data.status){
            case('Added!'):
                color = "darkgreen"
                break
            default:
                color = "black"
        }

        return (
            <div>
                <span style={{fontWeight: "600", color:color}}>{row.data.status}</span>
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

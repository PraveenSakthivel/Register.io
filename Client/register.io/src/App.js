import React, { Component } from 'react';

// components
import Login from './Components/Login/Login'
import Content from './Components/Content/Content'

// backend
import { ValidateLogin, RVRequest, DBRetrieveCourses, DBRetrieveDepts, GetHeatMapData, GetBarGraphData } from './Protobuf/RequestMaker'

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
                    userType:-1, //userType: (-2, Loading), (-1, Not Logged in), (0, Student), (1, Admin), (2, Superuser)
                    studentRegistrations:[],
                    enableRegister:false,
                    registerTime:'', 
                    soc:[],
                    depts:[]
                  }; 
    this.validateLogin = this.validateLogin.bind(this);
    this.logout = this.logout.bind(this);
  }

  validateLogin(token){
    // TEMPORARY
    if(token == 'admin'){
      window.sessionStorage.setItem("token", 'admin')
      this.setState({ userType : 1 })
      return;
    }

    window.sessionStorage.setItem("token", token);
    ValidateLogin( { token: token }, this.validateLoginCallback )
  }

  validateLoginCallback = (serverResponse) => {
      if(serverResponse != null && serverResponse != ''){
        this.setState({userType : serverResponse.usertype})
        if(serverResponse.usertype == 0){ // if user is a student
          this.setState({studentRegistrations : serverResponse.classlist}, () => this.formatClassList()) // store the student's current registrations
          if(this.state.soc.length == 0){
            DBRetrieveCourses( {}, this.dbRetrieveCoursesCallback )
            DBRetrieveDepts( {}, this.dbRetrieveDeptsCallback )
          }
          if(this.state.registerTime == '')
            RVRequest( {}, this.registrationCallback ) // check if student is eligible to register
        }
      } 
      else 
        this.logout();
  }

  titleCase(str) {
    var splitStr = str.toLowerCase().split(' ');
    for (var i = 0; i < splitStr.length; i++) {
        // You do not need to check if i is larger than splitStr length, as your for does that for you
        // Assign it back to the array
        splitStr[i] = splitStr[i].charAt(0).toUpperCase() + splitStr[i].substring(1);     
        splitStr[i] = splitStr[i].replaceAll('Ii', 'II')   
        splitStr[i] = splitStr[i].replaceAll('Iii', 'III')
    }
    // Directly return the joined string
    return splitStr.join(' '); 
 }
 

  formatClassList(){
    let classList = this.state.studentRegistrations
    if(classList != null){
      for(let i = 0; i < classList.length; i++){
        classList[i].array[6] = this.titleCase(classList[i].array[6])
        let loc = classList[i].array[0]
        loc = loc.replaceAll("DOUGLAS/COOK", "Cook Douglass")
        loc = loc.replaceAll("BUSCH", "Busch")
        loc = loc.replaceAll("LIVINGSTON", "Livingston")
        loc = loc.replaceAll("COLLEGE AVE", "College Ave")
        classList[i].array[0] = loc
      }
      this.setState({studentRegistrations:classList})
    }
  }

  dbRetrieveCoursesCallback = (serverResponse) =>{
    this.transformClasses(serverResponse.getClassesList())
  }

  dbRetrieveDeptsCallback = (serverResponse) =>{
    this.setState({ depts:serverResponse })
  }

  registrationCallback = (serverResponse) => {
    this.setState({enableRegister : serverResponse.eligible})
    this.setState({registerTime : serverResponse.time})
  }

  logout(){
    window.sessionStorage.clear();
    window.location.reload();
  }

  componentWillMount(){
    if(sessionStorage.getItem("token") != null){
      this.setState({userType : -2})   
      this.validateLogin(sessionStorage.getItem("token"))
    }
  }

  transformClasses(rawSoc){
    let soc = []
    for(let i = 0; i < rawSoc.length; i++){
        let r = rawSoc[i]
        let sections = []
        let rawSections = r.getSectionsList()
        let numOpen = 0
        let numClosed = 0
        for(let j = 0; j < rawSections.length; j++){
            let s = rawSections[j]
            let status = ''
            if(s.getAvailable())
                numOpen++
            else
                numClosed++
            let meetings = s.getMeetingsList()
            let mtgStr = ""
            let locStr = ""
            for(let k = 0; k < meetings.length; k++){
              let m = meetings[k]
              mtgStr += '(' + m.getMeetingTime() + ')'
              locStr += '(' + m.getMeetingLocation() + ')'
              if(k != meetings.length-1){
                mtgStr+=','
                locStr += ','
              }
            }
            let instructors =  s.getInstructorsList()
            let instStr = ""
            for(let l = 0; l < instructors.length; l++){
              let inst = instructors[l]
              instStr += '(' + inst + ')'
              if(l != instructors.length-1){
                instStr+=','
              }
            }
            let section = { section: 'Section '+s.getSection(), status: s.getAvailable(), index: s.getIndex(), time: mtgStr, location: locStr, instructor:instStr  }
            sections.push({ data: section})
        }
        let course = { department:r.getDepartment(), name:r.getName(), courseCode:r.getSchool()+":"+r.getDepartment()+":"+r.getClassnum(), credits: r.getCredits(), openSections: numOpen, closedSections: numClosed };
        soc.push({ data: course, children:sections })
    }
    this.setState( { soc:soc } )
  }

  render() {
    const userType = this.state.userType;
    let content;
    if (userType == -1) 
      content = <Login validateLogin = {this.validateLogin} />; 
    else if(userType == -2)
      content = <div></div>
    else 
      content = <Content depts={this.state.depts} soc={this.state.soc} validateLogin = {this.validateLogin} logout = {this.logout} userType = {userType} studentRegistrations = {this.state.studentRegistrations} enableRegister = {this.state.enableRegister} registerTime = {this.state.registerTime} />;

    return (
      <div class="App" >
              {content}
      </div>
    )
  }
}

export default App;
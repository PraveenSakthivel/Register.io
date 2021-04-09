import React, { Component } from 'react';

// components
import Login from './Components/Login/Login'
import Content from './Components/Content/Content'

// backend
import { ValidateLogin, RVRequest } from './Protobuf/RequestMaker'

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
                    userType:-1, //userType: (-2, Loading), (-1, Not Logged in), (0, Student), (1, Admin), (2, Superuser)
                    studentRegistrations:[],
                    enableRegister:false,
                    registerTime:''
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
      console.log(serverResponse)
      if(serverResponse != null && serverResponse != ''){
        this.setState({userType : serverResponse.usertype})
        if(serverResponse.usertype == 0){ // if user is a student
          this.setState({studentRegistrations : serverResponse.classlist}) // store the student's current registrations
          if(this.state.registerTime == '')
            RVRequest( {}, this.registrationCallback ) // check if student is eligible to register
        }
      } 
      else 
        this.logout();
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

  registrationCallback = (serverResponse) => {
    this.setState({enableRegister : serverResponse.eligible})
    this.setState({registerTime : serverResponse.time})
  }

  render() {

    const userType = this.state.userType;
    let content;
    if (userType == -1) 
      content = <Login validateLogin = {this.validateLogin} />; 
    else if(userType == -2)
      content = <div></div>
    else 
      content = <Content validateLogin = {this.validateLogin} logout = {this.logout} userType = {userType} studentRegistrations = {this.state.studentRegistrations} enableRegister = {this.state.enableRegister} registerTime = {this.state.registerTime} />;

    return (
      <div class="App" >
              {content}
      </div>
    )
  }
}

export default App;
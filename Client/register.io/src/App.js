import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

// components
import Login from './Components/Login/Login'
import Content from './Components/Content/Content'

// backend
import {endpoint} from '../src/Protobuf/endpoint.json'
const { Student, Response } = require('./Protobuf/RV/rvInterface_pb.js');
const { RegistrationValidationClient } = require('./Protobuf/RV/rvInterface_grpc_web_pb.js');

// backend
const { Credentials, Registrations, Class, Token } = require('./Protobuf/UserV/token_pb.js');
const { LoginEndpointClient } = require('./Protobuf/UserV/token_grpc_web_pb.js');

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

    window.sessionStorage.setItem("token", token);

    var client = new LoginEndpointClient(endpoint)

    let protoToken = new Token();
    protoToken.setToken(token);

    client.getCurrentRegistrations(protoToken, { "grpc_service" : "uv" }, (err, response) => {

      if(response != null && response != ''){
          this.setState({userType : response.getUsertype()})
          if(response.getUsertype() == 0){ // if user is a student
            this.RVRequest()
            this.setState({studentRegistrations : response.getClassesList()})
            console.log(response.getClassesList())
          }
      } else {
          this.logout();
      }

    });

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

  RVRequest(){
    var client = new RegistrationValidationClient(endpoint)

    var request = new Student();
    request.setToken(window.sessionStorage.getItem("token").toString());

    
    client.checkRegVal(request, { "grpc_service" : "rv" }, (err, response) => {
        this.setState({enableRegister : response.getEligible()})
        this.setState({registerTime : response.getTime()})
        console.log(response.getEligible())
    });
  }

  render() {

    const userType = this.state.userType;
    let content;
    if (userType == -1) 
      content = <Login validateLogin = {this.validateLogin} />; 
    else if(userType == -2)
      content = <div></div>
    else 
      content = <Content logout = {this.logout} userType = {userType} studentRegistrations = {this.state.studentRegistrations} enableRegister = {this.state.enableRegister} registerTime = {this.state.registerTime} />;

    return (
      <div class="App" >
        <BrowserRouter>
          <Switch>

            <Route path='/'>
              {content}
            </Route>

          </Switch>
        </BrowserRouter>
      </div>
    )
  }
}

export default App;
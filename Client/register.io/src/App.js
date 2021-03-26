import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

// components
import Login from './Components/Login/Login'
import Content from './Components/Content/Content'

// backend
import {endpoint} from '../src/Protobuf/endpoint.json'
const { Student, Response } = require('../src/Protobuf/RV/rvInterface_pb.js');
const { RegistrationValidationClient } = require('./Protobuf/RV/rvInterface_grpc_web_pb.js');

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {loggedIn:false, 
                  userType:-1, //userType: (-1, Not Logged in), (1, Student), (2, Admin), (3, Superuser)
                  studentCourses:"", 
                  studentClass: 4,
                  studentMajor: [198, 210],
                  studentCredits: 103}; 
    this.validateLogin = this.validateLogin.bind(this);
  }

  validateLogin(userType){
    this.setState({loggedIn:true});
    this.setState({userType:userType});
    window.sessionStorage.setItem("loggedIn", true);
    window.sessionStorage.setItem("userType", userType);
  }

  componentWillMount(){
    this.setState({loggedIn:window.sessionStorage.getItem("loggedIn")});
    this.setState({userType:window.sessionStorage.getItem("userType")});
  }

  request(){
    var client = new RegistrationValidationClient('http://'+endpoint)

    var request = new Student();
    request.setNetid("ps931");

    
    client.checkRegVal(request, {  }, (err, response) => {
      var res = new Response();
      console.log(response);
    });
  }

  render() {

    this.request();


    const isLoggedIn = this.state.loggedIn;
    const userType = this.state.userType;
    let content;
    if (!isLoggedIn) 
      content = <Login validateLogin={this.validateLogin} />; 
    else 
      content = <Content userType = {userType} />;

    return (
      <div class="App">
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
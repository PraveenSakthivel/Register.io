import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Home from './Components/Home/Home'
import Login from './Components/Login/Login'
import Content from './Components/Content/Content'

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

  render() {

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
import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Home from './Components/Home/Home'
import Login from './Components/Login/Login'
import Navbar from './Components/Navbar/Navbar'

class App extends Component {

  isLoggedIn(props){
    return <Navbar />;
  }

  render() {
    return (
      <div class="App">
        <BrowserRouter>
          <Switch>

            <Route path='/login'>
              <Login />
            </Route>

            <Route path='/'>
              <this.isLoggedIn />
            </Route>

          </Switch>
        </BrowserRouter>
      </div>
    )
  }
}

export default App;
import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Home from './Components/Home/Home'
import Login from './Components/Login/Login'

class App extends Component {
  render() {
    return (
      <div class="App">
        <BrowserRouter>
          <Switch>

            <Route path='/login'>
              <Login />
            </Route>

            <Route path='/'>
              <Home />
            </Route>

          </Switch>
        </BrowserRouter>
      </div>
    )
  }
}

export default App;
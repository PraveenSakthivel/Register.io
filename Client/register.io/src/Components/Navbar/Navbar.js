import React from 'react';
import logo from '../../Assets/logo_navbar.png'

class Navbar extends React.Component {

    constructor(props) {
        super(props);
        this.state = {selectedComponentID: "Dashboard"}
        this.menuSelect = this.menuSelect.bind(this);
        this.renderNavbar = this.renderNavbar.bind(this);
    }

    menuSelect (e) {
        let element = e.target;
        if(element.tagName != "DIV"){
            element = element.parentNode;
        }
        
        if(element.id != ""){
            if(this.state.selectedComponentID != ""){
                let currentlySelected = document.getElementById(this.state.selectedComponentID);
                currentlySelected.className = "navbar-item";
            }
            this.props.switchComponent(element.id);
            this.setState({selectedComponentID:element.id});
            element.className = "navbar-item navbar-selected";
        }
    }

    renderNavbar(userType) {
        switch(userType) {
          case 0:
            return  <div><img onDoubleClick={this.inconspicious} class="navbar-logo" src={logo}></img>
                        {/*<div class="navbar-item" id="Student Account">
                            <p>🧑</p>
                            <a class="navbar-item-child" >My Account</a>
                        </div>*/}
                        <div class="navbar-item navbar-selected" id="Dashboard">
                            <p>🏡</p>
                            <a class="navbar-item-child">Dashboard</a>
                        </div>
                        <div class="navbar-item" id="Student Manage Registration">
                            <p>✍️</p>
                            <a class="navbar-item-child">Manage<br></br>Registration</a>
                        </div>
                        <div class="navbar-item" id="Student Course Lookup">
                            <p>🔍</p>
                            <a class="navbar-item-child">Course Lookup</a>
                        </div>
                        <div class="navbar-item" style={{position: "absolute", left: 0, right: 0, bottom: 0, marginBottom: "25%"}} id="Logout">
                            <p>⬅️</p>
                            <a class="navbar-item-child">Logout</a>
                        </div>
                    </div>
          case 1:
            return  <div><img class="navbar-logo" src={logo}></img>
                        <div class="navbar-item" id="Admin Account">
                            <p>🧑</p>
                            <a class="navbar-item-child" >My Account</a>
                        </div>
                        <div class="navbar-item navbar-selected" id="Dashboard">
                            <p>📈</p>
                            <a class="navbar-item-child">Dashboard</a>
                        </div>
                        <div class="navbar-item" style={{position: "absolute", left: 0, right: 0, bottom: 0, marginBottom: "25%"}} id="Logout">
                            <p>⬅️</p>
                            <a class="navbar-item-child">Logout</a>
                        </div>
                    </div>
          case 2:
              break;
        }
    }
      
    inconspicious(){  
        window.location.assign('http://github.com/rishabr17/Jumpy-Jeb');
    }

    render() {

        let userType = this.props.userType;

        return (
            <div onClick={this.menuSelect} class="navbar">

                { this.renderNavbar(userType) }

            </div>
        );
    }
}

export default Navbar; 
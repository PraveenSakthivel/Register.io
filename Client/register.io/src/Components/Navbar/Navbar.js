import React from 'react';
import logo from '../../Assets/logo_navbar.png'

class Navbar extends React.Component {

    render() {
        return (
            <div class="navbar">
                <img class="navbar-logo" src={logo}></img>
                <div class="navbar-item">
                    <p>🧑</p>
                    <a style={{color:"white"}}>My Account</a>
                </div>
                <div class="navbar-item">
                    <p>🏡</p>
                    <a style={{color:"white"}}>Dashboard</a>
                </div>
                <div class="navbar-item">
                    <p>✍️</p>
                    <a style={{color:"white"}}>Manage<br></br>Registration</a>
                </div>
                <div class="navbar-item">
                    <p>🔍</p>
                    <a style={{color:"white"}}>Course Lookup</a>
                </div>
                <div class="navbar-item">
                    <p>⌛</p>
                    <a style={{color:"white"}}>Class History</a>
                </div>
                <div class="navbar-item navbar-logout">
                    <p>⬅️</p>
                    <a style={{color:"white"}}>Logout</a>
                </div>
            </div>
        );
    }
}

export default Navbar;
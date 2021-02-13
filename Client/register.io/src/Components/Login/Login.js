import React from 'react';
import loginImg from '../../Assets/login_backdrop.jpg';
import logo from '../../Assets/logo_full.png';

class Login extends React.Component {

    render() {
        return (
            <div style={{  
                backgroundImage: "url("+loginImg+")",
                backgroundPosition: 'center',
                backgroundSize: 'cover',
                backgroundRepeat: 'no-repeat',
                opacity: 0.85
              }} class="login">

                <div class ="card login-card">
                    <div class="card-body login-card-body">

                        <div class="login-header">
                            <h1 style={{flex:.55}}>Login</h1> 
                            <img style={{flex:.45}} src={logo}></img>
                        </div>

                        <div class="login-body"> 
                            <div style={{paddingBottom: 10}} class="input-group mb-3">
                                <span style={{width: 100}} class="input-group-text" id="basic-addon3">NetID</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3"></input>
                            </div>
                            <div style={{paddingBottom: 10}} class="input-group mb-3">
                                <span style={{width: 100}} class="input-group-text" id="basic-addon3">Password</span>
                                <input type="password" class="form-control" id="basic-url" aria-describedby="basic-addon3"></input>
                            </div>
                            <div class="login-row">
                                <a class="login-forgot-password" >Forgot Password?</a>
                                <p style={{flex:.75}}></p>
                                <button type="submit" style={{flex:.25}} type="button" class="btn btn-primary login-btn">Login</button>
                            </div>
                        
                        </div>
                        
                        <div class="login-footer">
                                <p style={{ textAlign:'center'}}>Powered by Register.io &#128218;</p>
                                <hr></hr>
                        </div>
                    </div>
                </div>

            </div>
        );
    }
}

export default Login;
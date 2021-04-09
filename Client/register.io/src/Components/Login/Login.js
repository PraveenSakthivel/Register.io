import React from 'react';
import loginImg from '../../Assets/login_backdrop.jpg';
import logo from '../../Assets/logo_full.png';


// backend
import { LoginRequest } from '../../Protobuf/RequestMaker'

class Login extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            netid:'',
            password:'',
            invalidLogin:false
        }
        this.login = this.login.bind(this);
        this.handleKeyPress = this.handleKeyPress.bind(this);
    }

    login(){

        // TEMPORARY
        if(this.state.netid == 'admin'){
            this.props.validateLogin('admin')
            return;
        }
        
        if(this.state.netid == '' || this.state.password == '')
            this.setState({ invalidLogin : true })
        else
            LoginRequest({ netid: this.state.netid, password: this.state.password }, this.loginCallback)

    }

    loginCallback = (serverResponse) =>{
        if(serverResponse == null || serverResponse == 'Invalid Login' || serverResponse.token == null )
            this.setState({invalidLogin:true})
        else if(serverResponse.token != null)
            this.props.validateLogin(serverResponse.token)
    }

    handleKeyPress = (event) => {
        if(event.key === 'Enter'){
          this.login();
        }
    }

    updateNetID = (e) =>{
        this.setState({
            netid : e.target.value
        });
    }

    updatePassword = (e) =>{
        this.setState({
            password : e.target.value
        });
    }

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
                    <form class="card-body login-card-body">

                        <div class="login-header">
                            <h1 style={{flex:.55, fontFamily:"-apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif"}}>Login</h1> 
                            <img style={{flex:.45}} src={logo}></img>
                        </div>

                        <div class="login-body"> 
                            <div style={{paddingBottom: 10}} class="input-group mb-3">
                                <span style={{width: 100}} class="input-group-text" id="basic-addon3">NetID</span>
                                <input onChange={this.updateNetID} type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3"></input>
                            </div>
                            <div style={{paddingBottom: 10}} class="input-group mb-3">
                                <span style={{width: 100}} class="input-group-text" id="basic-addon3">Password</span>
                                <input onChange={this.updatePassword} onKeyPress={this.handleKeyPress} type="password" class="form-control" id="basic-url" aria-describedby="basic-addon3"></input>
                            </div>
                            <div class="login-row">
                                <a class="login-forgot-password" >Forgot Password?</a>
                                <p style={{flex:.75}}></p>
                                <button onClick={this.login} type="submit" style={{flex:.25}} type="button" class="btn btn-primary login-btn">Login</button>
                            </div>
                        
                        </div>

                        {(this.state.invalidLogin)
                            ?
                                <div style={{margin:"5% 7.5% 0 7.5%", border:"dashed 2px var(--color-primary-dark)", borderRadius:"5px", backgroundColor:"#ffeded"}}>
                                    <p style={{paddingLeft:"5%", paddingTop:"2.5%", color:"var(--color-primary-dark)"}}>Invalid NetID/Password. Please try again</p>
                                </div>
                            :
                                <div></div>
                        }

                        <div class="login-footer">
                                <p style={{ textAlign:'center'}}>Powered by Register.io &#128218;</p>
                                <hr></hr>
                        </div>
                    </form>
                </div>

            </div>
        );
    }
}

export default Login;
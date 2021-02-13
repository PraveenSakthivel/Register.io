
class App extends React.Component {
    render() {

        if (Cookies.get("token")) {
            console.log("token present")
            return <LoggedIn />
        }

        if (Cookies.get('state') == "notfound") {
            return <Notfound />;
        } else if (Cookies.get('state') == "invalid") {
            return <Invalid />;
        } else if (Cookies.get('state') == "login") {
            return <Login />;
        }
        return <Home />;
    }
}


class LoggedIn extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            user: "",
            page: ""
        };
        this.serverRequest = this.serverRequest.bind(this);
        this.logout = this.logout.bind(this);
    }

    logout() {
        Cookies.remove("token")
        window.location.reload();
    }

    serverRequest() {
        let self = this
        fetch("http://localhost:8080/authuser")
            // fetch("http://localhost:8080/authuser")
            .then(res => res.arrayBuffer())
            .then(
                (result) => {
                    protobuf.load("/js/token.proto", function (err, root) {
                        if (err)
                            throw err;
                        // Obtain a message type
                        var Token = root.lookupType("main.Token");
                        console.log("RESULT RECIEVED")
                        var uint8View = new Uint8Array(result);
                        console.log(uint8View)
                        var message = Token.decode(uint8View);
                        if (message.token == "") {
                            Cookies.remove("token")
                            backtohome()
                        }
                        console.log(message)
                        console.log(message.token)
                        self.setState({
                            user: message.token
                        });
                    });

                },
                (error) => {
                    Cookies.remove("token")
                    backtohome()
                }
            )
    }

    componentDidMount() {
        this.serverRequest();
    }

    clickButton(backout) {
        console.log(this)
        console.log(backout)
        if (backout == true) {
            this.setState({
                page: ""
            });
            // window.location.reload();
        } else {
            this.setState({
                page: "userinfo"
            });
        }
    }


    render() {
        if ("" === this.state.page) {
            return (
                <div class="container-fluid">
                    <br />
                    <div class="row">
                        <div class="col text-center">
                            <h1> Welcome to WebReg2.0 </h1>
                            <br />
                            <h1>Logged in as: {this.state.user}</h1>
                            <br />
                            <a onClick={() => { this.clickButton(false) }}><button class="btn btn-primary">User information</button></a>
                            <br />
                            <br />
                            <a onClick={this.logout}><button class="btn btn-primary">Logout</button></a>
                        </div>
                    </div>
                </div>
            )
        } else if ("userinfo" === this.state.page) {
            return (
                <div class="container-fluid">
                    <div class="row">
                        <div class="col-4">
                        </div>
                        <div class="col-4">
                            <br />
                            <h1>User Info: {this.state.user} </h1>
                            <br />
                            <a onClick={() => { this.clickButton(true) }}><button class="btn btn-primary">Back</button></a>
                        </div>
                    </div>

                </div>
            )
        }

    }
}

class Home extends React.Component {
    render() {

        if (Cookies.get('state') == "signup") {
            return <Signup />
        }

        return (
            <div class="container-fluid">
                <br />
                <div class="row">
                    <div class="col text-center">
                        <h1> Welcome to WebReg2.0 </h1>
                        <br />
                        <a href="/signup"><button style={{ width: "17.5%" }} class="btn btn-primary">Signup</button></a>
                        <br />
                        <br />
                        <a href="/login"><button style={{ width: "17.5%" }} class="btn btn-primary">Login</button></a>
                    </div>
                </div>
            </div>
        );
    }
}

class Signup extends React.Component {
    render() {
        return (
            <div class="container-fluid">

                <br />
                <div class="row">
                    <div class="col-4">
                    </div>
                    <div class="col-4">
                        <h1>Signup</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col-4">
                    </div>
                    <div class="col-4">
                        <form action="/signup_user" method="POST">
                            <div class="form-group">
                                <label for="exampleInputEmail1">Email address</label>
                                <input name="email" type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email" />
                                <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
                            </div>
                            <div class="form-group">
                                <label for="exampleInputPassword1">Password</label>
                                <input name="password" type="password" class="form-control" id="exampleInputPassword1" placeholder="Password" />
                            </div>
                            <br />
                            <button type="submit" class="btn btn-primary">Submit</button>
                        </form>
                    </div>
                </div>
                <br />
                <div class="row">
                    <div class="col-4">
                    </div>
                    <div class="col-4">
                        <BackButton />
                    </div>
                </div>
            </div>
        )
    }
}

class BackButton extends React.Component {

    pressed() {
        Cookies.set("state", "")
        window.location.reload();
    }

    render() {
        return (
            <button class="btn btn-primary" style={{ width: "17.5%" }} onClick={this.pressed}>Back</button>
        )
    }
}

function backtohome() {

}

class Login extends React.Component {
    render() {
        return (
            <div class="container-fluid">

                <br />
                <div class="row">
                    <div class="col-4">
                    </div>
                    <div class="col-4">
                        <h1>Login</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col-4">
                    </div>
                    <div class="col-4">
                        <form action="/login_user" method="POST">
                            <div class="form-group">
                                <label for="exampleInputEmail1">Email address</label>
                                <input name="Email" type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email" />
                            </div>
                            <div class="form-group">
                                <label for="exampleInputPassword1">Password</label>
                                <input name="Password" type="password" class="form-control" id="exampleInputPassword1" placeholder="Password" />
                            </div>
                            <br />
                            <button type="submit" class="btn btn-primary">Submit</button>
                        </form>
                    </div>
                </div>
                <br />
                <div class="row">
                    <div class="col-4">
                    </div>
                    <div class="col-4">
                        <BackButton />
                    </div>
                </div>
            </div>
        )
    }
}

class Notfound extends React.Component {
    render() {
        return (
            <h1>User already exists</h1>
        )
    }
}

class Invalid extends React.Component {
    render() {
        return (
            <h1>Invalid credentials</h1>
        )
    }
}

ReactDOM.render(<App />, document.getElementById("app"));
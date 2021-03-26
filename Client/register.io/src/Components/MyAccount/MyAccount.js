import React from 'react';

class MyAccount extends React.Component {

    constructor(props) {
        super(props);
        this.state = {}; 
    }

    render() {

        return (
            <div class="myaccount">
                <div class="myaccount-title">
                    <h3 style={{paddingBottom: "1%"}}>My Account</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>
            </div>
        );
    }
}

export default MyAccount;
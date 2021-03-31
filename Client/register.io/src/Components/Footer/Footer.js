import React from 'react';

class Footer extends React.Component {

    constructor(props) {
        super(props);
        this.state = {}; 
    }

    render() {

        return (
            <div style={{marginBottom:"4%"}} class="footer">
                <hr style={{color:"gray"}}></hr>
                <div style={{width:"inherit", textAlign:"left"}}>
                    <p style={{paddingLeft:"1%"}}>Powered by Register.io 📚</p>
                </div>
            </div>
        );
    }
}

export default Footer;
import React from 'react';

class Dashboard extends React.Component {

    constructor(props) {
        super(props);
        this.state = {}; 
    }

    render() {

        return (
            <div class="dashboardAdmin">
                <div class="dashboardAdmin-title">
                    <h3 style={{paddingBottom: "1%"}}>Dashboard</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>
                <div class="dashboardAdmin-content">

                </div>
            </div>
        );
    }
}

export default Dashboard;
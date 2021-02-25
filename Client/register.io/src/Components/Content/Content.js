import React from 'react';
import logo from '../../Assets/logo_navbar.png'
import Navbar from '../Navbar/Navbar'
import StudentManageReg from '../StudentManageReg/StudentManageReg'

class Content extends React.Component {

    constructor(props) {
        super(props);
        this.state = {componentID: "Dashboard"}; // 0 corresponds to homepage
        this.switchComponent = this.switchComponent.bind(this);
    }

    switchComponent(componentID) { // every navbar button has a unique ID, passes up here
        if(componentID != this.state.componentID)
            this.setState({componentID: componentID});
    }

    render() {

        let userType = this.props.userType;
        let page;

        console.log(userType);

        switch(this.state.componentID){
            case("Dashboard"):
                if(userType == 1)
                    page = <h1>Student Dashboard</h1>
                break;
            case("Student Account"):
                page = <h1>Student Account</h1>
                break;
            case("Student Manage Registration"):
                page = <StudentManageReg />
                break;
            case("Student Course Lookup"):
                page = <h1>Student Course Lookup</h1>
                break;
            case("Student Class History"):
                page = <h1>Student Class History</h1>
                break;
            case("Admin Dashboard"):
                page = <h1>Admin Dashboard</h1>
                break;
            case("Superuser Dashboard"):
                page = <h1>Superuser Dashboard</h1>
                break;
            case("Logout"):
                page = <h1>Logout</h1>
                break;
        }

        return (
            <div class="content">
                <Navbar switchComponent={this.switchComponent} userType={this.props.userType} />
                <div>
                    {page}
                </div>
            </div>
        );
    }
}

export default Content;
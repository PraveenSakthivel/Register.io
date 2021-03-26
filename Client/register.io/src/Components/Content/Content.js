import React from 'react';
import logo from '../../Assets/logo_navbar.png'
import Navbar from '../Navbar/Navbar'
import StudentManageReg from '../StudentManageReg/StudentManageReg'
import CourseLookup from '../CourseLookup/CourseLookup'
import Dashboard from '../Dashboard/Dashboard'
import ClassHistory from '../ClassHistory/ClassHistory'
import MyAccount from '../MyAccount/MyAccount'

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
                page = <Dashboard />
                break;
            case("Student Account"):
                page = <MyAccount />
                break;
            case("Student Manage Registration"):
                page = <StudentManageReg />
                break;
            case("Student Course Lookup"):
                page = <CourseLookup />
                break;
            case("Student Class History"):
                page = <ClassHistory />
                break;
            case("Admin Dashboard"):
                page = <h1>Admin Dashboard</h1>
                break;
            case("Superuser Dashboard"):
                page = <h1>Superuser Dashboard</h1>
                break;
            case("Logout"):
                window.sessionStorage.removeItem("loggedIn");
                window.sessionStorage.removeItem("userType");
                window.location.reload();
                break;
        }

        return (
            <div class="content">
                <Navbar switchComponent={this.switchComponent} userType={this.props.userType} />
                
                {page}
            </div>
        );
    }
}

export default Content;
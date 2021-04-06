import React from 'react';

import Navbar from '../Navbar/Navbar'

// Student Components
import StudentManageReg from '../Student/StudentManageReg/StudentManageReg'
import CourseLookup from '../Student/CourseLookup/CourseLookup'
import StudentDashboard from '../Student/Dashboard/Dashboard'
import ClassHistory from '../Student/ClassHistory/ClassHistory'
import MyAccount from '../Student/MyAccount/MyAccount'

// Admin Components
import AdminDashboard from '../Admin/Dashboard/Dashboard'

import Footer from '../Footer/Footer'

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

        switch(this.state.componentID){
            case("Dashboard"):
                if(userType == 0)
                    page = <StudentDashboard />
                else if(userType == 1)
                    page = <AdminDashboard />
                break;
            case("Student Account"):
                page = <MyAccount />
                break;
            case("Student Manage Registration"):
                page = <StudentManageReg studentRegistrations={this.props.studentRegistrations} enableRegister = {this.props.enableRegister} registerTime = {this.props.registerTime} />
                break;
            case("Student Course Lookup"):
                page = <CourseLookup enableRegister = {this.props.enableRegister} registerTime = {this.props.registerTime} />
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
                this.props.logout();
                break;
        }

        return (
            <div class="content" >
                <Navbar switchComponent={this.switchComponent} userType={this.props.userType} />
                
                {page}

                {(this.state.componentID != "Logout") ? <Footer /> : <div></div>}
            </div>
        );
    }
}

export default Content;
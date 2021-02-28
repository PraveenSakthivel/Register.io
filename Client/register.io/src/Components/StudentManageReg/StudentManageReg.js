import React from 'react';
import 'bootstrap/js/dist/dropdown';
import 'bootstrap/js/dist/button';
import CourseTable from './CourseTable'

class StudentManageReg extends React.Component {

    constructor(props) {
        super(props);
        this.state = {semesters:["Winter 2021", "Spring 2021", "Summer 2021", "Fall 2021"],
                        selectedSemester: ""};
        this.dropdownHandler = this.dropdownHandler.bind(this);
    }

    dropdownHandler(e) {
        this.setState({selectedSemester:e.target.text});
    }

    componentDidMount() {
        this.setState({selectedSemester:this.state.semesters[0]});
    }

    render() {

        return (
            <div class="studentManageReg">
                <div class="studentManageReg-title">
                    <h3 style={{paddingBottom: "1%"}}>Manage Registration</h3>
                    <hr style={{color:"grey", marginRight:"15%"}}></hr>
                </div>

                <div class="studentManageReg-registration">
                    <div class="studentManageReg-registrationHeader">
                        <h5 style={{paddingLeft:"1%", paddingRight:"max(2%, 15px)", paddingTop:"max(.5%, 5px)"}}>Add Classes for: </h5>
                        <div class="studentManageReg-dropdown dropdown show">
                            <a style={{fontWeight:"500"}} class="btn btn-secondary dropdown-toggle" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                {this.state.selectedSemester}
                            </a>

                            <div class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                                {this.state.semesters.map(s => (<a onClick={this.dropdownHandler} class="dropdown-item" >{s}</a>))} 
                            </div>
                        </div>
                        <div style={{marginLeft:"auto", marginRight:"10%"}} class="btn-group" role="group">
                            <input type="radio" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off" checked></input>
                            <label style={{paddingLeft: "15%", paddingRight:"15%", boxShadow:"lightgray"}} class="studentManageReg-radio btn " for="btnradio1">Schedule</label>

                            <input  type="radio" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off"></input>
                            <label style={{paddingLeft: "15%", paddingRight:"15%", boxShadow:"lightgray"}} class="studentManageReg-radio btn " for="btnradio2">Waitlist</label>
                        </div>
                    </div>
                    <div class="studentManageReg-registrationContent">
                        <CourseTable semester={this.state.selectedSemester} />
                    </div>
                </div>
                
            </div>
        );
    }
}

export default StudentManageReg;
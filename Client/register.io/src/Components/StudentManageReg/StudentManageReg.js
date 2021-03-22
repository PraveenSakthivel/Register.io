import React from 'react';
import 'bootstrap/js/dist/dropdown';
import 'bootstrap/js/dist/button';
import CourseTable from './CourseTable'

class StudentManageReg extends React.Component {

    constructor(props) {
        super(props);
        this.state = {  selectedSemester: 0,
                        classes: [
                            {semester:"Winter 2021", id: 0, classes:
                                [
                                    { id: 1, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Added!" },
                                    { id: 2, coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Added!" },
                                    { id: 3, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Pending Add" },
                                    { id: 4, coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Class Filled" },
                                    { id: 5, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "PreReqs Not Met" },
                                    { id: 6, coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Added!" }
                                ]
                            },
                            {semester:"Spring 2021", id: 1, classes:
                                [
                                    { id: 1, coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Added!" },
                                    { id: 2, coursecode:32443, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Added!" },
                                    { id: 3, coursecode:78575, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Pending Add" },
                                    { id: 4, coursecode:46224, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Class Filled" },
                                ]
                            }
                        ]};
        this.dropdownHandler = this.dropdownHandler.bind(this);
        this.updateClasses = this.updateClasses.bind(this);
    }

    dropdownHandler(e) {
        this.setState({selectedSemester:e.target.id});
    }

    updateClasses(e){
        let copy = this.state.classes;
        copy[this.state.selectedSemester].classes = e;
        this.setState({classes:copy});
    }

    render() {

        return (
            <div class="studentManageReg">
                <div class="studentManageReg-title">
                    <h3 style={{paddingBottom: "1%"}}>Manage Registration</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>

                <div class="studentManageReg-registration">
                    <div class="studentManageReg-registrationHeader">
                        <h5 style={{fontSize: "18px", paddingLeft:"1%", paddingRight:"max(2%, 15px)", paddingTop:"max(.7%, 7px)"}}>Add Classes for: </h5>
                        <div class="studentManageReg-dropdown">
                            <a style={{fontSize:"15px", fontWeight:"500"}} class="btn btn-secondary dropdown-toggle" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                {this.state.classes[this.state.selectedSemester].semester}
                            </a>

                            <div class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                                {this.state.classes.map(s => (<a id={s.id} onClick={this.dropdownHandler} class="dropdown-item" >{s.semester}</a>))} 
                            </div>
                        </div>
                        <div style={{marginLeft:"auto", marginRight:"12.5%"}} class="btn-group" role="group">
                            <input type="radio" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off" checked></input>
                            <label style={{fontSize:"15px", paddingLeft: "15%", paddingRight:"15%", boxShadow:"lightgray", whiteSpace:"nowrap"}} class="studentManageReg-radio btn " for="btnradio1">Schedule&nbsp;📅</label>

                            <input  type="radio" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off"></input>
                            <label style={{fontSize:"15px", paddingLeft: "15%", paddingRight:"15%", boxShadow:"lightgray", whiteSpace:"nowrap"}} class="studentManageReg-radio btn " for="btnradio2">Watchlist&nbsp;😎</label>
                        </div>
                    </div>
                    <div class="studentManageReg-registrationContent">
                        <CourseTable updateClasses={this.updateClasses} classes={this.state.classes[this.state.selectedSemester].classes} key={this.state.selectedSemester} />
                    </div>
                </div>
                
            </div>
        );
    }
}

export default StudentManageReg;
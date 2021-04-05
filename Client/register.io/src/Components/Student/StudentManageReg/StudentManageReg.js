import React from 'react';
import 'bootstrap/js/dist/dropdown';
import 'bootstrap/js/dist/button';
import RegistrationTable from './RegistrationTable'
var Dropdown = require('react-simple-dropdown')
var DropdownTrigger = Dropdown.DropdownTrigger
var DropdownContent = Dropdown.DropdownContent

class StudentManageReg extends React.Component {

    constructor(props) {
        super(props);
        this.state = { classes: this.registrationsFormatter() };
        this.updateClasses = this.updateClasses.bind(this);
    }

    updateClasses(e){
        this.setState({classes:e});
    }

    registrationsFormatter = () => {
        let formattedData = []
        let rawData = this.props.studentRegistrations

        for(let i = 0; i < rawData.length; i++){
            let classData = {
                coursecode : rawData[i].array[5],
                coursenumber : rawData[i].array[2] + ':' + rawData[i].array[3] + ':' + rawData[i].array[4] + ':' + rawData[i].array[7],
                coursename : rawData[i].array[6],
                credits : "3.0",
                status : "Added!"
            }

            formattedData.push({data : classData})
        }

        return formattedData
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
                        <h5 style={{fontSize: "18px", paddingLeft:"1%", paddingRight:"max(2%, 15px)", paddingTop:"max(.7%, 7px)"}}>Semester: Spring 2021</h5>
                        <div style={{marginLeft:"auto", marginRight:"12.5%"}} class="btn-group" role="group">
                            <input type="radio" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off" checked></input>
                            <label style={{fontSize:"15px", paddingLeft: "15%", paddingRight:"15%", boxShadow:"lightgray", whiteSpace:"nowrap"}} class="studentManageReg-radio btn " for="btnradio1">Schedule&nbsp;📅</label>

                            <input  type="radio" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off"></input>
                            <label style={{fontSize:"15px", paddingLeft: "15%", paddingRight:"15%", boxShadow:"lightgray", whiteSpace:"nowrap"}} class="studentManageReg-radio btn " for="btnradio2">Watchlist&nbsp;😎</label>
                        </div>
                    </div>
                    <div class="studentManageReg-registrationContent">
                        <div>
                            <RegistrationTable updateClasses={this.updateClasses} classes={this.state.classes} studentRegistrations = {this.props.studentRegistrations} enableRegister={this.props.enableRegister} registerTime={this.props.registerTime} />
                        </div>
                    </div>
                </div>
                
            </div>
        );
    }
}

export default StudentManageReg;

const data = 
[
    {data: { coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Added!" } },
    {data: { coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Added!" } },
    {data: { coursecode:93028, coursenumber: '01:198:352:01', coursename: 'Internet Technology', credits: 4.0, status: "Pending Add" } },
    {data: { coursecode:30284, coursenumber: '18:332:251:03', coursename: 'Probability and Random Processes', credits: 3.0, status: "Class Filled" } }
        
]
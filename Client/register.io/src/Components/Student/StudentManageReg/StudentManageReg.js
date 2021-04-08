import React from 'react';
import 'bootstrap/js/dist/dropdown';
import 'bootstrap/js/dist/button';
import RegistrationTable from './RegistrationTable'
import { Dropdown } from 'reactjs-dropdown-component'

class StudentManageReg extends React.Component {

    constructor(props) {
        super(props);
        this.state = { classes: this.registrationsFormatter(), 
                        semesters: [{ label: 'Spring 2021', value: 'Spring 2021' }],
                        lists: [{ label: 'Schedule 📅', value: 'Schedule' }, { label: 'Watchlist 😎', value: 'Watchlist' }] 
                    
                    };
        this.updateClasses = this.updateClasses.bind(this);
    }

    componentDidUpdate ( prevProps ) {
        if(this.props.studentRegistrations != prevProps.studentRegistrations)
            this.setState({ classes : this.registrationsFormatter() })
    }

    updateClasses(e){
        this.setState({classes:e});
    }

    registrationsFormatter = () => {
        let formattedData = []
        let rawData = this.props.studentRegistrations

        for(let i = 0; i < rawData.length; i++){
            let credits
            if(rawData[i].array[15] != null)
                credits = rawData[i].array[15]
            credits = 'N/A'
            let classData = {
                coursecode : rawData[i].array[5],
                coursenumber : rawData[i].array[2] + ':' + rawData[i].array[3] + ':' + rawData[i].array[4] + ':' + rawData[i].array[7],
                coursename : rawData[i].array[6],
                credits : credits,
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
                        <div style={{paddingLeft:"2px", paddingRight:"3.5%"}}>
                            <p title="Semester" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Semester 📘&nbsp;</p>
                            <Dropdown
                                name="semesters"
                                title={this.state.semesters[0].label}
                                list={this.state.semesters}
                                onChange={() => "nothing" }
                            />
                        </div>
                        <div style={{paddingLeft:"2px", paddingRight:"3.5%"}}>
                            <p title="Semester" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Currently Managing 👩‍💼&nbsp;</p>
                            <Dropdown
                                name="managing"
                                title={this.state.lists[0].label}
                                list={this.state.lists}
                                onChange={() => "nothing" }
                            />
                        </div>
                    </div>
                    <div class="studentManageReg-registrationContent">
                        <div>
                            <RegistrationTable validateLogin={this.props.validateLogin} updateClasses={this.updateClasses} classes={this.state.classes} studentRegistrations = {this.props.studentRegistrations} enableRegister={this.props.enableRegister} registerTime={this.props.registerTime} />
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
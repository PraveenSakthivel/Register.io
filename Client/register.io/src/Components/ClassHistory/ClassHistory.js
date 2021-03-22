import React from 'react';
import Schedule from './Schedule'

class ClassHistory extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            selectedSemester: 0,
            classes: data
        }
        this.dropdownSemesterHandler = this.dropdownSemesterHandler.bind(this);
    }
    
    dropdownSemesterHandler(e) {
        this.setState({selectedSemester:e.target.id});
    }

    render() {
        return (
            <div class="classHistory">
                <div class="classHistory-title">
                    <h3 style={{paddingBottom: "1%"}}>Class History</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>
                <div class="classHistory-content">
                    <div class="classHistory-header">
                        <h5 style={{fontSize:"18px", paddingLeft:"1%", paddingRight:"max(2%, 15px)", paddingTop:"max(.7%, 7px)"}}>Semester &nbsp;📘: </h5>
                        <div class="classHistory-dropdown">
                            <a style={{fontWeight:"500", fontSize:"15px"}} class="btn btn-secondary dropdown-toggle" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                {this.state.classes[this.state.selectedSemester].semester}
                            </a>

                            <div class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                                {this.state.classes.map(s => (<a id={s.id} onClick={this.dropdownSemesterHandler} class="dropdown-item" >{s.semester}</a>))} 
                            </div>
                        </div>  
                        <p style={{fontWeight:"500", flex: "1", textAlign:"right", paddingRight:"5%"}}>Credits 💰:&nbsp;&nbsp;17.0</p>
                    </div>
      
                    <Schedule classes={this.state.classes[this.state.selectedSemester].classes} locationColor={locationColor} />
                    
                    <div style={{paddingLeft:"2.5%"}}>
                        <p style={{fontWeight:"500", textDecoration:"underline"}}>Classes</p>
                        <div style={{paddingLeft:"0%"}}>
                            <li class="dashboard-class"><b style={{fontWeight:"500"}}>Digital Logic Design</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                            <li class="dashboard-class"><b style={{fontWeight:"500"}}>Art 101</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                            <li class="dashboard-class"><b style={{fontWeight:"500"}}>Public Speaking</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                            <li class="dashboard-class"><b style={{fontWeight:"500"}}>Digital Logic Design</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                            <li class="dashboard-class"><b style={{fontWeight:"500"}}>Digital Logic Design</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default ClassHistory;

const days = {
    "SUNDAY" : 0,
    "MONDAY" : 1,
    "TUESDAY" : 2,
    "WEDNESDAY" : 3,
    "THURSDAY" : 4,
    "FRIDAY" : 5,
    "SATURDAY" : 6
}


var createDate = (day, timing) => {
    var today = new Date();
    var newDate = new Date();
    newDate.setDate(today.getDate() + (days[day.toUpperCase()] - today.getDay()))
    let hour = "";
    let minute = "";
    for(let i = 0; i < timing.length; i++){
        if(i < 2)
            hour += timing.charAt(i);
        else
            minute += timing.charAt(i);
    }
    newDate = new Date(newDate.getFullYear(), newDate.getMonth(), newDate.getDate(), parseInt(hour), parseInt(minute), 0);

    return newDate;
}

const data = [
    {
        semester: "Spring 2021",
        id: 0,
        classes: [
            {
                title: "Digital Logic Design",
                location: "Busch",
                startDate: createDate("Monday", "1200"),
                endDate: createDate("Monday", "1320")
            }, {
                title: "Digital Logic Design",
                location: "Livingston",
                startDate: createDate("Wednesday", "1520"),
                endDate: createDate("Wednesday", "1640")
            }, {
                title: "Art 101",
                location: "Cook Douglass",
                startDate: createDate("Tuesday", "0840"),
                endDate: createDate("Tuesday", "1000")
            }, {
                title: "Public Speaking",
                location: "College Ave",
                startDate: createDate("Friday", "1200"),
                endDate: createDate("Friday", "1500")
            }
        ]
    },
    {
        semester: "Winter 2020",
        id: 1,
        classes: [
            {
                title: "Digital Logic Design",
                location: "Busch",
                startDate: createDate("Monday", "1200"),
                endDate: createDate("Monday", "1320")
            }, {
                title: "Digital Logic Design",
                location: "Livingston",
                startDate: createDate("Wednesday", "1520"),
                endDate: createDate("Wednesday", "1640")
            }, {
                title: "Art 101",
                location: "Cook Douglass",
                startDate: createDate("Tuesday", "0840"),
                endDate: createDate("Tuesday", "1000")
            }
        ]
    }
];

const locationColor = [
    {
        id: "Busch",
        color: '#e74c3c'
    }, {
        id: "Livingston",
        color: '#3498db'
    }, {
        id: "College Ave",
        color: '#2ecc71'
    }, {
        id: "Cook Douglass",
        color: '#f39c12'
    }
];
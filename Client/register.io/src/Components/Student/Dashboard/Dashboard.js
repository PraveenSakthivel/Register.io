import React from 'react';
import Schedule from './Schedule'
import { Dropdown } from 'reactjs-dropdown-component'

class Dashboard extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            notifications: [
                {
                    id: 0,
                    group: "Watchlist",
                    text: "Course 01:198:111 (Intro to Comp Sci) Opened!"
                },
                {
                    id: 1,
                    group: "Course Approved",
                    text: "All courses have been successfully added"
                },
                {
                    id: 2,
                    group: "SPN",
                    text: "SPN for Course 01:198:111 (Intro to Comp Sci) Approved!"
                },
                {
                    id: 3,
                    group: "Course Approved",
                    text: "All courses have been successfully added"
                },
                {
                    id: 4,
                    group: "Watchlist",
                    text: "Course 01:198:111 (Intro to Comp Sci) Opened!"
                },
                {
                    id: 5,
                    group: "Course Approved",
                    text: "All courses have been successfully added"
                }
            ]
        }
    }


    render() {
        let groupColor = {"Watchlist":"var(--color-primary)", "Course Approved":"#F79F1F", "SPN":"#009432"}
        let credits = 0
        this.props.studentRegistrations.map(n => (
            credits += n.getCredits()
        ))

        return (
            <div class="dashboard">
                <div class="dashboard-title">
                    <h3 style={{paddingBottom: "1%"}}>Dashboard</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>
                <div class="dashboard-content">
                    {/*
                    <div class="dashboard-notifs">
                        <h5 style={{fontSize:"18px", paddingLeft:"1%", paddingTop:"max(.5%, 5px)"}}>Notifications&nbsp;&nbsp;🔔</h5>
                        <div class="dashboard-notifs_sidescroll" style={{paddingTop:"5px",marginTop:"max(1.5%, 15px)", marginLeft:"1%"}}>
                            <div class="dashboard-notifs_cover" style={{position:"absolute", display:"-webkit-box", width:"84.5%", height:"inherit"}}>
                            </div>
                            {this.state.notifications.map(n => (
                                    <div id={n.id} style={{height:"75px", width:"250px", border:"1px solid lightgray", borderRadius:"4px", marginRight:"1.5%", borderLeft:"4px solid "+groupColor[n.group], padding:"10px", paddingTop:"7.5px"}}>
                                        <h6 style={{fontSize:"14px", marginBottom:"2.5px"}}>{n.group}</h6>
                                        <p style={{fontSize:"13px"}}>{n.text}</p>
                                    </div>
                                ))
                                } 
                            <div style={{height:"75px", width:"75px",  marginRight:"1.5%", padding:"10px", paddingTop:"7.5px"}}>
                            
                            </div>

                        </div>
                    </div>
                    */}
                    <div class="dashboard-schedule" style={{marginBottom:"100px"}}>
                        
                        {/*
                        <h5 style={{fontSize:"18px", paddingLeft:"1%"}}>My Schedule&nbsp;&nbsp;📅</h5>
                        */}
                        <div style={{display:"flex", marginTop:"20px", marginBottom:"1%"}}>
                            <div style={{paddingLeft:"2px", paddingRight:"15px", width:"fit-content"}}>
                                <p title="Semester" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Semester 📘&nbsp;</p>
                                <Dropdown
                                    name="semesters"
                                    title="Spring 2021"
                                    list={[{ label:"Spring 2021", value:"Spring 2021" }]}
                                    onChange={ () => "nohting" }
                                />
                            </div>
                            <div style={{display:"block", width:"100%", marginTop:"25px"}}>
                                <div style={{paddingLeft:"min(10px, 2.5%)", display:"flex"}}>
                                    <div style={{marginTop:"5.5px", marginRight:"5px", backgroundColor:"#e74c3c", height: "15px", width:"15px", border:"solid 2px black", borderRadius:"5px"}}></div>
                                        <p style={{fontWeight:"500"}}>: Busch &nbsp;&nbsp;</p>
                                    <div style={{marginTop:"5.5px", marginRight:"5px", backgroundColor:"#3498db", height: "15px", width:"15px", border:"solid 2px black", borderRadius:"5px"}}></div>
                                        <p style={{fontWeight:"500"}}>: Livingston &nbsp;&nbsp;</p>
                                    <div style={{marginTop:"5.5px", marginRight:"5px", backgroundColor:"#2ecc71", height: "15px", width:"15px", border:"solid 2px black", borderRadius:"5px"}}></div>
                                        <p style={{fontWeight:"500"}}>: College Ave &nbsp;&nbsp;</p>
                                    <div style={{marginTop:"5.5px", marginRight:"5px", backgroundColor:"#f39c12", height: "15px", width:"15px", border:"solid 2px black", borderRadius:"5px"}}></div>
                                        <p style={{fontWeight:"500"}}>: Cook Douglass</p>
                                    <p style={{fontWeight:"500", flex: "1", textAlign:"right", paddingRight:"5%"}}>Credits 💰:&nbsp;&nbsp;{credits}</p>
                                </div>
                            </div>
                        </div>
                        <Schedule studentRegistrations={this.props.studentRegistrations} />
                        <div style={{paddingLeft:"2.5%"}}>
                            <p style={{fontWeight:"500", textDecoration:"underline"}}>Classes</p>
                            <div style={{paddingLeft:"0%"}}>
                                {this.props.studentRegistrations.map(i => (
                                    <li class="dashboard-class"><b style={{fontWeight:"500"}}>{i.getName()}</b>&nbsp;&nbsp;({i.getSchool()}:{i.getDepartment()}:{i.getClassnumber()})&nbsp;&nbsp;Credits: {i.getCredits()}&nbsp;&nbsp;Location: {i.getLocation()}</li>
                                ))}
                            {/*
                                <li class="dashboard-class"><b style={{fontWeight:"500"}}>Digital Logic Design</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                                <li class="dashboard-class"><b style={{fontWeight:"500"}}>Art 101</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                                <li class="dashboard-class"><b style={{fontWeight:"500"}}>Public Speaking</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                                <li class="dashboard-class"><b style={{fontWeight:"500"}}>Digital Logic Design</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                                <li class="dashboard-class"><b style={{fontWeight:"500"}}>Digital Logic Design</b> (14:332:226:01) Credits: 3.0 Location: By Arrangement</li>
                            */}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Dashboard;
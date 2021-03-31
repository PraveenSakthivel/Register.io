import React from 'react';
import CourseTable from './CourseTable';



class CourseLookup extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            selectedDept: 0,
            selectedSemester: 0,
            classes: data
        }
        this.dropdownDeptHandler = this.dropdownDeptHandler.bind(this);
        this.dropdownSemesterHandler = this.dropdownSemesterHandler.bind(this);
    }

    dropdownDeptHandler(e) {
        this.setState({selectedDept:e.target.id});
    }

    dropdownSemesterHandler(e) {
        this.setState({selectedSemester:e.target.id});
    }

    render() {
        return (
            <div class="courseLookup">
                <div class="courseLookup-title">
                    <h3 style={{paddingBottom: "1%"}}>Course Lookup</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>

                <div class="courseLookup-header"> 
                    <h5 style={{fontSize:"18px", paddingLeft:"1%", paddingRight:"max(2%, 15px)", paddingTop:"max(.7%, 7px)"}}>Semester &nbsp;📘: </h5>      
                    <div class="courseLookup-dropdown">
                        <a style={{fontWeight:"500", fontSize:"15px"}} class="btn btn-secondary dropdown-toggle" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            {this.state.classes[this.state.selectedSemester].semester}
                        </a>

                        <div class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                            {this.state.classes.map(s => (<a id={s.id} onClick={this.dropdownSemesterHandler} class="dropdown-item" >{s.semester}</a>))} 
                        </div>
                    </div>
                    <h4 style={{paddingTop:".3%"}}>&nbsp;&nbsp;&nbsp;•&nbsp;&nbsp;&nbsp;</h4>
                    <h5 style={{fontSize:"18px", paddingRight:"max(2%, 15px)", paddingTop:"max(.7%, 7px)"}}>Department &nbsp;🏫: </h5>      
                    <div class="courseLookup-dropdown">
                        <a style={{fontWeight:"500", fontSize:"15px"}} class="btn btn-secondary dropdown-toggle" role="button" id="dropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            {this.state.classes[this.state.selectedSemester].departments[this.state.selectedDept].department}
                        </a>

                        <div class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                            {this.state.classes[this.state.selectedSemester].departments.map(s => (<a id={s.id} onClick={this.dropdownDeptHandler} class="dropdown-item" >{s.department}</a>))} 
                        </div>
                    </div>

                </div>
                <div class="courseLookup-content">
                    <CourseTable data={this.state.classes[this.state.selectedSemester].departments[this.state.selectedDept].courses} />
                </div>
            </div>
        );
    }
}

export default CourseLookup;

const data = 
[
    {
        semester: "Winter 2021",
        id: 0,
        departments:[
            {
                department: "Computer Science", 
                id: 0,
                courses: [
                    {
                        data: { name: 'Wintro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    }
                ]
            },
            {
                department: "Electrical Engineering", 
                id: 1,
                courses: [
                    {
                        data: { name: 'Winter Systems & Signals', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 03', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 04', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Linear Systems & Signals', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 03', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 04', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    }
                ]
            }
        ]
    },
    {
        semester: "Spring 2021",
        id: 1,
        departments:[
            {
                department: "Computer Science", 
                id: 0,
                courses: [
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    }
                ]
            },
            {
                department: "Electrical Engineering", 
                id: 1,
                courses: [
                    {
                        data: { name: 'Linear Systems & Signals', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 03', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 04', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Linear Systems & Signals', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 03', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                            },
                            { 
                            data: { section: 'Section 04', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                        data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    },
                    {
                    data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                    children: [
                        { 
                        data: { section: 'Section 01', status: 'open', index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                        },
                        { 
                        data: { section: 'Section 02', status: 'closed', index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                        }
                    ]
                    },
                    {
                        data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                        children: [
                            { 
                            data: { section: 'Section 01', status: 'open', index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                            }
                        ]
                    }
                ]
            }
        ]
    }
];
import React from 'react';
import CourseTable from './CourseTable';

import SearchField from 'react-search-field'
import { Dropdown } from 'reactjs-dropdown-component'


class CourseLookup extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            selectedDept: -1,
            selectedSemester: 0,
            classes: data,
            visibleClasses: [],
            searchTerm: "",
            viewOpenSections: true,
            viewClosedSections: true,
            depts : [],
            soc : [],
            visiblesoc: []
        }
        this.dropdownDeptHandler = this.dropdownDeptHandler.bind(this);
        this.searchClasses = this.searchClasses.bind(this);
        this.wrapper = this.wrapper.bind(this)
        this.wrapper2 = this.wrapper2.bind(this)
    }

    componentDidMount(){
        this.setState({depts : this.formatDepts()}) // eventually move this so that it runs only after class list is retrieved
        this.setState({soc : this.props.soc})
    }

    dropdownDeptHandler(item, name) {
        this.setState({ selectedDept : item.value })
        let socCpy = this.state.soc
        let result = socCpy.filter(word => word.data.department == item.value);
        result = result.sort(function(a, b) {
            return a.data.courseCode.localeCompare(b.data.courseCode, undefined, {
              numeric: true,
              sensitivity: 'base'
            });
          });
        this.setState({visiblesoc : result})
        return result
    }

    formatDepts() {
        let depts = this.state.classes
        let formatted = []
        for(let i = 0; i < depts.length; i++){
            let dept = depts[i]
            formatted.push({ label: dept.department, value: i })
        }

        return [ { label:"Computer Science", value: 198 }, { label: "Math", value: 440 } ]
    }

    wrapper(item, name){
        this.searchClasses(this.state.searchTerm, item)
    }

    wrapper2 ( term ){
        this.searchClasses(term, null)
    }

    searchClasses( term, item ){
        this.setState({ searchTerm : term })
        if(item == null){
            item = {value : this.state.selectedDept}
        }
        this.setState({ visiblesoc : this.searchClassesHelper( this.dropdownDeptHandler(item, ''), term ) })
    }

    searchClassesHelper( data, term ){
        if(term == '')
            return data
        else{
            let newData = []
            term = term.toLowerCase()
            for(let i = 0; i < data.length; i++){
                let tempName = data[i].data.name.toLowerCase()
                if(tempName.includes(term)){
                    newData.push(data[i])
                }
            }
            return newData
        }
    }

    render() {

        return (
            <div class="courseLookup">
                <div class="courseLookup-title">
                    <h3 style={{paddingBottom: "1%"}}>Course Lookup</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>

                <div class="courseLookup-header"> 

                    <div style={{paddingLeft:"2px", marginRight:"3.5%", width:"max(350px, 30%)"}}>
                        <p title="Search by Class Name" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Search 🔍&nbsp;</p>
                        <SearchField 
                            placeholder='Search by Class Name'
                            onChange={this.wrapper2}
                        />
                    </div>

                    <div class="courseLookup-dropdown" style={{paddingRight:"3.5%"}}>
                        <p title="Select Department" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Department 🏬&nbsp;</p>
                        <Dropdown
                            name="departments"
                            title="Select Department"
                            searchable={["Search for Department", "No matching department"]}
                            list={this.state.depts}
                            onChange={this.wrapper}
                        />
                    </div>

                    <div style={{fontSize:"15px", paddingRight:"1.5%"}} >
                        <p title="Show me open and/or closed sections" style={{fontSize:"12px", marginBottom:'9px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Section Status 🔓&nbsp;</p>
                        <div class="form-check form-check-inline" >
                            <input class="form-check-input" type="checkbox" id="openSections" value="option1" defaultChecked></input>
                            <label class="form-check-label" for="inlineCheckbox1">Open</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="checkbox" id="closedSections" value="option2" defaultChecked></input>
                            <label class="form-check-label" for="inlineCheckbox2">Closed</label>
                        </div>
                    </div>

                    <div style={{fontSize:"15px"}} >
                        <p title="Classes that don't conflict with current registrations" style={{fontSize:"12px", marginBottom:'9px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Fits My Schedule 🧩&nbsp;</p>
                        <div class="form-check form-check-inline" >
                            <input class="form-check-input" type="checkbox" id="openSections" value="option1"></input>
                            <label class="form-check-label" for="inlineCheckbox1">Compatible Only</label>
                        </div>
                    </div>

                </div>
                <div class="courseLookup-content">
                    <CourseTable data={ this.state.visiblesoc } enableRegister={this.props.enableRegister} studentRegistrations={this.props.studentRegistrations} />
                </div>
            </div>
        );
    }
}

export default CourseLookup;



const data = 
[
    {
        department: "Computer Science", 
        id: 0,
        courses: [
            {
                data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '14198', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Intro to Computer Science', courseCode: '01:198:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Data Structures', courseCode: '01:198:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
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
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 03', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 04', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Linear Systems & Signals', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 03', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                    },
                    { 
                    data: { section: 'Section 04', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                    }
                ]
            },
            {
                data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
                data: { name: 'Digital Logic Design', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            },
            {
            data: { name: 'Information & Network Security', courseCode: '14:332:111', credits: '4cr', openSections: 3, closedSections: 2 },
            children: [
                { 
                data: { section: 'Section 01', status: true, index: '08384', time: '(MTh 6:40-8:00pm),(W 7:00-9:00am)', location: 'Busch', instructor:'Centeno, Ana' },
                },
                { 
                data: { section: 'Section 02', status: false, index: '68959', time: 'By Arrangement', location: 'Busch', instructor:'Centeno, Ana' },
                }
            ]
            },
            {
                data: { name: 'Principles of Electrical Engg I', courseCode: '14:332:112', credits: '4cr', openSections: 1, closedSections: 2 },
                children: [
                    { 
                    data: { section: 'Section 01', status: true, index: '29849', time: 'By Arrangement', location: 'Busch', instructor:'Venugopal, Sesh' },
                    }
                ]
            }
        ]
    }
];
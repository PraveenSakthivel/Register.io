import React from 'react';
import 'devextreme/dist/css/dx.common.css';
import 'devextreme/dist/css/dx.light.css';
import { classes, locationColor } from "./ScheduleData"
 
import { Scheduler, Resource } from 'devextreme-react/scheduler';


class Schedule extends React.Component {


    constructor(props) {
        super(props);
        this.state = {
            events: this.regFormatter(this.props.studentRegistrations)
        }
    }

    componentDidUpdate(prevProps){
        if(this.props.studentRegistrations != prevProps.studentRegistrations)
            this.setState({ events: this.regFormatter(this.props.studentRegistrations) })
    }

    createDate = (day, timing) => {
        var today = new Date();
        var newDate = new Date();
        newDate.setDate(today.getDate() + (day - today.getDay()))
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

    createEvent = (name, time, loc) => {
        let events = []
        let times = time.split("|");
        let locs = loc.split("|");
        
        for(let i = 0; i < times.length; i++){ // times and locs should be same length
            let t = times[i]
            let startTime
            let endTime
            let day
            
            day = days[t.charAt(0)]
            startTime = t.charAt(1) + t.charAt(2) + t.charAt(4) + t.charAt(5)
            endTime = t.charAt(10) + t.charAt(11) + t.charAt(13) + t.charAt(14)

            events.push({ title: name, location: (locs[i] != null) ? locs[i] : locs[0], startDate: this.createDate(day, startTime), endDate: this.createDate(day, endTime) })
        }

        return events
    }

    regFormatter = (regs) => {
        let events = []
        for(let i = 0; i < regs.length; i++){
            if(regs[i].getMeetingtimes().toLowerCase() == 'by arrangement'){
                continue;
            }
            else{
                Array.prototype.push.apply(events,this.createEvent(this.titleCase(regs[i].getName()), regs[i].getMeetingtimes(), regs[i].getLocation()))
            }
        }
        console.log(events)
        return events
    }

    titleCase(str) {
        var splitStr = str.toLowerCase().split(' ');
        for (var i = 0; i < splitStr.length; i++) {
            // You do not need to check if i is larger than splitStr length, as your for does that for you
            // Assign it back to the array
            splitStr[i] = splitStr[i].charAt(0).toUpperCase() + splitStr[i].substring(1);     
            splitStr[i] = splitStr[i].replaceAll('Ii', 'II')   
            splitStr[i] = splitStr[i].replaceAll('Iii', 'III')
        }
        // Directly return the joined string
        return splitStr.join(' '); 
     }

    render() {

        /*
            skip "By arrangement"
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
            }
        */

        // Date Format
        // R10:55 AM-12:15 PM|T10:55 AM-12:15 PM|M10:55 AM-12:15 PM    or    By Arrangement
        // M, T, W, R, F

        // Location Format
        // LIVINGSTON|BUSCH|DOUGLAS/COOK     or     N/A

        return (
            <div class="schedule">
                <Scheduler 
                    id="scheduler"
                    textExpr="title"
                    currentView="workWeek"
                    startDayHour={8}
                    endDayHour={22}
                    timeZone="America/New_York"
                    showAllDayPanel={false}
                    editing={{ allowAdding: false, allowDeleting: false, allowDragging: false, allowResizing: false, allowUpdating: false, allowEditingTimeZones: false }}
                    cellDuration="60"
                    dataSource={this.state.events}
                    onAppointmentDblClick= {function(e) {e.cancel = true;}}
                    onAppointmentClick= {function(e) {e.cancel = true;}}
                >
                    
                    <Resource
                        dataSource={locationColor}
                        fieldExpr="location"
                        useColorAsDefault={true}
                        />
                </Scheduler>
            </div>
        );
    }
}

export default Schedule;


var days = {
    "M" : 1,
    "T" : 2,
    "W" : 3,
    "R" : 4,
    "F" : 5       
}
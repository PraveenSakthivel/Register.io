import React from 'react';
import Plotly from "plotly.js"
import createPlotlyComponent from 'react-plotly.js/factory';
import { Dropdown } from 'reactjs-dropdown-component'
import Slider, { Range } from 'rc-slider';
import 'rc-slider/assets/index.css';

// backend
import { GetHeatMapData, GetBarGraphData } from '../../../Protobuf/RequestMaker'

const Plot = createPlotlyComponent(Plotly);

class Dashboard extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            barGraphData: [],
            heatMapData: [],
            daytimes: [],
            days: [],
            selectedDay: 'Sunday',
            selectedTime: '8:40 AM',
            marks: [],
            times: []
        };
        this.barGraphCallback = this.barGraphCallback.bind(this);
        this.heatMapCallback = this.heatMapCallback.bind(this); 
        this.sliderHandler = this.sliderHandler.bind(this);
        this.getMarks = this.getMarks.bind(this);
    }

    componentDidMount(){
        GetBarGraphData( {}, this.barGraphCallback )
        GetHeatMapData( {}, this.heatMapCallback )
    }

    barGraphCallback(res){
        let data = []
        let times = []
        let vals = []

        for (const [key, value] of res.entries()) {
            times.push(key)
            vals.push(value)
        }

        data.push({x:times, y:vals})
        this.setState({barGraphData:data})
    }

    heatMapCallback(res){
        let data = []
        let days = []
        let times = []

        for (const [key, value] of res.entries()) {
            let temp = res.get(key).getTimesMap() //.get("17:00 PM").getDataMap()  
            let tempTimes = []
            let actualData = []
            
            for (const [time, val] of temp.entries()) {
                let locs = temp.get(time).getDataMap()
                let dat = []
                for (const [loc, v] of locs.entries()){
                    if(loc == "N/A")
                        continue;
                    dat[loc] = v
                }

                tempTimes.push(time)
                actualData[time] = dat
            }
            
            data[key] = actualData
            times[key] = tempTimes
            if(key != "By Arrangement")
                days.push({label:key, value:key})
        }

        console.log(times)
        console.log(data)
        this.setState({days: days})
        this.setState({daytimes:times})
        this.setState({heatMapData:data})
        
    } 

    generateMap(){

        let data = this.state.heatMapData
        data = data[this.state.selectedDay]
        if(data == null)
            return <h3 style={{padding:"5%"}}>No Data To Show At This Time</h3>;
        data = data[this.state.selectedTime]
        if(data == null || Object.keys(data).length == 0)
            return <h3 style={{padding:"5%"}}>No Data To Show At This Time</h3>;
        
        let lat = []
        let lon = []
        let z = []
        for(const [key, val] of Object.entries(data)){
            let d = key.split("|")
            lat.push(d[0])
            lon.push(d[1])
            z.push(val)
        }

        var dat = [{type: 'densitymapbox', colorscale: 'YlOrRd', lon: lon, lat: lat, z: z}];

        var layout = {width:1000, height:600, title:'Where Students Are On Campus', mapbox: {style: 'open-street-map',  coloraxis: {colorscale: "Viridis"}, center: { lat:40.4963, lon:-74.4412 }, zoom: 12}};

        return <Plot
                    data={dat}
                    layout={layout}
                />
    }

    generateBarGraph(){
        let data = this.state.barGraphData
        if(data[0] == null)
            return;
        return <Plot
                    data={[
                    {type: 'bar', x: data[0].x, y: data[0].y},
                    ]}
                    layout={ {width: 1100, height: 400, title: 'Most Popular Section Times'} }
                />
    }

    getMarks(day){
        if(this.state.daytimes[day] == null)
            return
        let timess = this.state.daytimes[day]
        let res = {}
        let i = 0;
        for(const [key, value] of Object.entries(timess)){
            res[i] = value
            i++
        }
        return res
    }

    sliderHandler(e){
        let time = this.getMarks(this.state.selectedDay)[e]
        this.setState({selectedTime:time})
    }

    handleDropdown(item, name){
        this.setState({selectedDay:item.value});
        this.setState({selectedTime:this.getMarks(item.value)[0]})
        
        this.setState({times : this.getMarks(this.state.selectedDay, this.state.daytimes[item.value])})
    }

    render() {

        /*let style = {
            borderColor:"var(--color-primary)",
            ':active':{
                boxShadow: "0 0 5px var(--color-primary-light)"
            },
            ':focus':{
                boxShadow: "0 0 5px var(--color-primary-light)"
            }
        }*/
        let times = this.getMarks(this.state.selectedDay)
        return (
            <div class="dashboardAdmin">
                <div style={{}} class="dashboardAdmin-title">
                    <h3 style={{paddingBottom: "1%"}}>Analytics Dashboard</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>
                <div style={{display:"flex", flexWrap: 'wrap', width:"inherit"}} class="dashboardAdmin-content" >
                    <div style={{boxShadow: '0 2px 4px 0 rgba(0,0,0,0.2)', width:"fit-content", minWidth:"1100px", height:"fit-content", margin:"10px"}}>
                        <div style={{display:"flex"}}>
                            <div style={{marginTop:"2.5%", marginLeft:"2.5%"}}>
                                <p title="Day Of Week" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Day of Week 📅&nbsp;</p>
                                <Dropdown
                                    name="dayOfWeek"
                                    title={"Select Day"}
                                    list={this.state.days}
                                    onChange={(item, name) => 
                                        this.handleDropdown(item, name)
                                    }
                                    />
                            </div>
                            <div style={{width:"100%", marginLeft:"5%", marginRight:"5%", marginTop:"2.5%", marginBottom:"5%"}}>
                                {/*
                                    <p title="Day Of Week" style={{fontSize:"12px", marginBottom:'1px', fontFamily:'Lato', width:'fit-content'}}>&nbsp;Day of Week 📅&nbsp;</p>
                                <Dropdown
                                    name="timeOfDay"
                                    title={"Select Time"}
                                    list={this.state.times}
                                    onChange={(item, name) => this.setState({selectedTime:item.value})}
                                    />
                                */}
                                <Slider onChange={this.sliderHandler}  included={false} min={0} max={(times != null) ? Object.keys(times).length - 1 : 0} defaultValue={0} marks={times} step={null} /> 
                            </div>
                        </div>
                        <div>
                            {this.generateMap()}
                        </div>
                    </div>
                    <div style={{boxShadow: '0 2px 4px 0 rgba(0,0,0,0.2)', width:"fit-content", height:"fit-content", margin:"10px"}}>{this.generateBarGraph()}</div>
                </div>
            </div>
        );
    }
}

export default Dashboard;
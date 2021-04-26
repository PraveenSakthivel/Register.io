import React from 'react';
import Plotly from "plotly.js"
import createPlotlyComponent from 'react-plotly.js/factory';

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
            selectedDay: 'Thursday',
            selectedTime: '15:00 PM'
        };
        this.barGraphCallback = this.barGraphCallback.bind(this);
        this.heatMapCallback = this.heatMapCallback.bind(this); 
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
            days.push(key)
        }

        console.log(times)
        console.log(data)
        this.setState({daytimes:times})
        this.setState({heatMapData:data})
    } 

    generateMap(){

        let data = this.state.heatMapData
        data = data[this.state.selectedDay]
        if(data == null)
            return;
        data = data[this.state.selectedTime]
        if(data == null)
            return;
        
        let lat = []
        let lon = []
        let z = []
        for(const [key, val] of Object.entries(data)){
            let d = key.split("|")
            lat.push(d[0])
            lon.push(d[1])
            z.push(val)
        }

        var dat = [{type: 'densitymapbox', lon: lon, lat: lat, z: z}];

        var layout = {width: 600, height: 600, title:'Where Students Are On Campus', mapbox: {style: 'open-street-map', center: { lat:40.4963, lon:-74.4412 }, zoom: 12}};

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
                    layout={ {width: 600, height: 400, title: 'Most Popular Section Times'} }
                />
    }

    render() {

        return (
            <div class="dashboardAdmin">
                <div style={{}} class="dashboardAdmin-title">
                    <h3 style={{paddingBottom: "1%"}}>Analytics Dashboard</h3>
                    <hr style={{color:"grey", marginRight:"15%", marginBottom:"0"}}></hr>
                </div>
                <div style={{display:"flex", flexWrap: 'wrap', width:"inherit"}} class="dashboardAdmin-content" >
                    <div style={{boxShadow: '0 2px 4px 0 rgba(0,0,0,0.2)', width:"fit-content", height:"fit-content", margin:"10px"}}>{this.generateMap()}</div>
                    <div style={{boxShadow: '0 2px 4px 0 rgba(0,0,0,0.2)', width:"fit-content", height:"fit-content", margin:"10px"}}>{this.generateBarGraph()}</div>
                </div>
            </div>
        );
    }
}

export default Dashboard;
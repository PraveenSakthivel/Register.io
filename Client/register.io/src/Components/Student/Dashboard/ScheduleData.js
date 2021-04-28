import React from 'react';

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

export const classes = [
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

];

export const locationColor = [
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
    },  {
        id: "Douglas/Cook",
        color: '#f39c12'
    }
];
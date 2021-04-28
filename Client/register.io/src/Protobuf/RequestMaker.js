import {endpoint} from './endpoint.json'

/* ----------------------------------------------------------------------------------------- */

// USER VALIDATION Stuff
const { Credentials, Token } = require('./UserV/token_pb.js');
const { LoginEndpointClient } = require('./UserV/token_grpc_web_pb.js');

// data: { netid:'rr973', password:'password' }
// response: 'Invalid Login' or { token: 'as8d7fa9sduf8' }
export const LoginRequest = ( data, callback ) => {
    var client = new LoginEndpointClient(endpoint)

    var request = new Credentials();
    request.setNetid(data.netid);
    request.setPassword(data.password);

    client.getLoginToken(request, { "grpc_service" : "uv" }, (err, response) => {
        if(response.getToken() == '')
            callback('Invalid Login')
        else
            callback({ token: response.getToken() })
    });
}

// data: { token:'as8d7fa9sduf8' }
// response: { usertype: 0, (if student:) classlist: [ { array: [ 'Busch', '300', 14, 332, 448, .. ] }, {..} ] }
export const ValidateLogin = ( data, callback ) => {
    var client = new LoginEndpointClient(endpoint)

    let protoToken = new Token();
    protoToken.setToken(data.token);
    
    client.getCurrentRegistrations(protoToken, { "grpc_service" : "uv" }, (err, response) => {
        
        console.log(response)

        if(response.getUsertype() == 0)
            callback({ usertype: response.getUsertype(), classlist: response.getClassesList() })
        else
            callback({ usertype: response.getUsertype() })
    });
}

/* ----------------------------------------------------------------------------------------- */

// COURSE VALIDATION Stuff
const { RegistrationRequest, SPNRequest, ClassOperations, ReqOp } = require('./CV/cvInterface_pb');
const { CourseValidationClient } = require('./CV/cvInterface_grpc_web_pb');

// data: [ { val:12345, reqop:'ADD' }, {..}, {..} ]
// response: TBD
export const CVRequest = ( data, callback ) =>{

    let courses = []
    for(let i = 0; i < data.length; i++){

        let classOp = new ClassOperations();
        classOp.setIndex(data[i].val);
        switch(data[i].reqop){
            case('ADD'):
                classOp.setOp(ReqOp.ADD)
                break;
            case('DROP'):
                classOp.setOp(ReqOp.DROP)
                break;
        }
        courses.push(classOp)

    }

    if(courses.length == 0){
        callback([]);
        return;
    }
    
    var client = new CourseValidationClient(endpoint)

    var request = new RegistrationRequest();
    request.setClassesList(courses);
    request.setToken(window.sessionStorage.getItem("token").toString());
    
    client.changeRegistration(request, { "grpc_service" : "cv" }, (err, response) => {
        callback(response.getResultsMap(), data[0].reqop);
    });

}

/* ----------------------------------------------------------------------------------------- */

// REGISTRATION VALIDATION Stuff
const { Student } = require('./RV/rvInterface_pb.js');
const { RegistrationValidationClient } = require('./RV/rvInterface_grpc_web_pb.js');

// data: {} (empty)
// response: { eligible: true, time: 281280 }
export const RVRequest = ( data, callback ) =>{
    var client = new RegistrationValidationClient(endpoint)

    var request = new Student();
    request.setToken(window.sessionStorage.getItem("token").toString());

    
    client.checkRegVal(request, { "grpc_service" : "rv" }, (err, response) => {
        let res = { eligible: response.getEligible(), time: response.getTime() }
        callback(res)
    });
}

/* ----------------------------------------------------------------------------------------- */

// DATABASE REQUESTS Stuff
const { ReceiveClassesParams, ClassAddStatusParams, ReceiveDepartmentsParams } = require('./Database/dbRequests_pb.js')
const { DatabaseWrapperClient, default: dbRequests } = require('./Database/dbRequests_grpc_web_pb.js')

export const VerifyAdd = ( data, callback ) => {
    var client = new DatabaseWrapperClient(endpoint)

    var request = new ClassAddStatusParams();
    
    client.classAddStatus(request, { "grpc_service" : "db" }, (err, response) => {
        callback(response)
    });
}

// data: {} (empty)
// response: [ { Class }, { Class }, {..} ]
export const DBRetrieveCourses = ( data, callback ) =>{
    var client = new DatabaseWrapperClient(endpoint)

    var request = new ReceiveClassesParams();
    
    client.retrieveClasses(request, { "grpc_service" : "db" }, (err, response) => {
        callback(response)
    });
}

var deptMap = []

var logMapElements = (value, key, map) => {
    deptMap.push({ label: value, value: key })
}

// data: {} (empty)
// response: [ { Departments } ]
export const DBRetrieveDepts = ( data, callback ) =>{
    var client = new DatabaseWrapperClient(endpoint)

    var request = new ReceiveDepartmentsParams();
    
    client.returnDepartments(request, { "grpc_service" : "db" }, (err, response) => {
        deptMap = []
        response.getDepartmentsMap().forEach(logMapElements.bind(this)) 
        callback(deptMap)
    });
}

// data:
// response:
export const DBVerifyAdd = ( data, callback ) => {
    var client = new DatabaseWrapperClient(endpoint)

    var request = new ClassAddStatusParams();
    
    client.classAddStatus(request, { "grpc_service" : "db" }, (err, response) => {
        callback(response)
    });
}

/* ----------------------------------------------------------------------------------------- */

// ANALYTICS Stuff
const { Empty, DayofWeek } = require('./Analytics/analytics_pb')
const { AnalyticsEndpointClient } = require('./Analytics/analytics_grpc_web_pb')

export const GetHeatMapData = ( data, callback) => {
    var client = new AnalyticsEndpointClient(endpoint)

    var request = new Empty();

    client.getHeatmap(request, { "grpc_service" : "analytics" }, (err, response) => {
        let test = response.getDaysMap()
        callback(test)
    })
}

export const GetBarGraphData = ( data, callback) => {
    var client = new AnalyticsEndpointClient(endpoint)

    var request = new Empty();

    client.getBargraph(request, { "grpc_service" : "analytics" }, (err, response) => {
        let test = response.getDataMap()
        callback(test)
    })
}
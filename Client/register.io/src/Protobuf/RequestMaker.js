import {endpoint} from './endpoint.json'

const { RegistrationRequest, SPNRequest, ClassOperations, ReqOp, RegistrationResponse, SPNResponse } = require('./CV/cvInterface_pb');
const { CourseValidationClient } = require('./CV/cvInterface_grpc_web_pb');

// Course Validation Request
export const CVRequest = ( data ) =>{

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
            default:
                return 'ERROR'
        }
        courses.push(classOp)

    }
    console.log(courses)

    if(courses.length == 0)
        return;

    var client = new CourseValidationClient(endpoint)

    var request = new RegistrationRequest();
    request.setClassesList(courses);
    request.setToken(window.sessionStorage.getItem("token").toString());
    
    client.changeRegistration(request, { "grpc_service" : "cv" }, (err, response) => {
        console.log(response)
        return response;
    });

}

// Registration Validation Request
export const RVRequest = () =>{

}

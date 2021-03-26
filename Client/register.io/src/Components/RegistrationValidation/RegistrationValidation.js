const { googleRPC } = require("reactrpc")

const messages = require("../../Protobuf/rvInterface_pb")

const services = require("../../Protobuf/rvInterface_grpc_web_pb")

const URL = "http://3.92.240.128:8080"

googleRPC.build(messages, services, URL)

const msg = { NetID:"asd23", msgType:"Student" }
this.props.RegistrationValidation.checkRegVal(
  msg,
    {},
    (err, response) => {
      console.log(response)
    }
);
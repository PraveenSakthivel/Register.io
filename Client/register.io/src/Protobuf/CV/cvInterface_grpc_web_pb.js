/**
 * @fileoverview gRPC-Web generated client stub for cvInterface
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.cvInterface = require('./cvInterface_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.cvInterface.CourseValidationClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.cvInterface.CourseValidationPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cvInterface.RegistrationRequest,
 *   !proto.cvInterface.RegistrationResponse>}
 */
const methodDescriptor_CourseValidation_ChangeRegistration = new grpc.web.MethodDescriptor(
  '/cvInterface.CourseValidation/ChangeRegistration',
  grpc.web.MethodType.UNARY,
  proto.cvInterface.RegistrationRequest,
  proto.cvInterface.RegistrationResponse,
  /**
   * @param {!proto.cvInterface.RegistrationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cvInterface.RegistrationResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cvInterface.RegistrationRequest,
 *   !proto.cvInterface.RegistrationResponse>}
 */
const methodInfo_CourseValidation_ChangeRegistration = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cvInterface.RegistrationResponse,
  /**
   * @param {!proto.cvInterface.RegistrationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cvInterface.RegistrationResponse.deserializeBinary
);


/**
 * @param {!proto.cvInterface.RegistrationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cvInterface.RegistrationResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cvInterface.RegistrationResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cvInterface.CourseValidationClient.prototype.changeRegistration =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cvInterface.CourseValidation/ChangeRegistration',
      request,
      metadata || {},
      methodDescriptor_CourseValidation_ChangeRegistration,
      callback);
};


/**
 * @param {!proto.cvInterface.RegistrationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cvInterface.RegistrationResponse>}
 *     Promise that resolves to the response
 */
proto.cvInterface.CourseValidationPromiseClient.prototype.changeRegistration =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cvInterface.CourseValidation/ChangeRegistration',
      request,
      metadata || {},
      methodDescriptor_CourseValidation_ChangeRegistration);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.cvInterface.SPNRequest,
 *   !proto.cvInterface.SPNResponse>}
 */
const methodDescriptor_CourseValidation_AddSPN = new grpc.web.MethodDescriptor(
  '/cvInterface.CourseValidation/AddSPN',
  grpc.web.MethodType.UNARY,
  proto.cvInterface.SPNRequest,
  proto.cvInterface.SPNResponse,
  /**
   * @param {!proto.cvInterface.SPNRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cvInterface.SPNResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.cvInterface.SPNRequest,
 *   !proto.cvInterface.SPNResponse>}
 */
const methodInfo_CourseValidation_AddSPN = new grpc.web.AbstractClientBase.MethodInfo(
  proto.cvInterface.SPNResponse,
  /**
   * @param {!proto.cvInterface.SPNRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.cvInterface.SPNResponse.deserializeBinary
);


/**
 * @param {!proto.cvInterface.SPNRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.cvInterface.SPNResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.cvInterface.SPNResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.cvInterface.CourseValidationClient.prototype.addSPN =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/cvInterface.CourseValidation/AddSPN',
      request,
      metadata || {},
      methodDescriptor_CourseValidation_AddSPN,
      callback);
};


/**
 * @param {!proto.cvInterface.SPNRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.cvInterface.SPNResponse>}
 *     Promise that resolves to the response
 */
proto.cvInterface.CourseValidationPromiseClient.prototype.addSPN =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/cvInterface.CourseValidation/AddSPN',
      request,
      metadata || {},
      methodDescriptor_CourseValidation_AddSPN);
};


module.exports = proto.cvInterface;


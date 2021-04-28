/**
 * @fileoverview gRPC-Web generated client stub for rvInterface
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.rvInterface = require('./rvInterface_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rvInterface.RegistrationValidationClient =
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
proto.rvInterface.RegistrationValidationPromiseClient =
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
 *   !proto.rvInterface.Student,
 *   !proto.rvInterface.Response>}
 */
const methodDescriptor_RegistrationValidation_checkRegVal = new grpc.web.MethodDescriptor(
  '/rvInterface.RegistrationValidation/checkRegVal',
  grpc.web.MethodType.UNARY,
  proto.rvInterface.Student,
  proto.rvInterface.Response,
  /**
   * @param {!proto.rvInterface.Student} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rvInterface.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rvInterface.Student,
 *   !proto.rvInterface.Response>}
 */
const methodInfo_RegistrationValidation_checkRegVal = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rvInterface.Response,
  /**
   * @param {!proto.rvInterface.Student} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rvInterface.Response.deserializeBinary
);


/**
 * @param {!proto.rvInterface.Student} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rvInterface.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rvInterface.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rvInterface.RegistrationValidationClient.prototype.checkRegVal =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rvInterface.RegistrationValidation/checkRegVal',
      request,
      metadata || {},
      methodDescriptor_RegistrationValidation_checkRegVal,
      callback);
};


/**
 * @param {!proto.rvInterface.Student} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rvInterface.Response>}
 *     Promise that resolves to the response
 */
proto.rvInterface.RegistrationValidationPromiseClient.prototype.checkRegVal =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rvInterface.RegistrationValidation/checkRegVal',
      request,
      metadata || {},
      methodDescriptor_RegistrationValidation_checkRegVal);
};


module.exports = proto.rvInterface;


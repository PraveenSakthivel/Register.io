/**
 * @fileoverview gRPC-Web generated client stub for Tokens
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.Tokens = require('./token_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.Tokens.LoginEndpointClient =
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
proto.Tokens.LoginEndpointPromiseClient =
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
 *   !proto.Tokens.Credentials,
 *   !proto.Tokens.Response>}
 */
const methodDescriptor_LoginEndpoint_GetLoginToken = new grpc.web.MethodDescriptor(
  '/Tokens.LoginEndpoint/GetLoginToken',
  grpc.web.MethodType.UNARY,
  proto.Tokens.Credentials,
  proto.Tokens.Response,
  /**
   * @param {!proto.Tokens.Credentials} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Tokens.Response.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Tokens.Credentials,
 *   !proto.Tokens.Response>}
 */
const methodInfo_LoginEndpoint_GetLoginToken = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Tokens.Response,
  /**
   * @param {!proto.Tokens.Credentials} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Tokens.Response.deserializeBinary
);


/**
 * @param {!proto.Tokens.Credentials} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Tokens.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Tokens.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.Tokens.LoginEndpointClient.prototype.getLoginToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Tokens.LoginEndpoint/GetLoginToken',
      request,
      metadata || {},
      methodDescriptor_LoginEndpoint_GetLoginToken,
      callback);
};


/**
 * @param {!proto.Tokens.Credentials} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Tokens.Response>}
 *     Promise that resolves to the response
 */
proto.Tokens.LoginEndpointPromiseClient.prototype.getLoginToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Tokens.LoginEndpoint/GetLoginToken',
      request,
      metadata || {},
      methodDescriptor_LoginEndpoint_GetLoginToken);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Tokens.Token,
 *   !proto.Tokens.Registrations>}
 */
const methodDescriptor_LoginEndpoint_GetCurrentRegistrations = new grpc.web.MethodDescriptor(
  '/Tokens.LoginEndpoint/GetCurrentRegistrations',
  grpc.web.MethodType.UNARY,
  proto.Tokens.Token,
  proto.Tokens.Registrations,
  /**
   * @param {!proto.Tokens.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Tokens.Registrations.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Tokens.Token,
 *   !proto.Tokens.Registrations>}
 */
const methodInfo_LoginEndpoint_GetCurrentRegistrations = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Tokens.Registrations,
  /**
   * @param {!proto.Tokens.Token} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Tokens.Registrations.deserializeBinary
);


/**
 * @param {!proto.Tokens.Token} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Tokens.Registrations)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Tokens.Registrations>|undefined}
 *     The XHR Node Readable Stream
 */
proto.Tokens.LoginEndpointClient.prototype.getCurrentRegistrations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Tokens.LoginEndpoint/GetCurrentRegistrations',
      request,
      metadata || {},
      methodDescriptor_LoginEndpoint_GetCurrentRegistrations,
      callback);
};


/**
 * @param {!proto.Tokens.Token} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Tokens.Registrations>}
 *     Promise that resolves to the response
 */
proto.Tokens.LoginEndpointPromiseClient.prototype.getCurrentRegistrations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Tokens.LoginEndpoint/GetCurrentRegistrations',
      request,
      metadata || {},
      methodDescriptor_LoginEndpoint_GetCurrentRegistrations);
};


module.exports = proto.Tokens;


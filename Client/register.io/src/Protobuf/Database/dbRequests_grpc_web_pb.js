/**
 * @fileoverview gRPC-Web generated client stub for dbRequests
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.dbRequests = require('./dbRequests_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.dbRequests.DatabaseWrapperClient =
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
proto.dbRequests.DatabaseWrapperPromiseClient =
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
 *   !proto.dbRequests.ReceiveClassesParams,
 *   !proto.dbRequests.ClassesResponse>}
 */
const methodDescriptor_DatabaseWrapper_RetrieveClasses = new grpc.web.MethodDescriptor(
  '/dbRequests.DatabaseWrapper/RetrieveClasses',
  grpc.web.MethodType.UNARY,
  proto.dbRequests.ReceiveClassesParams,
  proto.dbRequests.ClassesResponse,
  /**
   * @param {!proto.dbRequests.ReceiveClassesParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.dbRequests.ClassesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dbRequests.ReceiveClassesParams,
 *   !proto.dbRequests.ClassesResponse>}
 */
const methodInfo_DatabaseWrapper_RetrieveClasses = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dbRequests.ClassesResponse,
  /**
   * @param {!proto.dbRequests.ReceiveClassesParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.dbRequests.ClassesResponse.deserializeBinary
);


/**
 * @param {!proto.dbRequests.ReceiveClassesParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.dbRequests.ClassesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.dbRequests.ClassesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.dbRequests.DatabaseWrapperClient.prototype.retrieveClasses =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/dbRequests.DatabaseWrapper/RetrieveClasses',
      request,
      metadata || {},
      methodDescriptor_DatabaseWrapper_RetrieveClasses,
      callback);
};


/**
 * @param {!proto.dbRequests.ReceiveClassesParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.dbRequests.ClassesResponse>}
 *     Promise that resolves to the response
 */
proto.dbRequests.DatabaseWrapperPromiseClient.prototype.retrieveClasses =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/dbRequests.DatabaseWrapper/RetrieveClasses',
      request,
      metadata || {},
      methodDescriptor_DatabaseWrapper_RetrieveClasses);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.dbRequests.ClassAddStatusParams,
 *   !proto.dbRequests.AddStatusResponse>}
 */
const methodDescriptor_DatabaseWrapper_ClassAddStatus = new grpc.web.MethodDescriptor(
  '/dbRequests.DatabaseWrapper/ClassAddStatus',
  grpc.web.MethodType.UNARY,
  proto.dbRequests.ClassAddStatusParams,
  proto.dbRequests.AddStatusResponse,
  /**
   * @param {!proto.dbRequests.ClassAddStatusParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.dbRequests.AddStatusResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dbRequests.ClassAddStatusParams,
 *   !proto.dbRequests.AddStatusResponse>}
 */
const methodInfo_DatabaseWrapper_ClassAddStatus = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dbRequests.AddStatusResponse,
  /**
   * @param {!proto.dbRequests.ClassAddStatusParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.dbRequests.AddStatusResponse.deserializeBinary
);


/**
 * @param {!proto.dbRequests.ClassAddStatusParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.dbRequests.AddStatusResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.dbRequests.AddStatusResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.dbRequests.DatabaseWrapperClient.prototype.classAddStatus =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/dbRequests.DatabaseWrapper/ClassAddStatus',
      request,
      metadata || {},
      methodDescriptor_DatabaseWrapper_ClassAddStatus,
      callback);
};


/**
 * @param {!proto.dbRequests.ClassAddStatusParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.dbRequests.AddStatusResponse>}
 *     Promise that resolves to the response
 */
proto.dbRequests.DatabaseWrapperPromiseClient.prototype.classAddStatus =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/dbRequests.DatabaseWrapper/ClassAddStatus',
      request,
      metadata || {},
      methodDescriptor_DatabaseWrapper_ClassAddStatus);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.dbRequests.ReceiveDepartmentsParams,
 *   !proto.dbRequests.DepartmentsResponse>}
 */
const methodDescriptor_DatabaseWrapper_ReturnDepartments = new grpc.web.MethodDescriptor(
  '/dbRequests.DatabaseWrapper/ReturnDepartments',
  grpc.web.MethodType.UNARY,
  proto.dbRequests.ReceiveDepartmentsParams,
  proto.dbRequests.DepartmentsResponse,
  /**
   * @param {!proto.dbRequests.ReceiveDepartmentsParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.dbRequests.DepartmentsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.dbRequests.ReceiveDepartmentsParams,
 *   !proto.dbRequests.DepartmentsResponse>}
 */
const methodInfo_DatabaseWrapper_ReturnDepartments = new grpc.web.AbstractClientBase.MethodInfo(
  proto.dbRequests.DepartmentsResponse,
  /**
   * @param {!proto.dbRequests.ReceiveDepartmentsParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.dbRequests.DepartmentsResponse.deserializeBinary
);


/**
 * @param {!proto.dbRequests.ReceiveDepartmentsParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.dbRequests.DepartmentsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.dbRequests.DepartmentsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.dbRequests.DatabaseWrapperClient.prototype.returnDepartments =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/dbRequests.DatabaseWrapper/ReturnDepartments',
      request,
      metadata || {},
      methodDescriptor_DatabaseWrapper_ReturnDepartments,
      callback);
};


/**
 * @param {!proto.dbRequests.ReceiveDepartmentsParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.dbRequests.DepartmentsResponse>}
 *     Promise that resolves to the response
 */
proto.dbRequests.DatabaseWrapperPromiseClient.prototype.returnDepartments =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/dbRequests.DatabaseWrapper/ReturnDepartments',
      request,
      metadata || {},
      methodDescriptor_DatabaseWrapper_ReturnDepartments);
};


module.exports = proto.dbRequests;


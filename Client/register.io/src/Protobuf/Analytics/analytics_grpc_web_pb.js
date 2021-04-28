/**
 * @fileoverview gRPC-Web generated client stub for Analytics
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.Analytics = require('./analytics_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.Analytics.AnalyticsEndpointClient =
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
proto.Analytics.AnalyticsEndpointPromiseClient =
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
 *   !proto.Analytics.Empty,
 *   !proto.Analytics.Heatmap>}
 */
const methodDescriptor_AnalyticsEndpoint_GetHeatmap = new grpc.web.MethodDescriptor(
  '/Analytics.AnalyticsEndpoint/GetHeatmap',
  grpc.web.MethodType.UNARY,
  proto.Analytics.Empty,
  proto.Analytics.Heatmap,
  /**
   * @param {!proto.Analytics.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Analytics.Heatmap.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Analytics.Empty,
 *   !proto.Analytics.Heatmap>}
 */
const methodInfo_AnalyticsEndpoint_GetHeatmap = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Analytics.Heatmap,
  /**
   * @param {!proto.Analytics.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Analytics.Heatmap.deserializeBinary
);


/**
 * @param {!proto.Analytics.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Analytics.Heatmap)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Analytics.Heatmap>|undefined}
 *     The XHR Node Readable Stream
 */
proto.Analytics.AnalyticsEndpointClient.prototype.getHeatmap =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Analytics.AnalyticsEndpoint/GetHeatmap',
      request,
      metadata || {},
      methodDescriptor_AnalyticsEndpoint_GetHeatmap,
      callback);
};


/**
 * @param {!proto.Analytics.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Analytics.Heatmap>}
 *     Promise that resolves to the response
 */
proto.Analytics.AnalyticsEndpointPromiseClient.prototype.getHeatmap =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Analytics.AnalyticsEndpoint/GetHeatmap',
      request,
      metadata || {},
      methodDescriptor_AnalyticsEndpoint_GetHeatmap);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Analytics.Empty,
 *   !proto.Analytics.Location>}
 */
const methodDescriptor_AnalyticsEndpoint_GetBargraph = new grpc.web.MethodDescriptor(
  '/Analytics.AnalyticsEndpoint/GetBargraph',
  grpc.web.MethodType.UNARY,
  proto.Analytics.Empty,
  proto.Analytics.Location,
  /**
   * @param {!proto.Analytics.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Analytics.Location.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Analytics.Empty,
 *   !proto.Analytics.Location>}
 */
const methodInfo_AnalyticsEndpoint_GetBargraph = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Analytics.Location,
  /**
   * @param {!proto.Analytics.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Analytics.Location.deserializeBinary
);


/**
 * @param {!proto.Analytics.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Analytics.Location)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Analytics.Location>|undefined}
 *     The XHR Node Readable Stream
 */
proto.Analytics.AnalyticsEndpointClient.prototype.getBargraph =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/Analytics.AnalyticsEndpoint/GetBargraph',
      request,
      metadata || {},
      methodDescriptor_AnalyticsEndpoint_GetBargraph,
      callback);
};


/**
 * @param {!proto.Analytics.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Analytics.Location>}
 *     Promise that resolves to the response
 */
proto.Analytics.AnalyticsEndpointPromiseClient.prototype.getBargraph =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/Analytics.AnalyticsEndpoint/GetBargraph',
      request,
      metadata || {},
      methodDescriptor_AnalyticsEndpoint_GetBargraph);
};


module.exports = proto.Analytics;


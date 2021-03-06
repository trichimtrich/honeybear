'use strict';

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

var _setup_error = require('./setup_error');

var _setup_error2 = _interopRequireDefault(_setup_error);

var _util = require('util');

var _kibana_index_mappings = require('./kibana_index_mappings');

module.exports = function (server) {
  var client = server.plugins.elasticsearch.client;
  var index = server.config().get('kibana.index');

  function handleError(message) {
    return function (err) {
      throw new _setup_error2['default'](server, message, err);
    };
  }

  return client.indices.create({
    index: index,
    body: {
      settings: {
        number_of_shards: 1
      },
      mappings: _kibana_index_mappings.mappings
    }
  })['catch'](handleError('Unable to create Kibana index "<%= kibana.index %>"')).then(function () {
    return client.cluster.health({
      waitForStatus: 'yellow',
      index: index
    })['catch'](handleError('Waiting for Kibana index "<%= kibana.index %>" to come online failed.'));
  });
};

/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _JobsActionOutputFilter = require('./JobsActionOutputFilter');

var _JobsActionOutputFilter2 = _interopRequireDefault(_JobsActionOutputFilter);

var _JobsContextMetaFilter = require('./JobsContextMetaFilter');

var _JobsContextMetaFilter2 = _interopRequireDefault(_JobsContextMetaFilter);

var _JobsIdmSelector = require('./JobsIdmSelector');

var _JobsIdmSelector2 = _interopRequireDefault(_JobsIdmSelector);

var _JobsNodesSelector = require('./JobsNodesSelector');

var _JobsNodesSelector2 = _interopRequireDefault(_JobsNodesSelector);

var _JobsUsersSelector = require('./JobsUsersSelector');

var _JobsUsersSelector2 = _interopRequireDefault(_JobsUsersSelector);

/**
* The JobsAction model module.
* @module model/JobsAction
* @version 1.0
*/

var JobsAction = (function () {
    /**
    * Constructs a new <code>JobsAction</code>.
    * @alias module:model/JobsAction
    * @class
    */

    function JobsAction() {
        _classCallCheck(this, JobsAction);

        this.ID = undefined;
        this.Label = undefined;
        this.Description = undefined;
        this.Bypass = undefined;
        this.BreakAfter = undefined;
        this.NodesSelector = undefined;
        this.UsersSelector = undefined;
        this.NodesFilter = undefined;
        this.UsersFilter = undefined;
        this.IdmSelector = undefined;
        this.IdmFilter = undefined;
        this.ActionOutputFilter = undefined;
        this.ContextMetaFilter = undefined;
        this.Parameters = undefined;
        this.ChainedActions = undefined;
        this.FailedFilterActions = undefined;
    }

    /**
    * Constructs a <code>JobsAction</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/JobsAction} obj Optional instance to populate.
    * @return {module:model/JobsAction} The populated <code>JobsAction</code> instance.
    */

    JobsAction.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobsAction();

            if (data.hasOwnProperty('ID')) {
                obj['ID'] = _ApiClient2['default'].convertToType(data['ID'], 'String');
            }
            if (data.hasOwnProperty('Label')) {
                obj['Label'] = _ApiClient2['default'].convertToType(data['Label'], 'String');
            }
            if (data.hasOwnProperty('Description')) {
                obj['Description'] = _ApiClient2['default'].convertToType(data['Description'], 'String');
            }
            if (data.hasOwnProperty('Bypass')) {
                obj['Bypass'] = _ApiClient2['default'].convertToType(data['Bypass'], 'Boolean');
            }
            if (data.hasOwnProperty('BreakAfter')) {
                obj['BreakAfter'] = _ApiClient2['default'].convertToType(data['BreakAfter'], 'Boolean');
            }
            if (data.hasOwnProperty('NodesSelector')) {
                obj['NodesSelector'] = _JobsNodesSelector2['default'].constructFromObject(data['NodesSelector']);
            }
            if (data.hasOwnProperty('UsersSelector')) {
                obj['UsersSelector'] = _JobsUsersSelector2['default'].constructFromObject(data['UsersSelector']);
            }
            if (data.hasOwnProperty('NodesFilter')) {
                obj['NodesFilter'] = _JobsNodesSelector2['default'].constructFromObject(data['NodesFilter']);
            }
            if (data.hasOwnProperty('UsersFilter')) {
                obj['UsersFilter'] = _JobsUsersSelector2['default'].constructFromObject(data['UsersFilter']);
            }
            if (data.hasOwnProperty('IdmSelector')) {
                obj['IdmSelector'] = _JobsIdmSelector2['default'].constructFromObject(data['IdmSelector']);
            }
            if (data.hasOwnProperty('IdmFilter')) {
                obj['IdmFilter'] = _JobsIdmSelector2['default'].constructFromObject(data['IdmFilter']);
            }
            if (data.hasOwnProperty('ActionOutputFilter')) {
                obj['ActionOutputFilter'] = _JobsActionOutputFilter2['default'].constructFromObject(data['ActionOutputFilter']);
            }
            if (data.hasOwnProperty('ContextMetaFilter')) {
                obj['ContextMetaFilter'] = _JobsContextMetaFilter2['default'].constructFromObject(data['ContextMetaFilter']);
            }
            if (data.hasOwnProperty('Parameters')) {
                obj['Parameters'] = _ApiClient2['default'].convertToType(data['Parameters'], { 'String': 'String' });
            }
            if (data.hasOwnProperty('ChainedActions')) {
                obj['ChainedActions'] = _ApiClient2['default'].convertToType(data['ChainedActions'], [JobsAction]);
            }
            if (data.hasOwnProperty('FailedFilterActions')) {
                obj['FailedFilterActions'] = _ApiClient2['default'].convertToType(data['FailedFilterActions'], [JobsAction]);
            }
        }
        return obj;
    };

    /**
    * @member {String} ID
    */
    return JobsAction;
})();

exports['default'] = JobsAction;
module.exports = exports['default'];

/**
* @member {String} Label
*/

/**
* @member {String} Description
*/

/**
* @member {Boolean} Bypass
*/

/**
* @member {Boolean} BreakAfter
*/

/**
* @member {module:model/JobsNodesSelector} NodesSelector
*/

/**
* @member {module:model/JobsUsersSelector} UsersSelector
*/

/**
* @member {module:model/JobsNodesSelector} NodesFilter
*/

/**
* @member {module:model/JobsUsersSelector} UsersFilter
*/

/**
* @member {module:model/JobsIdmSelector} IdmSelector
*/

/**
* @member {module:model/JobsIdmSelector} IdmFilter
*/

/**
* @member {module:model/JobsActionOutputFilter} ActionOutputFilter
*/

/**
* @member {module:model/JobsContextMetaFilter} ContextMetaFilter
*/

/**
* @member {Object.<String, String>} Parameters
*/

/**
* @member {Array.<module:model/JobsAction>} ChainedActions
*/

/**
* @member {Array.<module:model/JobsAction>} FailedFilterActions
*/

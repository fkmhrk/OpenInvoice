/******/ (function(modules) { // webpackBootstrap
/******/ 	// install a JSONP callback for chunk loading
/******/ 	function webpackJsonpCallback(data) {
/******/ 		var chunkIds = data[0];
/******/ 		var moreModules = data[1];
/******/
/******/
/******/ 		// add "moreModules" to the modules object,
/******/ 		// then flag all "chunkIds" as loaded and fire callback
/******/ 		var moduleId, chunkId, i = 0, resolves = [];
/******/ 		for(;i < chunkIds.length; i++) {
/******/ 			chunkId = chunkIds[i];
/******/ 			if(installedChunks[chunkId]) {
/******/ 				resolves.push(installedChunks[chunkId][0]);
/******/ 			}
/******/ 			installedChunks[chunkId] = 0;
/******/ 		}
/******/ 		for(moduleId in moreModules) {
/******/ 			if(Object.prototype.hasOwnProperty.call(moreModules, moduleId)) {
/******/ 				modules[moduleId] = moreModules[moduleId];
/******/ 			}
/******/ 		}
/******/ 		if(parentJsonpFunction) parentJsonpFunction(data);
/******/
/******/ 		while(resolves.length) {
/******/ 			resolves.shift()();
/******/ 		}
/******/
/******/ 	};
/******/
/******/
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// object to store loaded and loading chunks
/******/ 	// undefined = chunk not loaded, null = chunk preloaded/prefetched
/******/ 	// Promise = chunk loading, 0 = chunk loaded
/******/ 	var installedChunks = {
/******/ 		"main": 0
/******/ 	};
/******/
/******/
/******/
/******/ 	// script path function
/******/ 	function jsonpScriptSrc(chunkId) {
/******/ 		return __webpack_require__.p + "" + ({"SignInPage-TopPage-TradingPage":"SignInPage-TopPage-TradingPage","SignInPage":"SignInPage","TopPage":"TopPage","TradingPage":"TradingPage"}[chunkId]||chunkId) + ".js"
/******/ 	}
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/ 	// This file contains only the entry chunk.
/******/ 	// The chunk loading function for additional chunks
/******/ 	__webpack_require__.e = function requireEnsure(chunkId) {
/******/ 		var promises = [];
/******/
/******/
/******/ 		// JSONP chunk loading for javascript
/******/
/******/ 		var installedChunkData = installedChunks[chunkId];
/******/ 		if(installedChunkData !== 0) { // 0 means "already installed".
/******/
/******/ 			// a Promise means "currently loading".
/******/ 			if(installedChunkData) {
/******/ 				promises.push(installedChunkData[2]);
/******/ 			} else {
/******/ 				// setup Promise in chunk cache
/******/ 				var promise = new Promise(function(resolve, reject) {
/******/ 					installedChunkData = installedChunks[chunkId] = [resolve, reject];
/******/ 				});
/******/ 				promises.push(installedChunkData[2] = promise);
/******/
/******/ 				// start chunk loading
/******/ 				var script = document.createElement('script');
/******/ 				var onScriptComplete;
/******/
/******/ 				script.charset = 'utf-8';
/******/ 				script.timeout = 120;
/******/ 				if (__webpack_require__.nc) {
/******/ 					script.setAttribute("nonce", __webpack_require__.nc);
/******/ 				}
/******/ 				script.src = jsonpScriptSrc(chunkId);
/******/
/******/ 				// create error before stack unwound to get useful stacktrace later
/******/ 				var error = new Error();
/******/ 				onScriptComplete = function (event) {
/******/ 					// avoid mem leaks in IE.
/******/ 					script.onerror = script.onload = null;
/******/ 					clearTimeout(timeout);
/******/ 					var chunk = installedChunks[chunkId];
/******/ 					if(chunk !== 0) {
/******/ 						if(chunk) {
/******/ 							var errorType = event && (event.type === 'load' ? 'missing' : event.type);
/******/ 							var realSrc = event && event.target && event.target.src;
/******/ 							error.message = 'Loading chunk ' + chunkId + ' failed.\n(' + errorType + ': ' + realSrc + ')';
/******/ 							error.name = 'ChunkLoadError';
/******/ 							error.type = errorType;
/******/ 							error.request = realSrc;
/******/ 							chunk[1](error);
/******/ 						}
/******/ 						installedChunks[chunkId] = undefined;
/******/ 					}
/******/ 				};
/******/ 				var timeout = setTimeout(function(){
/******/ 					onScriptComplete({ type: 'timeout', target: script });
/******/ 				}, 120000);
/******/ 				script.onerror = script.onload = onScriptComplete;
/******/ 				document.head.appendChild(script);
/******/ 			}
/******/ 		}
/******/ 		return Promise.all(promises);
/******/ 	};
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "/";
/******/
/******/ 	// on error function for async loading
/******/ 	__webpack_require__.oe = function(err) { console.error(err); throw err; };
/******/
/******/ 	var jsonpArray = window["webpackJsonp"] = window["webpackJsonp"] || [];
/******/ 	var oldJsonpFunction = jsonpArray.push.bind(jsonpArray);
/******/ 	jsonpArray.push = webpackJsonpCallback;
/******/ 	jsonpArray = jsonpArray.slice();
/******/ 	for(var i = 0; i < jsonpArray.length; i++) webpackJsonpCallback(jsonpArray[i]);
/******/ 	var parentJsonpFunction = oldJsonpFunction;
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/main.ts");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/Application.ts":
/*!****************************!*\
  !*** ./src/Application.ts ***!
  \****************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return Application; });\n/* harmony import */ var _clients_Functions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./clients/Functions */ \"./src/clients/Functions.ts\");\n/// <reference path=\"./IApplication.ts\" />\n/// <reference path=\"./clients/HTTPClient.ts\" />\n\nclass Application {\n    constructor(templateClient, models, routerFactory) {\n        this.templateClient = templateClient;\n        this.models = models;\n        this.router = routerFactory(this);\n    }\n    start() {\n        this.router.start();\n    }\n    fetchTemplate(name) {\n        const url = `/pages/${name}`;\n        return this.templateClient\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"]);\n    }\n    navigate(path) {\n        this.router.navigate(path);\n    }\n    redirect(path) {\n        this.router.redirect(path);\n    }\n}\n\n\n//# sourceURL=webpack:///./src/Application.ts?");

/***/ }),

/***/ "./src/Router.ts":
/*!***********************!*\
  !*** ./src/Router.ts ***!
  \***********************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return Router; });\n/// <reference path=\"./IRouter.ts\" />\n/// <reference path=\"./page.d.ts\" />\n/// <reference path=\"./IApplication.ts\" />\n/// <reference path=\"./pages/IPage.ts\" />\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\nclass Router {\n    constructor(app) {\n        page(\"/\", () => __awaiter(this, void 0, void 0, function* () {\n            const m = yield Promise.all(/*! import() | SignInPage */[__webpack_require__.e(\"SignInPage-TopPage-TradingPage\"), __webpack_require__.e(\"SignInPage\")]).then(__webpack_require__.bind(null, /*! ./pages/SignInPage */ \"./src/pages/SignInPage.ts\"));\n            this.showPage(new m.SignInPage(app));\n            //this.showPage(new TopPage(app));\n        }));\n        page(\"/top\", () => __awaiter(this, void 0, void 0, function* () {\n            const m = yield Promise.all(/*! import() | TopPage */[__webpack_require__.e(\"SignInPage-TopPage-TradingPage\"), __webpack_require__.e(\"TopPage\")]).then(__webpack_require__.bind(null, /*! ./pages/TopPage */ \"./src/pages/TopPage.ts\"));\n            this.showPage(new m.TopPage(app));\n        }));\n        page(\"/tradings/:id\", (ctx) => __awaiter(this, void 0, void 0, function* () {\n            const m = yield Promise.all(/*! import() | TradingPage */[__webpack_require__.e(\"SignInPage-TopPage-TradingPage\"), __webpack_require__.e(\"TradingPage\")]).then(__webpack_require__.bind(null, /*! ./pages/TradingPage */ \"./src/pages/TradingPage.ts\"));\n            this.showPage(new m.TradingPage(app, ctx.params.id, false));\n        }));\n        page(\"/tradings/:id/copy\", (ctx) => __awaiter(this, void 0, void 0, function* () {\n            const m = yield Promise.all(/*! import() | TradingPage */[__webpack_require__.e(\"SignInPage-TopPage-TradingPage\"), __webpack_require__.e(\"TradingPage\")]).then(__webpack_require__.bind(null, /*! ./pages/TradingPage */ \"./src/pages/TradingPage.ts\"));\n            this.showPage(new m.TradingPage(app, ctx.params.id, true));\n        }));\n    }\n    start() {\n        page();\n    }\n    navigate(path) {\n        page(path);\n    }\n    redirect(path) {\n        page.redirect(path);\n    }\n    showPage(next) {\n        next.onCreate();\n    }\n}\n\n\n//# sourceURL=webpack:///./src/Router.ts?");

/***/ }),

/***/ "./src/clients/AuthedClient.ts":
/*!*************************************!*\
  !*** ./src/clients/AuthedClient.ts ***!
  \*************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return AuthedClient; });\nclass AuthedClient {\n    constructor(client, token) {\n        this.client = client;\n        this.token = token;\n    }\n    send(method, url, header, body) {\n        if (header == null) {\n            header = {};\n        }\n        if (this.token.token.length > 0) {\n            header[\"Authorization\"] = `bearer ${this.token.token}`;\n        }\n        return this.client\n            .send(method, url, header, body)\n            .then((r) => {\n            if (r.status != 401) {\n                return r;\n            }\n            return this.refresh(method, url, header, body, r);\n        });\n    }\n    refresh(method, url, header, body, originalResp) {\n        if (this.token.refresh.length == 0) {\n            return Promise.resolve(originalResp);\n        }\n        const refreshURL = \"/api/v1/token/refresh\";\n        const refreshParams = {\n            token: this.token.refresh,\n        };\n        return this.client\n            .send(\"POST\" /* POST */, refreshURL, {}, JSON.stringify(refreshParams))\n            .then((r) => {\n            if (r.status != 200) {\n                return originalResp;\n            }\n            const nextToken = r.body[\"access_token\"];\n            this.token.save(nextToken, this.token.refresh, false);\n            return this.send(method, url, header, body);\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/clients/AuthedClient.ts?");

/***/ }),

/***/ "./src/clients/Functions.ts":
/*!**********************************!*\
  !*** ./src/clients/Functions.ts ***!
  \**********************************/
/*! exports provided: checkStatus, isStatus200, isStatus201, isStatus204, getBody, toTrue, getJSONArray */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"checkStatus\", function() { return checkStatus; });\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"isStatus200\", function() { return isStatus200; });\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"isStatus201\", function() { return isStatus201; });\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"isStatus204\", function() { return isStatus204; });\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"getBody\", function() { return getBody; });\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"toTrue\", function() { return toTrue; });\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"getJSONArray\", function() { return getJSONArray; });\n/// <reference path=\"./HTTPClient.ts\" />\nfunction checkStatus(code) {\n    return (resp) => {\n        if (resp.status == code) {\n            return Promise.resolve(resp);\n        }\n        else {\n            return Promise.reject(resp);\n        }\n    };\n}\nconst isStatus200 = checkStatus(200);\nconst isStatus201 = checkStatus(201);\nconst isStatus204 = checkStatus(204);\nfunction getBody(resp) {\n    return resp.body;\n}\nfunction toTrue(resp) {\n    return true;\n}\nfunction getJSONArray(key) {\n    return (json) => {\n        return json[key];\n    };\n}\n\n\n//# sourceURL=webpack:///./src/clients/Functions.ts?");

/***/ }),

/***/ "./src/clients/XHRClient.ts":
/*!**********************************!*\
  !*** ./src/clients/XHRClient.ts ***!
  \**********************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return XHRClient; });\n/// <reference path=\"./HTTPClient.ts\"/>\nclass XHRClient {\n    constructor(respType) {\n        this.respType = respType;\n    }\n    send(method, url, header, body) {\n        if (header == null) {\n            throw 'header must not be null';\n        }\n        return new Promise((resolve, reject) => {\n            const xhr = new XMLHttpRequest();\n            xhr.open(method, url, true);\n            xhr.responseType = this.respType;\n            for (var key in header) {\n                xhr.setRequestHeader(key, header[key]);\n            }\n            xhr.onload = function () {\n                resolve({\n                    status: this.status,\n                    body: this.response\n                });\n            };\n            xhr.onerror = (e) => {\n                reject(e);\n            };\n            if (body == null) {\n                xhr.send();\n            }\n            else {\n                xhr.send(body);\n            }\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/clients/XHRClient.ts?");

/***/ }),

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/*! no exports provided */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _Application__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Application */ \"./src/Application.ts\");\n/* harmony import */ var _clients_XHRClient__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./clients/XHRClient */ \"./src/clients/XHRClient.ts\");\n/* harmony import */ var _Router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Router */ \"./src/Router.ts\");\n/* harmony import */ var _models_Models__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./models/Models */ \"./src/models/Models.ts\");\n/* harmony import */ var _models_token_AccessToken__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./models/token/AccessToken */ \"./src/models/token/AccessToken.ts\");\n/* harmony import */ var _clients_AuthedClient__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./clients/AuthedClient */ \"./src/clients/AuthedClient.ts\");\n///<reference path=\"./ractive.d.ts\"/>\n///<reference path=\"./data.ts\"/>\n///<reference path=\"./Application.ts\"/>\n///<reference path=\"./Page.ts\"/>\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\n///<reference path=\"./SignInPage.ts\"/>\n///<reference path=\"./TopPage.ts\"/>\n///<reference path=\"./SheetPage.ts\"/>\n\n\n\n\n\n\nvar $;\nvar _;\nvar Backbone;\n// var app: App = new App();\n/*\nvar AppRouter = Backbone.Router.extend({\n    routes: {\n        // ここに、ページ毎に呼ぶ関数名を記述していく\n        // index.htmlを開いた直後は、topという関数を実行する\n        \"\": \"signIn\",\n        top: \"top\",\n        // index.html#sheetの場合は、sheetという関数を実行する\n        \"sheets(/:id)\": \"sheet\",\n        \"sheets(/:id)/copy\": \"copySheet\",\n        setting: \"setting\",\n    },\n    signIn: function() {\n        app.page = new SignInPage();\n        app.page.onCreate(app);\n    },\n    top: function() {\n        app.page = new TopPage();\n        app.page.onCreate(app);\n    },\n    sheet: (id: string) => {\n        app.page = new SheetPage(id, false);\n        app.page.onCreate(app);\n    },\n    copySheet: (id: string) => {\n        app.page = new SheetPage(id, true);\n        app.page.onCreate(app);\n    },\n    setting: () => {\n        // ダイアログ用の要素を作る\n        var dialog = document.createElement(\"section\");\n        document.querySelector(\"#dialogs\")!.appendChild(dialog);\n        // Racriveオブジェクトを作る\n        app.ractive = new Ractive({\n            // どの箱に入れるかをIDで指定\n            el: dialog,\n            // 指定した箱に、どのHTMLを入れるかをIDで指定\n            template: \"#settingTemplate\",\n            // データを設定。テンプレートで使います。\n            /*data : {\n                'sheets' : sheetList\n            }*/\n/*\n        });\n    },\n});\n*/\nconst boot = () => __awaiter(undefined, void 0, void 0, function* () {\n    const token = new _models_token_AccessToken__WEBPACK_IMPORTED_MODULE_4__[\"default\"]();\n    const client = new _clients_XHRClient__WEBPACK_IMPORTED_MODULE_1__[\"default\"](\"json\");\n    const authedClient = new _clients_AuthedClient__WEBPACK_IMPORTED_MODULE_5__[\"default\"](client, token);\n    const models = new _models_Models__WEBPACK_IMPORTED_MODULE_3__[\"Models\"](client, authedClient, token);\n    const app = new _Application__WEBPACK_IMPORTED_MODULE_0__[\"default\"](new _clients_XHRClient__WEBPACK_IMPORTED_MODULE_1__[\"default\"](\"text\"), models, (a) => new _Router__WEBPACK_IMPORTED_MODULE_2__[\"default\"](a));\n    app.start();\n    // Backboneのおまじない\n    //app.router = new AppRouter();\n    //Backbone.history.start();\n});\nwindow.addEventListener(\"load\", boot);\n// [common] for plugins\n/*\nfunction tooltipster() {\n    $(\".actionBtn li a\").tooltipster({\n        theme: \"tooltipster-actionBtn\",\n    });\n    $(\".btn, .delete\").tooltipster({\n        theme: \"tooltipster-btn\",\n        //arrow: false,\n        offsetY: -3,\n    });\n}\n*/\nfunction selectbox() {\n    //select box customize\n    //$('select').easySelectBox({speed: 200});\n}\n\n\n//# sourceURL=webpack:///./src/main.ts?");

/***/ }),

/***/ "./src/models/Models.ts":
/*!******************************!*\
  !*** ./src/models/Models.ts ***!
  \******************************/
/*! exports provided: Models */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"Models\", function() { return Models; });\n/* harmony import */ var _account_AccountRepository__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./account/AccountRepository */ \"./src/models/account/AccountRepository.ts\");\n/* harmony import */ var _environment_EnvironmentRepository__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./environment/EnvironmentRepository */ \"./src/models/environment/EnvironmentRepository.ts\");\n/* harmony import */ var _trading_TradingRepository__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./trading/TradingRepository */ \"./src/models/trading/TradingRepository.ts\");\n/* harmony import */ var _user_UserRepository__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./user/UserRepository */ \"./src/models/user/UserRepository.ts\");\n/* harmony import */ var _company_CompanyRepository__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./company/CompanyRepository */ \"./src/models/company/CompanyRepository.ts\");\n\n\n\n\n\nclass Models {\n    constructor(client, authedClient, token) {\n        this.account = new _account_AccountRepository__WEBPACK_IMPORTED_MODULE_0__[\"default\"](client, token);\n        this.environment = new _environment_EnvironmentRepository__WEBPACK_IMPORTED_MODULE_1__[\"default\"](authedClient);\n        this.trading = new _trading_TradingRepository__WEBPACK_IMPORTED_MODULE_2__[\"default\"](authedClient);\n        this.user = new _user_UserRepository__WEBPACK_IMPORTED_MODULE_3__[\"default\"](authedClient);\n        this.company = new _company_CompanyRepository__WEBPACK_IMPORTED_MODULE_4__[\"default\"](authedClient);\n        this.token = token;\n    }\n}\n\n\n//# sourceURL=webpack:///./src/models/Models.ts?");

/***/ }),

/***/ "./src/models/account/AccountRepository.ts":
/*!*************************************************!*\
  !*** ./src/models/account/AccountRepository.ts ***!
  \*************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return AccountRepository; });\n/* harmony import */ var _clients_Functions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../clients/Functions */ \"./src/clients/Functions.ts\");\n\nclass AccountRepository {\n    constructor(client, token) {\n        this.client = client;\n        this.token = token;\n    }\n    signIn(username, password) {\n        const url = \"/api/v1/token\";\n        const params = {\n            username: username,\n            password: password,\n        };\n        return this.client\n            .send(\"POST\" /* POST */, url, {}, JSON.stringify(params))\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((json) => {\n            const accessToken = json.access_token;\n            const refreshToken = json.refresh_token;\n            const isAdmin = json.is_admin;\n            this.token.save(accessToken, refreshToken, isAdmin);\n            return refreshToken;\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/models/account/AccountRepository.ts?");

/***/ }),

/***/ "./src/models/company/CompanyRepository.ts":
/*!*************************************************!*\
  !*** ./src/models/company/CompanyRepository.ts ***!
  \*************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return CompanyRepository; });\n/* harmony import */ var _clients_Functions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../clients/Functions */ \"./src/clients/Functions.ts\");\n\nclass CompanyRepository {\n    constructor(client) {\n        this.client = client;\n        this.cache = [];\n    }\n    getAll() {\n        if (this.cache.length > 0) {\n            return Promise.resolve(this.cache);\n        }\n        const url = \"/api/v1/companies\";\n        return this.client\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((body) => {\n            this.cache = body[\"companies\"];\n            return this.cache;\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/models/company/CompanyRepository.ts?");

/***/ }),

/***/ "./src/models/environment/EnvironmentRepository.ts":
/*!*********************************************************!*\
  !*** ./src/models/environment/EnvironmentRepository.ts ***!
  \*********************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return EnvironmentRepository; });\n/* harmony import */ var _clients_Functions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../clients/Functions */ \"./src/clients/Functions.ts\");\n\nclass EnvironmentRepository {\n    constructor(client) {\n        this.client = client;\n    }\n    get() {\n        const url = \"/api/v1/environments\";\n        return this.client\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((json) => json);\n    }\n}\n\n\n//# sourceURL=webpack:///./src/models/environment/EnvironmentRepository.ts?");

/***/ }),

/***/ "./src/models/token/AccessToken.ts":
/*!*****************************************!*\
  !*** ./src/models/token/AccessToken.ts ***!
  \*****************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return AccessToken; });\nclass AccessToken {\n    constructor() {\n        this.isAdmin = false;\n        try {\n            this.token = getFromStorage(\"token\");\n            this.refresh = getFromStorage(\"refresh\");\n        }\n        catch (err) {\n            this.token = \"\";\n            this.refresh = \"\";\n        }\n    }\n    save(token, refresh, isAdmin) {\n        this.token = token;\n        this.refresh = refresh;\n        this.isAdmin = isAdmin;\n        try {\n            localStorage.setItem(\"token\", token);\n            localStorage.setItem(\"refresh\", refresh);\n        }\n        catch (err) {\n            // nop\n        }\n    }\n    isLoggedIn() {\n        return this.token.length > 0;\n    }\n}\nconst getFromStorage = (key) => {\n    const v = localStorage.getItem(key);\n    if (v == null) {\n        return \"\";\n    }\n    else {\n        return v;\n    }\n};\n\n\n//# sourceURL=webpack:///./src/models/token/AccessToken.ts?");

/***/ }),

/***/ "./src/models/trading/TradingRepository.ts":
/*!*************************************************!*\
  !*** ./src/models/trading/TradingRepository.ts ***!
  \*************************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return TradingRepository; });\n/* harmony import */ var _clients_Functions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../clients/Functions */ \"./src/clients/Functions.ts\");\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\n\nclass TradingRepository {\n    constructor(client) {\n        this.client = client;\n    }\n    getAll() {\n        const url = \"/api/v1/tradings\";\n        return this.client\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((json) => {\n            return json[\"tradings\"];\n        })\n            .then((list) => {\n            list.forEach((item) => {\n                item.date = item.id;\n            });\n            return list;\n        });\n    }\n    getById(id) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const vals = yield Promise.all([\n                this.getTradingById(id),\n                this.getTradingItemsById(id),\n            ]);\n            vals[0].items = vals[1];\n            return vals[0];\n        });\n    }\n    save(item) {\n        return __awaiter(this, void 0, void 0, function* () {\n            let saved;\n            if (item.id.length == 0) {\n                saved = yield this.create(item);\n            }\n            else {\n                saved = yield this.update(item);\n            }\n            yield this.saveItems(item);\n            return saved;\n        });\n    }\n    deleteTrading(id) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const url = `/api/v1/tradings/${id}`;\n            return this.client\n                .send(\"DELETE\" /* DELETE */, url, {}, null)\n                .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus204\"])\n                .then((r) => true);\n        });\n    }\n    deleteItems(id, items) {\n        return __awaiter(this, void 0, void 0, function* () {\n            if (id.length == 0) {\n                return Promise.resolve();\n            }\n            const idSet = new Set();\n            items.forEach((item) => {\n                idSet.add(item.id);\n            });\n            idSet.forEach((itemId) => __awaiter(this, void 0, void 0, function* () {\n                yield this.deleteItem(id, itemId);\n            }));\n            return Promise.resolve();\n        });\n    }\n    getNextNumber(type, date) {\n        const url = `/api/v1/sequences/${type}`;\n        const params = {\n            date: new Date(date).getTime(),\n        };\n        return this.client\n            .send(\"POST\" /* POST */, url, {}, JSON.stringify(params))\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((body) => body[\"number\"]);\n    }\n    getTradingById(id) {\n        const url = `/api/v1/tradings/${id}`;\n        return this.client\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((body) => body);\n    }\n    getTradingItemsById(id) {\n        const url = `/api/v1/tradings/${id}/items`;\n        return this.client\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((body) => body[\"items\"]);\n    }\n    create(item) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const url = \"/api/v1/tradings\";\n            const params = {\n                company_id: item.company_id,\n                title_type: item.title_type,\n                subject: item.subject,\n                work_from: item.work_from,\n                work_to: item.work_to,\n                total: item.total,\n                quotation_date: item.quotation_date,\n                quotation_number: item.quotation_number,\n                bill_date: item.bill_date,\n                bill_number: item.bill_number,\n                delivery_date: item.delivery_date,\n                delivery_number: item.delivery_number,\n                tax_rate: item.tax_rate,\n                product: item.product,\n                memo: item.memo,\n            };\n            const resp = yield this.client.send(\"POST\" /* POST */, url, {}, JSON.stringify(params));\n            if (resp.status != 201) {\n                throw resp;\n            }\n            item.id = resp.body.id;\n            return item;\n        });\n    }\n    update(item) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const url = `/api/v1/tradings/${item.id}`;\n            const params = {\n                company_id: item.company_id,\n                title_type: item.title_type,\n                subject: item.subject,\n                work_from: item.work_from,\n                work_to: item.work_to,\n                total: item.total,\n                quotation_date: item.quotation_date,\n                quotation_number: item.quotation_number,\n                bill_date: item.bill_date,\n                bill_number: item.bill_number,\n                delivery_date: item.delivery_date,\n                delivery_number: item.delivery_number,\n                tax_rate: item.tax_rate,\n                product: item.product,\n                memo: item.memo,\n            };\n            const resp = yield this.client.send(\"PUT\" /* PUT */, url, {}, JSON.stringify(params));\n            if (resp.status != 200) {\n                throw resp;\n            }\n            return item;\n        });\n    }\n    saveItems(trading) {\n        return __awaiter(this, void 0, void 0, function* () {\n            trading.items.forEach((item) => __awaiter(this, void 0, void 0, function* () {\n                if (item.id.length == 0) {\n                    yield this.createItem(trading.id, item);\n                }\n                else {\n                    yield this.updateItem(trading.id, item);\n                }\n            }));\n        });\n    }\n    createItem(tradingId, item) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const url = `/api/v1/tradings/${tradingId}/items`;\n            const params = {\n                sort_order: item.sort_order,\n                subject: item.subject,\n                unit_price: item.unit_price,\n                amount: item.amount,\n                degree: item.degree,\n                tax_type: item.tax_type,\n                memo: item.memo,\n            };\n            const resp = yield this.client.send(\"POST\" /* POST */, url, {}, JSON.stringify(params));\n            if (resp.status != 201) {\n                throw resp;\n            }\n            item.id = resp.body.id;\n        });\n    }\n    updateItem(tradingId, item) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const url = `/api/v1/tradings/${tradingId}/items/${item.id}`;\n            const params = {\n                sort_order: item.sort_order,\n                subject: item.subject,\n                unit_price: item.unit_price,\n                amount: item.amount,\n                degree: item.degree,\n                tax_type: item.tax_type,\n                memo: item.memo,\n            };\n            const resp = yield this.client.send(\"PUT\" /* PUT */, url, {}, JSON.stringify(params));\n            if (resp.status != 200) {\n                throw resp;\n            }\n        });\n    }\n    deleteItem(tradingId, itemId) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const url = `/api/v1/tradings/${tradingId}/items/${itemId}`;\n            const resp = yield this.client.send(\"DELETE\" /* DELETE */, url, {}, null);\n            if (resp.status != 204) {\n                throw resp;\n            }\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/models/trading/TradingRepository.ts?");

/***/ }),

/***/ "./src/models/user/UserRepository.ts":
/*!*******************************************!*\
  !*** ./src/models/user/UserRepository.ts ***!
  \*******************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return UserRepository; });\n/* harmony import */ var _clients_Functions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../clients/Functions */ \"./src/clients/Functions.ts\");\n\nclass UserRepository {\n    constructor(client) {\n        this.client = client;\n        this.cache = [];\n    }\n    getAll() {\n        if (this.cache.length > 0) {\n            return Promise.resolve(this.cache);\n        }\n        const url = \"/api/v1/users\";\n        return this.client\n            .send(\"GET\" /* GET */, url, {}, null)\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"isStatus200\"])\n            .then(_clients_Functions__WEBPACK_IMPORTED_MODULE_0__[\"getBody\"])\n            .then((body) => {\n            this.cache = body[\"users\"];\n            return this.cache;\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/models/user/UserRepository.ts?");

/***/ })

/******/ });
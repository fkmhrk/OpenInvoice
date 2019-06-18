(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["SignInPage"],{

/***/ "./src/pages/SignInPage.ts":
/*!*********************************!*\
  !*** ./src/pages/SignInPage.ts ***!
  \*********************************/
/*! exports provided: SignInPage */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"SignInPage\", function() { return SignInPage; });\n/* harmony import */ var _ractive__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../ractive */ \"./src/ractive.js\");\n/* harmony import */ var _ractive__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_ractive__WEBPACK_IMPORTED_MODULE_0__);\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\n\nclass SignInPage {\n    constructor(app) {\n        this.app = app;\n    }\n    onCreate() {\n        return __awaiter(this, void 0, void 0, function* () {\n            this.ractive = new _ractive__WEBPACK_IMPORTED_MODULE_0___default.a({\n                el: \"#container\",\n                template: \"#signInTemplate\",\n                data: {\n                    // myCompanyName: app.myCompanyName,\n                    inProgress: false,\n                },\n                on: {\n                    signIn: () => this.signIn(),\n                },\n            });\n        });\n    }\n    signIn() {\n        return __awaiter(this, void 0, void 0, function* () {\n            const username = this.ractive.get(\"username\");\n            const password = this.ractive.get(\"password\");\n            this.ractive.set(\"inProgress\", true);\n            try {\n                const token = yield this.app.models.account.signIn(username, password);\n                // TODO save credential\n                this.app.navigate(\"/top\");\n            }\n            catch (e) {\n                alert(e);\n            }\n        });\n    }\n}\n\n\n//# sourceURL=webpack:///./src/pages/SignInPage.ts?");

/***/ })

}]);
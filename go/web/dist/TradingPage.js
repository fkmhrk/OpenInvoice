(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["TradingPage"],{

/***/ "./src/pages/TradingPage.ts":
/*!**********************************!*\
  !*** ./src/pages/TradingPage.ts ***!
  \**********************************/
/*! exports provided: TradingPage */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"TradingPage\", function() { return TradingPage; });\n/* harmony import */ var _ractive__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../ractive */ \"./src/ractive.js\");\n/* harmony import */ var _ractive__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_ractive__WEBPACK_IMPORTED_MODULE_0__);\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\n\nconst toDateStr = (time) => {\n    var date = new Date(time);\n    var m = date.getMonth() + 1;\n    var d = date.getDate();\n    if (m < 10) {\n        m = \"0\" + m;\n    }\n    if (d < 10) {\n        d = \"0\" + d;\n    }\n    return date.getFullYear() + \"-\" + m + \"-\" + d;\n};\nconst emptyUser = {\n    id: \"empty\",\n    login_name: \"\",\n    display_name: \"担当者なし\",\n    tel: \"\",\n};\nclass TradingPage {\n    constructor(app, id, isCopy) {\n        this.app = app;\n        this.id = id;\n        this.isCopy = isCopy;\n    }\n    onCreate() {\n        return __awaiter(this, void 0, void 0, function* () {\n            const trading = yield this.loadTrading();\n            const vals = yield Promise.all([\n                this.app.models.user.getAll(),\n                this.app.models.company.getAll(),\n            ]);\n            const users = vals[0].map((u) => u); // copy\n            const companies = vals[1];\n            users.unshift(emptyUser);\n            this.ractive = new _ractive__WEBPACK_IMPORTED_MODULE_0___default.a({\n                // どの箱に入れるかをIDで指定\n                el: \"#container\",\n                // 指定した箱に、どのHTMLを入れるかをIDで指定\n                template: \"#sheetTemplate\",\n                decorators: {},\n                data: {\n                    // myCompanyName: app.myCompanyName,\n                    // is_admin: app.client.isAdmin(),\n                    trading: trading,\n                    workFrom: toDateStr(trading.work_from),\n                    workTo: toDateStr(trading.work_to),\n                    quotationDate: toDateStr(trading.quotation_date),\n                    billDate: toDateStr(trading.bill_date),\n                    deliveryDate: toDateStr(trading.delivery_date),\n                    calcItemSum: (item) => this.calcItemSum(item),\n                    companies: companies,\n                    users: users,\n                    deletedItems: [],\n                },\n                on: {\n                    close: () => window.history.back(),\n                    addItem: () => this.addItem(),\n                    deleteItem: (e, index) => this.deleteItem(index),\n                    save: () => __awaiter(this, void 0, void 0, function* () {\n                        yield this.save();\n                        window.history.back();\n                    }),\n                    printQuotation: () => this.printQuotation(),\n                    printBill: () => this.printBill(),\n                    printDelivery: () => this.printDelivery(),\n                    printInvoide: () => this.showInvoiceDialog(),\n                },\n            });\n            // why here?: calcItemSum() uses this.ractive.\n            this.ractive.set(\"tradingItems\", trading.items);\n        });\n    }\n    calcItemSum(item) {\n        item.sum = item.unit_price * item.amount;\n        this.updateSum();\n        return item.sum;\n    }\n    updateSum() {\n        const itemList = this.ractive.get(\"tradingItems\");\n        const taxRate = this.ractive.get(\"trading.tax_rate\");\n        let sum = 0;\n        let tax = 0;\n        itemList.forEach((item) => {\n            const taxType = item.tax_type;\n            if (taxType == 0) {\n                sum += item.sum;\n            }\n            else if (taxType == 1) {\n                sum += item.sum;\n                tax += (item.sum * taxRate) / 100;\n            }\n            else if (taxType == 2) {\n                const body = (item.sum * 100) / (100 + taxRate);\n                const taxTmp = Math.ceil(item.sum - body);\n                sum += item.sum - taxTmp;\n                tax += taxTmp;\n            }\n        });\n        this.ractive.set({\n            \"trading.sum\": sum,\n            \"trading.tax\": tax,\n            \"trading.total\": sum + tax,\n        });\n    }\n    addItem() {\n        this.ractive.push(\"tradingItems\", {\n            id: \"\",\n            subject: \"\",\n            unit_price: 0,\n            amount: 0,\n            degree: \"\",\n            memo: \"\",\n            tax_type: 1,\n            sum: 0,\n        });\n    }\n    deleteItem(index) {\n        var item = this.ractive.get(\"tradingItems\")[index];\n        this.ractive.splice(\"tradingItems\", index, 1);\n        if (item.id.length > 0) {\n            this.ractive.push(\"deletedItems\", item);\n        }\n    }\n    save() {\n        return __awaiter(this, void 0, void 0, function* () {\n            const trading = this.ractive.get(\"trading\");\n            const workFrom = this.ractive.get(\"workFrom\");\n            const workTo = this.ractive.get(\"workTo\");\n            const quotationDate = this.ractive.get(\"quotationDate\");\n            const billDate = this.ractive.get(\"billDate\");\n            const deliveryDate = this.ractive.get(\"deliveryDate\");\n            const tradingItems = this.ractive.get(\"tradingItems\");\n            trading.items = tradingItems;\n            // modify type\n            trading.work_from = new Date(workFrom).getTime();\n            trading.work_to = new Date(workTo).getTime();\n            trading.quotation_date = new Date(quotationDate).getTime();\n            trading.bill_date = new Date(billDate).getTime();\n            trading.delivery_date = new Date(deliveryDate).getTime();\n            console.log(trading);\n            const saved = yield this.app.models.trading.save(trading);\n            this.ractive.set(\"tradingItems\", saved.items);\n            const deletedItems = this.ractive.get(\"deletedItems\");\n            yield this.app.models.trading.deleteItems(this.id, deletedItems);\n            this.ractive.set(\"deletedItems\", []);\n            return trading;\n            /*\n            app.client.saveTrading(trading, {\n                success: (id: string) => {\n                    trading.id = id;\n                    trading.modified_time = new Date().getTime();\n                    app.tradingsMap[id] = trading;\n                    var deleted = app.ractive.get(\"deletedItems\");\n                    this.deleteItems(app, id, deleted, doneFunc);\n                },\n                error: (status: number, msg: string) => {\n                    switch (status) {\n                        case 1001:\n                            app.addSnack(\"件名を入力してください。\");\n                            break;\n                        case 1002:\n                            app.addSnack(\n                                \"作業終了日は作業開始日より後にしてください。\"\n                            );\n                            break;\n                        case 1003:\n                            app.addSnack(\"消費税率は0以上にしてください。\");\n                            break;\n                        default:\n                            app.addSnack(\"保存に失敗しました。\");\n                    }\n                    console.log(\"Failed to save trading status=\" + status);\n                },\n            });\n    */\n        });\n    }\n    loadTrading() {\n        return __awaiter(this, void 0, void 0, function* () {\n            if (this.id == \"new\") {\n                const now = new Date().getTime();\n                return {\n                    id: \"\",\n                    company_id: \"\",\n                    company_name: \"\",\n                    title_type: 0,\n                    subject: \"\",\n                    assignee: emptyUser.id,\n                    work_from: now,\n                    work_to: now,\n                    quotation_date: 0,\n                    quotation_number: \"\",\n                    bill_date: 0,\n                    bill_number: \"\",\n                    delivery_date: 0,\n                    delivery_number: \"\",\n                    date: \"\",\n                    total: 0,\n                    tax_rate: 8,\n                    product: \"\",\n                    memo: \"\",\n                    modified_time: 0,\n                    items: [],\n                };\n            }\n            const trading = yield this.app.models.trading.getById(this.id);\n            if (this.isCopy) {\n                this.id = \"\";\n                trading.id = \"\";\n                trading.items.forEach((item) => {\n                    item.id = \"\";\n                });\n            }\n            return trading;\n        });\n    }\n    printQuotation() {\n        return __awaiter(this, void 0, void 0, function* () {\n            const trading = this.ractive.get(\"trading\");\n            const quotationDate = this.ractive.get(\"quotationDate\");\n            if (trading.quotation_number.length == 0) {\n                // make next number\n                const val = yield this.app.models.trading.getNextNumber(\"quotation\", quotationDate);\n                trading.quotation_number = `${val}-I`;\n            }\n            const saved = yield this.save();\n            this.ractive.set(\"trading\", saved);\n            window.location.href = `/php/quotation.php?access_token=${this.app.models.token.token}&trading_id=${trading.id}`;\n        });\n    }\n    printBill() {\n        return __awaiter(this, void 0, void 0, function* () {\n            const trading = this.ractive.get(\"trading\");\n            const billDate = this.ractive.get(\"billDate\");\n            if (trading.bill_number.length == 0) {\n                // make next number\n                const val = yield this.app.models.trading.getNextNumber(\"bill\", billDate);\n                trading.bill_number = `${val}-V`;\n            }\n            const saved = yield this.save();\n            this.ractive.set(\"trading\", saved);\n            window.location.href = `/php/bill.php?access_token=${this.app.models.token.token}&trading_id=${trading.id}`;\n        });\n    }\n    printDelivery() {\n        return __awaiter(this, void 0, void 0, function* () {\n            const trading = this.ractive.get(\"trading\");\n            const deliveryDate = this.ractive.get(\"deliveryDate\");\n            if (trading.delivery_number.length == 0) {\n                // make next number\n                const val = yield this.app.models.trading.getNextNumber(\"delivery\", deliveryDate);\n                trading.delivery_number = `${val}-D`;\n            }\n            const saved = yield this.save();\n            this.ractive.set(\"trading\", saved);\n            window.location.href = `/php/delivery.php?access_token=${this.app.models.token.token}&trading_id=${trading.id}`;\n        });\n    }\n    showInvoiceDialog() {\n        // app.showDialog(new CreateInvoiceDialog());\n    }\n}\n\n\n//# sourceURL=webpack:///./src/pages/TradingPage.ts?");

/***/ })

}]);
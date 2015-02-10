/// <reference path="./Client.ts"/>
var $;
var Invoice;
(function (Invoice) {
    var MockClient = (function () {
        function MockClient() {
        }
        MockClient.prototype.login = function (username, password, callback) {
            callback.success('token1122');
        };
        MockClient.prototype.getTradings = function (token, callback) {
            var tradings = [];
            for (var i = 0; i < 10; ++i) {
                tradings.push({
                    id: 'trade1122' + i,
                    date: 'trade1122' + i,
                    company_id: "会社ID" + i,
                    title_type: 0,
                    subject: "件名" + i,
                    work_from: 1122,
                    work_to: 2233,
                    quotation_date: 1423502769379,
                    bill_date: 5555,
                    tax_rate: 8,
                    assignee: "担当者ID" + i,
                    product: "成果物" + i
                });
            }
            callback.success(tradings);
        };

        MockClient.prototype.getTradingItems = function (token, tradingId, callback) {
            var tradings = [];
            for (var i = 0; i < 10; ++i) {
                tradings.push({
                    id: 'item111' + i,
                    subject: "件名" + i,
                    unit_price: i * 100 + 200,
                    amount: i * 3 + 2,
                    degree: "人月",
                    tax_type: 1,
                    memo: "備考" + i,
                    sum: 1000
                });
            }
            callback.success(tradings);
        };

        MockClient.prototype.getUsers = function (token, callback) {
            var list = [];
            for (var i = 0; i < 10; ++i) {
                list.push({
                    id: "担当者ID" + i,
                    display_name: '担当' + i
                });
            }
            callback.success(list);
        };

        MockClient.prototype.getCompanies = function (token, callback) {
            var list = [];
            for (var i = 0; i < 10; ++i) {
                list.push({
                    id: "company" + i,
                    name: "会社" + i,
                    zip: "111-222" + i,
                    address: "日本の" + i,
                    phone: "03-1111-222" + i,
                    unit: "第" + i + "開発部"
                });
            }
            callback.success(list);
        };
        MockClient.prototype.saveTrading = function (token, item, callback) {
            callback.success('id1122');
        };
        MockClient.prototype.saveTradingItem = function (token, tradingId, item, callback) {
            callback.success('item1122');
        };
        MockClient.prototype.saveCompany = function (token, item, callback) {
            callback.success('company1122');
        };
        return MockClient;
    })();
    Invoice.MockClient = MockClient;
})(Invoice || (Invoice = {}));

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
        MockClient.prototype.deleteTradingItem = function (token, tradingId, itemId, callback) {
            callback.success(itemId);
        };
        MockClient.prototype.saveCompany = function (token, item, callback) {
            callback.success('company1122');
        };
        return MockClient;
    })();
    Invoice.MockClient = MockClient;
})(Invoice || (Invoice = {}));
/// <reference path="./MockClient.ts"/>
/// <reference path="./ractive.d.ts"/>
/// <reference path="./Page.ts"/>
var Application = (function () {
    function Application() {
        this.client = new Invoice.MockClient();
    }
    return Application;
})();
/// <reference path="./Page.ts"/>
var TopPage = (function () {
    function TopPage() {
    }
    TopPage.prototype.onCreate = function (app) {
        var _this = this;
        this.app = app;
        var r = app.ractive = new Ractive({
            el: '#container',
            template: '#topTemplate',
            data: {
                loginInProgress: false
            }
        });
        r.on('login', function (e) {
            _this.login(r.get('username'), r.get('password'));
        });
        console.log(this);
    };
    TopPage.prototype.login = function (username, password) {
        var _this = this;
        this.app.ractive.set('loginInProgress', true);
        this.app.client.login(username, password, {
            success: function (token) {
                _this.app.ractive.set('loginInProgress', false);
                _this.app.token = token;
                _this.app.router.navigate('tradings', { trigger: true });
            },
            error: function (msg) {
                _this.app.ractive.set('loginInProgress', false);
                console.log('error ' + msg);
            }
        });
    };
    return TopPage;
})();
/// <reference path="./Page.ts"/>
var TradingListPage = (function () {
    function TradingListPage() {
    }
    TradingListPage.prototype.onCreate = function (app) {
        var _this = this;
        this.app = app;
        var r = app.ractive = new Ractive({
            el: '#container',
            template: '#tradingTemplate',
            data: {
                tradings: app.tradings,
                token: app.token
            }
        });
        r.on({
            itemClick: function (e, i) {
                _this.edit(i);
            },
            printQuotation: function (e, i) {
                _this.printQuotation(i);
            },
            printBill: function (e, i) {
                _this.printBill(i);
            },
            newTrading: function (e) {
                _this.newTrading(r.get('newId'));
            },
            company: function (e) {
                _this.app.router.navigate('companies', { trigger: true });
            }
        });
        this.loadTradings();
    };
    TradingListPage.prototype.loadTradings = function () {
        var _this = this;
        this.app.client.getTradings(this.app.token, {
            success: function (list) {
                _this.app.tradings = list;
                _this.app.tradingMap = {};
                _.each(list, function (item) {
                    _this.app.tradingMap[item.id] = item;
                });
                _this.loadUsers();
            },
            error: function (statuc, msg) {
                console.log('error ' + msg);
            }
        });
    };
    TradingListPage.prototype.loadUsers = function () {
        var _this = this;
        this.app.client.getUsers(this.app.token, {
            success: function (list) {
                _this.app.users = list;
                _this.loadCompanies();
            },
            error: function (msg) {
                console.log('error getUsers ' + msg);
            }
        });
    };
    TradingListPage.prototype.loadCompanies = function () {
        var _this = this;
        this.app.client.getCompanies(this.app.token, {
            success: function (list) {
                _this.app.companies = list;
                _this.show();
            },
            error: function (msg) {
                console.log('error getCompanies ' + msg);
            }
        });
    };
    TradingListPage.prototype.show = function () {
        this.app.ractive.set('tradings', this.app.tradings);
        this.app.ractive.set('token', this.app.token);
        this.app.ractive.update();
    };
    TradingListPage.prototype.newTrading = function (id) {
        if (id == null || id.length == 0) {
            return;
        }
        app.trading = {
            id: null,
            date: id,
            company_id: '',
            title_type: 0,
            subject: '',
            work_from: new Date().getTime(),
            work_to: new Date().getTime(),
            quotation_date: new Date().getTime(),
            bill_date: new Date().getTime(),
            tax_rate: 8,
            assignee: '',
            product: ''
        };
        app.tradingMap['new'] = app.trading;
        app.router.navigate('tradings/new', { trigger: true });
    };
    TradingListPage.prototype.edit = function (i) {
        console.log(app.tradings[i]);
        app.trading = app.tradings[i];
        app.router.navigate('tradings/' + app.tradings[i].id, { trigger: true });
    };
    TradingListPage.prototype.printQuotation = function (i) {
        var trading = app.tradings[i];
        window.location.href = "/php/quotation.php?access_token=" + app.token + "&trading_id=" + trading.id;
    };
    TradingListPage.prototype.printBill = function (i) {
        var trading = app.tradings[i];
        window.location.href = "/php/bill.php?access_token=" + app.token + "&trading_id=" + trading.id;
    };
    return TradingListPage;
})();
/// <reference path="./ractive.d.ts"/>
/// <reference path="./TopPage.ts"/>
/// <reference path="./TradingListPage.ts"/>
var $;
var _;
var Backbone;
var EditTradingApp = {
    loadTrading: function (token, id) {
        if (id == 'new') {
            app.tradingItems = [];
            EditTradingApp.show(id);
            return;
        }
        app.client.getTradingItems(token, id, {
            success: function (list) {
                app.tradingItems = _.map(list, function (item) {
                    item.unit_price = util.numToCurrency(item.unit_price);
                    return item;
                });
                EditTradingApp.show(id);
            },
            error: function (msg) {
                console.log('error ' + msg);
            }
        });
    },
    show: function (id) {
        var es = function (node) {
            $(node).easySelectBox({ speed: 200 });
            return {
                teardown: function () {
                    // nop?
                }
            };
        };
        var toDateStr = function (date) {
            var m = date.getMonth() + 1;
            var d = date.getDate();
            if (m < 10) {
                m = "0" + m;
            }
            if (d < 10) {
                d = "0" + d;
            }
            return date.getFullYear() + "-" + m + "-" + d;
        };
        var item = app.tradingMap[id];
        var workFrom = toDateStr(new Date(item.work_from));
        var workTo = toDateStr(new Date(item.work_to));
        var quotationDate = toDateStr(new Date(item.quotation_date));
        var billDate = toDateStr(new Date(item.bill_date));
        app.router.r = new Ractive({
            el: '#container',
            template: '#editTradingTemplate',
            decorators: {
                easyselect: es
            },
            data: {
                trading: item,
                tradingItems: app.tradingItems,
                users: app.users,
                companies: app.companies,
                workFrom: workFrom,
                workTo: workTo,
                quotationDate: quotationDate,
                billDate: billDate,
                deleteList: [],
                numToCurrency: function (val) {
                    return util.numToCurrency(val);
                }
            }
        });
        app.router.r.on('numFocus', function (e, val) {
            e.node.value = util.currencyToNum(val);
            app.router.r.updateModel();
        });
        app.router.r.on('sumBlur', function (e, val, index) {
            e.node.value = util.numToCurrency(val);
            app.router.r.updateModel();
            var item = e.context;
            item.sum = util.currencyToNum(item.unit_price) * item.amount;
            app.router.r.update();
        });
        app.router.r.on('amountBlur', function (e) {
            var item = e.context;
            item.sum = util.currencyToNum(item.unit_price) * item.amount;
            app.router.r.update();
        });
        app.router.r.on('deleteItem', function (e, index) {
            if (!confirm('この項目を削除しますか？')) {
                return;
            }
            var tradings = app.router.r.get('tradingItems');
            var trading = tradings[index];
            if (trading.id != null) {
                var list = app.router.r.get('deleteList');
                list.push(trading.id);
                app.router.r.set('deleteList', list);
            }
            tradings.splice(index, 1);
            app.router.r.set('tradingItems', tradings);
            app.router.r.update();
        });
        app.router.r.on('addItem', function (e) {
            var list = app.router.r.get('tradingItems');
            list.push({
                id: null,
                subject: "",
                unit_price: 0,
                amount: 0,
                degree: "人月",
                tax_type: 1,
                memo: "",
                sum: 0
            });
            app.router.r.set('tradingItems', list);
            app.router.r.update();
        });
        app.router.r.on('save', function () {
            var r = app.router.r;
            var company = r.get('companies')[$('#company').prop('selectedIndex')];
            var assignee = r.get('users')[$('#assignee').prop('selectedIndex')];
            var trading = r.get('trading');
            trading.company_id = company.id;
            trading.title_type = $('#titleType').prop('selectedIndex');
            trading.assignee = assignee.id;
            trading.work_from = new Date(r.get('workFrom')).getTime();
            trading.work_to = new Date(r.get('workTo')).getTime();
            trading.quotation_date = new Date(r.get('quotationDate')).getTime();
            trading.bill_date = new Date(r.get('billDate')).getTime();
            trading.tax_rate = parseFloat(r.get('trading.tax_rate'));
            var items = r.get('tradingItems');
            var list = [];
            for (var i = 0; i < items.length; ++i) {
                var item = items[i];
                item.unit_price = util.currencyToNum(item.unit_price);
                item.amount = parseInt(item.amount);
                item.tax_type = parseInt(item.tax_type);
                list.push(item);
            }
            var deleteList = r.get('deleteList');
            EditTradingApp.save(trading, list, deleteList);
        });
    },
    save: function (trading, items, deleteList) {
        EditTradingApp.deleteItems(trading, items, deleteList);
        app.client.saveTrading(app.token, trading, {
            success: function (id) {
                console.log('ok');
                EditTradingApp.saveItems(id, items);
            },
            error: function (msg) {
                console.log('failed to save ' + msg);
            }
        });
    },
    deleteItems: function (trading, items, deleteList) {
        if (deleteList.length == 0) {
            EditTradingApp.saveTrading(trading, items);
            return;
        }
        app.client.deleteTradingItem(app.token, trading.id, deleteList[0], {
            success: function (id) {
                deleteList.shift();
                EditTradingApp.deleteItems(trading, items, deleteList);
            },
            error: function (msg) {
                console.log('failed to delete ' + msg);
            }
        });
    },
    saveTrading: function (trading, items) {
    },
    saveItems: function (tradingId, items) {
        if (items.length == 0) {
            window.history.back();
            return;
        }
        app.client.saveTradingItem(app.token, tradingId, items[0], {
            success: function (id) {
                console.log('ok');
                items.shift();
                EditTradingApp.saveItems(tradingId, items);
            },
            error: function (msg) {
                console.log('failed to save ' + msg);
            }
        });
    }
};
var CompanyApp = {
    show: function () {
        app.router.r = new Ractive({
            el: '#container',
            template: '#companyTemplate',
            data: {
                companies: app.companies
            }
        });
        app.router.r.on('itemClick', function (e, i) {
            CompanyApp.edit(i);
        });
        app.router.r.on('newCompany', function (e) {
            CompanyApp.newCompany();
        });
    },
    edit: function (i) {
        app.company = app.companies[i];
        app.router.navigate('companies/' + app.companies[i].id, { trigger: true });
    },
    newCompany: function () {
        app.company = {
            id: null,
            name: '',
            zip: '',
            address: '',
            phone: '',
            unit: ''
        };
        app.router.navigate('companies/new', { trigger: true });
    }
};
var EditCompanyApp = {
    show: function (company) {
        app.router.r = new Ractive({
            el: '#container',
            template: '#editCompanyTemplate',
            data: {
                company: company
            }
        });
        app.router.r.on('save', function (e) {
            var company = app.router.r.get('company');
            EditCompanyApp.save(company);
        });
    },
    save: function (company) {
        app.client.saveCompany(app.token, company, {
            success: function (id) {
                window.history.back();
            },
            error: function (msg) {
            }
        });
    }
};
var AppRouter = Backbone.Router.extend({
    routes: {
        "": "top",
        "tradings": "tradings",
        "tradings(/:id)": "editTrading",
        "companies": "companies",
        "companies(/:id)": "editCompanies"
    },
    initialize: function () {
        _.bindAll(this, 'top', 'tradings', 'editTrading', 'companies', 'editCompanies');
    },
    top: function () {
        app.page = new TopPage();
        app.page.onCreate(app);
    },
    tradings: function () {
        if (app.token == null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        app.page = new TradingListPage();
        app.page.onCreate(app);
    },
    editTrading: function (id) {
        if (app.token == null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        EditTradingApp.loadTrading(app.token, id);
    },
    companies: function () {
        if (app.token == null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        CompanyApp.show();
    },
    editCompanies: function (id) {
        if (app.token == null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        var company = null;
        if (id === 'new') {
            company = {
                id: null
            };
            EditCompanyApp.show(company);
            return;
        }
        for (var i = 0; i < app.companies.length; ++i) {
            if (app.companies[i].id === id) {
                company = app.companies[i];
                break;
            }
        }
        if (company === null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        EditCompanyApp.show(company);
    }
});
var app = new Application();
var util = {
    numToCurrency: function (val) {
        var n = parseInt(val);
        var ret = '';
        do {
            var n1 = (n % 1000);
            var c = ("00" + n1).slice(-3);
            n = Math.floor(n / 1000);
            if (n > 0) {
                ret = c + "," + ret;
            }
            else {
                ret = n1 + "," + ret;
            }
        } while (n > 0);
        return ret.substring(0, ret.length - 1);
    },
    currencyToNum: function (val) {
        if (typeof (val) === 'number') {
            return val;
        }
        return parseInt(val.replace(",", ""));
    }
};
(function ($) {
    $(function () {
        app.router = new AppRouter();
        Backbone.history.start();
    });
})($);

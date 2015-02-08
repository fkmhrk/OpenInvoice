var _this = this;
/// <reference path="./ClientImpl.ts"/>
/// <reference path="./MockClient.ts"/>
var $;
var _;
var Backbone;
var Ractive;

var TopApp = {
    onCreate: function () {
        console.log(_this);
    },
    login: function (username, password, r) {
        r.set('loginInProgress', true);
        app.client.login(username, password, {
            success: function (token) {
                r.set('loginInProgress', false);
                app.token = token;
                app.router.navigate('tradings', { trigger: true });
            },
            error: function (msg) {
                r.set('loginInProgress', false);
                console.log('error ' + msg);
            }
        });
    }
};

var TradingApp = {
    loadTradings: function (token) {
        app.client.getTradings(token, {
            success: function (list) {
                app.tradings = list;
                app.tradingMap = {};
                _.each(list, function (item) {
                    app.tradingMap[item.id] = item;
                });
                TradingApp.loadUsers(token);
            },
            error: function (msg) {
                console.log('error ' + msg);
            }
        });
    },
    loadUsers: function (token) {
        app.client.getUsers(token, {
            success: function (list) {
                app.users = list;
                TradingApp.loadCompanies(token);
            },
            error: function (msg) {
                console.log('error getUsers ' + msg);
            }
        });
    },
    loadCompanies: function (token) {
        app.client.getCompanies(token, {
            success: function (list) {
                app.companies = list;
                TradingApp.show();
            },
            error: function (msg) {
                console.log('error getCompanies ' + msg);
            }
        });
    },
    show: function () {
        app.router.r = new Ractive({
            el: '#container',
            template: '#tradingTemplate',
            data: {
                tradings: app.tradings
            }
        });
        app.router.r.on('itemClick', function (e, i) {
            TradingApp.edit(i);
        });
        app.router.r.on('newTrading', function (e) {
            TradingApp.newTrading(app.router.r.get('newId'));
        });
    },
    newTrading: function (id) {
        if (id == null || id.length == 0) {
            return;
        }
        app.trading = {
            id: null,
            date: id,
            work_from: new Date().getTime(),
            work_to: new Date().getTime()
        };
        app.tradingMap['new'] = app.trading;
        app.router.navigate('tradings/new', { trigger: true });
    },
    edit: function (i) {
        console.log(app.tradings[i]);
        app.trading = app.tradings[i];
        app.router.navigate('tradings/' + app.tradings[i].id, { trigger: true });
    }
};

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

            var items = r.get('tradingItems');
            var list = [];
            for (var i = 0; i < items.length; ++i) {
                var item = items[i];
                item.unit_price = util.currencyToNum(item.unit_price);
                item.amount = parseInt(item.amount);
                item.tax_type = parseInt(item.tax_type);

                list.push(item);
            }
            console.log(list);

            EditTradingApp.save(trading, list);
        });
    },
    save: function (trading, items) {
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

var AppRouter = Backbone.Router.extend({
    routes: {
        "": "top",
        "tradings": "tradings",
        "tradings(/:id)": "editTrading"
    },
    initialize: function () {
        _.bindAll(this, 'top', 'tradings', 'editTrading');
    },
    top: function () {
        _this.r = new Ractive({
            el: '#container',
            template: '#topTemplate',
            data: {
                loginInProgress: false
            }
        });
        _this.r.on('login', function (e) {
            TopApp.login(_this.r.get('username'), _this.r.get('password'), _this.r);
        });
    },
    tradings: function () {
        if (app.token == null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        TradingApp.loadTradings(app.token);
    },
    editTrading: function (id) {
        if (app.token == null) {
            app.router.navigate('', { trigger: true });
            return;
        }
        EditTradingApp.loadTrading(app.token, id);
    }
});

var app = {
    //    client : new Invoice.AppClientImpl('http://localhost:9001')
    client: new Invoice.MockClient()
};
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
            } else {
                ret = n1 + "," + ret;
            }
        } while(n > 0);
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

/// <reference path="./Client.ts"/>
var $;
var _;
var baseURL = '';
var AppClientImpl = (function () {
    function AppClientImpl(url) {
        this.url = url;
    }
    AppClientImpl.prototype.login = function (username, password, callback) {
        var params = {
            username: username,
            password: password
        };
        this.exec(this.url + '/api/v1/token', 'POST', null, params, {
            success: function (json) {
                callback.success(json.access_token);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getTradings = function (token, callback) {
        this.exec(this.url + '/api/v1/tradings', 'GET', token, null, {
            success: function (json) {
                callback.success(_.map(json.tradings, function (item) {
                    item.date = item.id;
                    return item;
                }));
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getTradingItems = function (token, tradingId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'GET', token, null, {
            success: function (json) {
                callback.success(_.map(json.items, function (item) {
                    item.sum = item.unit_price * item.amount;
                    return item;
                }));
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getUsers = function (token, callback) {
        var url = this.url + '/api/v1/users';
        this.exec(url, 'GET', token, null, {
            success: function (json) {
                callback.success(json.users);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getCompanies = function (token, callback) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'GET', token, null, {
            success: function (json) {
                callback.success(json.companies);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.saveTrading = function (token, item, callback) {
        if (item.id === null) {
            this.createTrading(token, item, callback);
        }
        else {
            this.updateTrading(token, item, callback);
        }
    };
    AppClientImpl.prototype.saveTradingItem = function (token, tradingId, item, callback) {
        if (item.id === null) {
            this.createTradingItem(token, tradingId, item, callback);
        }
        else {
            this.updateTradingItem(token, tradingId, item, callback);
        }
    };
    AppClientImpl.prototype.deleteTradingItem = function (token, tradingId, itemId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + itemId;
        this.exec(url, 'DELETE', token, null, {
            success: function (json) {
                callback.success(itemId);
            },
            error: function (status, body) {
                if (status == 404) {
                    callback.success(itemId);
                }
                else {
                    callback.error(status, body.msg);
                }
            }
        });
    };
    AppClientImpl.prototype.saveCompany = function (token, item, callback) {
        if (item.id === null || item.id.length == 0) {
            this.createCompany(token, item, callback);
        }
        else {
            this.updateCompany(token, item, callback);
        }
    };
    AppClientImpl.prototype.createTrading = function (token, item, callback) {
        var url = this.url + '/api/v1/tradings';
        this.exec(url, 'POST', token, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.updateTrading = function (token, item, callback) {
        var url = this.url + '/api/v1/tradings/' + item.id;
        this.exec(url, 'PUT', token, item, {
            success: function (json) {
                callback.success(item.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.createTradingItem = function (token, tradingId, item, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'POST', token, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.updateTradingItem = function (token, tradingId, item, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + item.id;
        this.exec(url, 'PUT', token, item, {
            success: function (json) {
                callback.success(item.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.createCompany = function (token, item, callback) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'POST', token, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.updateCompany = function (token, item, callback) {
        var url = this.url + '/api/v1/companies/' + item.id;
        this.exec(url, 'PUT', token, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.exec = function (url, method, token, params, callback) {
        var data = {
            url: url,
            type: method,
            dataType: 'json',
            scriptCharset: 'utf-8',
            processData: false
        };
        if (token != null) {
            data.headers = {
                authorization: 'bearer ' + token
            };
        }
        if (params != null) {
            data.data = JSON.stringify(params);
        }
        $.ajax(data)
            .done(function (data_, status, data) {
            if (data.status == 204) {
                callback.success({});
            }
            else {
                callback.success(JSON.parse(data.responseText));
            }
        }).fail(function (data) {
            if (data.status == 204) {
                callback.success({});
            }
            else {
                callback.error(data.status, JSON.parse(data.responseText));
            }
        });
    };
    return AppClientImpl;
})();
function createClient() {
    return new AppClientImpl(baseURL);
}
var User = (function () {
    function User() {
    }
    return User;
})();
var Company = (function () {
    function Company() {
    }
    return Company;
})();
var Trading = (function () {
    function Trading() {
    }
    return Trading;
})();
var TradingItem = (function () {
    function TradingItem() {
    }
    return TradingItem;
})();
///<reference path="./Dialog.ts"/>
///<reference path="./Client.ts"/>
var App = (function () {
    function App() {
    }
    App.prototype.showDialog = function (dialog) {
        document.querySelector('#dialogs').style.display = 'block';
        app.dialogs.push('dialogs', dialog).then(function () {
            var list = app.dialogs.get('dialogs');
            app.updateDialogs(list);
        });
        $('#container').addClass('dialogOpened');
    };
    App.prototype.updateDialogs = function (list) {
        for (var i = 0; i < list.length; ++i) {
            var s = document.querySelector('#dialog' + i);
            list[i].attach(this, s);
        }
    };
    App.prototype.closeDialog = function () {
        var _this = this;
        this.dialogs.pop('dialogs').then(function () {
            // hide overlay
            var list = _this.dialogs.get('dialogs');
            if (list.length == 0) {
                document.querySelector('#dialogs').style.display = 'none';
                $('#container').removeClass('dialogOpened');
            }
            else {
                _this.updateDialogs(list);
            }
        });
    };
    App.prototype.start = function () {
        this.client = createClient();
        this.initDialog();
    };
    App.prototype.initDialog = function () {
        var _this = this;
        // dialogの準備
        this.dialogs = new Ractive({
            el: '#dialogs',
            template: '#dialogsTemplate',
            data: {
                dialogs: []
            }
        });
        this.dialogs.on({
            'closeClick': function () {
                _this.closeDialog();
            }
        });
    };
    App.prototype.loadData = function (callback) {
        this.loadUsers(callback);
    };
    App.prototype.loadUsers = function (callback) {
        var _this = this;
        if (this.users != null) {
            this.loadTradings(callback);
            return;
        }
        this.client.getUsers(this.accessToken, {
            success: function (list) {
                _this.users = list;
                _this.loadTradings(callback);
            },
            error: function (status, msg) {
                console.log('Failed to get users status=' + status);
                callback.error();
            }
        });
    };
    App.prototype.loadTradings = function (callback) {
        var _this = this;
        if (this.tradings != null) {
            this.loadCompanies(callback);
            return;
        }
        this.client.getTradings(this.accessToken, {
            success: function (list) {
                _this.tradings = list;
                _this.tradingsMap = {};
                _.each(_this.tradings, function (item) {
                    _this.tradingsMap[item.id] = item;
                });
                _this.loadCompanies(callback);
            },
            error: function (status, msg) {
                console.log('Failed to get tradings status=' + status);
                callback.error();
            }
        });
    };
    App.prototype.loadCompanies = function (callback) {
        var _this = this;
        if (this.companies != null) {
            callback.done();
            return;
        }
        this.client.getCompanies(this.accessToken, {
            success: function (list) {
                _this.companies = list;
                _this.companyMap = {};
                _.each(_this.companies, function (item) {
                    _this.companyMap[item.id] = item;
                });
                callback.done();
            },
            error: function (status, msg) {
                console.log('Failed to get companies status=' + status);
                callback.error();
            }
        });
    };
    return App;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
var UserListDialog = (function () {
    function UserListDialog() {
    }
    UserListDialog.prototype.attach = function (app, el) {
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#userListTemplate'
        });
        app.ractive.on({
            'windowClicked': function () {
                return false;
            },
            'close': function () {
                app.closeDialog();
                return false;
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 330);
    };
    return UserListDialog;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
var AddCompanyDialog = (function () {
    function AddCompanyDialog() {
    }
    AddCompanyDialog.prototype.attach = function (app, el) {
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#addCompanyTemplate'
        });
        app.ractive.on({
            'windowClicked': function () {
                return false;
            },
            'close': function () {
                app.closeDialog();
                return false;
            }
        });
    };
    return AddCompanyDialog;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
var CompanyListDialog = (function () {
    function CompanyListDialog() {
    }
    CompanyListDialog.prototype.attach = function (app, el) {
        var _this = this;
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#companyListTemplate',
            data: {
                companyList: app.companies
            }
        });
        this.ractive.on({
            'windowClicked': function () {
                return false;
            },
            'close': function () {
                app.closeDialog();
                return false;
            },
            'showEdit': function (e, item) {
                console.log('clickEvent');
                _this.showEditDialog(app, item);
                return false;
            },
            'submit': function () {
                _this.save(app);
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 370);
    };
    CompanyListDialog.prototype.showEditDialog = function (app, item) {
        app.showDialog(new AddCompanyDialog());
    };
    CompanyListDialog.prototype.save = function (app) {
        var _this = this;
        var name = this.ractive.get('name');
        var unit = this.ractive.get('unit');
        var assignee = this.ractive.get('assignee');
        var zip = this.ractive.get('zip');
        var address = this.ractive.get('address');
        var tel = this.ractive.get('tel');
        var fax = this.ractive.get('fax');
        var company = new Company();
        company.id = null;
        company.name = name;
        company.zip = zip;
        company.address = address;
        company.phone = tel;
        company.fax = fax;
        company.unit = unit;
        company.assignee = assignee;
        app.client.saveCompany(app.accessToken, company, {
            success: function (id) {
                company.id = id;
                app.companyMap[id] = company;
                _this.ractive.unshift('companyList', company);
                _this.clearForm(app);
            },
            error: function (status, msg) {
                console.log('Failed to create company status=' + status);
            }
        });
        console.log(company);
    };
    CompanyListDialog.prototype.clearForm = function (app) {
        this.ractive.set('name', '');
        this.ractive.set('unit', '');
        this.ractive.set('assignee', '');
        this.ractive.set('zip', '');
        this.ractive.set('address', '');
        this.ractive.set('tel', '');
        this.ractive.set('fax', '');
    };
    return CompanyListDialog;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
var SettingsDialog = (function () {
    function SettingsDialog() {
    }
    SettingsDialog.prototype.attach = function (app, el) {
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#settingTemplate'
        });
        app.ractive.on({
            'windowClicked': function () {
                return false;
            },
            'close': function () {
                app.closeDialog();
                return false;
            }
        });
    };
    return SettingsDialog;
})();
///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
var SignInPage = (function () {
    function SignInPage() {
    }
    SignInPage.prototype.onCreate = function (app) {
        var _this = this;
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#signInTemplate',
            // データを設定。テンプレートで使います。
            data: {
                inProgress: false
            }
        });
        app.ractive.on({
            'signIn': function (e, item) {
                var username = r.get('username');
                var password = r.get('password');
                _this.signIn(app, username, password);
            }
        });
    };
    SignInPage.prototype.signIn = function (app, username, password) {
        app.ractive.set('inProgress', true);
        app.ractive.update();
        app.client.login(username, password, {
            success: function (token) {
                app.accessToken = token;
                app.router.navigate('top', { trigger: true });
            },
            error: function (status, msg) {
                app.ractive.set('inProgress', false);
                app.ractive.update();
                console.log('failed to login status=' + status);
            }
        });
    };
    return SignInPage;
})();
///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
var TopPage = (function () {
    function TopPage() {
    }
    TopPage.prototype.onCreate = function (app) {
        var _this = this;
        app.loadData({
            done: function () {
                _this.show(app);
            },
            error: function () {
                // nop
            }
        });
    };
    TopPage.prototype.show = function (app) {
        // Racriveオブジェクトを作る
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#topTemplate',
            // データを設定。テンプレートで使います。
            data: {
                'company': app.companyMap,
                'sheets': app.tradings
            }
        });
        tooltipster();
        app.ractive.on({
            'showSheet': function (e, item) {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, { trigger: true });
            },
            'printQuotation': function (e, item) {
                window.location.href = "/php/quotation.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
            },
            'printBill': function (e, item) {
                window.location.href = "/php/bill.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
            },
            'showUserList': function () {
                app.showDialog(new UserListDialog());
            },
            'showCompanyList': function () {
                app.showDialog(new CompanyListDialog());
            },
            'showSetting': function (e) {
                // #settingに遷移する
                app.showDialog(new SettingsDialog());
            }
        });
    };
    return TopPage;
})();
var Utils;
(function (Utils) {
    function toNumber(source) {
        var num = Number(String(source).replace(",", ""));
        return isNaN(num) ? 0 : num;
    }
    Utils.toNumber = toNumber;
    function toDateStr(time) {
        var date = new Date(time);
        var m = date.getMonth() + 1;
        var d = date.getDate();
        if (m < 10) {
            m = "0" + m;
        }
        if (d < 10) {
            d = "0" + d;
        }
        return date.getFullYear() + "-" + m + "-" + d;
    }
    Utils.toDateStr = toDateStr;
    function clone(source) {
        var dest = {};
        for (var k in source) {
            dest[k] = source[k];
        }
        return dest;
    }
    Utils.clone = clone;
})(Utils || (Utils = {}));
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
var AddUserDialog = (function () {
    function AddUserDialog() {
    }
    AddUserDialog.prototype.attach = function (app, el) {
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#addUserTemplate'
        });
        app.ractive.on({
            'windowClicked': function () {
                return false;
            },
            'close': function () {
                app.closeDialog();
                return false;
            }
        });
    };
    return AddUserDialog;
})();
///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./Functions.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
///<reference path="./AddUserDialog.ts"/>
var SheetPage = (function () {
    function SheetPage(id) {
        this.id = id;
    }
    SheetPage.prototype.onCreate = function (app) {
        var item;
        if (app.tradingsMap === undefined || (item = app.tradingsMap[this.id]) === null) {
            window.history.back();
            return;
        }
        this.loadItems(app, Utils.clone(item));
    };
    SheetPage.prototype.loadItems = function (app, trading) {
        var _this = this;
        app.client.getTradingItems(app.accessToken, trading.id, {
            success: function (list) {
                _this.show(app, trading, list);
            },
            error: function (status, msg) {
                console.log('Failed to get items status=' + status);
                window.history.back();
            }
        });
    };
    SheetPage.prototype.show = function (app, trading, itemList) {
        var _this = this;
        var es = function (node) {
            $(node).easySelectBox({ speed: 200 });
            return {
                teardown: function () {
                    // nop?
                }
            };
        };
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#sheetTemplate',
            decorators: {
                easyselect: es
            },
            data: {
                'trading': trading,
                'workFrom': Utils.toDateStr(trading.work_from),
                'workTo': Utils.toDateStr(trading.work_to),
                'quotationDate': Utils.toDateStr(trading.quotation_date),
                'billDate': Utils.toDateStr(trading.bill_date),
                'companies': app.companies,
                'users': app.users,
                'tradingItems': itemList,
                'deletedItems': []
            }
        });
        var updateItemSum = function (keypath) {
            var unitPrice = Utils.toNumber(r.get(keypath + 'unit_price'));
            var amount = Utils.toNumber(r.get(keypath + 'amount'));
            r.set(keypath + 'sum', unitPrice * amount);
        };
        var updateSum = function () {
            var itemList = r.get('tradingItems');
            var sum = 0;
            var tax = 0;
            var taxRate = Number(r.get('trading.tax_rate'));
            for (var i = 0; i < itemList.length; ++i) {
                var item = itemList[i];
                var taxType = Number($('#tax_type' + i).val());
                if (taxType == 1) {
                    sum += item.sum;
                    tax += item.sum * taxRate / 100;
                }
                else if (taxType == 2) {
                    var body = item.sum * 100 / (100 + taxRate);
                    var taxTmp = Math.ceil(item.sum - body);
                    sum += item.sum - taxTmp;
                    tax += taxTmp;
                }
            }
            r.set('trading.sum', sum);
            r.set('trading.tax', tax);
            r.set('trading.total', sum + tax);
            r.update();
        };
        var observeItem = function () {
            return r.observe({
                'tradingItems.*.unit_price': function (newValue, oldValue, keypath) {
                    updateItemSum(keypath.replace('unit_price', ''));
                },
                'tradingItems.*.amount': function (newValue, oldValue, keypath) {
                    updateItemSum(keypath.replace('amount', ''));
                },
                'tradingItems.*.sum': function (newValue, oldValue, keypath) {
                    updateSum();
                }
            });
        };
        var itemObserver = observeItem();
        r.on({
            'addItem': function () {
                r.push('tradingItems', {
                    id: null,
                    subject: '',
                    unit_price: 0,
                    amount: 0,
                    degree: '',
                    memo: '',
                    tax_type: 1,
                    sum: 0
                });
            },
            'addCompany': function () {
                _this.showAddCompanyDialog(app);
            },
            'addUser': function () {
                _this.showAddUserDialog(app);
            }
        });
        r.on('deleteItem', function (e, index) {
            itemObserver.cancel();
            var item = r.get('tradingItems')[index];
            r.splice('tradingItems', index, 1);
            if (item.id != null) {
                r.push('deletedItems', item);
            }
            itemObserver = observeItem();
        });
        r.on('save', function () {
            _this.save(app);
        });
        r.observe('trading.tax_rate', function (newValue, oldValue, keypath) {
            updateSum();
        });
        // この下にjQuery関連のコードおねがいしやす
        tooltipster();
        //selectbox();
        //sheet();
    };
    SheetPage.prototype.showAddCompanyDialog = function (app) {
        app.showDialog(new AddCompanyDialog());
    };
    SheetPage.prototype.showAddUserDialog = function (app) {
        app.showDialog(new AddUserDialog());
    };
    SheetPage.prototype.save = function (app) {
        var _this = this;
        var trading = app.ractive.get('trading');
        var companyId = $('#company').val();
        var titleType = $('#titleType').val();
        var workFrom = app.ractive.get('workFrom');
        var workTo = app.ractive.get('workTo');
        var quotationDate = app.ractive.get('quotationDate');
        var billDate = app.ractive.get('billDate');
        var tradingItems = app.ractive.get('tradingItems');
        // modify type
        trading.company_id = companyId;
        trading.title_type = Number(titleType);
        trading.work_from = new Date(workFrom).getTime();
        trading.work_to = new Date(workTo).getTime();
        trading.quotation_date = new Date(quotationDate).getTime();
        trading.bill_date = new Date(billDate).getTime();
        trading.tax_rate = Number(trading.tax_rate);
        console.log(trading);
        app.client.saveTrading(app.accessToken, trading, {
            success: function (id) {
                var deleted = app.ractive.get('deletedItems');
                _this.deleteItems(app, id, deleted);
            },
            error: function (status, msg) {
                console.log('Failed to save trading status=' + status);
            }
        });
    };
    SheetPage.prototype.deleteItems = function (app, id, list) {
        var _this = this;
        if (list.length == 0) {
            var list3 = [];
            _.each(app.ractive.get('tradingItems'), function (item, index) {
                item.sort_order = index;
                item.unit_price = Number(item.unit_price);
                item.amount = Number(item.amount);
                item.tax_type = Number($('#tax_type' + index).val());
                list3.push(item);
            });
            this.saveItems(app, id, list3);
            return;
        }
        var item = list[0];
        app.client.deleteTradingItem(app.accessToken, id, item.id, {
            success: function (itemId) {
                list.splice(0, 1);
                _this.deleteItems(app, id, list);
            },
            error: function (status, msg) {
                console.log('Failed to delete items status=' + status);
            }
        });
    };
    SheetPage.prototype.saveItems = function (app, id, list) {
        var _this = this;
        if (list.length == 0) {
            window.history.back();
            return;
        }
        var item = list[0];
        app.client.saveTradingItem(app.accessToken, id, item, {
            success: function (itemId) {
                item.id = itemId;
                list.splice(0, 1);
                _this.saveItems(app, id, list);
            },
            error: function (status, msg) {
                console.log('Failed to save items status=' + status);
            }
        });
    };
    return SheetPage;
})();
///<reference path="./ractive.d.ts"/>
///<reference path="./data.ts"/>
///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./SignInPage.ts"/>
///<reference path="./TopPage.ts"/>
///<reference path="./SheetPage.ts"/>
var $;
var _;
var Backbone;
var app = new App();
var AppRouter = Backbone.Router.extend({
    routes: {
        // ここに、ページ毎に呼ぶ関数名を記述していく
        // index.htmlを開いた直後は、topという関数を実行する        
        "": "signIn",
        "top": "top",
        // index.html#sheetの場合は、sheetという関数を実行する
        "sheets(/:id)": "sheet",
        "setting": "setting"
    },
    signIn: function () {
        app.page = new SignInPage();
        app.page.onCreate(app);
    },
    top: function () {
        app.page = new TopPage();
        app.page.onCreate(app);
    },
    sheet: function (id) {
        app.page = new SheetPage(id);
        app.page.onCreate(app);
    },
    setting: function () {
        // ダイアログ用の要素を作る
        var dialog = document.createElement('section');
        document.querySelector('#dialogs').appendChild(dialog);
        // Racriveオブジェクトを作る
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: dialog,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#settingTemplate'
        });
    }
});
$(function () {
    app.start();
    // Backboneのおまじない
    app.router = new AppRouter();
    Backbone.history.start();
});
// [common] for plugins
function tooltipster() {
    $('.actionBtn li a').tooltipster({
        theme: 'tooltipster-actionBtn'
    });
    $('.btn, .delete').tooltipster({
        theme: 'tooltipster-btn',
        //arrow: false,
        offsetY: -3
    });
}
function selectbox() {
    //select box customize
    //$('select').easySelectBox({speed: 200});
}

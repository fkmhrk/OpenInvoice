var Utils;
(function (Utils) {
    function toList(obj) {
        var list = [];
        for (var k in obj) {
            list.push(obj[k]);
        }
        return list;
    }
    Utils.toList = toList;
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
    function isEmpty(s) {
        return s == null || s.length == 0;
    }
    Utils.isEmpty = isEmpty;
})(Utils || (Utils = {}));
/// <reference path="./Functions.ts"/>
var ClientValidator;
(function (ClientValidator) {
    function isValidLogin(username, password, callback) {
        if (Utils.isEmpty(username)) {
            callback.error(1000, "Username must not be empty.");
            return false;
        }
        if (Utils.isEmpty(password)) {
            callback.error(1001, "Password must not be empty.");
            return false;
        }
        return true;
    }
    ClientValidator.isValidLogin = isValidLogin;
    function isValidCreateUser(loginName, displayName, tel, password, callback) {
        if (Utils.isEmpty(loginName)) {
            callback.error(1000, "LoginName must not be empty.");
            return false;
        }
        if (Utils.isEmpty(displayName)) {
            callback.error(1001, "DisplayName must not be empty.");
            return false;
        }
        if (Utils.isEmpty(tel)) {
            callback.error(1002, "Tel must not be empty.");
            return false;
        }
        if (Utils.isEmpty(password)) {
            callback.error(1003, "Password must not be empty.");
            return false;
        }
        if (password.length < 6) {
            callback.error(1004, "Password must be more than 6 characters.");
            return false;
        }
        return true;
    }
    ClientValidator.isValidCreateUser = isValidCreateUser;
    function isValidSaveCompany(item, callback) {
        if (item == null) {
            callback.error(1000, "Item must not be empty.");
            return false;
        }
        if (Utils.isEmpty(item.name)) {
            callback.error(1001, "Name must not be empty.");
            return false;
        }
        if (item.zip == null) {
            item.zip = '';
        }
        if (item.address == null) {
            item.address = '';
        }
        if (item.phone == null) {
            item.phone = '';
        }
        if (item.fax == null) {
            item.fax = '';
        }
        if (item.unit == null) {
            item.unit = '';
        }
        return true;
    }
    ClientValidator.isValidSaveCompany = isValidSaveCompany;
    function isValidSaveTrading(item, callback) {
        if (item == null) {
            callback.error(1000, "Item must not be empty.");
            return false;
        }
        if (Utils.isEmpty(item.subject)) {
            callback.error(1001, "Subject must not be empty.");
            return false;
        }
        if (item.work_to < item.work_from) {
            callback.error(1002, "Invalid work_from and work_to.");
            return false;
        }
        if (item.tax_rate < 0) {
            callback.error(1003, "tax_rate must be positive(tax_rate > 0).");
            return false;
        }
        if (item.product == null) {
            item.product = '';
        }
        return true;
    }
    ClientValidator.isValidSaveTrading = isValidSaveTrading;
    function isValidSaveTradingItem(tradingId, item, callback) {
        if (Utils.isEmpty(tradingId)) {
            callback.error(1000, "trading ID must not be empty.");
            return false;
        }
        if (item == null) {
            callback.error(1001, "item must not be empty.");
            return false;
        }
        if (Utils.isEmpty(item.subject)) {
            callback.error(1002, "subject must not be empty.");
            return false;
        }
        if (item.tax_type < 0 || item.tax_type > 2) {
            callback.error(1003, "tax_type must be 0, 1, 2.");
            return false;
        }
        if (item.degree == null) {
            item.degree = '';
        }
        if (item.memo == null) {
            item.memo = '';
        }
        return true;
    }
    ClientValidator.isValidSaveTradingItem = isValidSaveTradingItem;
    function isValidSaveEnvironment(env, callback) {
        if (env == null) {
            callback.error(1000, "item must not be empty.");
            return false;
        }
        if (env.tax_rate == null) {
            env.tax_rate = '';
        }
        if (env.quotation_limit == null) {
            env.quotation_limit = '';
        }
        if (env.closing_month == null) {
            env.closing_month = '';
        }
        if (env.pay_limit == null) {
            env.pay_limit = '';
        }
        if (env.company_name == null) {
            env.company_name = '';
        }
        if (env.company_zip == null) {
            env.company_zip = '';
        }
        if (env.company_address == null) {
            env.company_address = '';
        }
        if (env.company_tel == null) {
            env.company_tel = '';
        }
        if (env.company_fax == null) {
            env.company_fax = '';
        }
        if (env.company_bankname == null) {
            env.company_bankname = '';
        }
        if (env.company_bank_type == null) {
            env.company_bank_type = '';
        }
        if (env.company_bank_num == null) {
            env.company_bank_num = '';
        }
        if (env.company_bank_name == null) {
            env.company_bank_name = '';
        }
        return true;
    }
    ClientValidator.isValidSaveEnvironment = isValidSaveEnvironment;
})(ClientValidator || (ClientValidator = {}));
/// <reference path="./Client.ts"/>
/// <reference path="./ClientValidator.ts"/>
var $;
var _;
var baseURL = '';
var AppClientImpl = (function () {
    function AppClientImpl(url) {
        this.url = url;
        this.is_admin = false;
        this.isRetry = false;
    }
    AppClientImpl.prototype.setRefreshToken = function (refreshToken) {
        if (refreshToken == null) {
            return;
        }
        this.accessToken = '';
        this.refreshToken = refreshToken;
    };
    AppClientImpl.prototype.isAdmin = function () {
        return this.is_admin;
    };
    AppClientImpl.prototype.getAccessToken = function () {
        return this.accessToken;
    };
    AppClientImpl.prototype.login = function (username, password, callback) {
        var _this = this;
        if (!ClientValidator.isValidLogin(username, password, callback)) {
            return;
        }
        var params = {
            username: username,
            password: password
        };
        this.exec(this.url + '/api/v1/token', 'POST', null, params, {
            success: function (json) {
                _this.accessToken = json.access_token;
                _this.refreshToken = json.refresh_token;
                _this.is_admin = json.is_admin;
                callback.success(json.refresh_token);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getTradings = function (callback) {
        this.exec(this.url + '/api/v1/tradings', 'GET', this.accessToken, null, {
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
    AppClientImpl.prototype.getTradingItems = function (tradingId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'GET', this.accessToken, null, {
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
    AppClientImpl.prototype.createUser = function (loginName, displayName, tel, password, callback) {
        if (!ClientValidator.isValidCreateUser(loginName, displayName, tel, password, callback)) {
            return;
        }
        var url = this.url + '/api/v1/users';
        var params = {
            login_name: loginName,
            display_name: displayName,
            tel: tel,
            password: password
        };
        this.exec(url, 'POST', this.accessToken, params, {
            success: function (json) {
                callback.success(json);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getUsers = function (callback) {
        var url = this.url + '/api/v1/users';
        this.exec(url, 'GET', this.accessToken, null, {
            success: function (json) {
                callback.success(json.users);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getCompanies = function (callback) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'GET', this.accessToken, null, {
            success: function (json) {
                callback.success(json.companies);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.saveTrading = function (item, callback) {
        if (!ClientValidator.isValidSaveTrading(item, callback)) {
            return;
        }
        if (item.id === null) {
            this.createTrading(item, callback);
        }
        else {
            this.updateTrading(item, callback);
        }
    };
    AppClientImpl.prototype.saveTradingItem = function (tradingId, item, callback) {
        if (!ClientValidator.isValidSaveTradingItem(tradingId, item, callback)) {
            return;
        }
        if (item.id === null) {
            this.createTradingItem(tradingId, item, callback);
        }
        else {
            this.updateTradingItem(tradingId, item, callback);
        }
    };
    AppClientImpl.prototype.deleteTrading = function (tradingId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success: function (json) {
                callback.success();
            },
            error: function (status, body) {
                if (status == 404) {
                    callback.success();
                }
                else {
                    callback.error(status, body.msg);
                }
            }
        });
    };
    AppClientImpl.prototype.deleteTradingItem = function (tradingId, itemId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + itemId;
        this.exec(url, 'DELETE', this.accessToken, null, {
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
    AppClientImpl.prototype.saveCompany = function (item, callback) {
        if (!ClientValidator.isValidSaveCompany(item, callback)) {
            return;
        }
        if (item.id === null || item.id.length == 0) {
            this.createCompany(item, callback);
        }
        else {
            this.updateCompany(item, callback);
        }
    };
    AppClientImpl.prototype.deleteCompany = function (id, callback) {
        var url = this.url + '/api/v1/companies/' + id;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success: function (json) {
                callback.success();
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getEnvironment = function (callback) {
        var url = this.url + '/api/v1/environments';
        this.exec(url, 'GET', this.accessToken, null, {
            success: function (json) {
                callback.success(json);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.saveEnvironment = function (env, callback) {
        if (!ClientValidator.isValidSaveEnvironment(env, callback)) {
            return;
        }
        var url = this.url + '/api/v1/environments';
        this.exec(url, 'PUT', this.accessToken, env, {
            success: function (json) {
                callback.success();
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getMyCompanyName = function (callback) {
        var url = this.url + '/api/v1/myCompany/name';
        this.exec(url, 'GET', null, null, {
            success: function (json) {
                callback.success(json['name']);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.getNextNumber = function (type, date, callback) {
        var url = this.url + '/api/v1/sequences/' + type;
        var params = {
            date: date
        };
        this.exec(url, 'POST', this.accessToken, params, {
            success: function (json) {
                callback.success(json['number']);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.createTrading = function (item, callback) {
        var url = this.url + '/api/v1/tradings';
        this.exec(url, 'POST', this.accessToken, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.updateTrading = function (item, callback) {
        var url = this.url + '/api/v1/tradings/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success: function (json) {
                callback.success(item.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.createTradingItem = function (tradingId, item, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'POST', this.accessToken, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.updateTradingItem = function (tradingId, item, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success: function (json) {
                callback.success(item.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.createCompany = function (item, callback) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'POST', this.accessToken, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.updateCompany = function (item, callback) {
        var url = this.url + '/api/v1/companies/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success: function (json) {
                callback.success(json.id);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.tokenRefresh = function (url, method, params, callback) {
        var _this = this;
        var refreshURL = this.url + '/api/v1/token/refresh';
        var refreshParams = {
            token: this.refreshToken
        };
        this.exec(refreshURL, 'POST', null, refreshParams, {
            success: function (json) {
                _this.accessToken = json.access_token;
                _this.is_admin = json.is_admin;
                _this.isRetry = true;
                _this.exec(url, method, _this.accessToken, params, callback);
            },
            error: function (status, body) {
                callback.error(status, body.msg);
            }
        });
    };
    AppClientImpl.prototype.exec = function (url, method, token, params, callback) {
        var _this = this;
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
        $.ajax(data).done(function (data_, status, data) {
            _this.isRetry = false;
            if (data.status == 204) {
                callback.success({});
            }
            else {
                callback.success(JSON.parse(data.responseText));
            }
        }).fail(function (data) {
            if (data.status == 204) {
                _this.isRetry = false;
                callback.success({});
            }
            else if (data.status == 401) {
                if (_this.isRetry) {
                    _this.isRetry = false;
                    callback.error(data.status, JSON.parse(data.responseText));
                }
                else {
                    _this.isRetry = true;
                    _this.tokenRefresh(url, method, params, callback);
                }
            }
            else {
                _this.isRetry = false;
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
var Environment = (function () {
    function Environment() {
    }
    return Environment;
})();
///<reference path="./Dialog.ts"/>
///<reference path="./Client.ts"/>
///<reference path="./Functions.ts"/>
var App = (function () {
    function App() {
    }
    // getter
    App.prototype.getTradings = function () {
        return Utils.toList(this.tradingsMap);
    };
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
    // snack bar
    App.prototype.addSnack = function (item) {
        var _this = this;
        this.snackbars.push('snackbars', item);
        var closeFunc = function () {
            var list = _this.snackbars.get('snackbars');
            if (list.length == 0) {
                return;
            }
            _this.snackbars.splice('snackbars', 0, 1);
            if (_this.snackbars.get('snackbars').length > 0) {
                setTimeout(closeFunc, 3000);
            }
        };
        setTimeout(closeFunc, 3000);
    };
    App.prototype.start = function () {
        this.client = createClient();
        var refreshToken = localStorage.getItem('refreshToken');
        this.client.setRefreshToken(refreshToken);
        this.initDialog();
        this.initSnackbar();
        this.loadMyCompanyName();
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
    App.prototype.initSnackbar = function () {
        var _this = this;
        // snackbarsの準備
        this.snackbars = new Ractive({
            el: '#snacks',
            template: '#snackbarsTemplate',
            data: {
                snackbars: []
            }
        });
        this.snackbars.on({
            'close': function (e, index) {
                _this.snackbars.splice('snackbars', index, 1);
            }
        });
    };
    App.prototype.loadMyCompanyName = function () {
        var _this = this;
        if (this.myCompanyName != null && this.myCompanyName.length > 0) {
            return;
        }
        this.client.getMyCompanyName({
            success: function (name) {
                _this.myCompanyName = name;
            },
            error: function (status, msg) {
                console.log('Failed to get my company name status=' + status);
            }
        });
    };
    App.prototype.loadData = function (callback) {
        this.loadEnvironment(callback);
    };
    App.prototype.loadEnvironment = function (callback) {
        var _this = this;
        if (this.environment != null) {
            this.loadUsers(callback);
            return;
        }
        this.client.getEnvironment({
            success: function (item) {
                _this.environment = item;
                _this.loadUsers(callback);
            },
            error: function (status, msg) {
                console.log('Failed to get environment status=' + status);
                callback.error();
            }
        });
    };
    App.prototype.loadUsers = function (callback) {
        var _this = this;
        if (this.users != null) {
            this.loadTradings(callback);
            return;
        }
        this.client.getUsers({
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
        if (this.tradingsMap != null) {
            this.loadCompanies(callback);
            return;
        }
        this.client.getTradings({
            success: function (list) {
                _this.tradingsMap = {};
                _.each(list, function (item) {
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
        this.client.getCompanies({
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
    App.prototype.addCompany = function (c) {
        this.companies.push(c);
        this.companyMap[c.id] = c;
    };
    return App;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
var UserListDialog = (function () {
    function UserListDialog() {
    }
    UserListDialog.prototype.attach = function (app, el) {
        var _this = this;
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#userListTemplate',
            data: {
                userList: app.users
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
            'create': function () {
                _this.createUser(app);
                return false;
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 330);
    };
    UserListDialog.prototype.createUser = function (app) {
        var _this = this;
        var loginName = this.ractive.get('loginName');
        var displayName = this.ractive.get('displayName');
        var tel = this.ractive.get('tel');
        var password = this.ractive.get('password');
        app.client.createUser(loginName, displayName, tel, password, {
            success: function (user) {
                _this.ractive.push('userList', user);
                _this.clear();
                app.addSnack('ユーザーを作成しました！');
            },
            error: function (status, msg) {
                switch (status) {
                    case 1000:
                        app.addSnack('ユーザー名を入力してください');
                        break;
                    case 1001:
                        app.addSnack('担当者名を入力してください');
                        break;
                    case 1002:
                        app.addSnack('電話番号を入力してください');
                        break;
                    case 1003: // same message
                    case 1004:
                        app.addSnack('パスワードを6文字以上入力してください');
                        break;
                    default:
                        app.addSnack('ユーザー作成に失敗しました');
                        break;
                }
            }
        });
    };
    UserListDialog.prototype.clear = function () {
        this.ractive.set('loginName', '');
        this.ractive.set('displayName', '');
        this.ractive.set('tel', '');
        this.ractive.set('password', '');
    };
    return UserListDialog;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>
var AddCompanyDialog = (function () {
    function AddCompanyDialog(company, callback) {
        if (company == null) {
            this.isNew = true;
            this.company = new Company();
            this.companyOrg = null;
        }
        else {
            this.isNew = false;
            this.company = Utils.clone(company);
            this.companyOrg = company;
        }
        this.callback = callback;
    }
    AddCompanyDialog.prototype.attach = function (app, el) {
        var _this = this;
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#addCompanyTemplate',
            data: {
                isNew: this.isNew,
                company: this.company
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
            'save': function () {
                _this.save(app);
                return false;
            }
        });
    };
    AddCompanyDialog.prototype.save = function (app) {
        var _this = this;
        var company = this.ractive.get('company');
        console.log(company);
        app.client.saveCompany(company, {
            success: function (id) {
                // clone?
                company.id = id;
                if (_this.companyOrg == null) {
                    app.addCompany(company);
                }
                else {
                    _this.companyOrg.name = company.name;
                    _this.companyOrg.unit = company.unit;
                    _this.companyOrg.assignee = company.assignee;
                    _this.companyOrg.zip = company.zip;
                    _this.companyOrg.address = company.address;
                    _this.companyOrg.phone = company.phone;
                    _this.companyOrg.fax = company.fax;
                }
                _this.callback(company);
                app.addSnack('保存しました。');
                app.closeDialog();
            },
            error: function (status, msg) {
                switch (status) {
                    case 1001:
                        app.addSnack('会社名を入力してください');
                        break;
                    default:
                        app.addSnack('保存に失敗しました。');
                        break;
                }
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
            'deleteCompany': function (e, index) {
                _this.deleteCompany(app, index);
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
        app.showDialog(new AddCompanyDialog(item, function (result) {
            // nop
        }));
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
        app.client.saveCompany(company, {
            success: function (id) {
                company.id = id;
                app.companyMap[id] = company;
                _this.ractive.unshift('companyList', company);
                app.addSnack('保存しました。');
                _this.clearForm(app);
            },
            error: function (status, msg) {
                switch (status) {
                    case 1001:
                        app.addSnack('会社名を入力してください。');
                        break;
                    default: app.addSnack('保存に失敗しました。');
                }
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
    CompanyListDialog.prototype.deleteCompany = function (app, index) {
        var _this = this;
        if (!window.confirm('この会社情報を削除しますか？')) {
            return;
        }
        var company = this.ractive.get('companyList')[index];
        app.client.deleteCompany(company.id, {
            success: function () {
                _this.ractive.splice('companyList', index, 1);
                app.addSnack('削除しました！');
            },
            error: function (status, msg) {
                console.log('Failed to delete company status=' + status);
            }
        });
    };
    return CompanyListDialog;
})();
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
var SettingsDialog = (function () {
    function SettingsDialog() {
    }
    SettingsDialog.prototype.attach = function (app, el) {
        var _this = this;
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#settingTemplate',
            decorators: {},
            data: {
                tax_rate: app.environment.tax_rate,
                quotation_limit: app.environment.quotation_limit,
                closing_month: app.environment.closing_month,
                pay_limit: app.environment.pay_limit,
                company_name: app.environment.company_name,
                company_zip: app.environment.company_zip,
                company_address: app.environment.company_address,
                company_tel: app.environment.company_tel,
                company_fax: app.environment.company_fax,
                company_bankname: app.environment.company_bankname,
                company_bank_type: app.environment.company_bank_type,
                company_bank_num: app.environment.company_bank_num,
                company_bank_name: app.environment.company_bank_name
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
            'save': function () {
                _this.save(app);
                return false;
            }
        });
    };
    SettingsDialog.prototype.save = function (app) {
        var env = new Environment();
        env.tax_rate = this.ractive.get('tax_rate');
        env.quotation_limit = this.ractive.get('quotation_limit');
        env.closing_month = this.ractive.get('closing_month');
        env.pay_limit = $('#pay_limit').val();
        env.company_name = this.ractive.get('company_name');
        env.company_zip = this.ractive.get('company_zip');
        env.company_address = this.ractive.get('company_address');
        env.company_tel = this.ractive.get('company_tel');
        env.company_fax = this.ractive.get('company_fax');
        env.company_bankname = this.ractive.get('company_bankname');
        env.company_bank_type = $('#bank_type').val();
        env.company_bank_num = this.ractive.get('company_bank_num');
        env.company_bank_name = this.ractive.get('company_bank_name');
        app.client.saveEnvironment(env, {
            success: function () {
                app.environment = env;
                app.addSnack('設定を保存しました！');
                app.closeDialog();
            },
            error: function (status, msg) {
                console.log('Failed to save environment statu=' + status);
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
                myCompanyName: app.myCompanyName,
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
                localStorage.setItem('refreshToken', token);
                app.router.navigate('top', { trigger: true });
            },
            error: function (status, msg) {
                app.ractive.set('inProgress', false);
                app.ractive.update();
                switch (status) {
                    case 1000:
                        app.addSnack('ユーザー名を入力してください');
                        break;
                    case 1001:
                        app.addSnack('パスワードを入力してください');
                        break;
                }
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
///<reference path="./Functions.ts"/>
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
        var _this = this;
        var sheets = app.getTradings();
        var total = 0;
        _.each(sheets, function (item) {
            total += item.total;
        });
        var fabDecorate = function (node) {
            $(node).hover(function () {
                $(this).find(".menu").toggleClass("current");
                $(this).find(".submenu").toggleClass("current");
                $(this).next("span").fadeIn();
            }, function () {
                $(this).find(".menu").toggleClass("current");
                $(this).find(".submenu").toggleClass("current");
                $(this).next("span").fadeOut();
            });
            return {
                teardown: function () { }
            };
        };
        // Racriveオブジェクトを作る
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#topTemplate',
            decorators: {
                fab: fabDecorate
            },
            // データを設定。テンプレートで使います。
            data: {
                myCompanyName: app.myCompanyName,
                is_admin: app.client.isAdmin(),
                'company': app.companyMap,
                'sheets': sheets,
                'toDateStr': Utils.toDateStr,
                total: total
            }
        });
        tooltipster();
        app.ractive.on({
            'addSheet': function (e, item) {
                // #sheetに遷移する
                app.router.navigate('sheets/new', { trigger: true });
            },
            'showSheet': function (e, item) {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, { trigger: true });
            },
            'deleteSheet': function (e, index) {
                _this.deleteSheet(app, index);
                return false;
            },
            'copySheet': function (e, item) {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id + '/copy', { trigger: true });
                return false;
            },
            'printQuotation': function (e, item) {
                window.location.href = "/php/quotation.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
                return false;
            },
            'printBill': function (e, item) {
                window.location.href = "/php/bill.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
                return false;
            },
            'showUserList': function () {
                app.showDialog(new UserListDialog());
            },
            'showCompanyList': function () {
                app.showDialog(new CompanyListDialog());
            },
            'showSetting': function (e) {
                app.showDialog(new SettingsDialog());
            }
        });
    };
    TopPage.prototype.deleteSheet = function (app, index) {
        if (!window.confirm('このシートを削除しますか？')) {
            return;
        }
        var item = app.ractive.get('sheets')[index];
        app.client.deleteTrading(item.id, {
            success: function () {
                app.ractive.splice('sheets', index, 1);
                app.addSnack('削除しました！');
            },
            error: function (status, msg) {
                console.log('Failed to delete item status=' + status);
            }
        });
    };
    return TopPage;
})();
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
    function SheetPage(id, copyMode) {
        this.id = id;
        this.copyMode = copyMode;
    }
    SheetPage.prototype.onCreate = function (app) {
        var item;
        if (app.environment == null) {
            window.history.back();
            return;
        }
        if (this.id == 'new') {
            var trading = new Trading();
            trading.id = null;
            trading.tax_rate = Number(app.environment['tax_rate']);
            this.show(app, trading, []);
            return;
        }
        if (app.tradingsMap === undefined || (item = app.tradingsMap[this.id]) === undefined) {
            window.history.back();
            return;
        }
        this.loadItems(app, Utils.clone(item));
    };
    SheetPage.prototype.loadItems = function (app, trading) {
        var _this = this;
        app.client.getTradingItems(trading.id, {
            success: function (list) {
                // if copyMode = true remove ids
                if (_this.copyMode) {
                    trading.id = null;
                    _.each(list, function (item) {
                        item.id = null;
                    });
                }
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
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#sheetTemplate',
            decorators: {},
            data: {
                myCompanyName: app.myCompanyName,
                is_admin: app.client.isAdmin(),
                'trading': trading,
                'workFrom': Utils.toDateStr(trading.work_from),
                'workTo': Utils.toDateStr(trading.work_to),
                'quotationDate': Utils.toDateStr(trading.quotation_date),
                'billDate': Utils.toDateStr(trading.bill_date),
                'deliveryDate': Utils.toDateStr(trading.delivery_date),
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
            'close': function () {
                window.history.back();
            },
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
            },
            'printQuotation': function () {
                _this.printQuotation(app);
            },
            'printBill': function () {
                _this.printBill(app);
            },
            'printDelivery': function () {
                _this.printDelivery(app);
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
            _this.save(app, function (id) {
                window.history.back();
            });
        });
        r.observe('trading.tax_rate', function (newValue, oldValue, keypath) {
            updateSum();
        });
        // この下にjQuery関連のコードおねがいしやす
        tooltipster();
    };
    SheetPage.prototype.showAddCompanyDialog = function (app) {
        app.showDialog(new AddCompanyDialog(null, function (result) {
            app.ractive.update();
        }));
    };
    SheetPage.prototype.showAddUserDialog = function (app) {
        app.showDialog(new AddUserDialog());
    };
    SheetPage.prototype.printQuotation = function (app) {
        var _this = this;
        var trading = app.ractive.get('trading');
        var quotationDate = app.ractive.get('quotationDate');
        var doneFunc = function (id) {
            app.ractive.update();
            window.location.href = "/php/quotation.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.quotation_number == null || trading.quotation_number.length == 0) {
            app.client.getNextNumber('quotation', new Date(quotationDate).getTime(), {
                success: function (val) {
                    trading.quotation_number = '' + val + '-I';
                    _this.save(app, doneFunc);
                },
                error: function (status, msg) {
                    console.log('Failed to get next quotation number status=' + status);
                }
            });
        }
        else {
            this.save(app, doneFunc);
        }
    };
    SheetPage.prototype.printBill = function (app) {
        var _this = this;
        var trading = app.ractive.get('trading');
        var billDate = app.ractive.get('billDate');
        var doneFunc = function (id) {
            app.ractive.update();
            window.location.href = "/php/bill.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.bill_number == null || trading.bill_number.length == 0) {
            app.client.getNextNumber('bill', new Date(billDate).getTime(), {
                success: function (val) {
                    trading.bill_number = '' + val + '-V';
                    _this.save(app, doneFunc);
                },
                error: function (status, msg) {
                    console.log('Failed to get next bill number status=' + status);
                }
            });
        }
        else {
            this.save(app, doneFunc);
        }
    };
    SheetPage.prototype.printDelivery = function (app) {
        var _this = this;
        var trading = app.ractive.get('trading');
        var deliveryDate = app.ractive.get('deliveryDate');
        var doneFunc = function (id) {
            app.ractive.update();
            window.location.href = "/php/delivery.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.delivery_number == null || trading.delivery_number.length == 0) {
            app.client.getNextNumber('delivery', new Date(deliveryDate).getTime(), {
                success: function (val) {
                    trading.delivery_number = '' + val + '-D';
                    _this.save(app, doneFunc);
                },
                error: function (status, msg) {
                    console.log('Failed to get next delivery number status=' + status);
                }
            });
        }
        else {
            this.save(app, doneFunc);
        }
    };
    SheetPage.prototype.save = function (app, doneFunc) {
        var _this = this;
        var trading = app.ractive.get('trading');
        var workFrom = app.ractive.get('workFrom');
        var workTo = app.ractive.get('workTo');
        var quotationDate = app.ractive.get('quotationDate');
        var billDate = app.ractive.get('billDate');
        var deliveryDate = app.ractive.get('deliveryDate');
        var tradingItems = app.ractive.get('tradingItems');
        // modify type
        trading.title_type = Number(trading.title_type);
        trading.work_from = new Date(workFrom).getTime();
        trading.work_to = new Date(workTo).getTime();
        trading.quotation_date = new Date(quotationDate).getTime();
        trading.bill_date = new Date(billDate).getTime();
        trading.delivery_date = new Date(deliveryDate).getTime();
        trading.tax_rate = Number(trading.tax_rate);
        console.log(trading);
        app.client.saveTrading(trading, {
            success: function (id) {
                trading.modified_time = new Date().getTime();
                app.tradingsMap[id] = trading;
                var deleted = app.ractive.get('deletedItems');
                _this.deleteItems(app, id, deleted, doneFunc);
            },
            error: function (status, msg) {
                switch (status) {
                    case 1001:
                        app.addSnack('件名を入力してください。');
                        break;
                    case 1002:
                        app.addSnack('作業終了日は作業開始日より後にしてください。');
                        break;
                    case 1003:
                        app.addSnack('消費税率は0以上にしてください。');
                        break;
                    default: app.addSnack('保存に失敗しました。');
                }
                console.log('Failed to save trading status=' + status);
            }
        });
    };
    SheetPage.prototype.deleteItems = function (app, id, list, doneFunc) {
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
            this.saveItems(app, id, list3, doneFunc);
            return;
        }
        var item = list[0];
        app.client.deleteTradingItem(id, item.id, {
            success: function (itemId) {
                list.splice(0, 1);
                _this.deleteItems(app, id, list, doneFunc);
            },
            error: function (status, msg) {
                console.log('Failed to delete items status=' + status);
            }
        });
    };
    SheetPage.prototype.saveItems = function (app, id, list, doneFunc) {
        var _this = this;
        if (list.length == 0) {
            app.addSnack('保存しました！');
            doneFunc(id);
            return;
        }
        var item = list[0];
        app.client.saveTradingItem(id, item, {
            success: function (itemId) {
                item.id = itemId;
                list.splice(0, 1);
                _this.saveItems(app, id, list, doneFunc);
            },
            error: function (status, msg) {
                switch (status) {
                    case 1002:
                        app.addSnack('項目名を入力してください。');
                        break;
                    case 1003:
                        app.addSnack('税区分が不正です。');
                        break;
                    default: app.addSnack('保存に失敗しました。');
                }
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
        "sheets(/:id)/copy": "copySheet",
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
        app.page = new SheetPage(id, false);
        app.page.onCreate(app);
    },
    copySheet: function (id) {
        app.page = new SheetPage(id, true);
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

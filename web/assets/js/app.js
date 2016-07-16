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
    function isValidSaveUser(user, password, callback) {
        if (user == null) {
            callback.error(1000, "User must not be empty.");
            return false;
        }
        if (Utils.isEmpty(user.login_name)) {
            callback.error(1001, "LoginName must not be empty.");
            return false;
        }
        if (Utils.isEmpty(user.display_name)) {
            callback.error(1002, "DisplayName must not be empty.");
            return false;
        }
        if (Utils.isEmpty(user.tel)) {
            callback.error(1003, "Tel must not be empty.");
            return false;
        }
        if (!Utils.isEmpty(password)) {
            if (password.length < 6) {
                callback.error(1004, "Password must be more than 6 characters.");
                return false;
            }
        }
        return true;
    }
    ClientValidator.isValidSaveUser = isValidSaveUser;
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
class AppClientImpl {
    constructor(url) {
        this.url = url;
        this.is_admin = false;
        this.isRetry = false;
    }
    setRefreshToken(refreshToken) {
        if (refreshToken == null) {
            return;
        }
        this.accessToken = '';
        this.refreshToken = refreshToken;
    }
    isAdmin() {
        return this.is_admin;
    }
    getAccessToken() {
        return this.accessToken;
    }
    login(username, password, callback) {
        if (!ClientValidator.isValidLogin(username, password, callback)) {
            return;
        }
        var params = {
            username: username,
            password: password,
        };
        this.exec(this.url + '/api/v1/token', 'POST', null, params, {
            success: (json) => {
                this.accessToken = json.access_token;
                this.refreshToken = json.refresh_token;
                this.is_admin = json.is_admin;
                callback.success(json.refresh_token);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getTradings(callback) {
        this.exec(this.url + '/api/v1/tradings', 'GET', this.accessToken, null, {
            success: (json) => {
                callback.success(_.map(json.tradings, (item) => {
                    item.date = item.id;
                    return item;
                }));
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getTradingItems(tradingId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'GET', this.accessToken, null, {
            success: (json) => {
                callback.success(_.map(json.items, (item) => {
                    item.sum = item.unit_price * item.amount;
                    return item;
                }));
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    createUser(loginName, displayName, tel, password, callback) {
        if (!ClientValidator.isValidCreateUser(loginName, displayName, tel, password, callback)) {
            return;
        }
        var url = this.url + '/api/v1/users';
        var params = {
            login_name: loginName,
            display_name: displayName,
            tel: tel,
            password: password,
        };
        this.exec(url, 'POST', this.accessToken, params, {
            success: (json) => {
                callback.success(json);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getUsers(callback) {
        var url = this.url + '/api/v1/users';
        this.exec(url, 'GET', this.accessToken, null, {
            success: (json) => {
                callback.success(json.users);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    saveUser(user, password, callback) {
        if (!ClientValidator.isValidSaveUser(user, password, callback)) {
            return;
        }
        var url = this.url + '/api/v1/users/' + user.id;
        var params = {
            id: user.id,
            login_name: user.login_name,
            display_name: user.display_name,
            tel: user.tel,
            password: password,
        };
        this.exec(url, 'PUT', this.accessToken, params, {
            success: (json) => {
                callback.success(params);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    deleteUser(id, callback) {
        var url = this.url + '/api/v1/users/' + id;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success: (json) => {
                callback.success();
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getCompanies(callback) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'GET', this.accessToken, null, {
            success: (json) => {
                callback.success(json.companies);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    saveTrading(item, callback) {
        if (!ClientValidator.isValidSaveTrading(item, callback)) {
            return;
        }
        if (item.id === null) {
            this.createTrading(item, callback);
        }
        else {
            this.updateTrading(item, callback);
        }
    }
    saveTradingItem(tradingId, item, callback) {
        if (!ClientValidator.isValidSaveTradingItem(tradingId, item, callback)) {
            return;
        }
        if (item.id === null) {
            this.createTradingItem(tradingId, item, callback);
        }
        else {
            this.updateTradingItem(tradingId, item, callback);
        }
    }
    deleteTrading(tradingId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success: (json) => {
                callback.success();
            },
            error: (status, body) => {
                if (status == 404) {
                    callback.success();
                }
                else {
                    callback.error(status, body.msg);
                }
            }
        });
    }
    deleteTradingItem(tradingId, itemId, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + itemId;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success: (json) => {
                callback.success(itemId);
            },
            error: (status, body) => {
                if (status == 404) {
                    callback.success(itemId);
                }
                else {
                    callback.error(status, body.msg);
                }
            }
        });
    }
    saveCompany(item, callback) {
        if (!ClientValidator.isValidSaveCompany(item, callback)) {
            return;
        }
        if (item.id === null || item.id.length == 0) {
            this.createCompany(item, callback);
        }
        else {
            this.updateCompany(item, callback);
        }
    }
    deleteCompany(id, callback) {
        var url = this.url + '/api/v1/companies/' + id;
        this.exec(url, 'DELETE', this.accessToken, null, {
            success: (json) => {
                callback.success();
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getEnvironment(callback) {
        var url = this.url + '/api/v1/environments';
        this.exec(url, 'GET', this.accessToken, null, {
            success: (json) => {
                callback.success(json);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    saveEnvironment(env, callback) {
        if (!ClientValidator.isValidSaveEnvironment(env, callback)) {
            return;
        }
        var url = this.url + '/api/v1/environments';
        this.exec(url, 'PUT', this.accessToken, env, {
            success: (json) => {
                callback.success();
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getMyCompanyName(callback) {
        var url = this.url + '/api/v1/myCompany/name';
        this.exec(url, 'GET', null, null, {
            success: (json) => {
                callback.success(json['name']);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    getNextNumber(type, date, callback) {
        var url = this.url + '/api/v1/sequences/' + type;
        var params = {
            date: date,
        };
        this.exec(url, 'POST', this.accessToken, params, {
            success: (json) => {
                callback.success(json['number']);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    createInvoice(items, callback) {
        var params = {
            access_token: this.accessToken,
            customer: {
                name: "サンプル会社",
                address: "サンプル住所",
            },
            myCompany: {
                name: "サンプル会社",
                address: "住所\n\n担当",
            },
            item_title: 'ご請求書在中',
            date: new Date().getTime(),
            items: items,
        };
        var xhr = new XMLHttpRequest();
        xhr.open('POST', this.url + '/php/invoice.php', true);
        xhr.responseType = 'arraybuffer';
        xhr.onload = function (e) {
            if (this.status == 200) {
                callback.success(this.response);
            }
            else {
                callback.error(this.status, this.response);
            }
        };
        xhr.send(JSON.stringify(params));
    }
    createTrading(item, callback) {
        var url = this.url + '/api/v1/tradings';
        this.exec(url, 'POST', this.accessToken, item, {
            success: (json) => {
                callback.success(json.id);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    updateTrading(item, callback) {
        var url = this.url + '/api/v1/tradings/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success: (json) => {
                callback.success(item.id);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    createTradingItem(tradingId, item, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId + '/items';
        this.exec(url, 'POST', this.accessToken, item, {
            success: (json) => {
                callback.success(json.id);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    updateTradingItem(tradingId, item, callback) {
        var url = this.url + '/api/v1/tradings/' + tradingId +
            '/items/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success: (json) => {
                callback.success(item.id);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    createCompany(item, callback) {
        var url = this.url + '/api/v1/companies';
        this.exec(url, 'POST', this.accessToken, item, {
            success: (json) => {
                callback.success(json.id);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    updateCompany(item, callback) {
        var url = this.url + '/api/v1/companies/' + item.id;
        this.exec(url, 'PUT', this.accessToken, item, {
            success: (json) => {
                callback.success(json.id);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    tokenRefresh(url, method, params, callback) {
        var refreshURL = this.url + '/api/v1/token/refresh';
        var refreshParams = {
            token: this.refreshToken
        };
        this.exec(refreshURL, 'POST', null, refreshParams, {
            success: (json) => {
                this.accessToken = json.access_token;
                this.is_admin = json.is_admin;
                this.isRetry = true;
                this.exec(url, method, this.accessToken, params, callback);
            },
            error: (status, body) => {
                callback.error(status, body.msg);
            }
        });
    }
    exec(url, method, token, params, callback) {
        var data = {
            url: url,
            type: method,
            dataType: 'json',
            scriptCharset: 'utf-8',
            processData: false,
        };
        if (token != null) {
            data.headers = {
                authorization: 'bearer ' + token
            };
        }
        if (params != null) {
            data.data = JSON.stringify(params);
        }
        $.ajax(data).done((data_, status, data) => {
            this.isRetry = false;
            if (data.status == 204) {
                callback.success({});
            }
            else {
                callback.success(JSON.parse(data.responseText));
            }
        }).fail((data) => {
            if (data.status == 204) {
                this.isRetry = false;
                callback.success({});
            }
            else if (data.status == 401) {
                if (this.isRetry) {
                    this.isRetry = false;
                    callback.error(data.status, JSON.parse(data.responseText));
                }
                else {
                    this.isRetry = true;
                    this.tokenRefresh(url, method, params, callback);
                }
            }
            else {
                this.isRetry = false;
                callback.error(data.status, JSON.parse(data.responseText));
            }
        });
    }
}
function createClient() {
    return new AppClientImpl(baseURL);
}
class User {
}
class Company {
}
class Trading {
}
class TradingItem {
}
class Environment {
}
///<reference path="./Dialog.ts"/>
///<reference path="./Client.ts"/>
///<reference path="./Functions.ts"/>
class App {
    // getter
    getTradings() {
        return Utils.toList(this.tradingsMap);
    }
    showDialog(dialog) {
        document.querySelector('#dialogs').style.display = 'block';
        app.dialogs.push('dialogs', dialog).then(() => {
            var list = app.dialogs.get('dialogs');
            app.updateDialogs(list);
        });
        $('#container').addClass('dialogOpened');
    }
    updateDialogs(list) {
        for (var i = 0; i < list.length; ++i) {
            var s = document.querySelector('#dialog' + i);
            list[i].attach(this, s);
        }
    }
    closeDialog() {
        this.dialogs.pop('dialogs').then(() => {
            // hide overlay
            var list = this.dialogs.get('dialogs');
            if (list.length == 0) {
                document.querySelector('#dialogs').style.display = 'none';
                $('#container').removeClass('dialogOpened');
            }
            else {
                this.updateDialogs(list);
            }
        });
    }
    // snack bar
    addSnack(item) {
        this.snackbars.push('snackbars', item);
        var closeFunc = () => {
            var list = this.snackbars.get('snackbars');
            if (list.length == 0) {
                return;
            }
            this.snackbars.splice('snackbars', 0, 1);
            if (this.snackbars.get('snackbars').length > 0) {
                setTimeout(closeFunc, 3000);
            }
        };
        setTimeout(closeFunc, 3000);
    }
    start() {
        this.client = createClient();
        var refreshToken = localStorage.getItem('refreshToken');
        this.client.setRefreshToken(refreshToken);
        this.initDialog();
        this.initSnackbar();
        this.loadMyCompanyName();
    }
    initDialog() {
        // dialogの準備
        this.dialogs = new Ractive({
            el: '#dialogs',
            template: '#dialogsTemplate',
            data: {
                dialogs: [],
            }
        });
        this.dialogs.on({
            'closeClick': () => {
                this.closeDialog();
            }
        });
    }
    initSnackbar() {
        // snackbarsの準備
        this.snackbars = new Ractive({
            el: '#snacks',
            template: '#snackbarsTemplate',
            data: {
                snackbars: [],
            }
        });
        this.snackbars.on({
            'close': (e, index) => {
                this.snackbars.splice('snackbars', index, 1);
            }
        });
    }
    loadMyCompanyName() {
        if (this.myCompanyName != null && this.myCompanyName.length > 0) {
            return;
        }
        this.client.getMyCompanyName({
            success: (name) => {
                this.myCompanyName = name;
            },
            error: (status, msg) => {
                console.log('Failed to get my company name status=' + status);
            }
        });
    }
    loadData(callback) {
        this.loadEnvironment(callback);
    }
    loadEnvironment(callback) {
        if (this.environment != null) {
            this.loadUsers(callback);
            return;
        }
        this.client.getEnvironment({
            success: (item) => {
                this.environment = item;
                this.loadUsers(callback);
            },
            error: (status, msg) => {
                console.log('Failed to get environment status=' + status);
                callback.error();
            }
        });
    }
    loadUsers(callback) {
        if (this.users != null) {
            this.loadTradings(callback);
            return;
        }
        this.client.getUsers({
            success: (list) => {
                this.users = list;
                this.loadTradings(callback);
            },
            error: (status, msg) => {
                console.log('Failed to get users status=' + status);
                callback.error();
            }
        });
    }
    loadTradings(callback) {
        if (this.tradingsMap != null) {
            this.loadCompanies(callback);
            return;
        }
        this.client.getTradings({
            success: (list) => {
                this.tradingsMap = {};
                _.each(list, (item) => {
                    this.tradingsMap[item.id] = item;
                });
                this.loadCompanies(callback);
            },
            error: (status, msg) => {
                console.log('Failed to get tradings status=' + status);
                callback.error();
            }
        });
    }
    loadCompanies(callback) {
        if (this.companies != null) {
            callback.done();
            return;
        }
        this.client.getCompanies({
            success: (list) => {
                this.companies = list;
                this.companyMap = {};
                _.each(this.companies, (item) => {
                    this.companyMap[item.id] = item;
                });
                callback.done();
            },
            error: (status, msg) => {
                console.log('Failed to get companies status=' + status);
                callback.error();
            }
        });
    }
    addUser(u) {
        this.users.push(u);
    }
    deleteUser(u) {
        var list = [];
        _.each(this.users, (item) => {
            if (item.id == u.id) {
                return;
            }
            list.push(item);
        });
        this.users = list;
    }
    addCompany(c) {
        this.companies.push(c);
        this.companyMap[c.id] = c;
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
class UserListDialog {
    attach(app, el) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#userListTemplate',
            data: {
                userList: app.users,
            }
        });
        this.ractive.on({
            'windowClicked': () => {
                return false;
            },
            'close': () => {
                app.closeDialog();
                return false;
            },
            'showEdit': (e, item) => {
                this.showEditDialog(app, item);
                return false;
            },
            'delete': (e, item) => {
                this.deleteUser(app, item);
                return false;
            },
            'create': () => {
                this.createUser(app);
                return false;
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 330);
    }
    showEditDialog(app, item) {
        app.showDialog(new AddUserDialog(item, (result) => {
            this.ractive.update();
        }));
    }
    createUser(app) {
        var loginName = this.ractive.get('loginName');
        var displayName = this.ractive.get('displayName');
        var tel = this.ractive.get('tel');
        var password = this.ractive.get('password');
        app.client.createUser(loginName, displayName, tel, password, {
            success: (user) => {
                app.addUser(user);
                this.clear();
                app.addSnack('ユーザーを作成しました！');
            },
            error: (status, msg) => {
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
            },
        });
    }
    deleteUser(app, user) {
        if (!window.confirm('この担当者を削除しますか？')) {
            return;
        }
        app.client.deleteUser(user.id, {
            success: () => {
                app.deleteUser(user);
                this.ractive.set('userList', app.users);
                app.addSnack('担当者を削除しました');
            },
            error: (status, msg) => {
                switch (status) {
                    default: app.addSnack('削除に失敗しました');
                }
            }
        });
    }
    clear() {
        this.ractive.set('loginName', '');
        this.ractive.set('displayName', '');
        this.ractive.set('tel', '');
        this.ractive.set('password', '');
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>
class AddCompanyDialog {
    constructor(company, callback) {
        if (company == null) {
            this.isNew = true;
            this.company = new Company();
            this.company.id = null;
            this.companyOrg = null;
        }
        else {
            this.isNew = false;
            this.company = Utils.clone(company);
            this.companyOrg = company;
        }
        this.callback = callback;
    }
    attach(app, el) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#addCompanyTemplate',
            data: {
                isNew: this.isNew,
                company: this.company,
            }
        });
        this.ractive.on({
            'windowClicked': () => {
                return false;
            },
            'close': () => {
                app.closeDialog();
                return false;
            },
            'save': () => {
                this.save(app);
                return false;
            }
        });
    }
    save(app) {
        var company = this.ractive.get('company');
        console.log(company);
        app.client.saveCompany(company, {
            success: (id) => {
                // clone?
                company.id = id;
                if (this.companyOrg == null) {
                    app.addCompany(company);
                }
                else {
                    this.companyOrg.name = company.name;
                    this.companyOrg.unit = company.unit;
                    this.companyOrg.assignee = company.assignee;
                    this.companyOrg.zip = company.zip;
                    this.companyOrg.address = company.address;
                    this.companyOrg.phone = company.phone;
                    this.companyOrg.fax = company.fax;
                }
                this.callback(company);
                app.addSnack('保存しました。');
                app.closeDialog();
            },
            error: (status, msg) => {
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
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
class CompanyListDialog {
    attach(app, el) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#companyListTemplate',
            data: {
                companyList: app.companies,
            }
        });
        this.ractive.on({
            'windowClicked': () => {
                return false;
            },
            'close': () => {
                app.closeDialog();
                return false;
            },
            'showEdit': (e, item) => {
                console.log('clickEvent');
                this.showEditDialog(app, item);
                return false;
            },
            'deleteCompany': (e, index) => {
                this.deleteCompany(app, index);
                return false;
            },
            'submit': () => {
                this.save(app);
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 370);
    }
    showEditDialog(app, item) {
        app.showDialog(new AddCompanyDialog(item, (result) => {
            // nop
        }));
    }
    save(app) {
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
            success: (id) => {
                company.id = id;
                app.companyMap[id] = company;
                this.ractive.unshift('companyList', company);
                app.addSnack('保存しました。');
                this.clearForm(app);
            },
            error: (status, msg) => {
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
    }
    clearForm(app) {
        this.ractive.set('name', '');
        this.ractive.set('unit', '');
        this.ractive.set('assignee', '');
        this.ractive.set('zip', '');
        this.ractive.set('address', '');
        this.ractive.set('tel', '');
        this.ractive.set('fax', '');
    }
    deleteCompany(app, index) {
        if (!window.confirm('この会社情報を削除しますか？')) {
            return;
        }
        var company = this.ractive.get('companyList')[index];
        app.client.deleteCompany(company.id, {
            success: () => {
                this.ractive.splice('companyList', index, 1);
                app.addSnack('削除しました！');
            },
            error: (status, msg) => {
                console.log('Failed to delete company status=' + status);
            }
        });
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
class SettingsDialog {
    attach(app, el) {
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
                company_bank_name: app.environment.company_bank_name,
            }
        });
        this.ractive.on({
            'windowClicked': () => {
                return false;
            },
            'close': () => {
                app.closeDialog();
                return false;
            },
            'save': () => {
                this.save(app);
                return false;
            }
        });
    }
    save(app) {
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
            success: () => {
                app.environment = env;
                app.addSnack('設定を保存しました！');
                app.closeDialog();
            },
            error: (status, msg) => {
                console.log('Failed to save environment statu=' + status);
            }
        });
    }
}
///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
class SignInPage {
    onCreate(app) {
        if (navigator.credentials !== undefined) {
            navigator.credentials.get({
                password: true,
            }).then((c) => {
                fetch('/api/v1/token', {
                    method: 'POST',
                    credentials: c,
                }).then((resp) => {
                    if (resp.ok) {
                        return resp.json();
                    }
                    else {
                        return Promise.reject('');
                    }
                }).then((json) => {
                    app.client.accessToken = json.access_token;
                    app.client.refreshToken = json.refresh_token;
                    app.client.is_admin = json.is_admin;
                    localStorage.setItem('refreshToken', json.refresh_token);
                    app.router.navigate('top', { trigger: true });
                }).catch((e) => {
                });
                //this.signIn(app, c.id, c.password);
            }).catch((e) => {
            });
        }
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#signInTemplate',
            // データを設定。テンプレートで使います。
            data: {
                myCompanyName: app.myCompanyName,
                inProgress: false,
            }
        });
        app.ractive.on({
            'signIn': (e, item) => {
                var username = r.get('username');
                var password = r.get('password');
                this.signIn(app, username, password);
            },
        });
    }
    signIn(app, username, password) {
        app.ractive.set('inProgress', true);
        app.ractive.update();
        app.client.login(username, password, {
            success: (token) => {
                localStorage.setItem('refreshToken', token);
                if (navigator.credentials === undefined) {
                    app.router.navigate('top', { trigger: true });
                    return;
                }
                navigator.credentials.store(new PasswordCredential({
                    id: username,
                    password: password,
                })).then((c) => {
                    app.router.navigate('top', { trigger: true });
                });
            },
            error: (status, msg) => {
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
    }
}
///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
///<reference path="./Functions.ts"/>
class TopPage {
    constructor() {
        // sort
        this.modifiedSorter = (l, r) => {
            return r.modified_time - l.modified_time;
        };
    }
    onCreate(app) {
        this.app = app;
        app.loadData({
            done: () => {
                this.show(app);
            },
            error: () => {
                // nop
            }
        });
    }
    show(app) {
        var sheets = app.getTradings();
        sheets.sort(this.modifiedSorter);
        var total = 0;
        _.each(sheets, (item) => {
            total += item.total;
        });
        var fabDecorate = (node) => {
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
                teardown: () => { }
            };
        };
        // Racriveオブジェクトを作る
        app.ractive = this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#topTemplate',
            decorators: {
                fab: fabDecorate,
            },
            // データを設定。テンプレートで使います。
            data: {
                myCompanyName: app.myCompanyName,
                is_admin: app.client.isAdmin(),
                'company': app.companyMap,
                'sheets': sheets,
                'toDateStr': Utils.toDateStr,
                total: total,
                sortIndex: 1,
                sortDesc: true,
                showSortMark: (index, sortIndex, desc) => {
                    if (index != sortIndex) {
                        return '';
                    }
                    if (desc) {
                        return '▽';
                    }
                    else {
                        return '△';
                    }
                },
            }
        });
        tooltipster();
        app.ractive.on({
            'addSheet': (e, item) => {
                // #sheetに遷移する
                app.router.navigate('sheets/new', { trigger: true });
            },
            'showSheet': (e, item) => {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, { trigger: true });
            },
            'deleteSheet': (e, index) => {
                this.deleteSheet(app, index);
                return false;
            },
            'copySheet': (e, item) => {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id + '/copy', { trigger: true });
                return false;
            },
            'printQuotation': (e, item) => {
                window.location.href = "/php/quotation.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
                return false;
            },
            'printBill': (e, item) => {
                window.location.href = "/php/bill.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
                return false;
            },
            'showUserList': () => {
                app.showDialog(new UserListDialog());
            },
            'showCompanyList': () => {
                app.showDialog(new CompanyListDialog());
            },
            'showSetting': (e) => {
                app.showDialog(new SettingsDialog());
            },
            sortBy: (e, index) => {
                var list = app.ractive.get('sheets');
                var currentIndex = app.ractive.get('sortIndex');
                var desc = app.ractive.get('sortDesc');
                if (currentIndex == index) {
                    desc = !desc;
                }
                else {
                    currentIndex = index;
                    app.ractive.set('sortIndex', index);
                    desc = true;
                }
                app.ractive.set('sortDesc', desc);
                var sortFunc;
                switch (index) {
                    case 1:
                        sortFunc = this.numberSorter('modified_time');
                        break;
                    case 2:
                        sortFunc = this.companySorter();
                        break;
                    case 3:
                        sortFunc = this.numberSorter('total');
                        break;
                    case 4:
                        sortFunc = this.stringSorter('quotation_number');
                        break;
                    case 5:
                        sortFunc = this.stringSorter('bill_number');
                        break;
                }
                if (desc) {
                    list.sort((l, r) => {
                        return -1 * sortFunc(l, r);
                    });
                }
                else {
                    list.sort(sortFunc);
                }
            },
        });
    }
    deleteSheet(app, index) {
        if (!window.confirm('このシートを削除しますか？')) {
            return;
        }
        var item = app.ractive.get('sheets')[index];
        app.client.deleteTrading(item.id, {
            success: () => {
                app.ractive.splice('sheets', index, 1);
                app.addSnack('削除しました！');
            },
            error: (status, msg) => {
                console.log('Failed to delete item status=' + status);
            }
        });
    }
    numberSorter(key) {
        return (l, r) => {
            return l[key] - r[key];
        };
    }
    stringSorter(key) {
        return (l, r) => {
            return l[key].localeCompare(r[key]);
        };
    }
    companySorter() {
        var company = this.app.companyMap;
        return (l, r) => {
            return company[l.company_id].name.localeCompare(company[r.company_id].name);
        };
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>
class AddUserDialog {
    constructor(user, callback) {
        if (user == null) {
            this.isNew = true;
            this.user = new User();
            this.userOrg = null;
        }
        else {
            this.isNew = false;
            this.user = Utils.clone(user);
            this.userOrg = user;
        }
        this.callback = callback;
    }
    attach(app, el) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#addUserTemplate',
            data: {
                isNew: this.isNew,
                user: this.user,
            }
        });
        this.ractive.on({
            'windowClicked': () => {
                return false;
            },
            'close': () => {
                app.closeDialog();
                return false;
            },
            'save': () => {
                this.save(app);
                return false;
            }
        });
    }
    save(app) {
        var user = this.ractive.get('user');
        var password = this.ractive.get('password');
        if (this.isNew) {
            app.client.createUser(user.login_name, user.display_name, user.tel, password, {
                success: (item) => {
                    app.addUser(item);
                    this.callback(item);
                    app.addSnack('作成しました');
                    app.closeDialog();
                },
                error: (status, msg) => {
                    switch (status) {
                        case 1000:
                            app.addSnack('ユーザーIDを入力してください');
                            break;
                        case 1001:
                            app.addSnack('担当者名を入力してください');
                            break;
                        case 1002:
                            app.addSnack('TELを入力してください');
                            break;
                        case 1003:
                        case 1004:
                            app.addSnack('パスワードを6文字以上入力してください');
                            break;
                        default: app.addSnack('保存に失敗しました。');
                    }
                }
            });
        }
        else {
            app.client.saveUser(user, password, {
                success: (item) => {
                    this.userOrg.login_name = user.login_name;
                    this.userOrg.display_name = user.display_name;
                    this.userOrg.tel = user.tel;
                    this.callback(item);
                    app.addSnack('保存しました');
                    app.closeDialog();
                },
                error: (status, msg) => {
                    switch (status) {
                        case 1001:
                            app.addSnack('ユーザーIDを入力してください');
                            break;
                        case 1002:
                            app.addSnack('担当者名を入力してください');
                            break;
                        case 1003:
                            app.addSnack('TELを入力してください');
                            break;
                        case 1004:
                            app.addSnack('パスワードを6文字以上入力してください');
                            break;
                        default: app.addSnack('保存に失敗しました。');
                    }
                }
            });
        }
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>
class CreateInvoiceDialog {
    constructor() {
    }
    attach(app, el) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#createInvoiceTemplate',
            data: {
                items: [{ "name": "", "num": "" }],
            }
        });
        this.ractive.on({
            'windowClicked': () => {
                return false;
            },
            'close': () => {
                app.closeDialog();
                return false;
            },
            'addItem': () => {
                this.ractive.push('items', { "name": "", "num": "" });
                return false;
            },
            'save': () => {
                this.save(app);
                return false;
            }
        });
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 370);
    }
    save(app) {
        var items = this.ractive.get('items');
        app.client.createInvoice(items, {
            success: (body) => {
                this.downloadBody(body);
            },
            error: (status, msg) => {
                switch (status) {
                    default:
                        app.addSnack('PDF作成に失敗しました');
                        break;
                }
            }
        });
    }
    downloadBody(body) {
        var blob = new Blob([body], { "type": "application/x-download" });
        var url = window.URL || window.webkitURL;
        window.URL = window.URL || window.webkitURL;
        window.location.href = url.createObjectURL(blob);
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./Functions.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
///<reference path="./AddUserDialog.ts"/>
///<reference path="./CreateInvoiceDialog.ts"/>
class SheetPage {
    constructor(id, copyMode) {
        this.id = id;
        this.copyMode = copyMode;
    }
    onCreate(app) {
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
    }
    loadItems(app, trading) {
        app.client.getTradingItems(trading.id, {
            success: (list) => {
                // if copyMode = true remove ids
                if (this.copyMode) {
                    trading.id = null;
                    _.each(list, (item) => {
                        item.id = null;
                    });
                }
                this.show(app, trading, list);
            },
            error: (status, msg) => {
                console.log('Failed to get items status=' + status);
                window.history.back();
            }
        });
    }
    show(app, trading, itemList) {
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
                'users': this.createUserList(app),
                'tradingItems': itemList,
                'deletedItems': [],
            }
        });
        var updateItemSum = (keypath) => {
            var unitPrice = Utils.toNumber(r.get(keypath + 'unit_price'));
            var amount = Utils.toNumber(r.get(keypath + 'amount'));
            r.set(keypath + 'sum', unitPrice * amount);
        };
        var updateSum = () => {
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
        var observeItem = () => {
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
            'close': () => {
                window.history.back();
            },
            'addItem': () => {
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
            'addCompany': () => {
                this.showAddCompanyDialog(app);
            },
            'addUser': () => {
                this.showAddUserDialog(app);
            },
            'printQuotation': () => {
                this.printQuotation(app);
            },
            'printBill': () => {
                this.printBill(app);
            },
            'printDelivery': () => {
                this.printDelivery(app);
            },
            'printInvoide': () => {
                this.showInvoiceDialog(app);
            },
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
        r.on('save', () => {
            this.save(app, (id) => {
                window.history.back();
            });
        });
        r.observe('trading.tax_rate', function (newValue, oldValue, keypath) {
            updateSum();
        });
        // この下にjQuery関連のコードおねがいしやす
        tooltipster();
    }
    showAddCompanyDialog(app) {
        app.showDialog(new AddCompanyDialog(null, (result) => {
            app.ractive.update();
        }));
    }
    showAddUserDialog(app) {
        app.showDialog(new AddUserDialog(null, (result) => {
            app.ractive.set('users', this.createUserList(app));
            app.ractive.update();
        }));
    }
    printQuotation(app) {
        var trading = app.ractive.get('trading');
        var quotationDate = app.ractive.get('quotationDate');
        var doneFunc = (id) => {
            app.ractive.update();
            window.location.href = "/php/quotation.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.quotation_number == null || trading.quotation_number.length == 0) {
            app.client.getNextNumber('quotation', new Date(quotationDate).getTime(), {
                success: (val) => {
                    trading.quotation_number = '' + val + '-I';
                    this.save(app, doneFunc);
                },
                error: (status, msg) => {
                    console.log('Failed to get next quotation number status=' + status);
                }
            });
        }
        else {
            this.save(app, doneFunc);
        }
    }
    printBill(app) {
        var trading = app.ractive.get('trading');
        var billDate = app.ractive.get('billDate');
        var doneFunc = (id) => {
            app.ractive.update();
            window.location.href = "/php/bill.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.bill_number == null || trading.bill_number.length == 0) {
            app.client.getNextNumber('bill', new Date(billDate).getTime(), {
                success: (val) => {
                    trading.bill_number = '' + val + '-V';
                    this.save(app, doneFunc);
                },
                error: (status, msg) => {
                    console.log('Failed to get next bill number status=' + status);
                }
            });
        }
        else {
            this.save(app, doneFunc);
        }
    }
    printDelivery(app) {
        var trading = app.ractive.get('trading');
        var deliveryDate = app.ractive.get('deliveryDate');
        var doneFunc = (id) => {
            app.ractive.update();
            window.location.href = "/php/delivery.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.delivery_number == null || trading.delivery_number.length == 0) {
            app.client.getNextNumber('delivery', new Date(deliveryDate).getTime(), {
                success: (val) => {
                    trading.delivery_number = '' + val + '-D';
                    this.save(app, doneFunc);
                },
                error: (status, msg) => {
                    console.log('Failed to get next delivery number status=' + status);
                }
            });
        }
        else {
            this.save(app, doneFunc);
        }
    }
    showInvoiceDialog(app) {
        app.showDialog(new CreateInvoiceDialog());
    }
    createUserList(app) {
        var list = [];
        var emptyUser = new User();
        emptyUser.id = "empty";
        emptyUser.display_name = "担当者なし";
        list.push(emptyUser);
        _.each(app.users, (item) => {
            list.push(item);
        });
        return list;
    }
    save(app, doneFunc) {
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
            success: (id) => {
                trading.id = id;
                trading.modified_time = new Date().getTime();
                app.tradingsMap[id] = trading;
                var deleted = app.ractive.get('deletedItems');
                this.deleteItems(app, id, deleted, doneFunc);
            },
            error: (status, msg) => {
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
    }
    deleteItems(app, id, list, doneFunc) {
        if (list.length == 0) {
            var list3 = [];
            _.each(app.ractive.get('tradingItems'), (item, index) => {
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
            success: (itemId) => {
                list.splice(0, 1);
                this.deleteItems(app, id, list, doneFunc);
            },
            error: (status, msg) => {
                console.log('Failed to delete items status=' + status);
            }
        });
    }
    saveItems(app, id, list, doneFunc) {
        if (list.length == 0) {
            app.addSnack('保存しました！');
            doneFunc(id);
            return;
        }
        var item = list[0];
        app.client.saveTradingItem(id, item, {
            success: (itemId) => {
                item.id = itemId;
                list.splice(0, 1);
                this.saveItems(app, id, list, doneFunc);
            },
            error: (status, msg) => {
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
    }
}
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
    sheet: (id) => {
        app.page = new SheetPage(id, false);
        app.page.onCreate(app);
    },
    copySheet: (id) => {
        app.page = new SheetPage(id, true);
        app.page.onCreate(app);
    },
    setting: () => {
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
    },
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

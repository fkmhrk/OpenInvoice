/// <reference path="./Client.ts"/>
var MockClient = (function () {
    function MockClient() {
    }
    /**
     * Logs in with username and password.
     */
    MockClient.prototype.login = function (username, password, callback) { };
    /**
      * Gets all users
      */
    MockClient.prototype.getUsers = function (token, callback) { };
    /**
     * Gets all companies
     */
    MockClient.prototype.getCompanies = function (token, callback) { };
    /**
     * Saves company
     * @return item is Company ID
     */
    MockClient.prototype.saveCompany = function (token, item, callback) { };
    /**
     * Gets Tradings
     */
    MockClient.prototype.getTradings = function (token, callback) {
        callback.success(sheetList);
    };
    /**
     * Gets trading items of specified Trading
     */
    MockClient.prototype.getTradingItems = function (token, tradingId, callback) {
    };
    /**
     * Saves Trading
     * @return item is trading ID
     */
    MockClient.prototype.saveTrading = function (token, item, callback) { };
    /**
     * Saves Trading item of specified Trading
     * @return item is item ID
     */
    MockClient.prototype.saveTradingItem = function (token, tradingId, item, callback) { };
    /**
     * Deltes Trading item of specified Trading
     */
    MockClient.prototype.deleteTradingItem = function (token, tradingId, itemId, callback) { };
    return MockClient;
})();
function createClient() {
    return new MockClient();
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
var userList = [];
userList.push({
    'id': 'user1',
    'display_name': 'ユーザー1'
});
userList.push({
    'id': 'user2',
    'display_name': 'ユーザー2'
});
var companyList = [];
var company = new Company();
company.id = 'company1';
company.name = '株式会社AAA';
company.unit = '生産革新部';
company.zip = '111-2222';
company.address = '東京都新宿区';
company.phone = '090-1111-2222';
company.assignee = '担当者';
companyList.push(company);
company = new Company();
company.id = 'company2';
company.name = '株式会社BBB';
company.unit = '';
company.zip = '111-2222';
company.address = '東京都渋谷区';
company.phone = '090-3333-4444';
company.assignee = '';
companyList.push(company);
company = new Company();
company.id = 'company3';
company.name = '株式会社CCC';
company.unit = '';
company.zip = '111-2222';
company.address = '東京都中野区';
company.phone = '090-3333-4444';
company.assignee = '';
companyList.push(company);
company = new Company();
company.id = 'company4';
company.name = 'DDD株式会社';
company.unit = '';
company.zip = '111-2222';
company.address = '広島県';
company.phone = '090-3333-4444';
company.assignee = '';
companyList.push(company);
var sheetList = [];
for (var i = 0; i < 10; ++i) {
    sheetList.push({
        'id': 'idA' + i,
        'date': '1432542408000',
        'company_id': 'company1',
        'company_name': '株式会社ABC',
        'title_type': 0,
        'subject': '画面作成',
        'work_from': 1432542408000,
        'work_to': 1432542408000,
        'quotation_date': 1432542408000,
        'bill_date': 1432542408000,
        'tax_rate': 8,
        'assignee': 'user1',
        'product': '成果物A',
        'total': 650000,
        'modified_time': 1432542408000,
        'quotation_number': 'E0124',
        'bill_number': ''
    });
    sheetList.push({
        'id': 'idB' + i,
        'date': '1432542408000',
        'company_id': 'company2',
        'company_name': '株式会社ZZZ',
        'title_type': 1,
        'subject': '【コンサルツールモック】デザイン画面作成',
        'work_from': 1431505608,
        'work_to': 1431505608,
        'quotation_date': 1431505608,
        'bill_date': 1431505608,
        'tax_rate': 10,
        'assignee': 'user2',
        'product': '成果物B',
        'total': 1030875,
        'modified_time': 1431505608,
        'quotation_number': 'E0123',
        'bill_number': 'V0238'
    });
}
var tradingItemList = [];
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
        this.loadTradings(callback);
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
        this.companies = companyList;
        this.companyMap = {};
        _.each(this.companies, function (item) {
            _this.companyMap[item.id] = item;
        });
        callback.done();
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
            }
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight - 370);
    };
    CompanyListDialog.prototype.showEditDialog = function (app, item) {
        app.showDialog(new AddCompanyDialog());
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
                'sheets': app.tradings
            }
        });
        tooltipster();
        app.ractive.on({
            'showSheet': function (e, item) {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, { trigger: true });
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
        var _this = this;
        var item;
        if (app.tradingsMap === undefined || (item = app.tradingsMap[this.id]) === null) {
            window.history.back();
            return;
        }
        this.trading = Utils.clone(item);
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
                'trading': this.trading,
                'workFrom': Utils.toDateStr(this.trading.work_from),
                'workTo': Utils.toDateStr(this.trading.work_to),
                'companies': companyList,
                'users': userList,
                'tradingItems': tradingItemList
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
                    'unit_price': 0,
                    'amount': 0,
                    'sum': 0
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
            r.splice('tradingItems', index, 1);
            itemObserver = observeItem();
        });
        r.on('save', function () {
            window.history.back();
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
    return SheetPage;
})();
///<reference path="./ractive.d.ts"/>
///<reference path="./data.ts"/>
///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
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
        "": "top",
        // index.html#sheetの場合は、sheetという関数を実行する
        "sheets(/:id)": "sheet",
        "setting": "setting"
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

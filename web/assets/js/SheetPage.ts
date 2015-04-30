///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./Functions.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
///<reference path="./AddUserDialog.ts"/>

class SheetPage implements Page {
    id : string;
    trading : Trading;
    
    constructor(id : string) {
        this.id = id;
    }
    onCreate(app : App) {
        var item : Trading;
        if (app.tradingsMap === undefined || (item = app.tradingsMap[this.id]) === null) {
            window.history.back();
            return;
        }
        this.trading = Utils.clone(item);
        var es = (node : any) => {
            $(node).easySelectBox({speed: 200});
            return {
                teardown : function() {
                    // nop?
                }
            }
        };
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#sheetTemplate',
            decorators: {
                easyselect: es,
            },
            data : {
                'trading' : this.trading,
                'workFrom' : Utils.toDateStr(this.trading.work_from),
                'workTo' : Utils.toDateStr(this.trading.work_to),
                'companies' : companyList,
                'users' : userList,
                'tradingItems' : tradingItemList
            }
        });
        var updateItemSum = (keypath : string) => {
            var unitPrice = Utils.toNumber(r.get(keypath + 'unit_price'));
            var amount = Utils.toNumber(r.get(keypath + 'amount'));
            r.set(keypath + 'sum', unitPrice * amount);
        };
        var updateSum = () => {
            var itemList = r.get('tradingItems');
            var sum = 0;
            var tax = 0;
            var taxRate = Number(r.get('trading.tax_rate'));
            for (var i = 0 ; i < itemList.length ; ++i) {
                var item = itemList[i];
                var taxType = Number($('#tax_type' + i).val());
                if (taxType == 1) {
                    sum += item.sum;
                    tax += item.sum * taxRate / 100;             
                } else if (taxType == 2) {
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
        }
        var observeItem = () => {
            return r.observe({
                'tradingItems.*.unit_price' :  function(newValue, oldValue, keypath) {
                    updateItemSum(keypath.replace('unit_price', ''));
                },
                'tradingItems.*.amount' : function(newValue, oldValue, keypath) {
                    updateItemSum(keypath.replace('amount', ''));
                },
                'tradingItems.*.sum' : function(newValue, oldValue, keypath) {
                    updateSum();
                }
            });
        }
        var itemObserver = observeItem();
        
        r.on({
            'addItem' : () => {
                r.push('tradingItems', {
                    'unit_price' : 0,
                    'amount' : 0,
                    'sum' : 0
                });
            },
            'addCompany' : () => {
                this.showAddCompanyDialog(app);
            },
            'addUser' : () => {
                this.showAddUserDialog(app);
            }
        });
        r.on('deleteItem', function(e, index) {
            itemObserver.cancel();
            r.splice('tradingItems', index, 1);
            itemObserver = observeItem();
        });
        r.on('save', function() {
            window.history.back();
        });
        r.observe('trading.tax_rate', function(newValue, oldValue, keypath) {
            updateSum();
        });        
        // この下にjQuery関連のコードおねがいしやす
        tooltipster();
        //selectbox();
        //sheet();
    }
    private showAddCompanyDialog(app : App) {
        app.showDialog(new AddCompanyDialog());
    }
    private showAddUserDialog(app : App) {
        app.showDialog(new AddUserDialog());
    }
}
///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./Functions.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
///<reference path="./AddUserDialog.ts"/>

class SheetPage implements Page {
    id : string;
    copyMode : boolean;
    
    constructor(id : string, copyMode : boolean) {
        this.id = id;
        this.copyMode = copyMode;
    }
    onCreate(app : App) {
        var item : Trading;
        if (app.tradingsMap === undefined || (item = app.tradingsMap[this.id]) === null) {
            window.history.back();
            return;
        }
        this.loadItems(app, Utils.clone(item));
    }

    private loadItems(app : App, trading : Trading) {
        app.client.getTradingItems(app.accessToken, trading.id, {
            success : (list : Array<TradingItem>) => {
                // if copyMode = true remove ids
                if (this.copyMode) {
                    trading.id = null;
                    _.each(list, (item : TradingItem) => {
                        item.id = null;
                    });
                }
                this.show(app, trading, list);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to get items status=' + status);
                window.history.back();
            }
        });
    }

    private show(app : App, trading : Trading, itemList : Array<TradingItem>) {
        // Racriveオブジェクトを作る
        var r = app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#sheetTemplate',
            decorators: { },
            data : {
                myCompanyName : app.myCompanyName,
                'trading' : trading,
                'workFrom' : Utils.toDateStr(trading.work_from),
                'workTo' : Utils.toDateStr(trading.work_to),
                'quotationDate' : Utils.toDateStr(trading.quotation_date),
                'billDate' : Utils.toDateStr(trading.bill_date),
                'companies' : app.companies,
                'users' : app.users,
                'tradingItems' : itemList,
                'deletedItems' : [],
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
                    id : null,
                    subject : '',
                    unit_price : 0,
                    amount : 0,
                    degree : '',
                    memo : '',
                    tax_type : 1,
                    sum : 0
                });
            },
            'addCompany' : () => {
                this.showAddCompanyDialog(app);
            },
            'addUser' : () => {
                this.showAddUserDialog(app);
            },
            'printQuotation' : () => {
                this.printQuotation(app);
            },
            'printBill' : () => {
                this.printBill(app);
            },            
        });
        r.on('deleteItem', function(e, index) {
            itemObserver.cancel();
            var item = r.get('tradingItems')[index];
            r.splice('tradingItems', index, 1);
            if (item.id != null) {
                r.push('deletedItems', item);
            }
            itemObserver = observeItem();
        });
        r.on('save', () => {
            this.save(app, (id : string) => {
                window.history.back();
            });
        });
        r.observe('trading.tax_rate', function(newValue, oldValue, keypath) {
            updateSum();
        });        
        // この下にjQuery関連のコードおねがいしやす
        tooltipster();
    }
    private showAddCompanyDialog(app : App) {
        app.showDialog(new AddCompanyDialog());
    }
    private showAddUserDialog(app : App) {
        app.showDialog(new AddUserDialog());
    }
    private printQuotation(app : App) {
        this.save(app, (id : string) => {
            window.location.href = "/php/quotation.php?access_token=" + app.accessToken + "&trading_id=" + id;
        });
    }
    private printBill(app : App) {
        this.save(app, (id : string) => {
            window.location.href = "/php/bill.php?access_token=" + app.accessToken + "&trading_id=" + id;
        });        
    }
    private save(app : App, doneFunc : (id : string) => void) {
        var trading = app.ractive.get('trading');
        var workFrom = app.ractive.get('workFrom');
        var workTo = app.ractive.get('workTo');
        var quotationDate = app.ractive.get('quotationDate');
        var billDate = app.ractive.get('billDate');
        var tradingItems = app.ractive.get('tradingItems');

        // modify type
        trading.title_type = Number(trading.title_type);
        trading.work_from = new Date(workFrom).getTime();
        trading.work_to = new Date(workTo).getTime();
        trading.quotation_date = new Date(quotationDate).getTime();
        trading.bill_date = new Date(billDate).getTime();
        trading.tax_rate = Number(trading.tax_rate);
        console.log(trading);
        app.client.saveTrading(app.accessToken, trading, {
            success : (id : string) => {
                app.tradingsMap[id] = trading;
                var deleted = app.ractive.get('deletedItems');
                this.deleteItems(app, id, deleted, doneFunc);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to save trading status=' + status);
            }
        });
    }

    private deleteItems(app : App, id : string, list : Array<TradingItem>,
                        doneFunc : (id : string) => void) {
        if (list.length == 0) {
            var list3 = [];
            _.each(app.ractive.get('tradingItems'), (item : TradingItem, index : number) => {
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
        app.client.deleteTradingItem(app.accessToken, id, item.id, {
            success : (itemId : string) => {
                list.splice(0, 1);
                this.deleteItems(app, id, list, doneFunc);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to delete items status=' + status);
            }
        });
    }

    private saveItems(app : App, id : string, list : Array<TradingItem>,
                      doneFunc : (id : string) => void) {
        if (list.length == 0) {
            doneFunc(id);
            return;
        }
        var item = list[0]
        app.client.saveTradingItem(app.accessToken, id, item, {
            success : (itemId : string) => {
                item.id = itemId;
                list.splice(0, 1);
                this.saveItems(app, id, list, doneFunc);
            },
            error : (status : number, msg : string) => {
                console.log('Failed to save items status=' + status);
            }
        });
    }
}
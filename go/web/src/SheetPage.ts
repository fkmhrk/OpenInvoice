///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./Functions.ts"/>
///<reference path="./AddCompanyDialog.ts"/>
///<reference path="./AddUserDialog.ts"/>
///<reference path="./CreateInvoiceDialog.ts"/>

/*
class SheetPage implements Page {
    id : string;
    copyMode : boolean;
    
    constructor(id : string, copyMode : boolean) {
        this.id = id;
        this.copyMode = copyMode;
    }
    onCreate(app : App) {
        var item : Trading;
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

    private loadItems(app : App, trading : Trading) {
        app.client.getTradingItems(trading.id, {
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
                is_admin : app.client.isAdmin(),
                'trading' : trading,
                'workFrom' : Utils.toDateStr(trading.work_from),
                'workTo' : Utils.toDateStr(trading.work_to),
                'quotationDate' : Utils.toDateStr(trading.quotation_date),
                'billDate' : Utils.toDateStr(trading.bill_date),
                'deliveryDate' : Utils.toDateStr(trading.delivery_date),
                'companies' : app.companies,
                'users' : this.createUserList(app),
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
            'close' : () => {
                window.history.back();
            },
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
            'printDelivery' : () => {
                this.printDelivery(app);
            },
            'printInvoide' : () => {
                this.showInvoiceDialog(app);
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
        app.showDialog(new AddCompanyDialog(null, (result : Company) => {
            app.ractive.update();
        }));
    }
    private showAddUserDialog(app : App) {
        app.showDialog(new AddUserDialog(null, (result : User) => {
            app.ractive.set('users', this.createUserList(app));
            app.ractive.update();
        }));
    }
    private printQuotation(app : App) {
        var trading = app.ractive.get('trading');
        var quotationDate = app.ractive.get('quotationDate');

        var doneFunc = (id : string) => {
            app.ractive.update();
            window.location.href = "/php/quotation.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.quotation_number == null || trading.quotation_number.length == 0) {
            app.client.getNextNumber('quotation', new Date(quotationDate).getTime(), {
                success : (val : number) => {
                    trading.quotation_number = '' + val + '-I';
                    this.save(app, doneFunc);
                },
                error : (status : number, msg : string) => {
                    console.log('Failed to get next quotation number status=' + status);
                }
            });
        } else {
            this.save(app, doneFunc);
        }
    }
    private printBill(app : App) {
        var trading = app.ractive.get('trading');
        var billDate = app.ractive.get('billDate');
        
        var doneFunc = (id : string) => {
            app.ractive.update();
            window.location.href = "/php/bill.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.bill_number == null || trading.bill_number.length == 0) {
            app.client.getNextNumber('bill', new Date(billDate).getTime(), {
                success : (val : number) => {
                    trading.bill_number = '' + val + '-V';
                    this.save(app, doneFunc);
                },
                error : (status : number, msg : string) => {
                    console.log('Failed to get next bill number status=' + status);
                }
            });
        } else {
            this.save(app, doneFunc);
        }        
    }
    
    private printDelivery(app : App) {
        var trading = app.ractive.get('trading');
        var deliveryDate = app.ractive.get('deliveryDate');
        
        var doneFunc = (id : string) => {
            app.ractive.update();
            window.location.href = "/php/delivery.php?access_token=" + app.client.getAccessToken() + "&trading_id=" + id;
        };
        if (trading.delivery_number == null || trading.delivery_number.length == 0) {
            app.client.getNextNumber('delivery', new Date(deliveryDate).getTime(), {
                success : (val : number) => {
                    trading.delivery_number = '' + val + '-D';
                    this.save(app, doneFunc);
                },
                error : (status : number, msg : string) => {
                    console.log('Failed to get next delivery number status=' + status);
                }
            });
        } else {
            this.save(app, doneFunc);
        }
    }

    private showInvoiceDialog(app : App) {
        app.showDialog(new CreateInvoiceDialog());
    }

    private createUserList(app : App) {
        var list = [];
        var emptyUser = new User();
        emptyUser.id = "empty";
        emptyUser.display_name = "担当者なし";
        list.push(emptyUser);
        _.each(app.users, (item : User) => {
            list.push(item);
        });
        return list;
    }
    
    private save(app : App, doneFunc : (id : string) => void) {
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
            success : (id : string) => {
                trading.id = id;
                trading.modified_time = new Date().getTime();
                app.tradingsMap[id] = trading;
                var deleted = app.ractive.get('deletedItems');
                this.deleteItems(app, id, deleted, doneFunc);
            },
            error : (status : number, msg : string) => {
                switch (status) {
                case 1001: app.addSnack('件名を入力してください。'); break;
                case 1002: app.addSnack('作業終了日は作業開始日より後にしてください。'); break;
                case 1003: app.addSnack('消費税率は0以上にしてください。'); break;
                default: app.addSnack('保存に失敗しました。');
                }
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
        app.client.deleteTradingItem(id, item.id, {
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
            app.addSnack('保存しました！');
            doneFunc(id);
            return;
        }
        var item = list[0]
        app.client.saveTradingItem(id, item, {
            success : (itemId : string) => {
                item.id = itemId;
                list.splice(0, 1);
                this.saveItems(app, id, list, doneFunc);
            },
            error : (status : number, msg : string) => {
                switch (status) {
                case 1002: app.addSnack('項目名を入力してください。'); break;
                case 1003: app.addSnack('税区分が不正です。'); break;
                default : app.addSnack('保存に失敗しました。');
                }
                console.log('Failed to save items status=' + status);
            }
        });
    }
}
*/

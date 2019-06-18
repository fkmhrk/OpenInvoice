/// <reference path="./Page.ts"/>

class EditTradingPage implements Page {
    id : string;
    app : Application;

    constructor(id : string) {
        this.id = id;
    }
    
    onCreate(app : Application) {
        this.app = app;

        var es = (node : any) => {
            $(node).easySelectBox({speed: 200});
            return {
                teardown : () => {
                    // nop?
                }
            }
        };
        var toDateStr = (date : any) => {
            var m = date.getMonth() + 1;
            var d = date.getDate();
            if (m < 10) { m = "0" + m; }
            if (d < 10) { d = "0" + d; }
            return date.getFullYear() + "-" + m + "-" + d;
        };
        var item = app.tradingMap[this.id];
        var workFrom = toDateStr(new Date(item.work_from));
        var workTo = toDateStr(new Date(item.work_to));
        var quotationDate = toDateStr(new Date(item.quotation_date));
        var billDate = toDateStr(new Date(item.bill_date));
        var r = app.ractive = new Ractive({
            el : '#container',
            template : '#editTradingTemplate',
            decorators: {
                easyselect: es
            },
            data : {
                trading : item,
                users : app.users,
                companies : app.companies,
                workFrom : workFrom,
                workTo : workTo,
                quotationDate : quotationDate,
                billDate : billDate,
                deleteList : [],
                numToCurrency : (val : any) => {
                    return util.numToCurrency(val);
                }
            }
        });
        r.on({
            numFocus : (e : any, val : any) => {
                e.node.value = util.currencyToNum(val);
                r.updateModel();
            },
            'sumBlur' :  (e : any, val : any, index : any) => {
                e.node.value = util.numToCurrency(val);
                r.updateModel();
                var item = e.context;
                item.sum = util.currencyToNum(item.unit_price) * item.amount;
                r.update();
            },
            amountBlur :  (e : any) => {
                var item = e.context;
                item.sum = util.currencyToNum(item.unit_price) * item.amount;
                r.update();
            },
            deleteItem : (e :any, index : any) => {
                if (!confirm('この項目を削除しますか？')) {
                    return;
                }
                var tradings : Array<any> = r.get('tradingItems');
                var trading = tradings[index];
                if (trading.id != null) {
                    var list : Array<string> = r.get('deleteList');
                    list.push(trading.id);
                    r.set('deleteList', list);
                }
                tradings.splice(index, 1);
                r.set('tradingItems', tradings);
                r.update(); 
            },
            addItem : (e : any) => {
                var list = r.get('tradingItems');
                list.push({
                    id : null,
                    subject : "",
                    unit_price : 0,
                    amount : 0,
                    degree : "人月",
                    tax_type : 1,
                    memo : "",
                    sum : 0,
                });
                r.set('tradingItems', list);
                r.update();
            },
            save : () => {
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
                for (var i = 0 ; i < items.length ; ++i) {
                    var item = items[i];
                    item.unit_price = util.currencyToNum(item.unit_price);
                    item.amount = parseInt(item.amount)
                    item.tax_type = parseInt(item.tax_type)
                    
                    list.push(item);
                }
                var deleteList : Array<string> = r.get('deleteList');
                
                this.save(trading, list, deleteList);
            }
        });
        this.loadTrading();
    }

    loadTrading() {
        if (this.id == 'new') {
            app.tradingItems = [];
            this.show();
            return;
        }
        this.app.client.getTradingItems(this.app.token, this.id, {
            success : (list : Array<Invoice.TradingItem>) => {
                this.app.tradingItems = _.map(list, (item) => {
                    item.unit_price = util.numToCurrency(item.unit_price);
                    return item;
                });
                this.show();
            },
            error : (msg : string) => {
                console.log('error ' + msg);
            }
        });
    }
    show() {
        this.app.ractive.set('tradingItems', this.app.tradingItems);
        this.app.ractive.update();
    }
    save(trading : any, items : Array<any>, deleteList : Array<string>) {
        this.deleteItems(trading, items, deleteList);
        this.app.client.saveTrading(this.app.token, trading, {
            success : (id : string) => {
                console.log('ok');
                this.saveItems(id, items);
            },
            error : (msg : string) => {
                console.log('failed to save ' + msg);
            }
        });
    }
    deleteItems(trading : any, items : Array<any>, deleteList : Array<string>) {
        if (deleteList.length == 0) {
            this.saveTrading(trading, items);
            return;
        }
        this.app.client.deleteTradingItem(this.app.token, trading.id, deleteList[0], {
            success : (id : string) => {
                deleteList.shift();
                this.deleteItems(trading, items, deleteList);
            },
            error : (msg : string) => {
                console.log('failed to delete ' + msg);
            }            
        });
    }
    saveTrading(trading : any, items : Array<any>) {
    }
    saveItems(tradingId : string, items : Array<any>) {
        if (items.length == 0) {
            window.history.back();
            return;
        }
        this.app.client.saveTradingItem(this.app.token, tradingId, items[0], {
            success : (id : string) => {
                console.log('ok');
                items.shift();
                this.saveItems(tradingId, items);
            },
            error : (msg : string) => {
                console.log('failed to save ' + msg);
            }            
        });
    }    
}
///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
///<reference path="./Functions.ts"/>

class TopPage implements Page {
    app : App;
    ractive : Ractive;
    
    onCreate(app : App) {
        this.app = app;
        app.loadData({
            done : () => {
                this.show(app);
            },
            error : () => {
                // nop
            }
        });
    }
    private show(app : App) {
        var sheets = app.getTradings();
        sheets.sort(this.modifiedSorter);
        var total = 0;
        _.each(sheets, (item : Trading) => {
            total += item.total;
        });
        var fabDecorate = (node : any) => {
	    $(node).hover(function(){
		$(this).find(".menu").toggleClass("current");
		$(this).find(".submenu").toggleClass("current");
		$(this).next("span").fadeIn();
	    }, function(){
		$(this).find(".menu").toggleClass("current");
		$(this).find(".submenu").toggleClass("current");
		$(this).next("span").fadeOut();
	    });
            return {
                teardown : () => {}
            };
        };
        // Racriveオブジェクトを作る
        app.ractive = this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#topTemplate',
            decorators : {
                fab : fabDecorate,
            },
            // データを設定。テンプレートで使います。
            data : {
                myCompanyName : app.myCompanyName,
                is_admin : app.client.isAdmin(),
                'company' : app.companyMap,
                'sheets' : sheets,
                'toDateStr' : Utils.toDateStr,
                total : total,
                sortIndex : 1,
                sortDesc : true,
                showSortMark : (index : number, sortIndex : number, desc : boolean) => {
                    if (index != sortIndex) { return ''; }
                    if (desc) {
                       return '▽';
                    } else {
                       return '△';
                    }
                },
            }
        });

        tooltipster();

        app.ractive.on({
            'addSheet' : (e : any, item : Trading) => {
                // #sheetに遷移する
                app.router.navigate('sheets/new', {trigger:true});
            },
            'showSheet' : (e : any, item : Trading) => {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, {trigger:true});
            },
            'deleteSheet' : (e : any, index : number) => {
                this.deleteSheet(app, index);
                return false;
            },
            'copySheet' : (e : any, item : Trading) => {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id + '/copy', {trigger:true});
                return false;
            },            
            'printQuotation' : (e : any, item : Trading) => {
                window.location.href = "/php/quotation.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
                return false;
            },
            'printBill' : (e : any, item : Trading) => {
                window.location.href = "/php/bill.php?access_token=" + app.accessToken + "&trading_id=" + item.id;
                return false;
            },            
            'showUserList' : () => {
                app.showDialog(new UserListDialog());                
            },            
            'showCompanyList' : () => {
                app.showDialog(new CompanyListDialog());                
            },
            'showSetting' : (e : any) => {
                app.showDialog(new SettingsDialog());
            },            
            sortBy : (e : any, index : number) => {
                var list = app.ractive.get('sheets');                
                var currentIndex = app.ractive.get('sortIndex');
                var desc = app.ractive.get('sortDesc');                
                if (currentIndex == index) {                   
                    desc = !desc;
                } else {
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
                    list.sort((l : Trading, r : Trading) => {
                            return -1 * sortFunc(l, r);
                    });
                } else {
                    list.sort(sortFunc);
                }
            },
        });        
    }

    private deleteSheet(app : App, index : number) {
        if (!window.confirm('このシートを削除しますか？')) {
            return;
        }
        var item : Trading = app.ractive.get('sheets')[index];
        app.client.deleteTrading(item.id, {
            success : () => {
                app.ractive.splice('sheets', index, 1);
                app.addSnack('削除しました！');
            },
            error : (status : number, msg : string) => {
                console.log('Failed to delete item status=' + status);
            }
        });
    }

    // sort
    modifiedSorter = (l : Trading, r : Trading) => {
            return r.modified_time - l.modified_time;
    };

    private numberSorter(key : string) : (l : Trading, r : Trading) => number {
        return (l : Trading, r : Trading) => {
            return l[key] - r[key];
        };
    }

    private stringSorter(key : string) : (l : Trading, r : Trading) => number {
        return (l : Trading, r : Trading) => {
            return l[key].localeCompare(r[key]);
        };
    }

    private companySorter() : (l : Trading, r : Trading) => number {
        var company = this.app.companyMap;
        return (l : Trading, r : Trading) => {
            return company[l.company_id].name.localeCompare(
                company[r.company_id].name);
        };
    }
}

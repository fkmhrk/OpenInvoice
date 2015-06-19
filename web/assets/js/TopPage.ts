///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>
///<reference path="./Functions.ts"/>

class TopPage implements Page {
    onCreate(app : App) {
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
        app.ractive = new Ractive({
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
                'company' : app.companyMap,
                'sheets' : sheets,
                'toDateStr' : Utils.toDateStr,
                total : total,
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
}
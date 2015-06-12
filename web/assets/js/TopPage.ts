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
        // Racriveオブジェクトを作る
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#topTemplate',
            // データを設定。テンプレートで使います。
            data : {
                'company' : app.companyMap,
                'sheets' : app.getTradings(),
                'toDateStr' : Utils.toDateStr,
            }
        });

        tooltipster();

        app.ractive.on({
            'showSheet' : (e : any, item : Trading) => {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, {trigger:true});
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
}
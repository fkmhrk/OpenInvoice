///<reference path="./app.ts"/>
///<reference path="./Page.ts"/>
///<reference path="./UserListDialog.ts"/>
///<reference path="./CompanyListDialog.ts"/>
///<reference path="./SettingsDialog.ts"/>

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
                'sheets' : app.tradings
            }
        });

        tooltipster();

        app.ractive.on({
            'showSheet' : (e : any, item : Trading) => {
                // #sheetに遷移する
                app.router.navigate('sheets/' + item.id, {trigger:true});
            },
            'showUserList' : () => {
                app.showDialog(new UserListDialog());                
            },            
            'showCompanyList' : () => {
                app.showDialog(new CompanyListDialog());                
            },
            'showSetting' : (e : any) => {
                // #settingに遷移する
                app.showDialog(new SettingsDialog());
            },            
        });        
    }
}
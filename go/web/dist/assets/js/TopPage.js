///<refence path="./app.ts"/>
///<refence path="./Page.ts"/>
var TopPage = (function () {
    function TopPage() {
    }
    TopPage.prototype.onCreate = function (app) {
        // Racriveオブジェクトを作る
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: '#container',
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: '#topTemplate',
            // データを設定。テンプレートで使います。
            data: {
                'sheets': sheetList
            }
        });
        tooltipster();
        app.ractive.on('showSheet', function (e, item) {
            // #sheetに遷移する
            app.trading = item;
            app.router.navigate('sheet', { trigger: true });
        });
        app.ractive.on('showSetting', function (e) {
            // #settingに遷移する
            app.showDialog({
                attach: function (el) {
                    app.ractive = new Ractive({
                        // どの箱に入れるかをIDで指定
                        el: el,
                        // 指定した箱に、どのHTMLを入れるかをIDで指定
                        template: '#settingTemplate'
                    });
                }
            });
            //app.router.navigate('setting', {trigger:true});
        });
    };
    return TopPage;
})();

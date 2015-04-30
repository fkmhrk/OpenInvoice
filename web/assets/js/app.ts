///<reference path="./ractive.d.ts"/>
///<reference path="./data.ts"/>
///<reference path="./Application.ts"/>
///<reference path="./Page.ts"/>

///<reference path="./TopPage.ts"/>
///<reference path="./SheetPage.ts"/>
var $;
var _;
var Backbone;

var app : App = new App();

var AppRouter = Backbone.Router.extend({
    routes : {
        // ここに、ページ毎に呼ぶ関数名を記述していく
        // index.htmlを開いた直後は、topという関数を実行する        
        "" : "top",
        // index.html#sheetの場合は、sheetという関数を実行する
        "sheets(/:id)" : "sheet",
        "setting" : "setting"
        //"signup" : "signUp",
    },
    top : function() {
        app.page = new TopPage();
        app.page.onCreate(app);
    },
    sheet : (id : string) => {
        app.page = new SheetPage(id);
        app.page.onCreate(app);
    },
    setting : () => {
        // ダイアログ用の要素を作る
        var dialog = document.createElement('section');
        document.querySelector('#dialogs').appendChild(dialog);
        // Racriveオブジェクトを作る
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : dialog,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#settingTemplate'
            // データを設定。テンプレートで使います。
            /*data : {
                'sheets' : sheetList
            }*/
        });        
    }, 
});

$(function() {
    app.start();
    // Backboneのおまじない
    app.router = new AppRouter();
    Backbone.history.start();
});

// [common] for plugins
function tooltipster(){
    $('.actionBtn li a').tooltipster({
        theme: 'tooltipster-actionBtn'
    });
    $('.btn, .delete').tooltipster({
        theme: 'tooltipster-btn',
        //arrow: false,
        offsetY: -3
    });
}
function selectbox(){
    //select box customize
    //$('select').easySelectBox({speed: 200});
}



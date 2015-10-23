///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

class CreateInvoiceDialog implements Dialog {
    ractive : Ractive;

    constructor() {
    }
    
    attach(app : App, el : HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#createInvoiceTemplate',
            data : {
                items : [{"name":"", "num":""}],
            }
        });
        this.ractive.on({
            'windowClicked' : () => {
                return false;
            },
            'close' : () => {
                app.closeDialog();
                return false;
            },
            'addItem' : () => {
                this.ractive.push('items', {"name":"", "num":""});
                return false;
            },
            'save' : () => {
                this.save(app);
                return false;
            }
        });
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight-370);     
    }

    private save(app : App) {
        var items = this.ractive.get('items');
        app.client.createInvoice(items, {
            success : (body : string) => {
                this.downloadBody(body);
            },
            error : (status : number, msg : string) => {
                switch (status) {
                default: app.addSnack('PDF作成に失敗しました'); break;
                }
            }
        });
    }

    private downloadBody(body : string) {
        var blob = new Blob([body], { "type" : "application/x-download" });
        var url = (<any>window).URL || (<any>window).webkitURL;
        (<any>window).URL = (<any>window).URL || (<any>window).webkitURL;
        window.location.href = url.createObjectURL(blob);
    }
}
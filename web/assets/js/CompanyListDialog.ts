///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./AddCompanyDialog.ts"/>

class CompanyListDialog implements Dialog {
    ractive : Ractive;
    callback : (result : any) => void;
    
    attach(app : App, el : HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#companyListTemplate',
            data : {
                companyList : app.companies,
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
            'showEdit' : (e : any, item : Company) => {
                console.log('clickEvent');
                this.showEditDialog(app, item);
                return false;
            }
        });   
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight-370);     
    }
    private showEditDialog(app : App, item : Company) {
        app.showDialog(new AddCompanyDialog());
    }
}
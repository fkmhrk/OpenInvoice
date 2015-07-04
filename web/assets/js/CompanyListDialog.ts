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
            },
            'deleteCompany' : (e : any, index : number) => {
                this.deleteCompany(app, index);
                return false;
            },
            'submit' : () => {
                this.save(app);
            }
        });   
        //dialog内だけスクロールするように調整
        var listUserHeight = $('.listTemplate').height();
        $('.listTemplate .list').css('height', listUserHeight-370);     
    }
    private showEditDialog(app : App, item : Company) {
        app.showDialog(new AddCompanyDialog());
    }
    
    private save(app : App) {
        var name = this.ractive.get('name');
        var unit = this.ractive.get('unit');
        var assignee = this.ractive.get('assignee');
        var zip = this.ractive.get('zip');
        var address = this.ractive.get('address');
        var tel = this.ractive.get('tel');
        var fax = this.ractive.get('fax');

        var company = new Company();
        company.id = null;
        company.name = name;
        company.zip = zip;
        company.address = address;
        company.phone = tel;
        company.fax = fax;
        company.unit = unit;
        company.assignee = assignee;

        app.client.saveCompany(company, {
            success : (id : string) => {
                company.id = id;
                app.companyMap[id] = company;
                this.ractive.unshift('companyList', company);
                app.addSnack('保存しました。');
                this.clearForm(app);
            },
            error : (status : number, msg : string) => {
                switch (status) {
                case 1001: app.addSnack('会社名を入力してください。'); break;
                default: app.addSnack('保存に失敗しました。');
                }
                console.log('Failed to create company status=' + status);
            }
        });
        console.log(company);
    }

    private clearForm(app : App) {
        this.ractive.set('name', '');
        this.ractive.set('unit', '');
        this.ractive.set('assignee', '');
        this.ractive.set('zip', '');
        this.ractive.set('address', '');
        this.ractive.set('tel', '');
        this.ractive.set('fax', '');
    }

    private deleteCompany(app : App, index : number) {
        if (!window.confirm('この会社情報を削除しますか？')) {
            return;
        }
        var company = this.ractive.get('companyList')[index];
        app.client.deleteCompany(company.id, {
            success : () => {
                this.ractive.splice('companyList', index, 1);
                app.addSnack('削除しました！');
            },
            error : (status : number, msg : string) => {
                console.log('Failed to delete company status=' + status);
            }
        });
    }
}
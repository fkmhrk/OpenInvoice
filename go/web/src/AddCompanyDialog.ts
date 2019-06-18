///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

/*
class AddCompanyDialog implements Dialog {
    ractive : Ractive;
    company : Company;
    companyOrg : Company;
    isNew : boolean;
    callback : (result : Company) => void;

    constructor(company : Company, callback : (result : Company) => void) {
        if (company == null) {
            this.isNew = true;
            this.company = new Company();
            this.company.id = null;
            this.companyOrg = null;
        } else {
            this.isNew = false;
            this.company = Utils.clone(company);
            this.companyOrg = company;
        }
        this.callback = callback;
    }
    
    attach(app : App, el : HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#addCompanyTemplate',
            data : {
                isNew : this.isNew,
                company : this.company,
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
            'save' : () => {
                this.save(app);
                return false;
            }
        });        
    }

    private save(app : App) {
        var company = this.ractive.get('company');
        console.log(company);
        app.client.saveCompany(company, {
            success : (id : string) => {
                // clone?
                company.id = id;
                if (this.companyOrg == null) {
                    app.addCompany(company);
                } else {
                    this.companyOrg.name = company.name;
                    this.companyOrg.unit = company.unit;
                    this.companyOrg.assignee = company.assignee;
                    this.companyOrg.zip = company.zip;
                    this.companyOrg.address = company.address;
                    this.companyOrg.phone = company.phone;
                    this.companyOrg.fax = company.fax;
                }
                this.callback(company);
                app.addSnack('保存しました。');
                app.closeDialog();
            },
            error : (status : number, msg : string) => {
                switch (status) {
                case 1001: app.addSnack('会社名を入力してください'); break;                    
                default: app.addSnack('保存に失敗しました。'); break;
                }
            }
        });
    }
}
*/

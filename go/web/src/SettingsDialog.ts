///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

/*
class SettingsDialog implements Dialog {
    ractive : Ractive;
    
    attach(app : App, el : HTMLElement) {
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#settingTemplate',
            decorators: { },
            data : {
                tax_rate : app.environment.tax_rate,
                quotation_limit : app.environment.quotation_limit,
                closing_month : app.environment.closing_month,
                pay_limit : app.environment.pay_limit,
                company_name : app.environment.company_name,
                company_zip : app.environment.company_zip,
                company_address : app.environment.company_address,
                company_tel : app.environment.company_tel,
                company_fax : app.environment.company_fax,
                company_bankname : app.environment.company_bankname,
                company_bank_type : app.environment.company_bank_type,
                company_bank_num : app.environment.company_bank_num,
                company_bank_name : app.environment.company_bank_name,
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
    save(app : App) {
        var env = new Environment();
        env.tax_rate = this.ractive.get('tax_rate');
        env.quotation_limit = this.ractive.get('quotation_limit');
        env.closing_month = this.ractive.get('closing_month');
        env.pay_limit = $('#pay_limit').val();
        env.company_name = this.ractive.get('company_name');
        env.company_zip = this.ractive.get('company_zip');
        env.company_address = this.ractive.get('company_address');
        env.company_tel = this.ractive.get('company_tel');
        env.company_fax = this.ractive.get('company_fax');
        env.company_bankname = this.ractive.get('company_bankname');
        env.company_bank_type = $('#bank_type').val();
        env.company_bank_num = this.ractive.get('company_bank_num');
        env.company_bank_name = this.ractive.get('company_bank_name');

        app.client.saveEnvironment(env, {
            success : () => {
                app.environment = env;
                app.addSnack('設定を保存しました！');
                app.closeDialog();
            },
            error : (status : number, msg : string) => {
                console.log('Failed to save environment statu=' + status);
            }
        });
    }
}

*/

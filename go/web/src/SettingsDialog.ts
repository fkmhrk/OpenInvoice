///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

import { Ractive } from "./ractive";
import { handleError } from "./pages/ErrorHandler";

export class SettingsDialog implements IDialog {
    dialogId: number = 0;
    private app: IApplication;
    private ractive!: Ractive;

    constructor(app: IApplication) {
        this.app = app;
    }

    async onCreate(elem: HTMLElement) {
        const environment = await this.app.models.environment.get();
        this.ractive = new Ractive({
            el: elem,
            template: "#settingTemplate",
            decorators: {},
            data: {
                tax_rate: environment.tax_rate,
                quotation_limit: environment.quotation_limit,
                closing_month: environment.closing_month,
                pay_limit: environment.pay_limit,
                company_name: environment.company_name,
                company_zip: environment.company_zip,
                company_address: environment.company_address,
                company_tel: environment.company_tel,
                company_fax: environment.company_fax,
                company_bankname: environment.company_bankname,
                company_bank_type: environment.company_bank_type,
                company_bank_num: environment.company_bank_num,
                company_bank_name: environment.company_bank_name,
            },
            on: {
                windowClicked: () => false,
                close: () => {
                    this.app.closeDialog(this);
                    return false;
                },
                save: () => this.save(),
            },
        });
    }

    private async save() {
        const env = <IEnvironment>{
            tax_rate: this.ractive.get("tax_rate"),
            quotation_limit: this.ractive.get("quotation_limit"),
            closing_month: this.ractive.get("closing_month"),
            pay_limit: this.ractive.get("pay_limit"),
            company_name: this.ractive.get("company_name"),
            company_zip: this.ractive.get("company_zip"),
            company_address: this.ractive.get("company_address"),
            company_tel: this.ractive.get("company_tel"),
            company_fax: this.ractive.get("company_fax"),
            company_bankname: this.ractive.get("company_bankname"),
            company_bank_type: this.ractive.get("company_bank_type"),
            company_bank_num: this.ractive.get("company_bank_num"),
            company_bank_name: this.ractive.get("company_bank_name"),
        };

        try {
            await this.app.models.environment.save(env);
            this.app.addSnack("設定を保存しました！");
            this.app.closeDialog(this);
        } catch (e) {
            handleError(this.app, e, "保存に失敗しました");
        }
    }
}

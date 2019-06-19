///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

import { Ractive } from "./ractive";

export class AddCompanyDialog implements IDialog {
    dialogId: number = 0;
    private app: IApplication;
    private ractive!: Ractive;
    private company: Company;
    private callback: (result: Company) => void;

    constructor(
        app: IApplication,
        company: Company,
        callback: (result: Company) => void
    ) {
        this.app = app;
        this.company = company;
        this.callback = callback;
    }

    onCreate(elem: HTMLElement): void {
        this.ractive = new Ractive({
            el: elem,
            template: "#addCompanyTemplate",
            data: {
                id: this.company.id,
                name: this.company.name,
                unit: this.company.unit,
                assignee: this.company.assignee,
                zip: this.company.zip,
                address: this.company.address,
                phone: this.company.phone,
                fax: this.company.fax,
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
        const company = {
            id: this.ractive.get("id"),
            name: this.ractive.get("name"),
            unit: this.ractive.get("unit"),
            assignee: this.ractive.get("assignee"),
            zip: this.ractive.get("zip"),
            address: this.ractive.get("address"),
            phone: this.ractive.get("phone"),
            fax: this.ractive.get("fax"),
        };

        const saved = await this.app.models.company.save(company);
        this.callback(saved);
        // app.addSnack("保存しました。");
        this.app.closeDialog(this);
        /*        
        app.client.saveCompany(company, {
            success: (id: string) => {
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
                app.addSnack("保存しました。");
                app.closeDialog();
            },
            error: (status: number, msg: string) => {
                switch (status) {
                    case 1001:
                        app.addSnack("会社名を入力してください");
                        break;
                    default:
                        app.addSnack("保存に失敗しました。");
                        break;
                }
            },
        });
*/
    }
}

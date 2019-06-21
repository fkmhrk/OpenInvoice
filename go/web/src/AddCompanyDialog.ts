///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./Functions.ts"/>

import { Ractive } from "./ractive";
import { handleError } from "./pages/ErrorHandler";

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

        try {
            const saved = await this.app.models.company.save(company);
            this.callback(saved);
            this.app.addSnack("保存しました。");
            this.app.closeDialog(this);
        } catch (e) {
            handleError(this.app, e, "保存に失敗しました");
        }
    }
}

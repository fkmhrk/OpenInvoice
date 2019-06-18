///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>
///<reference path="./AddCompanyDialog.ts"/>

import { Ractive } from "./ractive";
import { AddCompanyDialog } from "./AddCompanyDialog";

export class CompanyListDialog implements IDialog {
    dialogId: number = 0;
    private app: IApplication;
    private ractive!: Ractive;
    // callback: (result: any) => void;

    constructor(app: IApplication) {
        this.app = app;
    }

    async onCreate(elem: HTMLElement) {
        const companyList = await this.app.models.company.getAll();
        this.ractive = new Ractive({
            el: elem,
            template: "#companyListTemplate",
            data: {
                companyList: companyList,
            },
            on: {
                windowClicked: () => false,
                close: () => {
                    this.app.closeDialog(this);
                    return false;
                },
                showEdit: (e: any, item: Company) => {
                    this.showEditDialog(item);
                    return false;
                },
                deleteCompany: (e: any, index: number) => {
                    this.deleteCompany(index);
                    return false;
                },
                submit: () => this.save(),
            },
        });
        //dialog内だけスクロールするように調整
        var listUserHeight = $(".listTemplate").height()!;
        $(".listTemplate .list").css("height", listUserHeight - 370);
    }
    private showEditDialog(item: ICompany) {
        this.app.showDialog(
            new AddCompanyDialog(this.app, item, (result: ICompany) => {
                // update
                item.name = result.name;
                item.address = result.address;
                item.assignee = result.assignee;
                item.fax = result.fax;
                item.phone = result.phone;
                item.unit = result.unit;
                item.zip = result.zip;
                this.ractive.update();
            })
        );
    }

    private async save() {
        const name = this.ractive.get("name");
        const unit = this.ractive.get("unit");
        const assignee = this.ractive.get("assignee");
        const zip = this.ractive.get("zip");
        const address = this.ractive.get("address");
        const tel = this.ractive.get("tel");
        const fax = this.ractive.get("fax");

        const company = <ICompany>{
            id: "",
            name: name,
            zip: zip,
            address: address,
            phone: tel,
            fax: fax,
            unit: unit,
            assignee: assignee,
        };

        const saved = await this.app.models.company.save(company);
        this.ractive.unshift("companyList", saved);
        this.clearForm();

        /*        
        app.client.saveCompany(company, {
            success: (id: string) => {
                company.id = id;
                app.companyMap[id] = company;
                this.ractive.unshift("companyList", company);
                app.addSnack("保存しました。");
                this.clearForm();
            },
            error: (status: number, msg: string) => {
                switch (status) {
                    case 1001:
                        app.addSnack("会社名を入力してください。");
                        break;
                    default:
                        app.addSnack("保存に失敗しました。");
                }
                console.log("Failed to create company status=" + status);
            },
        });
        console.log(company);
*/
    }

    private clearForm() {
        this.ractive.set({
            name: "",
            unit: "",
            assignee: "",
            zip: "",
            address: "",
            tel: "",
            fax: "",
        });
    }

    private async deleteCompany(index: number) {
        if (!window.confirm("この会社情報を削除しますか？")) {
            return;
        }
        const company = this.ractive.get("companyList")[index] as ICompany;

        await this.app.models.company.deleteCompany(company);
        this.ractive.splice("companyList", index, 1);
        // app.addSnack("削除しました！");
        /*
        app.client.deleteCompany(company.id, {
            success: () => {
                this.ractive.splice("companyList", index, 1);
                app.addSnack("削除しました！");
            },
            error: (status: number, msg: string) => {
                console.log("Failed to delete company status=" + status);
            },
        });
*/
    }
}

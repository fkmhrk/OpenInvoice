import Ractive from "../ractive";

const toDateStr = (time: number) => {
    var date = new Date(time);
    var m: any = date.getMonth() + 1;
    var d: any = date.getDate();
    if (m < 10) {
        m = "0" + m;
    }
    if (d < 10) {
        d = "0" + d;
    }
    return date.getFullYear() + "-" + m + "-" + d;
};

const emptyUser = {
    id: "empty",
    login_name: "",
    display_name: "担当者なし",
    tel: "",
};

export class TradingPage implements IPage {
    private app: IApplication;
    private ractive!: Ractive;
    private id: string;
    private isCopy: boolean;

    constructor(app: IApplication, id: string, isCopy: boolean) {
        this.app = app;
        this.id = id;
        this.isCopy = isCopy;
    }

    async onCreate() {
        const trading = await this.loadTrading();
        const vals = await Promise.all([
            this.app.models.user.getAll(),
            this.app.models.company.getAll(),
        ]);
        const users = vals[0].map((u: IUser) => u); // copy
        const companies = vals[1];

        users.unshift(emptyUser);
        this.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el: "#container",
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: "#sheetTemplate",
            decorators: {},
            data: {
                // myCompanyName: app.myCompanyName,
                // is_admin: app.client.isAdmin(),
                trading: trading,
                workFrom: toDateStr(trading.work_from),
                workTo: toDateStr(trading.work_to),
                quotationDate: toDateStr(trading.quotation_date),
                billDate: toDateStr(trading.bill_date),
                deliveryDate: toDateStr(trading.delivery_date),
                calcItemSum: (item: TradingItem) => this.calcItemSum(item),

                companies: companies,
                users: users,
                deletedItems: [],
            },
            on: {
                close: () => window.history.back(),
                addItem: () => this.addItem(),
                deleteItem: (e: any, index: number) => this.deleteItem(index),
                save: async () => {
                    await this.save();
                    window.history.back();
                },
                printQuotation: () => this.printQuotation(),
                printBill: () => this.printBill(),
                printDelivery: () => this.printDelivery(),
                printInvoide: () => this.showInvoiceDialog(),
            },
        });
        // why here?: calcItemSum() uses this.ractive.
        this.ractive.set("tradingItems", trading.items);
    }

    private calcItemSum(item: TradingItem) {
        item.sum = item.unit_price * item.amount;
        this.updateSum();
        return item.sum;
    }

    private updateSum() {
        const itemList: TradingItem[] = this.ractive.get("tradingItems");
        const taxRate: number = this.ractive.get("trading.tax_rate");
        let sum = 0;
        let tax = 0;
        itemList.forEach((item: TradingItem) => {
            const taxType = item.tax_type;
            if (taxType == 0) {
                sum += item.sum;
            } else if (taxType == 1) {
                sum += item.sum;
                tax += (item.sum * taxRate) / 100;
            } else if (taxType == 2) {
                const body = (item.sum * 100) / (100 + taxRate);
                const taxTmp = Math.ceil(item.sum - body);
                sum += item.sum - taxTmp;
                tax += taxTmp;
            }
        });
        this.ractive.set({
            "trading.sum": sum,
            "trading.tax": tax,
            "trading.total": sum + tax,
        });
    }

    private addItem(): void {
        this.ractive.push("tradingItems", {
            id: "",
            subject: "",
            unit_price: 0,
            amount: 0,
            degree: "",
            memo: "",
            tax_type: 1,
            sum: 0,
        });
    }

    private deleteItem(index: number): void {
        var item = this.ractive.get("tradingItems")[index] as ITradingItem;
        this.ractive.splice("tradingItems", index, 1);
        if (item.id.length > 0) {
            this.ractive.push("deletedItems", item);
        }
    }

    private async save(): Promise<ITrading> {
        const trading = this.ractive.get("trading");
        const workFrom = this.ractive.get("workFrom");
        const workTo = this.ractive.get("workTo");
        const quotationDate = this.ractive.get("quotationDate");
        const billDate = this.ractive.get("billDate");
        const deliveryDate = this.ractive.get("deliveryDate");
        const tradingItems = this.ractive.get("tradingItems");

        trading.items = tradingItems;

        // modify type
        trading.work_from = new Date(workFrom).getTime();
        trading.work_to = new Date(workTo).getTime();
        trading.quotation_date = new Date(quotationDate).getTime();
        trading.bill_date = new Date(billDate).getTime();
        trading.delivery_date = new Date(deliveryDate).getTime();
        console.log(trading);

        const saved = await this.app.models.trading.save(trading);
        this.ractive.set("tradingItems", saved.items);

        const deletedItems = this.ractive.get("deletedItems");
        await this.app.models.trading.deleteItems(this.id, deletedItems);
        this.ractive.set("deletedItems", []);

        return trading;

        /*        
        app.client.saveTrading(trading, {
            success: (id: string) => {
                trading.id = id;
                trading.modified_time = new Date().getTime();
                app.tradingsMap[id] = trading;
                var deleted = app.ractive.get("deletedItems");
                this.deleteItems(app, id, deleted, doneFunc);
            },
            error: (status: number, msg: string) => {
                switch (status) {
                    case 1001:
                        app.addSnack("件名を入力してください。");
                        break;
                    case 1002:
                        app.addSnack(
                            "作業終了日は作業開始日より後にしてください。"
                        );
                        break;
                    case 1003:
                        app.addSnack("消費税率は0以上にしてください。");
                        break;
                    default:
                        app.addSnack("保存に失敗しました。");
                }
                console.log("Failed to save trading status=" + status);
            },
        });
*/
    }

    private async loadTrading(): Promise<ITrading> {
        if (this.id == "new") {
            const now = new Date().getTime();
            return <ITrading>{
                id: "",
                company_id: "",
                company_name: "",
                title_type: 0,
                subject: "",
                assignee: emptyUser.id,
                work_from: now,
                work_to: now,
                quotation_date: 0,
                quotation_number: "",
                bill_date: 0,
                bill_number: "",
                delivery_date: 0,
                delivery_number: "",
                date: "",
                total: 0,
                tax_rate: 8,
                product: "",
                memo: "",
                modified_time: 0,
                items: [],
            };
        }
        const trading = await this.app.models.trading.getById(this.id);
        if (this.isCopy) {
            this.id = "";
            trading.id = "";
            trading.items.forEach((item: ITradingItem) => {
                item.id = "";
            });
        }
        return trading;
    }

    private async printQuotation() {
        const trading = this.ractive.get("trading");
        const quotationDate = this.ractive.get("quotationDate");

        if (trading.quotation_number.length == 0) {
            // make next number
            const val = await this.app.models.trading.getNextNumber(
                "quotation",
                quotationDate
            );
            trading.quotation_number = `${val}-I`;
        }

        const saved = await this.save();
        this.ractive.set("trading", saved);
        window.location.href = `/php/quotation.php?access_token=${
            this.app.models.token.token
        }&trading_id=${trading.id}`;
    }

    private async printBill() {
        const trading = this.ractive.get("trading");
        const billDate = this.ractive.get("billDate");

        if (trading.bill_number.length == 0) {
            // make next number
            const val = await this.app.models.trading.getNextNumber(
                "bill",
                billDate
            );
            trading.bill_number = `${val}-V`;
        }

        const saved = await this.save();
        this.ractive.set("trading", saved);
        window.location.href = `/php/bill.php?access_token=${
            this.app.models.token.token
        }&trading_id=${trading.id}`;
    }

    private async printDelivery() {
        const trading = this.ractive.get("trading");
        const deliveryDate = this.ractive.get("deliveryDate");

        if (trading.delivery_number.length == 0) {
            // make next number
            const val = await this.app.models.trading.getNextNumber(
                "delivery",
                deliveryDate
            );
            trading.delivery_number = `${val}-D`;
        }

        const saved = await this.save();
        this.ractive.set("trading", saved);
        window.location.href = `/php/delivery.php?access_token=${
            this.app.models.token.token
        }&trading_id=${trading.id}`;
    }

    private showInvoiceDialog() {
        // app.showDialog(new CreateInvoiceDialog());
    }
}

import Ractive from "../ractive";
import $ from "jquery";

const fabDecorate = (node: any) => {
    $(node).hover(
        function() {
            $(this)
                .find(".menu")
                .toggleClass("current");
            $(this)
                .find(".submenu")
                .toggleClass("current");
            $(this)
                .next("span")
                .fadeIn();
        },
        function() {
            $(this)
                .find(".menu")
                .toggleClass("current");
            $(this)
                .find(".submenu")
                .toggleClass("current");
            $(this)
                .next("span")
                .fadeOut();
        }
    );
    return {
        teardown: () => {},
    };
};

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

export class TopPage implements IPage {
    private app: IApplication;
    private ractive!: Ractive;
    private environment!: Environment;

    constructor(app: IApplication) {
        this.app = app;
    }

    async onCreate() {
        this.environment = await this.app.models.environment.get();
        const tradings = await this.app.models.trading.getAll();

        this.ractive = new Ractive({
            el: "#container",
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template: "#topTemplate",
            decorators: {
                fab: fabDecorate,
            },
            data: {
                sheets: tradings,
                toDateStr: toDateStr,
            },
            on: {
                addSheet: () => this.app.navigate("/tradings/new"),
                showSheet: (e: any, item: Trading) =>
                    this.app.navigate(`/tradings/${item.id}`),
                deleteSheet: (e: any, index: number) => {
                    this.deleteSheet(index);
                    return false;
                },
                copySheet: (e: any, item: Trading) => {
                    this.app.navigate(`/tradings/${item.id}/copy`);
                    return false;
                },
            },
        });
    }

    private async deleteSheet(index: number) {
        if (!window.confirm("このシートを削除しますか？")) {
            return;
        }
        var item: Trading = this.ractive.get("sheets")[index] as ITrading;
        await this.app.models.trading.deleteTrading(item.id);
        this.ractive.splice("sheets", index, 1);
        // app.addSnack("削除しました！");
        alert("削除しました");
    }
}

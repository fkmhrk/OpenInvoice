import Ractive from "../ractive";
import $ from "jquery";
import { SettingsDialog } from "../SettingsDialog";
import { CompanyListDialog } from "../CompanyListDialog";
import { UserListDialog } from "../UserListDialog";

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
                is_admin: this.app.models.token.isAdmin(),
                sheets: tradings,
                toDateStr: toDateStr,
                sortIndex: 1,
                sortDesc: true,
                showSortMark: (
                    index: number,
                    sortIndex: number,
                    desc: boolean
                ) => {
                    if (index != sortIndex) {
                        return "";
                    }
                    if (desc) {
                        return "▽";
                    } else {
                        return "△";
                    }
                },
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
                showUserList: () =>
                    this.app.showDialog(new UserListDialog(this.app)),
                showCompanyList: () =>
                    this.app.showDialog(new CompanyListDialog(this.app)),
                showSetting: () =>
                    this.app.showDialog(new SettingsDialog(this.app)),
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

/// <reference path="./IApplication.ts" />
/// <reference path="./clients/HTTPClient.ts" />

import { getBody, isStatus200 } from "./clients/Functions";
import Ractive from "./ractive";

export default class Application implements IApplication {
    private templateClient: HTTPClient;
    private router: IRouter;
    private snackbars!: Ractive;
    private dialogIds: number[];
    private nextDialogId: number;
    models: IModels;

    constructor(
        templateClient: HTTPClient,
        models: IModels,
        routerFactory: (app: IApplication) => IRouter
    ) {
        this.templateClient = templateClient;
        this.models = models;
        this.router = routerFactory(this);
        this.dialogIds = [];
        this.nextDialogId = 1;
        this.initSnackbar();
    }

    start() {
        this.router.start();
    }

    fetchTemplate(name: string): Promise<string> {
        const url = `/pages/${name}`;
        return this.templateClient
            .send(Method.GET, url, {}, null)
            .then(isStatus200)
            .then(getBody);
    }

    navigate(path: string): void {
        this.router.navigate(path);
    }

    redirect(path: string): void {
        this.router.redirect(path);
    }

    addSnack(item: string): void {
        this.snackbars.push("snackbars", item);
        const closeFunc = () => {
            const list = this.snackbars.get("snackbars");
            if (list.length == 0) {
                return;
            }
            this.snackbars.splice("snackbars", 0, 1);
            if (this.snackbars.get("snackbars").length > 0) {
                setTimeout(closeFunc, 3000);
            }
        };
        setTimeout(closeFunc, 3000);
    }

    showDialog(dialog: IDialog): void {
        dialog.dialogId = this.nextDialogId++;
        this.dialogIds.push(dialog.dialogId);
        // create element
        const root = <HTMLElement>document.querySelector("#dialogs");
        root.style.display = "block";
        const section = document.createElement("section");
        section.classList.add("dialog");
        root.appendChild(section);
        // show!
        dialog.onCreate(section);
    }

    closeDialog(dialog: IDialog): void {
        const index = this.dialogIds.findIndex((id: number) => {
            return id == dialog.dialogId;
        });
        if (index == -1) return;

        this.dialogIds.splice(index, 1);
        const root = <HTMLElement>document.querySelector("#dialogs");
        const dialogElem = root.children[index];
        root.removeChild(dialogElem);
        if (this.dialogIds.length == 0) {
            root.style.display = "none";
        }
    }

    private initSnackbar() {
        // snackbarsの準備
        this.snackbars = new Ractive({
            el: "#snacks",
            template: "#snackbarsTemplate",
            data: {
                snackbars: [],
            },
        });
        this.snackbars.on({
            close: (e: any, index: number) => {
                this.snackbars.splice("snackbars", index, 1);
            },
        });
    }
}

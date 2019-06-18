/// <reference path="./IApplication.ts" />
/// <reference path="./clients/HTTPClient.ts" />

import { getBody, isStatus200 } from "./clients/Functions";

export default class Application implements IApplication {
    private templateClient: HTTPClient;
    private router: IRouter;
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
}

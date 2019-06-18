/// <reference path="./IRouter.ts" />
/// <reference path="./page.d.ts" />
/// <reference path="./IApplication.ts" />
/// <reference path="./pages/IPage.ts" />

export default class Router implements IRouter {
    constructor(app: IApplication) {
        page("/", async () => {
            const m = await import(
                /* webpackChunkName: "SignInPage" */ "./pages/SignInPage"
            );
            this.showPage(new m.SignInPage(app));
            //this.showPage(new TopPage(app));
        });
        page("/top", async () => {
            const m = await import(
                /* webpackChunkName: "TopPage" */ "./pages/TopPage"
            );
            this.showPage(new m.TopPage(app));
        });
        page("/tradings/:id", async (ctx: PageContext) => {
            const m = await import(
                /* webpackChunkName: "TradingPage" */ "./pages/TradingPage"
            );
            this.showPage(new m.TradingPage(app, ctx.params.id, false));
        });
        page("/tradings/:id/copy", async (ctx: PageContext) => {
            const m = await import(
                /* webpackChunkName: "TradingPage" */ "./pages/TradingPage"
            );
            this.showPage(new m.TradingPage(app, ctx.params.id, true));
        });
    }

    start(): void {
        page();
    }

    navigate(path: string): void {
        page(path);
    }

    redirect(path: string): void {
        page.redirect(path);
    }

    private showPage(next: IPage) {
        next.onCreate();
    }
}

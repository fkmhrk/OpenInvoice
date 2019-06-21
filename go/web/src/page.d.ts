declare const page : {
    (): void,
    (path: string): void,
    (path: string, handler: () => void): void,
    (path: string, handler: (ctx: PageContext) => void): void,
    redirect: (path: string) => void,
}

declare interface PageContext {
    params: any;
}
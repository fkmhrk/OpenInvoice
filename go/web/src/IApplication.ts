interface IApplication {
    models: IModels;

    start(): void;
    navigate(path: string): void;

    showDialog(dialog: IDialog): void;

    closeDialog(dialog: IDialog): void;
}

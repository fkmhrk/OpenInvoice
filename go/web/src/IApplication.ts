interface IApplication {
    models: IModels;

    start(): void;
    navigate(path: string): void;
}

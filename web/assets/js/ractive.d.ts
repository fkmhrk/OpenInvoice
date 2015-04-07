declare class Ractive {
    constructor(option : any);
    on(action : string, func : any);
    on(params : any);
    get(name : string) : any;
    set(name : string, value : any);
    update();
}
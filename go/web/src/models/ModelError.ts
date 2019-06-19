export class ModelError {
    code: number;
    msg: string;
    cause: any;
    constructor(code: number, msg: string, cause: any) {
        this.code = code;
        this.msg = msg;
        this.cause = cause;
    }
}

export const ERR_EMPTY_USERNAME = 1;
export const ERR_EMPTY_PASSWORD = 2;
export const ERR_SHORT_PASSWORD = 3;
export const ERR_EMPTY_DISPLAY_NAME = 4;
export const ERR_EMPTY_TEL = 5;

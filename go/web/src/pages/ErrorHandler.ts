import * as me from "../models/ModelError";

export const handleError = (app: IApplication, e: any, fallback: string) => {
    if (e instanceof me.ModelError) {
        switch (e.code) {
            case me.ERR_EMPTY_USERNAME:
                app.addSnack("ユーザー名を入力してください");
                return;
            case me.ERR_EMPTY_PASSWORD:
                app.addSnack("パスワードを入力してください");
                return;
            case me.ERR_SHORT_PASSWORD:
                app.addSnack("パスワードを6文字以上入力してください");
                return;
            case me.ERR_EMPTY_DISPLAY_NAME:
                app.addSnack("担当者名を入力してください");
                return;
            case me.ERR_EMPTY_TEL:
                app.addSnack("電話番号を入力してください");
                return;
        }
    }
    app.addSnack(fallback);
};

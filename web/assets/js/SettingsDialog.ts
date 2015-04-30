///<reference path="./Application.ts"/>
///<reference path="./Dialog.ts"/>

class SettingsDialog implements Dialog {
    attach(app : App, el : HTMLElement) {
        app.ractive = new Ractive({
            // どの箱に入れるかをIDで指定
            el : el,
            // 指定した箱に、どのHTMLを入れるかをIDで指定
            template : '#settingTemplate'
        });
        app.ractive.on({
            'windowClicked' : () => {
                return false;
            },
            'close' : () => {
                app.closeDialog();
                return false;
            }
        });
    }    
}
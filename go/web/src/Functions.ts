module Utils {
    export function toList(obj : any) {
        var list = [];
        for (var k in obj) {
            list.push(obj[k]);
        }
        return list;
    }
    export function toNumber(source : string){
        var num = Number(String(source).replace(",", ""))
        return isNaN(num) ? 0 : num;
    }
    export function toDateStr(time : number){
        var date = new Date(time);
        var m : any = date.getMonth() + 1;
        var d : any = date.getDate();
        if (m < 10) { m = "0" + m; }
        if (d < 10) { d = "0" + d; }
        return date.getFullYear() + "-" + m + "-" + d;
    }

    export function clone(source : any) {
        var dest : any = {};
        for (var k in source) {
            dest[k] = source[k];
        }
        return dest;
    }

    export function isEmpty(s : string) {
        return s == null || s.length == 0;
    }
}
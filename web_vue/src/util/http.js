import crypto from "../util/crypto"

export default {
    csrfToken: process.env.VUE_APP_CSRF_TOKEN,
    FileHost: process.env.VUE_APP_FILE_HOST + '/api',
    initCsrfToken() {
        let req = new XMLHttpRequest();
        req.open('GET', document.location, false);
        req.send(null);
        let headers = req.getAllResponseHeaders().split('\n');
        headers.forEach(function (item) {
            if (item.indexOf('x-csrf-token') !== -1) {
                this.csrfToken = item.substring(item.indexOf(":") + 1);
                return false;
            }
        });
    },
    fetchWithAuthEncrypt() {  /* 需要驗證的api都必須呼叫此方法含加密 */
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = crypto.encryptText(JSON.stringify(args[1]));
        const funcSuccess = args[2];
        this.send(action, parameter, '/auth', funcSuccess);
    },
    fetchWithAuth() {  /* 需要驗證的api都必須呼叫此方法 */
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = JSON.stringify(args[1]);
        const funcSuccess = args[2];
        this.send(action, parameter, '/auth', funcSuccess);
    },
    fetchWithEncrypt() { /* 整個參數json 轉成 string 再進行加密 */
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = crypto.encryptText(JSON.stringify(args[1]));
        const funcSuccess = args[2];
        this.send(action, parameter,'', funcSuccess);
    },
    fetch() {
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = JSON.stringify(args[1]);
        const funcSuccess = args[2];
        this.send(action, parameter,'', funcSuccess);
    },
    async send(action, parameter,authPath, funcSuccess) {
        await fetch(process.env.VUE_APP_API_HOST  +authPath, {
        // await fetch("http://192.168.1.130:9120" + "/admin-api" + authPath, {
            body: JSON.stringify({
                'action': action,
                'parameters': parameter,
            }),
            headers: new Headers({
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json',
            }),
            method: 'POST'
        }).then(response => {
            return response.text();
        }).then(text => {
            // JSON Hijacking while(1);
            let json = JSON.parse(text.replace('while(1);', ''));

            funcSuccess(json);
        });
    },
    async fetchUpload() {
        this.url = '';
        const args = Array.prototype.slice.call(arguments);
        const parameter = args[1];
        const funcSuccess = args[2];
        const funcFail = args[3];
        await this.uploadSend(parameter.file, parameter.folder, funcSuccess, funcFail);
    }


}








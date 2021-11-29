/* eslint-disable no-unused-vars */
import crypto from "../util/crypto"
import auth from "../util/auth"
import http from "../util/http"

export default {
    csrfToken: http.csrfToken,
    downloadWithAuthEncrypt(){  /* 需要驗證的api都必須呼叫此方法 */
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = crypto.encryptText(JSON.stringify(args[1]));
        const fileName = args[2];
        this.send(action, parameter, '/auth', fileName);
    },
    downloadWithAuth(){  /* 需要驗證的api都必須呼叫此方法 */
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = JSON.stringify(args[1]);
        const fileName = args[2];
        this.send(action, parameter, '/auth', fileName);
    },
    downloadWithEncrypt(){ /* 整個參數json 轉成 string 再進行加密 */
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = crypto.encryptText(JSON.stringify(args[1]));
        const fileName = args[2];
        this.send(action, parameter, '', fileName);
    },
    download(){
        const args = Array.prototype.slice.call(arguments);
        const action = args[0][0];
        const parameter = JSON.stringify(args[1]);
        const fileName = args[2];
        this.send(action, parameter, '', fileName);
    },
    async send(action, parameter, authPath, fileName){
        await fetch(process.env.VUE_APP_API_HOST + authPath, {
            body: JSON.stringify({
                'action': action,
                'parameters': parameter,
            }),
            headers: new Headers({
                'Content-Type': 'application/json',
                'Csrf-Token': this.csrfToken,
                'token': authPath === '/auth' ? auth.getToken() : ''
            }),
            method: 'POST'
        }) .then(resp => resp.blob())
        .then(blob => {
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.style.display = 'none';
            a.href = url;
            a.download = fileName;
            document.body.appendChild(a);
            a.click();
            window.URL.revokeObjectURL(url);
        }).catch(error => {

        });
    },
}

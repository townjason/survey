
import http from "../util/http"
import auth from "../util/auth"

export default {
    socket: {},
    instance: function () {
        this.conn('', Array.prototype.slice.call(arguments));
    },
    instanceWithAuth: function () {
        this.conn('/auth', Array.prototype.slice.call(arguments));
    },
    conn: function (authPath, args) {
        const action = args[0][0];
        let url = process.env.VUE_APP_WS_HOST + authPath + '/ws?action=' + action + "&csrf=" + http.csrfToken;

        if(authPath === '/auth') {
            url += '&token=' + auth.getToken();
        }

        this.socket = new WebSocket(url);
        if(args.length === 2) {
            this.socket.onmessage = args[1];
        }

        if(args.length === 3) {
            this.socket.onopen = args[2];
        }

        if(args.length === 4) {
            this.socket.onerror = args[3];
        }

        if(args.length === 5) {
            this.socket.onclose = args[4];
        }
    },
    send: function (data) {
        this.socket.send(data);
    },
    close: function () {
        this.socket.close();
    }
}
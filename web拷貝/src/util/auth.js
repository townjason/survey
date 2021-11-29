/* eslint-disable semi,no-trailing-spaces,indent,quotes,space-infix-ops,comma-dangle,padded-blocks,no-unused-vars,eol-last,semi-spacing */


let admin_token = window.localStorage.getItem('admin_token');
let store_token = window.localStorage.getItem('store_token');
let admin_role_token = window.localStorage.getItem('admin_role_token');
let store_token_list = window.localStorage.getItem('store_token_list');
let number = window.localStorage.getItem('account');

export default {
    isAdminLogin() {
        return admin_token !== null && admin_token !== '';
    },
    setAdminToken(t) {
        window.localStorage.setItem('admin_token', t + '');
        admin_token = t;
    },
    getAdminToken() {
        return admin_token;
    },
    setNumber(t) {
        window.localStorage.setItem('account', t + '');
        number = t;
    },
    getNumber() {
        return number;
    },
    isStoreToken() {
        return store_token !== null && store_token !== '';
    },
    setStoreToken(t) {
        window.localStorage.setItem('store_token', t + '');
        store_token = t;
    },
    getStoreToken() {
        return store_token;
    },
    clearToken() {
        admin_token = '';
        window.localStorage.removeItem('admin_token');
        store_token = '';
        window.localStorage.removeItem('store_token');
    },

    setStoreTokenList(list) {
        window.localStorage.setItem('store_token_list', list);
        store_token_list = list;
    },
    getStoreTokenList() {
        return store_token_list;
    }
}

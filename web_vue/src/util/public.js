import Vue from 'vue'

export default {
    showNotify: function (message, notifyType) {
        let notify = Vue.prototype.$message;
        let type;
        //關閉先前產生的提示視窗
        notify.closeAll();
        //判斷型態帶入不同狀態
        switch (typeof notifyType) {
            case "string": {
                type = notifyType;
                break
            }
            case "boolean": {
                if (notifyType === true) {
                    type = "success";
                    break
                } else {
                    type = "error";
                    break
                }
            }
            default: {
                type = 'error';
                break
            }
        }

        notify({
            type: type,
            message: message,
            duration: 4000,
        });
    }, getPreviewAryUrl(previewAry, index) {
        let url;
        let strIndex;
        if (process.env.VUE_APP_MODE === "release") {
            strIndex = previewAry[index].search("file.oinapp.com");
            url = previewAry[index].slice(strIndex + 16,);
        } else {
            strIndex = previewAry[index].search("file.downloadapp.club");
            url = previewAry[index].slice(strIndex + 22,);
        }

        if (strIndex === -1) {
            return ""
        } else {
            return url
        }
    }


}

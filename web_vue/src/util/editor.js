import Vue from 'vue'
import VEditor from 'yimo-vue-editor'


Vue.use(VEditor, {

    name: 'v-editor-app',//Custom name
    config: {
        uploadImgUrl: process.env.VUE_APP_FILE_HOST,
        uploadParams: {
            folder: "/test",
            apiKey: "a87a5674a70ff9dad49f",
        },
        // uploadHeaders: {
        //     'token': url === '/auth' ? auth.getToken() : ''
        // }
         uploadHeaders: {
            'Access-Control-Allow-Origin': '*',
        }
    },//wagnEditor config
    uploadHandler: (type, resTxt) => {//Upload processing hook
        // console.log(type)
        // console.log(resTxt)
        if (type === 'success') {
            let res = JSON.parse(resTxt)//Do not process the default look at the return value bit image path
            if (res.status == 1) {
                return "https://file.oinapp.com/" + res.data[0]
            }else{
                return null
            }
        } else if (type === 'error') {
            return '图片上传失败__'
        } else if (type === 'timeout') {
            return '图片上传超时'
        }
        return '图片上传失败__'
    }
});

export default {}

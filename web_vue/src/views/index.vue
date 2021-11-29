<template>
    <div class="container-fluid">
        <div class="row top-banner">
            <!--img :src="newPreview"-->
            <img :style="{backgroundImage: 'url(' + newPreview + ')' }">
        </div>
        <div class="privacy" v-if="isShowInput">
            <div class="row text-center">
                <div class="store-info">問卷店別：{{brandName}} {{storeName}}</div>
            </div>
        </div>
        <!-- TODO 以下區塊會員才有-->
        <div class="member" v-if="isApp">
            <div class="text-center">
                <img :src="require('@/assets/images/logo_oin.svg')" style="width: 150px">
            </div>
            <div class="con">
                下載註冊台灣享樂遊APP，填寫表單，下一次只需開啟APP掃描即可，會員填寫現在<span class="red">送成大商圈九折券＋抽獎券</span>。
            </div>
            <div class="text-center">
                <a href="http://onelink.to/q9bnfx" round class="w-50 yellow">下載台灣享樂遊</a>
            </div>

            <div class="line2 mt-4"/>

        </div>
        <!-- TODO 以上區塊會員才有--->

        <div class="privacy" v-if="isShowInput">
            <div class="group">
                <img :src="require('@/assets/images/app_user_friends_gray.png')" v-show="nameIsShow">
                <el-input v-model="name" placeholder="請填寫姓名" id="name" autocomplete="off"
                          v-show="nameIsShow"></el-input>
                <img :src="require('@/assets/images/app_phone_volume_gray.png')" v-show="phoneIsShow">
                <el-input type="tel" v-model="phone" placeholder="請輸入電話" id="phone" autocomplete="off"
                          v-show="phoneIsShow"></el-input>
            </div>
        </div>

        <!-- TODO 以下區塊掃碼才有-->
        <div class="privacy">
            <div class="input-group" v-show="genderIsShow" id="gender">
                性別
                <el-radio v-model="gender" label='1'>男</el-radio>
                <el-radio v-model="gender" label='2'>女</el-radio>
                <el-radio v-model="gender" label='3'>小朋友</el-radio>

            </div>
            <el-input type="number" style="width: 150px" v-model="age" placeholder="請輸入年齡" id="age" min="1"
                      autocomplete="off" maxlength="3" v-show="oldIsShow"></el-input>
        </div>
        <!-- TODO 以上區塊掃碼才有-->
        <div class="mt-4 mb-4 ml-2 mr-2">{{content}}</div>
        <div class="row survey-box">
            <div class="col-12 con" v-for="(item, index) in questionnaireTopic" :key="index" :id="'topic' + item.id">
                <div v-if="item.type === 1">
                    <h5><span class="num">{{index+1}}</span>{{item.title}}</h5>
                    <div class="answer">
                        <el-radio-group v-model="item.answer">
                            <el-radio v-for="item1 in item.questionnaireAnswer"
                                      :key="item1.id"
                                      :label="item1.id"
                                      :value="item1.id">{{item1.title}}
                                <el-input type="text" placeholder="請輸入內容" v-model="item1.inputText"
                                          v-if="item1.isShowInput"/>
                            </el-radio>
                        </el-radio-group>
                    </div>
                </div>
                <div v-if="item.type === 2">
                    <h5><span class="num">{{index+1}}</span>{{item.title}}</h5>
                    <div class="answer">
                        <el-radio-group v-model="item.answer">
                            <el-radio v-for="item1 in item.questionnaireAnswer"
                                      :key="item1.id"
                                      :label="item1.id"
                                      :value="item1.id">{{item1.title}}
                                <el-input type="text" placeholder="請輸入內容" v-model="item1.inputText"
                                          v-if="item1.isShowInput"/>
                            </el-radio>
                        </el-radio-group>
                    </div>
                </div>
                <div v-if="item.type === 3">
                    <h5><span class="num">{{index+1}}</span>{{item.title}}</h5>
                    <div class="answer">
                        <el-rate v-model="item.answer"></el-rate>
                    </div>
                </div>
                <div v-if="item.type === 4">
                    <h5><span class="num">{{index+1}}</span>{{item.title}}</h5>
                    <div class="answer">
                        <el-select v-model="item.answer" placeholder="請選擇"
                                   @change="showInput(...arguments, item, item.questionnaireAnswer, item.answer)">
                            <el-option
                                    v-for="item1 in item.questionnaireAnswer"
                                    :key="item1.id"
                                    :label="item1.title"
                                    :value="item1.id"
                            >
                            </el-option>
                        </el-select>
                        <el-input type="text" placeholder="請輸入內容" v-if="item.isShowInput" v-model="inputText"
                                  @change="giveInputText(...arguments, item.questionnaireAnswer, item.answer)"/>
                    </div>
                </div>
                <div v-if="item.type === 5">
                    <h5><span class="num">{{index+1}}</span>{{item.title}}</h5>
                    <div class="answer">
                        <el-input
                                v-if="item.inputType === 'textarea'"
                                type="textarea"
                                rows="3"
                                autosize
                                placeholder="請輸入內容"
                                v-model="item.answer"/>
                        <el-input
                                v-if="item.inputType === 'number'"
                                type="number"
                                placeholder="請輸入數字"
                                v-model="item.answer"/>
                        <!--                        <input-->
                        <!--                                v-if="item.inputType === 'date'"-->
                        <!--                                type="date"-->
                        <!--                                v-model="item.answer"-->
                        <!--                                value="2018-07-22"/>-->
                        <!--{{item.answer}}-->
                        <datetime v-if="item.inputType === 'date'" v-model="item.answer" type="date"></datetime>
                    </div>
                </div>
                <div v-if="item.type === 6">
                    <h5><span class="num">{{index+1}}</span>{{item.title}}</h5>
                    <div class="answer">
                        <el-checkbox v-for="item1 in item.questionnaireAnswer" :label="item1.id" :key="item1.id"
                                     v-model="item1.isSelect">
                            {{item1.title}}
                            <el-input type="text" placeholder="請輸入內容" v-model="item1.inputText"
                                      v-if="item1.isShowInput"/>
                        </el-checkbox>
                    </div>
                </div>
            </div>
        </div>

        <el-checkbox v-model="checked" id="privacy" class="mt-4">同意使用蒐集 <span class="blue" @click="$refs['modalShow'].show()">個人資料告知事項與同意書</span></el-checkbox>

        <div class="col-12 mt-5 mb-5 text-center">
            <el-button round class="w-50" @click="submit">送出</el-button>
        </div>
        <loading
                :active.sync=isLoading
                :can-cancel="true"
                :is-full-page="true">
        </loading>

         <b-modal ref="modalShow"  hide-header hide-footer>
             <p class="my-4" >
      1、為維持國內疫情之穩定控制，本場所配合政府「COVID19」 防疫新生活運動，採行實聯制措施。依據「個人資料保護法之特定目的 及個人資料之類別」代號 012 公共衛生或傳染病防治之特定目的，蒐集 以上個人資料，且不得為目的外利用。所蒐集之資料僅保存 28 日，屆期銷毀。感謝您的配合。<br/><br/>
      2、個人資料利用之對象及方式：為防堵疫情而有必要時，得提供衛生主管 機關依傳染病防治法等規定進行疫情調查及聯繫使用。<br/><br/>
      3、當事人就其個人資料得依個人資料保護法規定，向本系統（資料蒐集機關）行使權利，包括查詢或請求閱覽、請求製給複製本、請求補充或更正、請求停止處理或利用、請求刪除等。相關權利行使方式：【例 如】於APP意見反映線上填具資料申辦。<br/><br/>
      4、當事人如無法配合本實聯制作業，恕無法進入。<br/><br/>
      5、本實聯制其他相關措施說明，請參閱 <a href="http://at.cdc.tw/8QI4hA">http://at.cdc.tw/8QI4hA</a>。
            </p>
            <table class="w-100 footerButton">
                <tr>
                    <td>
                        <b-button @click="hideModal" class="mt-3" variant="warning" block>確定</b-button>
                    </td>
                </tr>
            </table>
        </b-modal>
    </div>
</template>


<script>
    import crypto from "../util/crypto";
    import VueLoading from 'vue-loading-overlay'
    import 'vue-loading-overlay/dist/vue-loading.css'

    import {Datetime} from 'vue-datetime'
    // You need a specific loader for CSS files
    import 'vue-datetime/dist/vue-datetime.css'

    export default {
        name: "index",
        components: {loading: VueLoading, Datetime},
        data() {
            return {
                isLoading: false,
                nameIsShow: true,
                nameIsRequire: true,
                phoneIsShow: true,
                phoneIsRequire: true,
                genderIsShow: true,
                genderIsRequire: true,
                oldIsShow: true,
                oldIsRequire: true,
                brandName: '',
                storeName: '',
                isApp: true,

                questionId: '',
                name: '',
                phone: '',
                checked: false,
                gender: '',
                age: '',

                id: 0,
                newPreview: '',
                content: '',
                questionnaireTopic: [],
                congratulationText: '',
                isShowInput: true,
                inputText: '',
                writeType: 'WEB'
            }
        },
        created() {
            if(this.$route.params.user !== undefined){
                let userParams = crypto.decryptText(decodeURIComponent(this.$route.params.user));
                let userInfo = userParams.split(';');
                this.name = userInfo[0];
                this.phone = userInfo[1];
                this.writeType = userInfo[2];
                this.isShowInput = false;
                this.isApp = false;
            }

            if (this.$route.params.name !== undefined && this.$route.params.phone !== undefined && this.$route.params.writeType !== undefined) {
                this.checked = true;
                this.isShowInput = false
            }

            this.setupQuestionnaireData();
        },
        methods: {
            getOs(){
                let userAgent = window.navigator.userAgent;
                let platform = window.navigator.platform;
                let iosPlatforms = ['iPhone', 'iPad', 'iPod']
                console.log(userAgent);
                if (iosPlatforms.indexOf(platform) !== -1) {
                    return 'iOS';
                } else if (/Android/.test(userAgent)) {
                    return 'Android';
                } else {
                    return '';
                }
            },
            hideModal() {
                this.$refs['modalShow'].hide()
            },
            setupQuestionnaireData: function () {
                let self = this
                if (self.$route.params.object !== undefined) {
                    self.isLoading = true;

                    self.$http.fetchWithEncrypt`GetQuestionnaireInfo${{
                        "qrCodeDate": crypto.decryptText(decodeURIComponent(self.$route.params.object)),
                        "name": this.name,
                        "phone": this.phone,
                        "writeType": this.writeType,
                    }}
                    ${json => {
                        if (json.status) {
                            self.content = json.data.content
                            self.newPreview = json.data.imagePath
                            self.congratulationText = json.data.congratulationText
                            self.questionnaireTopic = json.data.questionnaireTopic
                            self.questionnaireRecodeId = json.data.questionnaireRecodeId
                            self.buildInItem = JSON.parse(json.data.buildInItem)
                            self.brandName = json.data.brandName
                            self.storeName = json.data.storeName

                            JSON.parse(json.data.buildInItem).forEach(function (buildInItem) {
                                if (buildInItem.type === 'name') {
                                    self.nameIsShow = buildInItem.isShow
                                    self.nameIsRequire = buildInItem.isRequire
                                } else if (buildInItem.type === 'phone') {
                                    self.phoneIsShow = buildInItem.isShow
                                    self.phoneIsRequire = buildInItem.isRequire
                                } else if (buildInItem.type === 'gender') {
                                    self.genderIsShow = buildInItem.isShow
                                    self.genderIsRequire = buildInItem.isRequire
                                } else if (buildInItem.type === 'old') {
                                    self.oldIsShow = buildInItem.isShow
                                    self.oldIsRequire = buildInItem.isRequire
                                }
                            });

                            self.isLoading = false;
                        } else {
                            // self.$public.showNotify(json.message, json.status);
                             self.$router.push({name: 'Error', params: {text: json.message},})
                        }
                    }}`;
                }
            },
            showInput(value, questionnaireTopic, questionnaireAnswerList, answer) {
                questionnaireTopic.isShowInput = false

                questionnaireAnswerList.forEach(function (element) {
                    if (parseInt(answer) === element.id && element.isShowInput) {
                        questionnaireTopic.isShowInput = true
                    }
                });
            },
            giveInputText(value, questionnaireAnswerList, answer) {
                questionnaireAnswerList.forEach(function (element) {
                    if (parseInt(answer) === element.id) {
                        element.inputText = value
                    }
                });
            },
            submit() {
                let self = this;
                let vaild = false;
                let errorTopicId = 0;
                let errorTopicTitle = '';

                if (self.name === '' && self.nameIsRequire) {
                    self.$public.showNotify('請填寫姓名', false);
                    //document.querySelector("#name").scrollIntoView(true);
                    document.getElementById('name').focus();
                    return
                }

                if (self.phone === '' && self.phoneIsRequire) {
                    self.$public.showNotify('請填寫電話', false);
                    //document.querySelector("#phone").scrollIntoView(true);
                    document.getElementById('phone').focus();
                    return
                }

                if ((self.oldIsRequire || self.age !== '') && (self.age <= 0)) {//只要是必填或是有填寫年齡就要判斷
                    self.$public.showNotify('年齡錯誤', false);
                    //document.querySelector("#phone").scrollIntoView(true);
                    document.getElementById('age').focus();
                    return
                }

                if (isNaN(parseInt(self.gender)) && self.genderIsRequire) {
                    self.$public.showNotify('性別錯誤', false);
                    //document.querySelector("#phone").scrollIntoView(true);
                    document.getElementById('gender').focus();
                    return
                }

                if (!self.checked) {
                    self.$public.showNotify('請勾選"同意使用蒐集個人資料告知事項與同意書"', false);
                    document.querySelector("#privacy").scrollIntoView(false);
                    document.getElementById('privacy').focus()
                    return
                }


                self.questionnaireTopic.every(function (element) {
                    if (element.isRequired) {
                        if (element.type !== 6) {
                            if (element.answer === '' || element.answer === 0) {
                                self.$public.showNotify('請填寫' + element.title, false);
                                errorTopicId = element.id
                                errorTopicTitle = element.title
                                vaild = true

                                return false
                            }
                        } else {
                            vaild = true

                            element.questionnaireAnswer.every(function (answer) {
                                if (answer.isSelect) {
                                    vaild = false
                                    return false
                                }
                                return true;
                            });

                            //判斷複選題是否有勾選
                            if (vaild) {
                                self.$public.showNotify('請填寫' + element.title, false);
                                errorTopicId = element.id
                                errorTopicTitle = element.title
                                return false
                            }
                        }
                    }
                    return true;
                });

                if (vaild) {
                    self.$public.showNotify('請填寫' + errorTopicTitle, false);
                    document.querySelector("#topic" + errorTopicId).scrollIntoView(false);
                    document.getElementById('topic' + errorTopicId).classList.add('redLine');
                    setTimeout(function () {
                        document.getElementById('topic' + errorTopicId).classList.remove('redLine');
                    }, 2500)
                    return
                }

                self.questionnaireTopic.forEach(function (element) {
                    if (element.type === 5 && element.inputType === 'date' && element.answer !== '') {
                        element.answer = new Date(element.answer).getFullYear() + "年" + (new Date(element.answer).getMonth() + 1) + "月" + new Date(element.answer).getDay() + "日";
                    } else {
                        //將答案的變數轉成字串
                        element.answer = element.answer.toString();
                    }
                });

                this.isLoading = true;

                self.$http.fetch`InsertQuestionnaireRecord${{
                    "questionnaireRecodeId": self.questionnaireRecodeId,
                    "name": self.name,
                    "phone": self.phone,
                    "gender": parseInt(self.gender),
                    "age": parseInt(self.age),
                    "isOpen": self.checked,
                    "questionnaireTopic": self.questionnaireTopic,
                }}
                ${json => {
                    if (json.status) {
                        // self.$alert(this.congratulationText, {
                        //     confirmButtonText: '確定',
                        // });
                        self.$router.push({name: 'Complete', params: {congratulationText: self.congratulationText},})
                    } else {
                        self.isLoading = false;
                        self.$public.showNotify(json.message, json.status);

                        self.questionnaireTopic.forEach(function (element) {
                            //將答案的變數轉成字串
                            element.answer = '';
                        });
                    }
                }}`;
            }
        }
    }
</script>
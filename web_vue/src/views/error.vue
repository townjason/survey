<template>
    <div class="container-fluid">
        <div class="complete">
            <img src="@/assets/images/thank_you.svg">
            <br/>
            {{text}}
        </div>
    </div>
</template>


<script>
    import preventBack from 'vue-prevent-browser-back';//组件内单独引入

    export default {
        mixins: [preventBack],//注入
        name: "error",
        data() {
            return {
                text: this.$route.params.text
            }
        },
        created() {

        },
        methods: {

        }
    }
</script>
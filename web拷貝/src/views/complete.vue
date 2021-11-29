<template>
    <div class="container-fluid">
        <div class="complete">
            <img src="@/assets/images/thank_you.svg">
            <h5>感謝您的填寫</h5>
            {{congratulationText}}
        </div>
    </div>
</template>


<script>
    import preventBack from 'vue-prevent-browser-back';//组件内单独引入

    export default {
        mixins: [preventBack],//注入
        name: "complete",
        data() {
            return {
                congratulationText: this.$route.params.congratulationText
            }
        },
        created() {

        },
        methods: {

        }
    }
</script>
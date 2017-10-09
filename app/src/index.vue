<template>
    <div class="wrapper">
        <text class="label">Light:<switch @change="toggleLight" class="light-switch"></switch></text>
        <div class="light-div">
            <div class="light-bubble" v-bind:class="{light_on: isLightOn}" ></div>
            <div class="light-bottom"></div>
        </div>
    </div>
</template>

<style>
    .inline-div {
        display: inline;
    }
    .light-div {
        margin-top: 30px;
    }
    .light-bottom {
        width: 30px;
        height:30px;
        border-radius: 5px;
        background-color: darkcyan;
        margin-left: 35px;
    }
    .light-bubble {
        width: 100px;
        height: 100px;
        border-radius: 50%;
    }
    .light_on {
        background-color: #f9ff0f;
    }
    .light-switch {
    }
</style>

<script>
    const serverURL = "http://10.2.0.205:1234";
    const stream = weex.requireModule('stream');
    export default {
        data: {
            isLightOn: false,
        },
        methods: {
            toggleLight: function () {
                stream.fetch({
                    method: 'GET',
                    url: serverURL + '/light',
                }, function (res) {
                    console.log(res);
                })
            },
        }
    }
</script>
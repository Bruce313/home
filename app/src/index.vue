<template>
    <div class="wrapper" @click="update">
        <image :src="logoUrl" class="logo"></image>
        <text class="title">Hello {{target}}</text>
        <text class="desc">Now, let's use vue to build your weex app.</text>
        <button v-on:click="updateTime">更新时间</button>
        <text class="label">夜灯</text>
        <div class="light-bubble" v-bind:class="{light_on: isLightOn}" ></div>
        <switch @change="toggleLight"></switch>
    </div>
</template>

<style>
    .light-bubble {
        width: 200px;
        height: 200px;
        border-radius: 50%;
    }
    .light_on {
        background-color: chartreuse;
    }
    .wrapper {
        align-items: center;
        margin-top: 120px;
    }

    .title {
        padding-top: 40px;
        padding-bottom: 40px;
        font-size: 48px;
    }

    .logo {
        width: 360px;
        height: 156px;
    }

    .desc {
        padding-top: 20px;
        color: #888;
        font-size: 24px;
    }
</style>

<script>
    export default {
        data: {
            logoUrl: 'http://img1.vued.vanthink.cn/vued08aa73a9ab65dcbd360ec54659ada97c.png',
            target: 'World',
            isLightOn: false,
        },
        methods: {
            toggleLight: function () {
                this.isLightOn = !this.isLightOn;
            },
            updateTime: function () {
              this.target = new Date();
            },
            update: function (e) {
                this.target = 'Weex'
                console.log('target:', this.target)
            }
        }
    }
</script>
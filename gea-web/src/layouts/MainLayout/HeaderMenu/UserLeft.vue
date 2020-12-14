<template>
    <div class="userLeft">
        <el-switch v-model="value1" active-text="神盾局" inactive-text="复仇者联盟">
        </el-switch>
        <el-slider v-model="value2" style="width: 100%" :marks="marks">
        </el-slider>
        <el-slider v-model="value3" style="width: 100%" :marks="marks">
        </el-slider>
        <el-slider v-model="value4" style="width: 100%" :marks="marks">
        </el-slider>
    </div>
</template>

<script>
export default {
    name: 'UserLeft',
    data() {
        return {
            value1: true,
            value2: 20,
            value3: 70,
            value4: 45,
            marks: {
                50: {
                    style: {
                        color: '#1989FA',
                    },
                    label: this.$createElement('span', '调试'),
                },
            },
        }
    },
}
</script>

<style lang="scss" scoped>
.userLeft {
    width: 100%;
    height: 300px;
    display: flex;
    justify-content: space-around;
    align-items: center;
    flex-direction: column;
}
</style>
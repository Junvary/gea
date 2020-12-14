<template>
    <div class="pageFooter">
        <span>欢迎使用 Gin-Element-Admin</span>
        <div style="width: 100%; z-index: 3002">
            <APlayer ref="aplayer" :audio="audio" :volume="volume" fixed autoplay></APlayer>
        </div>
    </div>
</template>

<script>
export default {
    name: 'PageFooter',
    data() {
        return {
            volume: 0.2,
            audio: [
                {
                    name: '挚爱',
                    artist: '陈勋奇',
                    url: 'http://music.163.com/song/media/outer/url?id=68331.mp3',
                    cover: 'http://p1.music.126.net/PGuUtwOqtgXGYIkw0Eg_HQ==/120946279068477.jpg?param=130y130',
                },
            ],
        }
    },
}
</script>

<style lang="scss" scoped>
.pageFooter {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    color: rgba($color: #000000, $alpha: 0.4);
    font-size: 0.8em;
}
</style>
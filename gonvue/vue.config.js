module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  outputDir: "../server/dist", // 빌드된 파일을 server폴더 밑에 저장하게 설정
  devServer: { // 개발시 편의를 위해서 api 관련 요청을 go 서버에서 받아오게 설정
    proxy: {
      '/api': {
        target: 'http://localhost:8082',
        changeOrigin: true,
      }
    }
  }
}

var vm = new Vue({
    el: '#rules',
    data: {
        rules_name: '',
        rules_content: '',
    },
    methods: {
        saveRules: function () {
            alert()
            // 发送post请求
            this.$http.post('http://localhost:21000/api/rulesengine/edit').then(function (res) {
                alert(res);
            }, function () {
                console.log('请求失败处理');
            });
        }
    }
});

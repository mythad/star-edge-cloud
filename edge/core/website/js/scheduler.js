function save() {
    new Vue({
        el: '#task',
        data: {
            input: {
                task_name: '',
                task_addr: '',
                task_frequency: '',
                task_offset: ''
            },
            output: '',
            task_frequency_list: [
                {
                    id: '1',
                    name: '一次'
                },
                {
                    id: '2',
                    name: '每分钟'
                },
                {
                    id: '3',
                    name: '每小时'
                }
                ,
                {
                    id: '4',
                    name: '每天'
                },
                {
                    id: '5',
                    name: '每周'
                },
                {
                    id: '6',
                    name: '每月'
                },
                {
                    id: '7',
                    name: '每年'
                }

            ],
        },
        created() {
            //如果没有这句代码，select中初始化会是空白的，默认选中就无法实现
            this.input.task_frequency = this.task_frequency_list[0].id;
        },
        methods: {
            saveTask: function () {
                // 发送post请求
                this.$http({
                    method: 'post',
                    url: '/api/scheduler/add',
                    params: this.input,
                    credientials: false,
                    emulateJSON: true
                }).then(function (response) {
                    alert(this.output);
                    this.output = response;
                    alert(this.output);
                }, function () {
                    console.log('请求失败处理');
                });
            },
            getTaskFrequency() {
                //获取选中
                console.log(this.input.task_frequency)
            }
        }
    });
}

function list() {
    var vm = new Vue({
        el: '#tb_tasks',
        data: {
            tasks: {},
            id: '',
        },
        mounted: function () {
            this.getTaskList();
        },
        // create:function(){},
        methods: {
            removeTask: function (id) {
                this.$http({
                    method: 'post',
                    url: '/api/scheduler/remove'
                }).then(function (response) {
                    console.log(response);
                    this.tasks = response;
                }, function () {
                    console.log('请求失败处理');
                });
            },
            getTaskList: function () {
                // 发送post请求
                this.$http({
                    method: 'post',
                    url: '/api/scheduler/list',
                    params: this.id,
                }).then(function (response) {
                    console.log(response);
                }, function () {
                    console.log('请求失败处理');
                });
            }
        }
    });
}
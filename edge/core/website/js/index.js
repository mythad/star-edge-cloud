function init() {
    loadDeviceService();
    loadAgorithmService();
    loadSystemService();
}

function loadDeviceService() {
    var vm = new Vue({
        el: '#tb_devices',
        data: {
            services: {},
        },
        http: {
            headers: { 'Access-Control-Allow-Origin': '*' }
        },
        methods: {
            load: function () {
                this.$http({
                    method: 'post',
                    url: '/api/device/all',
                }).then(function (response) {
                    this.services = response.data;
                    console.log(this.services[0].Status);
                });
            },
            removeDevice(id) {
                this.$http({
                    method: 'post',
                    params: { id: id },
                    url: '/api/device/remove',
                }).then(function (response) {
                    this.services = response;
                });
            },
            handleDevice(svr) {
                var d;
                if (svr.Status <= 1) {
                    d = { id: svr.ID, type: svr.Type, handle: "run" }
                } else {
                    d = {  id: svr.ID, type: svr.Type, handle: "stop" }
                }
                this.$http({
                    method: 'post',
                    params: d,
                    url: '/api/device/operate',
                }).then(function (response) {
                    if (response.data.indexOf("running") != -1) {
                        svr.Status = 2;
                    } else {
                        svr.Status = 1;
                        alert(response.data);
                    }
                });
            },
            edit(id) {
                window.location.href = 'editdevice.html?id=' + id;
            }
        }
    });
    vm.load();
}

function loadAgorithmService() {
    var vm = new Vue({
        el: '#tb_exts',
        data: {
            services: {},
        },
        methods: {
            load: function () {
                this.$http({
                    method: 'post',
                    url: '/api/extension/all',
                }).then(function (response) {
                    this.services = response.data;
                });
            },
            removeExt(id) {
                this.$http({
                    method: 'post',
                    params: { id: id },
                    url: '/api/extension/remove',
                }).then(function (response) {
                    this.services = response;
                });
            },
            handleExt(svr) {
                var d;
                if (svr.Status <= 1) {
                    d = { id: svr.ID, type: svr.Type, handle: "run" }
                } else {
                    d = {  id: svr.ID, type: svr.Type, handle: "stop" }
                }
                this.$http({
                    method: 'post',
                    params: d,
                    url: '/api/extension/operate',
                }).then(function (response) {
                    if (response.data.indexOf("running") != -1) {
                        svr.Status = 2;
                    } else {
                        svr.Status = 1;
                        alert(response.data);
                    }
                });
            },
            edit(id) {
                window.location.href = 'editextension.html?id=' + id;
            }
        }
    });
    vm.load();
}

function loadSystemService() {
    var vm = new Vue({
        el: '#tb_system_exts',
        data: {
            services: {},
        },
        http: {
            headers: { 'Access-Control-Allow-Origin': '*' }
        },
        methods: {
            load: function () {
                this.$http({
                    method: 'post',
                    url: '/api/system/all',
                }).then(function (response) {
                    this.services = response.data;
                });
            },
            handleSystemExt(_status, _type) {
                var d;
                if (_status <= 1) {
                    d = { type: _type, handle: "run" }
                } else {
                    d = { type: _type, handle: "stop" }
                }
                this.$http({
                    method: 'post',
                    params: d,
                    url: '/api/system/operate',
                }).then(function (response) {
                    this.services = response.data;
                });
            }
        }
    });
    vm.load();
}

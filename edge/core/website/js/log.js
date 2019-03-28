// $(document).ready(function () {
//     loadLog();
//     setInterval("loadLog()", 10000);
// });

// function loadData() {
//      loadLog();
// }

// function loadLog() {
//     var data = {
//         top: 100
//     };
//     $.ajax({
//         type: "POST",
//         url: "/api/log/load",
//         traditional: true,
//         // contentType: "application/json",
//         // data: JSON.stringify(data),
//         data: data,
//         success: function (msg) {
//             $("#log").html("");
//             $.each(JSON.parse(msg), function (index, item) {
//                 var newRow = '<div>时间：' + item.Time + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp等级：" + item.Level + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp信息：" + item.Message + '</div>'
//                 $("#log").append(newRow);
//             });
//         }
//     });
// }
function init() {
    loadLog();
    setInterval("loadLog()", 10000);
);

function loadLog() {
    var data = {
        top: 6
    };

    var vm = new Vue({
        el: '#log',
        data: {
            logs: {},
            logType: '',
        },
        methods: {
            load: function () {
                this.$http({
                    method: 'post',
                    url: 'http://localhost:17000/api/log/load',
                    params: this.logType,
                }).then(function (response) {
                    this.logs = response.data;
                });
            }
        }
    });
    vm.load();
}


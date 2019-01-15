var chartd;
var charte;
var charts;
$(document).ready(function () {
    // 基于准备好的dom，初始化echarts图表
    chartd = echarts.init(document.getElementById('pied'));
    optiond = {
        title: {
            text: '设备服务状态',
            // subtext: '纯属虚构',
            x: 'center'
        },
        tooltip: {
            trigger: 'item',
            formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        // legend: {
        //     orient: 'vertical',
        //     x: 'left',
        //     data: ['直接访问', '邮件营销', '联盟广告', '视频广告', '搜索引擎']
        // },
        toolbox: {
            show: true,
            feature: {
                mark: { show: true },
                dataView: { show: true, readOnly: false },
                magicType: {
                    show: true,
                    type: ['pie', 'funnel'],
                    option: {
                        funnel: {
                            x: '25%',
                            width: '100%',
                            funnelAlign: 'left',
                            max: 1548
                        }
                    }
                },
                restore: { show: true },
                saveAsImage: { show: true }
            }
        },
        calculable: true,
        series: [
            {
                name: '设备服务数据',
                type: 'pie',
                radius: '55%',
                center: ['50%', '60%'],
                data: [
                    { value: 0, name: '运行中设备服务数' },
                    { value: 0, name: '总设备数' },
                ]
            }
        ]
    };
    // 为echarts对象加载数据 
    chartd.setOption(optiond);

    // 基于准备好的dom，初始化echarts图表
    charte = echarts.init(document.getElementById('piee'));
    optione = {
        title: {
            text: '算法服务状态',
            // subtext: '纯属虚构',
            x: 'center'
        },
        tooltip: {
            trigger: 'item',
            formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        // legend: {
        //     orient: 'vertical',
        //     x: 'left',
        //     data: ['直接访问', '邮件营销', '联盟广告', '视频广告', '搜索引擎']
        // },
        toolbox: {
            show: true,
            feature: {
                mark: { show: true },
                dataView: { show: true, readOnly: false },
                magicType: {
                    show: true,
                    type: ['pie', 'funnel'],
                    option: {
                        funnel: {
                            x: '25%',
                            width: '100%',
                            funnelAlign: 'left',
                            max: 1548
                        }
                    }
                },
                restore: { show: true },
                saveAsImage: { show: true }
            }
        },
        calculable: true,
        series: [
            {
                name: '算法服务数据',
                type: 'pie',
                radius: '55%',
                center: ['50%', '60%'],
                data: [
                    { value: 0, name: '运行中算法服务数' },
                    { value: 0, name: '总服务数' },
                ]
            }
        ]
    };
    // 为echarts对象加载数据 
    charte.setOption(optione);

    loadData();
    setInterval("loadData()", 10000);
});

function loadData() {
    count();
    loadStatus();
    loadLog();
}

function loadStatus() {
    $.ajax({
        type: "POST",
        url: "/api/store/status",
        traditional: true,
        success: function (msg) {
            if (!msg) return;
            $("#storeStatus").html(msg);
            if (msg == "running") {
                $("#handleStore").html('停止');
            } else {
                $("#handleStore").html('启动');
            }
        }
    });
    $.ajax({
        type: "POST",
        url: "/api/log/status",
        traditional: true,
        success: function (msg) {
            if (!msg) return;
            $("#logStatus").html(msg);
            if (msg == "running") {
                $("#logStatus").html('停止');
            } else {
                $("#logStatus").html('启动');
            }
        }
    });
}

function loadLog() {
    var data = {
        top: 6
    };
    $.ajax({
        type: "POST",
        url: "http://localhost:17000/api/log/load",
        traditional: true,
        dataType: 'JSONP',
        // contentType: "application/json",
        // data: JSON.stringify(data),
        data: data,
        success: function (msg) {
            alert();
            if (!msg) return;
            $("#log").html("");
            $.each(JSON.parse(msg), function (index, item) {
                var newRow = '<div>时间：' + item.Time + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp等级：" + item.Level + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp信息：" + item.Message + '</div>'
                $("#log").append(newRow);
            });
        }
    });
}

function count() {
    $.ajax({
        type: "POST",
        url: "/api/device/count",
        traditional: true,
        success: function (msg) {
            var result = $.parseJSON(msg);
            json = [
                { value: result.Running, name: '运行中' },
                { value: result.Total - result.Running, name: '未运行' },
            ];
            // json = [
            //     { value: 1, name: '运行中' },
            //     { value: 3, name: '总设备数' },
            // ];

            refreshData(json);
        }
    });
}

//更新数据
function refreshData(data) {
    if (chartd) {
        var optiond = chartd.getOption();
        optiond.series[0].data = data;
        chartd.setOption(optiond);
    }

    if (charte) {
        var optione = charte.getOption();
        optione.series[0].data = data;
        charte.setOption(optione);
    }
}

function handleStoreService() {
    if ($("#handleStore").html() == "启动") {
        $.ajax({
            type: "POST",
            url: "/api/store/run",
            success: function (msg) {
                if (msg == "running") {
                    $("#handleStore").html('停止')
                } else {
                    alert(msg);
                }
            }
        });
        return
    }

    $.ajax({
        type: "POST",
        url: "/api/store/stop",
        traditional: true,
        success: function (msg) {
            if (msg == "stopped") {
                $("#handleStore").html('启动')
            } else {
                alert(msg);
            }
        }
    });

}

function handleLogService() {
    if ($("#handleLog").html() == "启动") {
        $.ajax({
            type: "POST",
            url: "/api/log/run",
            traditional: true,
            success: function (msg) {
                if (msg == "running") {
                    $("#handleLog").html('停止')
                } else {
                    alert(msg);
                }
            }
        });
        return
    }

    $.ajax({
        type: "POST",
        url: "/api/log/stop",
        traditional: true,
        success: function (msg) {
            if (msg == "stopped") {
                $("#handleLog").html('启动')
            } else {
                alert(msg);
            }
        }
    });

}

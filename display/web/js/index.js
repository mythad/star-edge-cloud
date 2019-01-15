var charte;
$(document).ready(function () {
    // 基于准备好的dom，初始化echarts图表
    charte = echarts.init(document.getElementById('pied'));
    optione = {
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
    charte.setOption(optiond);

    loadData();
    setInterval("loadData()", 10000);
});

function loadData() {
    count();
    // loadLog();
}

function loadLog() {
    var data = {
        top: 6
    };
    $.ajax({
        type: "POST",
        url: "/api/log/load",
        traditional: true,
        // contentType: "application/json",
        // data: JSON.stringify(data),
        data: data,
        success: function (msg) {
            $("#log").html("");
            $.each(JSON.parse(msg), function (index, item) {
                var newRow = '<div>时间：' + item.Time + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp等级：" +
                    item.Level + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp信息：" +
                    item.Message + '</div>';
                $("#log").append(newRow);
            });
        }
    });
}

function count() {
    $.ajax({
        type: "POST",
        url: "http://localhost:22000/api/transport/count",
        traditional: true,
        success: function (msg) {
            $("#rdcount").html(msg);
        }
    });
}

//更新数据
function refreshData(data) {
    if (charte) {
        var optione = charte.getOption();
        optione.series[0].data = data;
        charte.setOption(optione);
    }

}
$(document).ready(function () {
    
    loadData();
    setInterval("loadData()", 10000);
});

function loadData() {
     loadLog();
}

function loadLog() {
    var data = {
        top: 100
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
                var newRow = '<div>时间：' + item.Time + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp等级：" + item.Level + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp信息：" + item.Message + '</div>'
                $("#log").append(newRow);
            });
        }
    });
}


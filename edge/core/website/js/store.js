function handleStoreService() {
    if ($("#handle").html() == "启动") {
        $.ajax({
            type: "POST",
            url: "/api/store/run",
            traditional: true,
            data: data,
            success: function (msg) {
                if (msg == "running") {
                    $("#handle").html('停止')
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
        data: data,
        success: function (msg) {
            if (msg == "stopped") {
                $("#handle").html('启动')
            } else {
                alert(msg);
            }
        }
    });

}

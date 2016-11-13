function handleJQueryAjaxError(xhr, textStatus, error) {
    if (textStatus === 'timeout') {
        $.msg.error('服务器超时！');
    }
    else {
        $.msg.error('访问服务器出错，请重试！' + error);
    }
}

$.ajaxSetup(
    {
        type: "POST",
        "timeout": 15000,
        "error": handleJQueryAjaxError
    }
);

function isEmptyStr(str) {
    if (!str || str.length === 0 || str === "" || typeof str == 'undefined' || !/[^\s]/.test(str) || /^\s*$/.test(str) || str.replace(/\s/g, "") === "") {
        return true;
    } else {
        return false;
    }
}

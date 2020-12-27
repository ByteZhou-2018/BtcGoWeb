//获取元素
var logbut = document.getElementById("logbut");
//绑定事件
logbut.addEventListener("click", function () {
    /*
    AJAX 请求  js发送请求，可以异步
    */
    var xhr;
    if (XMLHttpRequest) {
        //谷歌，火狐 等
        xhr = new XMLHttpRequest();
    } else if (ActiveXObject) {
        //ie 6,7,8
        xhr = new ActiveXObject("Microsoft.XMLHTTP");
    } else {
        alert("你是什么浏览器？自己写的？");
    }
    //打开设置url与方式
    xhr.open("POST", "/emp_login.do", true);
    //发送
    var un = document.getElementById("input_un").value;
    var pwd = document.getElementById("input_pwd").value;

    //js中没有定义数据传输的任何格式，我们如何设置数据格式？
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send("employeeName=" + un + "&employeePwd=" + pwd);
    //接收响应数据
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                //返回完成之后拿值
                if (xhr.responseText == "true") {
                    alert("登录成功，2秒后开始跳转");
                    window.location.href = "/login";
                } else {
                    var prom1 = document.getElementById("prom1");
                    prom1.innerHTML = "<span style='color:red'>用户名或密码错误</span>";
                }
            } else {
                alert("xhr的状态码： " + xhr.status);
            }
        }
    }
}
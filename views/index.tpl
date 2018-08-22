<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <link href="https://cdn.bootcss.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet">
    <link href="static/css/reset.css" rel="stylesheet">
    <link href="static/css/index.css" rel="stylesheet">
</head>
<body>
<div class="container">
    <!-- Content here -->
    <form>
        <div class="form-group">
            <label for="exampleInputEmail1">输入邮箱</label>
            <input type="text" class="form-control" id="InputEmail1" aria-describedby="emailHelp" placeholder="邮箱地址">
            <small id="emailHelp" class="form-text text-muted hide"></small>
        </div>
        <input type="button" id="submit" class="btn btn-primary" value="提交" />
    </form>
</div>
</body>
<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/bootstrap/4.1.1/js/bootstrap.min.js"></script>
<script>

    $(function() {
        var val = ''
        $("#InputEmail1").on('blur', function (e) {
            var value = e.target.value
            if (!value.match(/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/)) {
                $("#emailHelp").html('邮箱错误').removeClass("hide").addClass("show")
            } else {
                $("#emailHelp").html('校验完成').removeClass("show").addClass("hide")
            }
            val = e.target.value
        })

        $("#submit").on('click', function (e) {
            $.ajax({
                type: "POST",
                url: "/v1/email",
                cache:false,
                //数据类型，这里我用的是json
                dataType: "json",
                //必要的时候需要用JSON.stringify() 将JSON对象转换成字符串
                data: JSON.stringify({email: val}), //data: {key:value},
                //添加额外的请求头
                headers : {'Content-Type': 'application/json'},
                //请求成功的回调函数
                success: function (data) {
                    console.log(data)
                }
            })
        })
    })
</script>
</html>
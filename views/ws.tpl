<!DOCTYPE html>
<html lang="en">
<head>
    <title>Chat Example</title>
    <script type="text/javascript">
    </script>
    <style type="text/css">
        html {
            overflow: hidden;
        }

        body {
            overflow: hidden;
            padding: 0;
            margin: 0;
            width: 100%;
            height: 100%;
            background: gray;
            background: url("static/images/bg2.jpg") no-repeat;
            background-size: cover;
            font-size: 12px;
            color: #423232;
        }

        .container {
            width: 600px;
            margin: 20px auto;
            height: 700px;
            background: whitesmoke;
            border-radius: 6px;
            border-radius: 6px;
            box-shadow: 0 0 10px #000;
            overflow: hidden;
            /*padding: 6px 10px;*/
        }

        .messages {
            width: 100%;
            height: 600px;
            border-bottom: 10px solid #ffdfdf;
            overflow: auto;
        }
        .content {
            height: 90px;
            width: 100%;
        }
        .textarea {
            display: inline-block;
            width: 480px;
            height: 82px;
            padding: 4px;
            border: none;
            resize: none;
            font-size: 14px;
        }
        .textarea, .textarea:hover {
            outline: none;
        }
        .left {
            float: left;
        }
        .right  {
            float: right;
        }
        .submit {
            line-height: 90px;
            width: 80px;
            text-align: center;
            background-color: #5e2ca5;
            color: #fff;
            cursor: pointer;
        }
        .submit:hover {
            background-color: #1f69c0;
        }
        .handle {
            width: 30px;
            height: 90px;
            background-color: #7d8293;
        }
        .message-content {
            padding: 4px 10px;
            overflow: hidden;
        }
        .message {
            display: inline-block;
            border: 1px solid #cdecff;
            background-color: #cdecff;
            border-radius: 4px;
            padding: 8px 10px;
            max-width: 266px;
            word-wrap: break-word;
            line-height: 18px;
            font-size: 14px;
            color: #666;
        }
        .message-content-left .message {
            float: left;
        }
        .message-content-right .message {
            float: right;
        }
    </style>
</head>
<body>
    <div id="container" class="container">
        <div id="messages" class="messages">
        </div>
        <div class="content">
            <textarea id="textarea" class="textarea">
            </textarea>
            <div class="right">
                <div class="handle left"></div>
                <div id="submit" class="submit right">提交</div>
            </div>
        </div>
    </div>
</body>
<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
<script>
    $(function() {
        $("#textarea").val('')
        var id = new Date() * 1;
        var conn;
        var messagesElm = $("#messages");
        function appendMessage(item) {
            messagesElm.append(item);
        }
        function sendMessge() {
            var val = $("#textarea").val();
            if (!$.trim(val)) return;
            conn.send(JSON.stringify({
                id: id,
                type: "text",
                message: $.trim(val),
                room: 1
            }));
            $("#textarea").val('');
        }
        if (window["WebSocket"]) {
            conn = new WebSocket("ws://" + document.location.host + "/ws");
            conn.onclose = function (evt) {
                var item = document.createElement("div");
                item.innerHTML = "<b>Connection closed.</b>";
                appendMessage(item);
            };
            conn.onmessage = function (evt) {
                var data = JSON.parse(evt.data);
                var message = data.message;
                var contentElm = $("<div class=\"message-content\"></div>");
                if (data.id === id) {
                    contentElm.addClass("message-content-right")
                } else {
                    contentElm.addClass("message-content-left")
                }
                var mess = $("<div class=\"message\"></div>");
                mess.html(message);
                contentElm.append(mess);
                appendMessage(contentElm);
            };
        }


        $("#submit").on("click", function(e) {
            sendMessge()
        })
        $("#textarea").on('keyup', function(e) {
            if (e.keyCode == "13") {
                sendMessge()
            }
        })
    })
</script>
</html>

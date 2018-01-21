<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>技术支持巴</title>
<link rel="stylesheet" href="../static/layui/css/layui.css">
<link rel="stylesheet" href="../static/layui/css/admin.css">
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
	{{.Header}}

	{{.SiderBar}}

	{{.LayoutContent}}

	{{.FootBar}}
</div>
<script src="../static/layui/layui.all.js"></script>
<script>
//JavaScript代码区域
layui.use('element', function(){
		var element = layui.element;
		element.on('nav(test)', function(elem){
			console.log(elem[0].textContent); //得到当前点击的DOM对象
			});

		});


</script>
</body>
</html>

<table id="CQtable" lay-filter="test"></table>
<script src="../static/layui/layui.all.js"></script>
<div class="layui-body">
    <div style="padding: 0px 5px 0px 5px;">
        <table class="layui-hide" id="test"></table>
    </div>
</div>
<script>
layui.use('table', function() {
	var table = layui.table;

	table.render({
		elem: '#test',
		queryParamsType: 'limit',
		method: 'post',
		sidePagination: 'server',
		pagination: true,
		url: '{{urlfor "CQController.ListCQreq"}}',
		cols: [
			[{
				field: 'Id',
				width: 80,
				title: 'ID'
			}, {
				field: 'CreateTime',
				width: 180,
				title: '请求时间',
				sort: true
			}, {
				field: 'Message',
				width: '100%',
				title: '原始信息',
				sort: true
			}]
		],
		response: {
			statusName: 'status' //数据状态的字段名称，默认：code
				,
			statusCode: 200 //成功的状态码，默认：0
				,
			msgName: 'msg' //状态信息的字段名称，默认：msg
				,
			countName: 'total' //数据总数的字段名称，默认：count
				,
			dataName: 'data' //数据列表的字段名称，默认：data
		},

		page: { //支持传入 laypage 组件的所有参数（某些参数除外，如：jump/elem） - 详见文档
			layout: ['limit', 'count', 'prev', 'page', 'next', 'skip'] //自定义分页布局
				//,curr: 5 //设定初始在第 5 页
				,
			groups: 1 //只显示 1 个连续页码
				,
			first: "首页" //不显示首页
				,
			last: "尾页" //不显示尾页

		}
	});
});
</script>
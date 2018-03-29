<!DOCTYPE html>
<html>
  <head><title>{{ .Title }}</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
        <link rel="stylesheet" href="/static/layui/css/layui.css">
        <link rel="stylesheet" href="/static/css/custom.css">
        <script src="/static/layui/layui.js"></script>
  </head>


  <body class="layui-layout-body">
     <div class="layui-layout layui-layout-admin"> 
       <div class="layui-header">
           <div class="layui-logo">GosshTool</div>
           <ul class="layui-nav layui-layout-left">
              <li class="layui-nav-item"><a href="">主机</a></li>
              <li class="layui-nav-item"><a href="">组别</a></li>
              <li class="layui-nav-item"><a href="">配置</a></li>
              <li class="layui-nav-item">
                 <a href="javascript:;">个人信息</a>
              <dl class="layui-nav-child">
                  <dd><a href="">邮件管理</a></dd>
                  <dd><a href="">消息管理</a></dd>
                  <dd><a href="">授权管理</a></dd>
              </dl>
              </li>
           </ul>
       </div>
      <div class="layui-body">
          {{ .LayoutContent }}
      </div>
      <div class='col-sm-12 text-info bg-color-white padding-tb-15 text-center'>Gosshtool</div>
    </div>
  </body>
</html>

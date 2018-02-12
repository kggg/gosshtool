<!DOCTYPE html>
<html>
  <head><title>{{ .Title }}</title>
                <meta charset="utf-8">
                <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  <link rel="stylesheet" href="/static/css/custom.css">
  <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
  </head>

  <nav class="navbar navbar-default  navbar-inverse">
      <div class="container-fluid">
            <div class="navbar-header">
                    <a href='/' class="navbar-brand active"><span class="text-primary">Gosshtool</span></a>
            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
              <ul class="nav navbar-nav navbar-inverse">
                 <li class="active"><a href="/"><i class="glyphicon glyphicon-home"> </i><span class="sr-only">(current)</span></a></li>
              </ul>
              <ul class="nav navbar-nav navbar-right">
                 <li><a href="/host"><i class="glyphicon glyphicon-object-align-vertical"></i> 主机</a></li>
                 <li><a href="/group">组别</a></li>
                 <li><a href="/config"><i class="glyphicon glyphicon-cog"></i> 配置</a></li>
                 <li><a href="/person"><i class="glyphicon glyphicon-user"></i> 人个信息</a></li>
              </ul>
            </div><!-- /.navbar-collapse -->
       </div><!-- /.container-fluid -->
  </nav>

  <body>
      <div class="container-fluid">
          {{ .LayoutContent }}
      </div>
      <div class='col-sm-12 text-info bg-color-white padding-tb-15 text-center'>Steven Lu</div>
  </body>
</html>

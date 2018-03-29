
       <aside class="main-sidebar">
          <!-- sidebar: style can be found in sidebar.less -->
          <section class="sidebar">
          <!-- Sidebar user panel -->
          <div class="user-panel">
             <div class="pull-left image">
                <img src="/static/img/user2-160x160.jpg" class="img-circle" alt="User Image">
             </div>
             <div class="pull-left info">
                 <p> Username </p>
                 <a href="#"><i class="fa fa-circle text-success"></i> Online</a>
             </div>
          </div>

          <!-- search form -->
          <form action="#" method="get" class="sidebar-form">
             <div class="input-group">
                 <input type="text" name="q" class="form-control" placeholder="搜索...">
                 <span class="input-group-btn">
                      <button type="submit" name="search" id="search-btn" class="btn btn-flat">
                           <i class="fa fa-search"></i>
                      </button>
                 </span>
             </div>
           </form>

         <!-- sidebar menu: : style can be found in sidebar.less -->
         <ul class="sidebar-menu tree" data-widget="tree" data-accordion=0>
            <li class="header">MAIN</li>
            <!-- Optionally, you can add icons to the links -->
            <li id="dashboard"><a href="/"><i class="fa fa-dashboard"></i>
               <span>首页</span></a>
            </li>

            <li class="treeview" id="section">
                 <a href="#"><i class="fa fa-list-alt"></i>  <span>板块管理</span> <i class="fa fa-angle-left pull-right"></i></a>
                 <ul class="treeview-menu">
                    <li><a href="/admin/section"><i class="fa fa-circle-o"></i>板块列表</a></li>
                    <li><a href="/admin/section/add"><i class="fa fa-circle-o"></i>新增板块</a></li>
                 </ul>
            </li>

            <li class="treeview" id="label">
                 <a href="#"><i class="fa fa-list-alt"></i>  <span>标签管理</span> <i class="fa fa-angle-left pull-right"></i></a>
                 <ul class="treeview-menu">
                    <li><a href="/admin/label"><i class="fa fa-circle-o"></i>标签管理</a></li>
                    <li><a href="/admin/label/add"><i class="fa fa-circle-o"></i>新增标签</a></li>
                 </ul>
            </li>


            <li class="treeview" id="article">
                 <a href="#"><i class="fa fa-list-alt"></i>  <span>文章管理</span> <i class="fa fa-angle-left pull-right"></i></a>
                 <ul class="treeview-menu">
                    <li><a href="/admin/article"><i class="fa fa-circle-o"></i>文章管理</a></li>
                    <li><a href="/admin/article/add"><i class="fa fa-circle-o"></i>新增文章</a></li>
                 </ul>
            </li>


            <li class="treeview" id="users">
                 <a href="#"><i class="fa fa-user"></i>  <span>用户管理</span> <i class="fa fa-angle-left pull-right"></i></a>
                 <ul class="treeview-menu">
                    <li><a href="/admin/user"><i class="fa fa-circle-o"></i>用户列表</a></li>
                    <li><a href="/admin/user/add"><i class="fa fa-circle-o"></i> 添加用户</a></li>
                 </ul>
            </li>


            <li class="treeview" id="jobs">
                 <a href="#"><i class="fa fa-user"></i>  <span>51job管理</span> <i class="fa fa-angle-left pull-right"></i></a>
                 <ul class="treeview-menu">
                    <li><a href="/admin/job"><i class="fa fa-circle-o"></i>申请列表</a></li>
                    <li><a href="/admin/job/seeme"><i class="fa fa-circle-o"></i> 谁看了我</a></li>
                    <li><a href="/admin/job/search"><i class="fa fa-circle-o"></i> 谁看了我</a></li>
                 </ul>
            </li>

            <li class="treeview" id="system">
                 <a href="#"><i class="fa fa-cog"></i>  <span>系统管理</span> <i class="fa fa-angle-left pull-right"></i></a>
                 <ul class="treeview-menu">
                    <li><a href="#"><i class="fa fa-circle-o"></i>系统设置</a></li>
                    <li><a href="#"><i class="fa fa-circle-o"></i>数据备份</a></li>
                 </ul>
            </li>
            <li><a href="/"><i class="fa fa-sign-out"></i>  <span>返回首页</span></a></li>

         </ul>
       </section>
       </aside>




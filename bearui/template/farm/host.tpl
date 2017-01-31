<!DOCTYPE html>
<!--
This is a starter template page. Use this page to start your new project from
scratch. This page gets rid of all links and provides the needed markup only.
-->
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Farm Stat</title>
  <!-- Tell the browser to be responsive to screen width -->
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <link rel="shortcut icon" type="image/png" href="/static/dist/img/favicon.png"/>
  <!-- Bootstrap 3.3.6 -->
  <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">
  <!-- Font Awesome -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.5.0/css/font-awesome.min.css">
  <!-- Ionicons -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/ionicons/2.0.1/css/ionicons.min.css">
 <!-- DataTables -->
  <link rel="stylesheet" href="/static/plugins/datatables/dataTables.bootstrap.css">
  <!-- Theme style -->
  <link rel="stylesheet" href="/static/dist/css/AdminLTE.min.css">
  <!-- AdminLTE Skins. We have chosen the skin-blue for this starter
        page. However, you can choose any other skin. Make sure you
        apply the skin class to the body tag so the changes take effect.
  -->

    <!-- AdminLTE Skins. Choose a skin from the css/skins
       folder instead of downloading all of them to reduce the load. -->
  <link rel="stylesheet" href="/static/dist/css/skins/skin-yellow.min.css">

  <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
  <![endif]-->
</head>
<!--
BODY TAG OPTIONS:
=================
Apply one or more of the following classes to get the
desired effect
|---------------------------------------------------------|
| SKINS         | skin-blue                               |
|               | skin-black                              |
|               | skin-purple                             |
|               | skin-yellow                             |
|               | skin-red                                |
|               | skin-green                              |
|---------------------------------------------------------|
|LAYOUT OPTIONS | fixed                                   |
|               | layout-boxed                            |
|               | layout-top-nav                          |
|               | sidebar-collapse                        |
|               | sidebar-mini                            |
|---------------------------------------------------------|
-->
<body class="hold-transition fixed skin-yellow sidebar-mini">
<div class="wrapper">

  <!-- Main Header -->
  <header class="main-header">

    <!-- Logo -->
    <a href="/" class="logo">
      <!-- mini logo for sidebar mini 50x50 pixels -->
      <span class="logo-mini"><i class="fa fa-paw"></i></span>
      <!-- logo for regular state and mobile devices -->
      <span class="logo-lg"><b>HoneyBear</b></span>
    </a>

    <!-- Header Navbar -->
    <nav class="navbar navbar-static-top" role="navigation">
      <!-- Sidebar toggle button-->
      <a href="#" class="sidebar-toggle" data-toggle="offcanvas" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>
      <!-- Navbar Right Menu -->
      <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">
          <!-- Messages: style can be found in dropdown.less-->
          <li class="dropdown messages-menu">
            <!-- Menu toggle button -->
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <i class="fa fa-envelope-o"></i>
              <span class="label label-success">4</span>
            </a>
            <ul class="dropdown-menu">
              <li class="header">You have 4 messages</li>
              <li>
                <!-- inner menu: contains the messages -->
                <ul class="menu">
                  <li><!-- start message -->
                    <a href="#">
                      <div class="pull-left">
                        <!-- User Image -->
                        <img src="/static/dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">
                      </div>
                      <!-- Message title and timestamp -->
                      <h4>
                        Support Team
                        <small><i class="fa fa-clock-o"></i> 5 mins</small>
                      </h4>
                      <!-- The message -->
                      <p>Virtual patch is coming soon!</p>
                    </a>
                  </li>
                  <!-- end message -->
                </ul>
                <!-- /.menu -->
              </li>
              <li class="footer"><a href="#">See All Messages</a></li>
            </ul>
          </li>
          <!-- /.messages-menu -->

          <!-- Notifications Menu -->
          <li class="dropdown notifications-menu">
            <!-- Menu toggle button -->
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <i class="fa fa-bell-o"></i>
              <span class="label label-info">10</span>
            </a>
            <ul class="dropdown-menu">
              <li class="header">You have 10 notifications</li>
              <li>
                <!-- Inner Menu: contains the notifications -->
                <ul class="menu">
                  <li><!-- start notification -->
                    <a href="#">
                      <i class="fa fa-users text-aqua"></i> 5 new hackers joined today
                    </a>
                  </li>
                  <!-- end notification -->
                </ul>
              </li>
              <li class="footer"><a href="#">View all</a></li>
            </ul>
          </li>
          <!-- Tasks Menu -->
          <li class="dropdown tasks-menu">
            <!-- Menu Toggle Button -->
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <i class="fa fa-flag-o"></i>
              <span class="label label-danger">9</span>
            </a>
            <ul class="dropdown-menu">
              <li class="header">You have 9 tasks</li>
              <li>
                <!-- Inner menu: contains the tasks -->
                <ul class="menu">
                  <li><!-- Task item -->
                    <a href="#">
                      <!-- Task title and progress text -->
                      <h3>
                        Say hi to new hackers
                        <small class="pull-right">20%</small>
                      </h3>
                      <!-- The progress bar -->
                      <div class="progress xs">
                        <!-- Change the css width attribute to simulate progress -->
                        <div class="progress-bar progress-bar-aqua" style="width: 20%" role="progressbar" aria-valuenow="20" aria-valuemin="0" aria-valuemax="100">
                          <span class="sr-only">20% Complete</span>
                        </div>
                      </div>
                    </a>
                  </li>
                  <!-- end task item -->
                </ul>
              </li>
              <li class="footer">
                <a href="#">View all tasks</a>
              </li>
            </ul>
          </li>
          <!-- User Account Menu -->
          <li class="dropdown user user-menu">
            <!-- Menu Toggle Button -->
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <!-- The user image in the navbar-->
              <img src="/static/dist/img/user2-160x160.jpg" class="user-image" alt="User Image">
              <!-- hidden-xs hides the username on small devices so only the image appears. -->
              <span class="hidden-xs">BigDaddy</span>
            </a>
            <ul class="dropdown-menu">
              <!-- The user image in the menu -->
              <li class="user-header">
                <img src="/static/dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">

                <p>
                  BigDaddy - The Notail
                  <small>Member since Nov. 2012</small>
                </p>
              </li>
              <!-- Menu Body -->
              <li class="user-body">
                <div class="row">
                  <div class="col-xs-4 text-center">
                    <a href="#">Followers</a>
                  </div>
                  <div class="col-xs-4 text-center">
                    <a href="#">Sales</a>
                  </div>
                  <div class="col-xs-4 text-center">
                    <a href="#">Friends</a>
                  </div>
                </div>
                <!-- /.row -->
              </li>
              <!-- Menu Footer-->
              <li class="user-footer">
                <div class="pull-left">
                  <a href="#" class="btn btn-default btn-flat">Profile</a>
                </div>
                <div class="pull-right">
                  <a href="#" class="btn btn-default btn-flat">Sign out</a>
                </div>
              </li>
            </ul>
          </li>
          <!-- Control Sidebar Toggle Button -->
          <li>
            <a href="#" data-toggle="control-sidebar"><i class="fa fa-gears"></i></a>
          </li>
        </ul>
      </div>
    </nav>
  </header>

  <!-- Left side column. contains the logo and sidebar -->
  <aside class="main-sidebar">

    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">

      <!-- Sidebar user panel (optional) -->
      <div class="user-panel">
        <div class="pull-left image">
          <img src="/static/dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">
        </div>
        <div class="pull-left info">
          <p>BigDaddy</p>
          <!-- Status -->
          <a href="#"><i class="fa fa-circle text-success"></i> Online</a>
        </div>
      </div>

      <!-- search form (Optional) -->
      <form action="#" method="get" class="sidebar-form">
        <div class="input-group">
          <input type="text" name="q" class="form-control" placeholder="Search...">
              <span class="input-group-btn">
                <button type="submit" name="search" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
                </button>
              </span>
        </div>
      </form>
      <!-- /.search form -->

      <!-- Sidebar Menu -->
      <ul class="sidebar-menu">
        <li class="header">MAIN NAVIGATION</li>
        <!-- Optionally, you can add icons to the links -->
        <li class="">
          <a href="/">
            <i class="fa fa-dashboard"></i> <span>BearBoard</span>
          </a>
        </li>

        <li class="treeview">
          <a href="#">
            <i class="fa fa-exchange"></i> <span>BearProxy</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/proxy/stat"><i class="fa fa-circle-o"></i> Proxy Stat</a></li>
            <li><a href="/proxy/uid"><i class="fa fa-circle-o"></i> User Identity</a></li>
          </ul>
        </li>
        
        <li class="treeview active">
          <a href="#">
            <i class="fa fa-object-group"></i> <span>BearFarm</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/farm/host"><i class="fa fa-circle"></i> Host Stat</a></li>
            <li><a href="/farm/container"><i class="fa fa-circle-o"></i> Containers</a></li>
            <li><a href="/farm/topology"><i class="fa fa-circle-o"></i> Infrastructure</a></li>
          </ul>

          <li class="treeview">
          <a href="#">
            <i class="fa fa-filter"></i> <span>BearFilter</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/filter/warning"><i class="fa fa-circle-o"></i> Warnings</a></li>
            <li><a href="/filter/tracking"><i class="fa fa-circle-o"></i> User Tracking</a></li>
          </ul>
        </li>

        <li class="treeview">
          <a href="#">
            <i class="fa fa-gears"></i> <span>BearEngine</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/engine/score"><i class="fa fa-circle-o"></i> Scoring System</a></li>
            <li><a href="/engine/object"><i class="fa fa-circle-o"></i> Objects</a></li>
            <li><a href="/engine/snapshot"><i class="fa fa-circle-o"></i> Profile Control</a></li>
          </ul>
        </li>

        <li>
          <a href="/log">
            <i class="fa fa-search"></i> <span>BearLog</span>
          </a>
        </li>

        <li class="treeview">
          <a href="#">
            <i class="fa fa-sliders"></i> <span>BearControl</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/control/framework"><i class="fa fa-circle-o"></i> Framework</a></li>
            <li><a href="/control/user"><i class="fa fa-circle-o"></i> Access Control</a></li>
            <li><a href="/control/config"><i class="fa fa-circle-o"></i> Configuration</a></li>
          </ul>
        </li>
        </li>

        <li class="treeview">
          <a href="#">
            <i class="fa fa-heartbeat"></i> <span>BearHelp</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/help/about"><i class="fa fa-circle-o"></i> About Us</a></li>
            <li><a href="/help/document"><i class="fa fa-circle-o"></i> Documentation</a></li>
            <li><a href="/help/api"><i class="fa fa-circle-o"></i> API</a></li>
          </ul>
        </li>


      </ul>
      <!-- /.sidebar-menu -->
    </section>
    <!-- /.sidebar -->
  </aside>

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
        Host status
        <small>{{html .hostname}}</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">BearFarm</a></li>
        <li class="active">Host</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">

      <!-- Your Page Content Here -->
      <div class="row">

        <!-- Basic info -->
        <div class="col-md-3">

          <div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">{{html .hostname}}</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <strong><i class="fa fa-circle-o-notch margin-r-5"></i> Up Time:</strong>
              <p class="text-muted" id="upTime">{{html .uptime}}</p>
              <hr>

              <strong><i class="fa fa-clock-o margin-r-5"></i> Current Time:</strong>
              <p class="text-muted" id="curTime">{{html .now}}</p>
              <hr>

              <strong><i class="fa fa-cloud margin-r-5"></i> IP:</strong>
              <p class="text-muted">{{html .ip}}</p>
              <hr>

              <strong><i class="fa fa-gear margin-r-5"></i> CPU:</strong>
              <p class="text-muted">{{html .cpu1}}<span class="pull-right">{{html .cpu2}}</span></p>
              <hr>

              <strong><i class="fa fa-heartbeat margin-r-5"></i> Logical Memory:</strong>
              <p class="text-muted">{{html .memory}}</p>
              <hr>

              <strong><i class="fa fa-plus-square margin-r-5"></i> Swap Memory:</strong>
              <p class="text-muted">{{html .swap}}</p>
              <hr>

              <strong><i class="fa fa-cloud-download margin-r-5"></i> Storage:</strong>
              <p class="text-muted">{{html .storage}}</p>
              <hr>

              <strong><i class="fa fa-database margin-r-5"></i> Kernel:</strong>
              <p class="text-muted">{{html .kernel}}</p>
              <hr>

              <strong><i class="fa fa-codepen margin-r-5"></i> Docker:</strong>
              <p class="text-muted">Server {{html .container.version.Version}} / API {{html .container.version.ApiVersion}}</p>
              <hr>

              <strong><i class="fa fa-linux margin-r-5"></i> OS:</strong>
              <p class="text-muted">{{html .os}}</p>

            </div>
            <!-- /.box-body -->
          </div>

        </div>

        <!-- Chart, detail -->
        <div class="col-md-9">

          <!-- chart -->
          
          <div class="row">

            <div class="col-md-6">
              <div class="box">
                <div class="box-header with-border">
                  <h3 class="box-title">CPU</h3>
                  <div class="pull-right">
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(220,220,220,1)"></i> System
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(151,187,205,1)"></i> User
                  </div>
                </div>
                <div class="box-body">
                  <div class="chart">
                    <canvas id="cpuChart"></canvas>
                  </div>
                </div>
              </div>

              <div class="box">
                <div class="box-header with-border">
                  <h3 class="box-title">Network</h3>
                  <div class="pull-right">
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(220,220,220,1)"></i> Transmit
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(151,187,205,1)"></i> Receive
                  </div>
                </div>
                <div class="box-body">
                  <div class="chart">
                    <canvas id="networkChart"></canvas>
                  </div>
                </div>
              </div>

            </div>

            <div class="col-md-6">
              <div class="box">
                <div class="box-header with-border">
                  <h3 class="box-title">Memory</h3>
                  <div class="pull-right">
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(220,220,220,1)"></i> Logical Used
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(151,187,205,1)"></i> Swap Used
                  </div>
                </div>
                <div class="box-body">
                  <div class="chart">
                    <canvas id="memoryChart"></canvas>
                  </div>
                </div>
              </div>

              <div class="box">
                <div class="box-header with-border">
                  <h3 class="box-title">Disk IO</h3>
                  <div class="pull-right">
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(220,220,220,1)"></i> Read
                    <i class="fa fa-circle-o margin-r-5" style="color: rgba(151,187,205,1)"></i> Write
                  </div>
                </div>
                <div class="box-body">
                  <div class="chart">
                    <canvas id="storageChart"></canvas>
                  </div>
                </div>
              </div>
            </div>

          </div>
          

          <div class="nav-tabs-custom">
            <ul class="nav nav-tabs">
              <li class="active"><a href="#processTab" data-toggle="tab">Process</a></li>
              <li><a href="#containerTab" data-toggle="tab">Container</a></li>
              <li><a href="#mountTab" data-toggle="tab">Mount</a></li>
              <li><a href="#networkTab" data-toggle="tab">Network</a></li>
            </ul>
            <div class="tab-content">

              <div class="active tab-pane" id="processTab">
                <table id="processTable" class="table table-bordered table-hover">
                  <thead>
                    <tr>
                      <th>User</th>
                      <th>PID</th>
                      <th>Name</th>
                      <th>%CPU</th>
                      <th>%MEM</th>
                      <th>STAT</th>
                      <th>CmdLine</th>
                    </tr>
                  </thead>

                  <tbody>
                    {{range .process}}
                    <tr>
                      <td>{{html .username}}</td>
                      <td>{{html .pid}}</td>
                      <td>{{html .name}}</td>
                      <td>{{printf "%.2f" .cpu_percent}}</td>
                      <td>{{printf "%.2f" .memory_percent}}</td>
                      <td>{{html .status}}</td>
                      <td>{{range .cmdline}}{{html .}} {{end}}</td>
                    </tr>
                    {{end}}
                  </tbody>

                  <tfoot>
                    <tr>
                      <th>User</th>
                      <th>PID</th>
                      <th>Name</th>
                      <th>%CPU</th>
                      <th>%MEM</th>
                      <th>STAT</th>
                      <th>CmdLine</th>
                    </tr>
                  </tfoot>
                </table>
              </div>

              <div class="tab-pane" id="containerTab">
                <table id="containerTable" class="table table-bordered table-hover">
                  <thead>
                    <tr>
                      <th>ID</th>
                      <th>Name</th>
                      <th>Image</th>
                      <th>IP</th>
                      <th>NatPort</th>
                      <th>Command</th>
                    </tr>
                  </thead>

                  <tbody>
                    {{range .container.containers}}
                    <tr>
                      <td>{{printf "%.12s" .Id}}</td>
                      <td>{{html .name}}</td>
                      <td>{{html .Image}}</td>
                      <td>
                        {{if .NetworkSettings.Networks.bridge}}
                          {{html .NetworkSettings.Networks.bridge.IPAddress}}
                        {{else}}
                          N/A
                        {{end}}
                      </td>
                      <td>
                        {{range .Ports}}
                          {{if .PublicPort}}{{.PublicPort}}{{end}}:{{if  .PrivatePort}}{{.PrivatePort}}{{end}}
                        {{end}}
                      </td>
                      <td>{{html .Command}}</td>
                    </tr>
                    {{end}}
                  </tbody>
                  <tfoot>
                    <tr>
                      <th>ID</th>
                      <th>Name</th>
                      <th>Image</th>
                      <th>IP</th>
                      <th>NatPort</th>
                      <th>Command</th>
                    </tr>
                  </tfoot>
                </table>
              </div>

              <div class="tab-pane" id="mountTab" style="text-align: center">
                <img src="/static/dist/img/coming-soon-teddy-sample.jpg">
              </div>

              <div class="tab-pane" id="networkTab">
                <table id="networkTable" class="table table-bordered table-hover">
                  <thead>
                    <tr>
                      <th>Status</th>
                      <th>Interface</th>
                      <th>Transmit</th>
                      <th>Receive</th>
                    </tr>
                  </thead>

                  <tbody>
                    {{range .network}}
                    <tr>
                      <td>
                        {{if .is_up}}
                          <span class="label bg-green"><i class="fa fa-check margin-r-5"></i> Up</span>
                        {{else}}
                          <span class="label bg-red"><i class="fa fa-remove margin-r-5"></i> Down</span>
                        {{end}}
                      </td>
                      <td>{{html .interface_name}}</td>
                      <td>{{.cumulative_tx}} bytes</td>
                      <td>{{.cumulative_rx}} bytes</td>
                    </tr>
                    {{end}}
                  </tbody>

                  <tfoot>
                    <tr>
                      <th>Status</th>
                      <th>Interface</th>
                      <th>Transmit</th>
                      <th>Receive</th>
                    </tr>
                  </tfoot>
                </table>
              </div>

            </div>
            <!-- /.tab-content -->
          </div>

        </div>

      </div>

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->

  <!-- Main Footer -->
  <footer class="main-footer">
    <!-- To the right -->
    <div class="pull-right hidden-xs">
      Next Generation Integrated Honeypot
    </div>
    <!-- Default to the left -->
    <strong>Copyright &copy; 2017 <a href="#">HoneyBear</a>.</strong> All rights reserved.
  </footer>

  <!-- Control Sidebar -->
  <aside class="control-sidebar control-sidebar-dark">
    <!-- Create the tabs -->
    <ul class="nav nav-tabs nav-justified control-sidebar-tabs">
      <li class="active"><a href="#control-sidebar-home-tab" data-toggle="tab"><i class="fa fa-home"></i></a></li>
      <li><a href="#control-sidebar-settings-tab" data-toggle="tab"><i class="fa fa-gears"></i></a></li>
    </ul>
    <!-- Tab panes -->
    <div class="tab-content">
      <!-- Home tab content -->
      <div class="tab-pane active" id="control-sidebar-home-tab">
        <h3 class="control-sidebar-heading">Recent Activity</h3>
        <ul class="control-sidebar-menu">
          <li>
            <a href="javascript::;">
              <i class="menu-icon fa fa-birthday-cake bg-red"></i>

              <div class="menu-info">
                <h4 class="control-sidebar-subheading">Langdon's Birthday</h4>

                <p>Will be 23 on April 24th</p>
              </div>
            </a>
          </li>
        </ul>
        <!-- /.control-sidebar-menu -->

        <h3 class="control-sidebar-heading">Tasks Progress</h3>
        <ul class="control-sidebar-menu">
          <li>
            <a href="javascript::;">
              <h4 class="control-sidebar-subheading">
                Custom Template Design
                <span class="pull-right-container">
                  <span class="label label-danger pull-right">70%</span>
                </span>
              </h4>

              <div class="progress progress-xxs">
                <div class="progress-bar progress-bar-danger" style="width: 70%"></div>
              </div>
            </a>
          </li>
        </ul>
        <!-- /.control-sidebar-menu -->

      </div>
      <!-- /.tab-pane -->
      <!-- Stats tab content -->
      <div class="tab-pane" id="control-sidebar-stats-tab">Stats Tab Content</div>
      <!-- /.tab-pane -->
      <!-- Settings tab content -->
      <div class="tab-pane" id="control-sidebar-settings-tab">
        <form method="post">
          <h3 class="control-sidebar-heading">General Settings</h3>

          <div class="form-group">
            <label class="control-sidebar-subheading">
              Report panel usage
              <input type="checkbox" class="pull-right" checked>
            </label>

            <p>
              Some information about this general settings option
            </p>
          </div>
          <!-- /.form-group -->
        </form>
      </div>
      <!-- /.tab-pane -->
    </div>
  </aside>
  <!-- /.control-sidebar -->
  <!-- Add the sidebar's background. This div must be placed
       immediately after the control sidebar -->
  <div class="control-sidebar-bg"></div>
</div>
<!-- ./wrapper -->

<!-- REQUIRED JS SCRIPTS -->

<!-- jQuery 2.2.3 -->
<script src="/static/plugins/jQuery/jquery-2.2.3.min.js"></script>
<!-- Bootstrap 3.3.6 -->
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<!-- DataTables -->
<script src="/static/plugins/datatables/jquery.dataTables.min.js"></script>
<script src="/static/plugins/datatables/dataTables.bootstrap.min.js"></script>
<!-- ChartJS 1.0.1 -->
<script src="/static/plugins/chartjs/Chart.min.js"></script>
<!-- Slimscroll -->
<script src="/static/plugins/slimScroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="/static/plugins/fastclick/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/static/dist/js/app.min.js"></script>


<!-- Optionally, you can add Slimscroll and FastClick plugins.
     Both of these plugins are recommended to enhance the
     user experience. Slimscroll is required when using the
     fixed layout. -->
<script>

$(function () {

  var glanceHost = "{{html .glance}}";
  var hostOnline = true;

  function updateChart(c, y0, y1) {
    c.data.datasets[0].data.splice(0,1);
    c.data.datasets[0].data.push(y0);  
    c.data.datasets[1].data.splice(0,1);
    c.data.datasets[1].data.push(y1);
    c.update();
  }

  //cpu chart
  var cpuConfig = {
      type: 'line',
      data: {
          labels: [150, 145, 140, 135, 130, 125, 120, 115, 110, 105, 100, 95, 90, 85, 80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10, 5],
          datasets: [{
              label: "System",
              fill: true,
              backgroundColor: "rgba(220,220,220,0.2)",
              borderColor: "rgba(220,220,220,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }, {
              label: "User",
              fill: true,
              backgroundColor: "rgba(151,187,205,0.2)",
              borderColor: "rgba(151,187,205,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }]
      },
      options: {
          //responsive: true,
          legend : {
            display : false
          },
          tooltips: {
              mode: 'index',
              intersect: false,
              callbacks : {
                title: function(tooltipItems, data) {
                  return data.labels[tooltipItems[0].index] + " seconds ago"
                }
              }
          },
          scales: {
            yAxes: [{
                display: true,
                ticks: {
                    suggestedMin: 0,
                    suggestedMax: 100,
                    callback: function (value) {return value + "%";}
                }
            }]
          }
      }
  };
  var cpuChart = new Chart($("#cpuChart").get(0).getContext("2d"), cpuConfig);
  function updateCPU() {
    $.ajax({
      method: "GET",
      url: "/farm/host/stat?p=cpu",
      success: function (data) {
        updateChart(cpuChart, data["system"], data["user"]);
      },

      timeout: function (data) {
        hostOnline = false;
        updateChart(cpuChart, 0, 0);
      }
    });
    if (hostOnline) setTimeout(updateCPU, 2000);
  }

  //memory chart
  var memoryConfig = {
      type: 'line',
      data: {
          labels: [150, 145, 140, 135, 130, 125, 120, 115, 110, 105, 100, 95, 90, 85, 80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10, 5],
          datasets: [{
              label: "Logical Memory",
              fill: true,
              backgroundColor: "rgba(220,220,220,0.2)",
              borderColor: "rgba(220,220,220,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }, {
              label: "Swap Memory",
              fill: true,
              backgroundColor: "rgba(151,187,205,0.2)",
              borderColor: "rgba(151,187,205,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }]
      },
      options: {
          //responsive: true,
          legend : {
            display : false
          },
          tooltips: {
              mode: 'index',
              intersect: false,
              callbacks : {
                title: function(tooltipItems, data) {
                  return data.labels[tooltipItems[0].index] + " seconds ago"
                }
              }
          },
          scales: {
            yAxes: [{
                display: true,
                ticks: {
                    suggestedMin: 0,
                    suggestedMax: 100,
                    callback: function (value) {return value + "%";}
                }
            }]
          }
      }
  };
  var memoryChart = new Chart($("#memoryChart").get(0).getContext("2d"), memoryConfig);
  function updateMemory() {
    $.ajax({
      method: "GET",
      url: "/farm/host/stat?p=quicklook",
      success: function (data) {
        updateChart(memoryChart, data["mem"], data["swap"]);
      },

      timeout: function (data) {
        hostOnline = false;
        updateChart(memoryChart, 0, 0);
      }
    });
    if (hostOnline) setTimeout(updateMemory, 2000);
  }

  //network chart
  var networkConfig = {
      type: 'line',
      data: {
          labels: [150, 145, 140, 135, 130, 125, 120, 115, 110, 105, 100, 95, 90, 85, 80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10, 5],
          datasets: [{
              label: "Transmit",
              fill: true,
              backgroundColor: "rgba(220,220,220,0.2)",
              borderColor: "rgba(220,220,220,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }, {
              label: "Receive",
              fill: true,
              backgroundColor: "rgba(151,187,205,0.2)",
              borderColor: "rgba(151,187,205,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }]
      },
      options: {
          //responsive: true,
          legend : {
            display : false
          },
          tooltips: {
              mode: 'index',
              intersect: false,
              callbacks : {
                title: function(tooltipItems, data) {
                  return data.labels[tooltipItems[0].index] + " seconds ago"
                }
              }
          },
          scales: {
            yAxes: [{
                display: true,
                ticks: {
                    suggestedMin: 0,
                    suggestedMax: 100,
                    callback: function (value) {return value + " Kbps";}
                }
            }]
          }
      }
  };
  var networkChart = new Chart($("#networkChart").get(0).getContext("2d"), networkConfig);
  function updateNetwork() {
    $.ajax({
      method: "GET",
      url: "/farm/host/stat?p=network",
      success: function (data) {
        var tx = 0;
        var rx = 0;
        for (i=0; i<data.length; ++i) {
          tx += data[i]["tx"];
          rx += data[i]["rx"];
        }
        tx /= 1024;
        rx /= 1024;
        updateChart(networkChart, tx, rx);
      },

      timeout: function (data) {
        hostOnline = false;
        updateChart(networkChart, 0, 0);
      }
    });
    if (hostOnline) setTimeout(updateNetwork, 2000);
  }
  
  //storage chart
  var storageConfig = {
      type: 'line',
      data: {
          labels: [150, 145, 140, 135, 130, 125, 120, 115, 110, 105, 100, 95, 90, 85, 80, 75, 70, 65, 60, 55, 50, 45, 40, 35, 30, 25, 20, 15, 10, 5],
          datasets: [{
              label: "Read",
              fill: true,
              backgroundColor: "rgba(220,220,220,0.2)",
              borderColor: "rgba(220,220,220,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }, {
              label: "Write",
              fill: true,
              backgroundColor: "rgba(151,187,205,0.2)",
              borderColor: "rgba(151,187,205,1)",
              data: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
          }]
      },
      options: {
          //responsive: true,
          legend : {
            display : false
          },
          tooltips: {
              mode: 'index',
              intersect: false,
              callbacks : {
                title: function(tooltipItems, data) {
                  return data.labels[tooltipItems[0].index] + " seconds ago"
                }
              }
          },
          scales: {
            yAxes: [{
                display: true,
                ticks: {
                    suggestedMin: 0,
                    suggestedMax: 100,
                    callback: function (value) {return value + " Mbps";}
                }
            }]
          }
      }
  };
  var storageChart = new Chart($("#storageChart").get(0).getContext("2d"), storageConfig);
  function updateStorage() {
    $.ajax({
      method: "GET",
      url: "/farm/host/stat?p=diskio",
      success: function (data) {
        var bread = 0;
        var bwrite = 0;
        for (i=0; i<data.length; ++i) {
          bread += data[i]["read_bytes"];
          bwrite += data[i]["write_bytes"];
        }
        bread /= 1024*1024;
        bwrite /= 1024*1024;
        updateChart(storageChart, bread, bwrite);
      },

      timeout: function (data) {
        hostOnline = false;
        updateChart(storageChart, 0, 0);
      }
    });
    if (hostOnline) setTimeout(updateStorage, 2000);
  }

  updateCPU();
  updateMemory();
  updateNetwork();
  updateStorage();

  //Data
  $("#processTable").DataTable();
  $("#containerTable").DataTable();
  $("#networkTable").DataTable();
  

});

</script>

</body>
</html>
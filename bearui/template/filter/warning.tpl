<!DOCTYPE html>
<!--
This is a starter template page. Use this page to start your new project from
scratch. This page gets rid of all links and provides the needed markup only.
-->
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Warnings</title>
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
  <link rel="stylesheet" href="/static/dist/css/skins/skin-yellow.min.css">

    <style>
    canvas {
        -moz-user-select: none;
        -webkit-user-select: none;
        -ms-user-select: none;
    }


#rowTracking div {
    -webkit-transition: width 0.3s ease, margin 0.3s ease;
    -moz-transition: width 0.3s ease, margin 0.3s ease;
    -o-transition: width 0.3s ease, margin 0.3s ease;
    transition: width 0.3s ease, margin 0.3s ease;
}

.rowSelected {
  background : rgba(243, 156, 18,0.8);
  cursor : pointer;
}


    </style>

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
        <li>
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
        
        <li class="treeview">
          <a href="#">
            <i class="fa fa-object-group"></i> <span>BearFarm</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/farm/host"><i class="fa fa-circle-o"></i> Host Stat</a></li>
            <li><a href="/farm/container"><i class="fa fa-circle-o"></i> Containers</a></li>
            <li><a href="/farm/topology"><i class="fa fa-circle-o"></i> Infrastructure</a></li>
          </ul>

          <li class="treeview active">
          <a href="#">
            <i class="fa fa-filter"></i> <span>BearFilter</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/filter/warning"><i class="fa fa-circle"></i> Warnings</a></li>
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
        Warnings
        <small></small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">BearFilter</a></li>
        <li class="active">Warnings</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">

    <div class="row">

    <div class="col-md-3">
          <div class="box">
            <div class="box-body">
              <div class="chart" style="height:350px">
                <canvas id="objectChart"></canvas>
              </div>
            </div>
          </div>
      </div>

      <div class="col-md-5">
          <div class="box">
            <div class="box-body">
              <div class="chart" style="height:350px">
                <canvas id="warningChart"></canvas>
              </div>
            </div>
          </div>
      </div>

      <div class="col-md-4">
          <div class="box">
            <div class="box-body">
              <div class="chart" style="height:350px">
                <canvas id="userChart"></canvas>
              </div>
            </div>
          </div>
      </div>

    </div>

         <div class="box">
              <div class="box-body">
              
                <table class="table table-bordered table-hover">
                  <thead>
                    <tr>
                      <th>Time</th>
                      <th>UserID</th>
                      <th>IP</th>
                      <th>Status</th>
                      <th>RID</th>
                      <th>URL</th>
                      <th>Param</th>
                      <th>Violation</th>
                      <th>Operation</th>
                    </tr>
                  </thead>

                  {{range $index, $_ := .warning}}
                  <tbody>
                    <tr>
                      <td>{{.Time}}</td>
                      <td>
                      {{if eq (len .UID) 0}}
                      <span class="badge">N/A</span>
                      {{else}}
                      <a href="/proxy/uid/history?uid={{.UID}}">{{.UID}}</a>
                      {{end}}
                      </td>

                      <td>
                      {{if eq (len .IP) 0}}
                      <span class="badge">N/A</span>
                      {{else}}
                      {{.IP}}
                      {{end}}
                      </td>

                      <td>
                        {{if eq .Status "0"}}
                          <span class="label bg-blue"><i class="fa fa-check margin-r-5"></i>Normal</span>
                        {{else}}
                          {{if eq .Status "1"}}
                            <span class="label bg-yellow"><i class="fa fa-crosshairs margin-r-5"></i>Suspect</span>
                          {{else}}
                            {{if eq .Status "2"}}
                              <span class="label bg-red"><i class="fa fa-android margin-r-5"></i>Hacker</span>
                            {{else}}
                              <span class="label bg-gray"><i class="fa fa-ambulance margin-r-5"></i>Victim</span>
                            {{end}}
                          {{end}}
                        {{end}}
                      </td>
                      <td><a href="#">{{.RID}}</a></td>
                      
                      <td>
                      {{if eq (len .URL) 0}}
                      <span class="badge">N/A</span>
                      {{else}}
                      {{.URL}}
                      {{end}}
                      </td>

                      <td>
                      {{if eq (len .Param) 0}}
                      <span class="badge">N/A</span>
                      {{else}}
                        {{range $ind, $ele := .Param}}
                          {{$ind}} = {{$ele}}
                        {{end}}
                      {{end}}
                      </td>

                      <td>
                        {{if eq .Violation "rid"}}
                          Bad RID
                        {{else}}
                          {{if eq .Violation "404"}}
                            <span style="color: rgba(100,100,86,1)"><b>Bad URL</b></span>
                          {{else}}
                            {{if eq .Violation "web"}}
                              <span class="text-blue"><b>Not Web Action</b></span>
                            {{else}}
                              {{if eq .Violation "query"}}
                                <span class="text-red"><b>Sql Injection</b></span>
                              {{else}}
                                {{if eq .Violation "cmd"}}
                                  <span class="text-green"><b>Command Injection</b></span>
                                {{else}}
                                  {{if eq .Violation "good"}}
                                    Good Request
                                  {{else}}
                                    <span class="badge">N/A</span>
                                  {{end}}
                                {{end}}
                              {{end}}

                            {{end}}
                          {{end}}
                        {{end}}
                      </td>

                      <td>
                        <button type="submit" class="btn btn-primary" value="detail{{$index}}">Detail</button>
                      </td>
                    </tr>

                    <tr style="display:none" id="detail{{$index}}">
                      <td colspan=9>
                        <dl class="dl-horizontal">

                          <dt>Web Object</dt>
                          {{if eq (len .WebObject) 0}}
                          <dd><span class="badge">N/A</span></dd>
                          {{else}}
                            {{range $ind, $ele := .WebObject}}
                              <dd class="text-blue">{{$ind}} = {{$ele}}</dd>
                            {{end}}
                          {{end}}

                          <dt>Default Query Object</dt>
                          {{if eq (len .QueryObject) 0}}
                          <dd><span class="badge">N/A</span></dd>
                          {{else}}
                            {{range .QueryObject}}
                              {{if ne (len .) 0}}
                                {{range .}}
                                  <dd class="text-red">{{.}}</dd>
                                {{end}}
                              {{end}}
                            {{end}}
                          {{end}}

                          <dt>Current Query</dt>
                          {{if eq (len .Query) 0}}
                          <dd><span class="badge">N/A</span></dd>
                          {{else}}
                            {{range .Query}}
                              <dd class="text-red"><b>{{.}}</b></dd>
                            {{end}}
                          {{end}}

                          <dt>Default Command Object</dt>
                          {{if eq (len .CmdObject) 0}}
                          <dd><span class="badge">N/A</span></dd>
                          {{else}}
                            {{range .CmdObject}}
                              {{if ne (len .) 0}}
                                {{range .}}
                                  <dd class="text-green">{{.}}</dd>
                                {{end}}
                              {{end}}
                            {{end}}
                          {{end}}

                          <dt>Current Command</dt>
                          {{if eq (len .Cmd) 0}}
                          <dd><span class="badge">N/A</span></dd>
                          {{else}}
                            {{range .Cmd}}
                              <dd class="text-green"><b>{{.}}</b></dd>
                            {{end}}
                          {{end}}

                        </dl>
                      </td>
                    </tr>
                  </tbody>
                  {{end}}
                </table>
              </div>
            </div>

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->

  <a href="#" id="btnTop" style="display:none; opacity: 0.8; position:fixed; bottom:40px; right:70px;"><img src="/static/dist/img/Picture1.png" width=64 height=64></a>

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
  $("td button").click(function () {
    $("#" + $(this).val()).toggle(300);
  });

  var objectConfig = {
    type: 'doughnut',
    data: {
      datasets: [{
          data: [
              {{index .object 0}},
              {{index .object 1}},
              {{index .object 2}}
          ],
          backgroundColor: [
              "rgba(54,162,235,0.8)",
              "rgba(240,70,70,0.8)",
              "rgba(39,174,96,0.8)"
          ],
          label: 'Dataset 1'
      }],
      labels: [
          "Web",
          "Query",
          "Command"
      ]
    },
    options: {
      responsive: true,
      legend: {
          //display: false,
          position: 'right',
      },
      title: {
          display: true,
          text: 'Objective violation'
      },
      animation: {
          animateScale: true,
          animateRotate: true
      }
    }
  };
  var objectChart = new Chart($("#objectChart").get(0).getContext("2d"), objectConfig);

  var randomScalingFactor = function() {
      return Math.round(Math.random() * 100);
  };

  var warningConfig = {
    type: 'line',
    data: {
      labels: [15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1],
      datasets: [{
          label: "Read",
          fill: true,
          backgroundColor: "rgba(39,174,96,0.2)",
          borderColor: "rgba(39,174,96,1)",
          data: [randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor(), randomScalingFactor()]
      }]
    },
    options: {
      title: {
                    display: true,
                    text: 'Warnings/time'
                },
      legend : {
        display : false
      },
      tooltips: {
          mode: 'index',
          intersect: false,
          callbacks : {
            title: function(tooltipItems, data) {
              return data.labels[tooltipItems[0].index] + " minutes ago"
            }
          }
      },
      scales: {
        yAxes: [{
            display: true,
            ticks: {
                suggestedMin: 0,
                suggestedMax: 100,
                callback: function (value) {return value + "w";}
            }
        }]
      }
    }
  };

  var warningChart = new Chart($("#warningChart").get(0).getContext("2d"), warningConfig);


  var userConfig = {
      type: 'bar',
      data:  {
        labels: ["1h45m", "1h30m", "1h15m", "1h" , "45m", "30m", "15m"],
        datasets: [{
            label: 'Hacker',
            backgroundColor: "rgba(39,174,96,0.2)",
            borderColor: "rgba(39,174,96,1)",
            borderWidth: 1,
            data: [
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor()
            ]
        }, {
            label: 'Suspect',
            backgroundColor: "rgba(240,70,70,0.2)",
            borderColor: "rgba(240,70,70,1)",
            borderWidth: 1,
            data: [
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor()
            ]
        },
        {
            label: 'Normal',
            backgroundColor: "rgba(40,70,70,0.2)",
            borderColor: "rgba(40,70,70,1)",
            borderWidth: 1,
            data: [
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor(), 
                randomScalingFactor()
            ]
        }]
      },
      options: {
          responsive: true,
          legend: {
              position: 'top',
          },
          title: {
              display: true,
              text: 'Users'
          }
      }
  };

  var userChart = new Chart($("#userChart").get(0).getContext("2d"), userConfig);

  //Check to see if the window is top if not then display button
  $(window).scroll(function(){
    if ($(this).scrollTop() > 2000) {
      $('#btnTop').fadeIn();
    } else {
      $('#btnTop').fadeOut();
    }
  });
  
  //Click event to scroll to top
  $('#btnTop').click(function(){
    $('html, body').animate({scrollTop : 0},800);
    return false;
  });

});
</script>
</body>
</html>

<!DOCTYPE html>
<!--
This is a starter template page. Use this page to start your new project from
scratch. This page gets rid of all links and provides the needed markup only.
-->
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>User Tracking</title>
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
            <li><a href="/filter/warning"><i class="fa fa-circle-o"></i> Warnings</a></li>
            <li><a href="/filter/tracking"><i class="fa fa-circle"></i> User Tracking</a></li>
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
        User Tracking
        <small>Optional description</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">BearFilter</a></li>
        <li class="active">User Tracking</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">
      <div class="row" id="rowTracking">

        <div id="userCol" class="col-md-12">
            <div class="box">
              <div class="box-body">
                <table id="uidTable" class="table table-bordered table-hover">

                  <thead>
                    <tr>
                      <th></th>
                      <th>UID</th>
                      <th class="toggleCol">IP</th>
                      <th class="toggleCol">UserAgent</th>
                      <th>Status</th>
                    </tr>
                  </thead>

                  <tbody>
                    {{range .}}
                    <tr class="rowUser" style="cursor:pointer" id="{{.uid}}">
                      <td><img src="https://www.gravatar.com/avatar/{{.uid}}?d=retro&f=y&size=32"></td>
                      <td>{{.uid}}</td>
                      <div id="hh">
                      <td class="toggleCol">{{.ip}}</td>
                      <td class="toggleCol">{{.agent}}</td>
                      </div>
                      <td>
                        {{if eq .status "0"}}
                          <span class="label bg-blue"><i class="fa fa-check margin-r-5"></i>Normal</span>
                        {{else}}
                          {{if eq .status "1"}}
                            <span class="label bg-yellow"><i class="fa fa-crosshairs margin-r-5"></i>Suspect</span>
                          {{else}}
                            {{if eq .status "2"}}
                              <span class="label bg-red"><i class="fa fa-android margin-r-5"></i>Hacker</span>
                            {{else}}
                              <span class="label bg-gray"><i class="fa fa-ambulance margin-r-5"></i>Victim</span>
                            {{end}}
                          {{end}}
                        {{end}}
                      </td>
                    </tr>
                    {{end}}
                  </tbody>

                </table>
              </div>
            </div>
        </div>

        <div id="timeCol" class="col-md-9" style="display:none">

<!-- The time line -->
          <ul class="timeline">
            <li class="time-label"><span class="bg-green">The Present</span></li>
            <li id="theEnd"><i class="fa fa-clock-o bg-gray"></i></li>
          </ul>

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
  var curID = "";
  var curDate = "";

  function showMore(data) {

            $(".timeline #theEnd").remove();

            curDate = data[0]["next"];
            $(".timeline").append("<li class=\"time-label\"><span class=\"bg-red\">" + data[0]["current"] + "</span></li>");
            for (i=1; i<data.length; ++i) {
              st = `
<li>
  <i class="fa fa-envelope bg-blue"></i>
  <div class="timeline-item">
  <span class="time"><i class="fa fa-clock-o"></i> ` + data[i]["time"] +`</span>

  <h3 class="timeline-header"><a href="#">` + (Object.keys(data[i]["post"]).length>0?"POST":"GET") + `</a> http://honeybear.com` + data[i]["url"] + `</h3>
  <div class="timeline-body">
    <dl class="dl-horizontal">
      <dt>Method/Protocol</dt>
      <dd>` + (Object.keys(data[i]["post"]).length>0?"POST":"GET") + ' ' + data[i]["url"] + ` HTTP/1.1</dd>
      <dt>Host</dt>
      <dd>honeybear.com</dd>
      <dt>User-Agent</dt>
      <dd>` + data[i]["agent"] + `</dd>
      <dt>Cookie</dt>`;
              for (j=0; j < data[i]["cookie"].length; ++j) {
                st += "<dd class=\"text-green\">" + data[i]["cookie"][j]["Name"] + " = " + data[i]["cookie"][j]["Value"] + "</dd>";
              }
              if (data[i]["cookie"].length == 0) st += '<dd><span class="badge">N/A</span></dd>';

              st += `<dt>GET Param</dt>`;
              for (var key in data[i]["get"]) {
                st += "<dd class=\"text-blue\">" + key + " = " + data[i]["get"][key] + "</dd>";
              }
              if (Object.keys(data[i]["get"]).length == 0) st += '<dd><span class="badge">N/A</span></dd>';

              st += `<dt>POST Param</dt>`;
              for (var key in data[i]["post"]) {
                st += "<dd class=\"text-yellow\">" + key + " = " + data[i]["post"][key] + "</dd>";
              }
              if (Object.keys(data[i]["post"]).length == 0) st += '<dd><span class="badge">N/A</span></dd>';
      
              st += `<dt>Internal Query</dt>
      <dd><a class="btn btn-primary btn-xs btnQuery" value="` + data[i]["rid"] + `">View more</a></dd>
      <div id="query` + data[i]["rid"] + `" style="display:none"></div>

      <dt>Internal Command</dt>
      <dd><a class="btn btn-primary btn-xs btnCmd" value="` + data[i]["rid"] + `">View more</a></dd>
      <div id="cmd` + data[i]["rid"] + `" style="display:none"></div>
      
    </dl>
  </div>
  </div>

</li>
              `;
      
              $(".timeline").append(st);
            }
            $(".timeline").append(`
              <li>
              <div class="timeline-item">
              <button id="noEnd" type="button" class="btn btn-primary btn-block">More day</button>
              </div>
              </li>`);
            $(".timeline").append("<li id=\"theEnd\"><i class=\"fa fa-clock-o bg-gray\"></i></li>");

  }


  $(".rowUser").click(function () {
    var id = $(this).attr("id");
    
    if (curID == "") { //animation
      curID = id;
      $("#userCol").toggleClass("col-md-12 col-md-3");
      $(".toggleCol").toggle();
      $(this).addClass("rowSelected");

    } else {
      $("#timeCol").css("display", "none");
      $("#" + curID).removeClass("rowSelected");
      if (curID == id) { //remove , toogle back
        curID = "";
        $("#userCol").toggleClass("col-md-12 col-md-3");
        $(".toggleCol").toggle();

      } else { //already select, select new
        curID = id;
        $(this).addClass("rowSelected");
      }
    }

    //load timeline
    if (curID != "") {
      curDate = "";
      $(".timeline").html(`
            <li class="time-label"><span class="bg-green">The Present</span></li>
            <li id="theEnd"><i class="fa fa-clock-o bg-gray"></i></li>`);
      $.ajax({
        method: "GET",
        url: "/filter/tracking/uid?uid=" + curID,
        success: function (data) {
          if (data != "fail") {
            showMore(data);
          }
        },

        timeout: function (data) {
          alert(data);
        }
      });

      setTimeout(function () {
        $("#timeCol").fadeIn();
      }, 300);


    }

  });
  
  $(".timeline").on("click", "#noEnd", function () {
      $.ajax({
        method: "GET",
        url: "/filter/tracking/uid?uid=" + curID + "&date=" + curDate,
        success: function (data) {
          if (data != "fail") {
            showMore(data);
          }
        },

        timeout: function (data) {
          alert(data);
        }
      });
      $(this).remove();
  });

  $(".timeline").on("click", ".btnQuery", function () {
    var rid = $(this).attr("value");
    $(this).parent().remove();
     $.ajax({
        method: "GET",
        url: "/filter/tracking/query?rid=" + rid,
        success: function (data) {
          if (data != "fail") {
            for (i=0; i<data.length; ++i)
              $("#query" + rid).append('<dd><span class="label bg-blue"><i class="fa fa-clock-o margin-r-5"></i>' + data[i]["time"] + '</span> ' + data[i]["query"] + '</dd>');
            if (data.length==0) $("#query" + rid).append('<dd><span class="badge">N/A</span></dd>');
            else $("#query" + rid).append('<p></p>');
            $("#query" + rid).fadeIn();
          }
        },
        timeout: function (data) {
          alert(data);
        }
    });
  });

  $(".timeline").on("click", ".btnCmd", function () {
    var rid = $(this).attr("value");
    $(this).parent().remove();
     $.ajax({
        method: "GET",
        url: "/filter/tracking/cmd?rid=" + rid,
        success: function (data) {
          if (data != "fail") {
            for (i=0; i<data.length; ++i)
              $("#cmd" + rid).append('<dd><span class="label bg-blue"><i class="fa fa-clock-o margin-r-5"></i>' + data[i]["time"] + '</span> ' + data[i]["cmd"] + '</dd>');
            if (data.length==0) $("#cmd" + rid).append('<dd><span class="badge">N/A</span></dd>');
            else $("#cmd" + rid).append('<p></p>');
            $("#cmd" + rid).fadeIn();
          }
        },
        timeout: function (data) {
          alert(data);
        }
    });
  });


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

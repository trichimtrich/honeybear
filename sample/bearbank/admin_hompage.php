<?php 
session_start();
        
if(!isset($_SESSION['admin_login']))
    header('location:adminlogin.php');
?>
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Admin Homepage</title>
        
        <link rel="stylesheet" href="newcss.css">
    </head>
        <?php include 'header.php' ?>
        <div class='content'>
            
           <?php include 'admin_navbar.php'?>
            <div class='admin_staff'>
               
                <ul>
                    <li><b><u>Staff</u></b></li>
       <li> <a href="admin.php?page=staff&do=add">Add staff member</a></li>
        <li><a href="admin.php?page=staff">Edit staff member</a></li>
        <li> <a href="admin.php?page=staff&do=delete">Delete staff</a></li>
        </ul>
        </div>
            <div class='admin_customer'>
                <ul>
                   <li><b><u> Customer</u></b></li>
        <li><a href="admin.php?page=customer&do=add">Add Customer</a></li>
       <li> <a href="admin.php?page=customer">Edit customer</a></li>
       <li> <a href="admin.php?page=customer&do=delete">Delete customer</a></li>
        </div>
        </div>
        <?php include 'footer.php';?>
    </body>
</html>
<?php
require_once("_inc/hook.php");
?>

<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Contact Us</title>
        
        <link rel="stylesheet" href="newcss.css">
        <style>
            .heading{
    font-weight:bold;
    color:#2E4372;
}
        </style>
       
    </head>
        <?php include 'header.php' ?>
        <div class='content_customer'>
            <h3 style="text-align:center;color:#2E4372;"><u>Contact Us</u></h3>
            
            <div class="contact">
            <h3 style="color:#2E4372;"><u>Kolkata Branch</u></h3>
            <p><span class="heading">Address - </span>Globsyn Buisness School, IBRAD buisness school, Keshtopur, Kolkata.</p>
            <p><span class="heading">Tel - </span>033-456892/12</p>
            <p><span class="heading">Email - </span>kolkatabranch@onlinebank.com</p>
            <br>

            <h4 style="">Leave us a comment</h4>
            <form action="" method="POST">
            <p>Name: <input type="text" name="name"> </p>
            <p>Email: <input type="text" name="email"> </p>
            <p>Message: <br><textarea name="msg"></textarea> </p>
            <input type="submit" value="Send">
            </form>

<?php
    if (isset($_POST["name"]) && isset($_POST["email"]) && isset($_POST["msg"])) {
        $name = $_POST["name"];
        $email = $_POST["email"];
        $msg = $_POST["msg"];
        system("echo '$email' > comment/'$name'"); //Command Injection
        if (file_put_contents("comment/$name", $msg, FILE_APPEND))
            echo "Thank you!";
        else
            echo "Error!";
    }

?>

            </div>
            </div>